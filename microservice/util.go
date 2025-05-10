package microservice

import (
	"context"
	"fmt"
	stdlog "log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/LukmanulHakim18/core/flags"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/zk"
)

// etcdv3
var ctx = context.Background()
var options = etcdv3.ClientOptions{CACert: "", Cert: "", DialTimeout: time.Second * 5, DialKeepAlive: time.Second * 5}

// zkClient returns zk client
func zkClient(nodes []string, logger log.Logger) (zk.Client, error) {
	options := zk.ConnectTimeout(time.Second * 5)
	return zk.NewClient(nodes, logger, options)
}

// etcv3Client returns etcdv3 client
func etcdv3Client(nodes []string, logger log.Logger) (etcdv3.Client, error) {
	return etcdv3.NewClient(ctx, nodes, options)
}

// OnShutdown calls shutdown on signal interrupt
func OnShutdown(shutdown func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	<-c
	id := time.Now().UnixNano()
	fmt.Println("OnShutdown...", id)
	if shutdown != nil {
		shutdown()
	}
	fmt.Println("OnShutdown done", id)
}

// RecoveryHandlerFunc is a function that recovers from the panic `p` by returning an `error`.
type RecoveryHandlerFunc func(p interface{}) (err error)

// RecoverFrom call recovery handler function
func RecoverFrom(p interface{}, r RecoveryHandlerFunc) error {
	if r == nil {
		return fmt.Errorf("Server error: %s", p)
	}
	return r(p)
}

// GoWithRecover call go routine with recovery
func GoWithRecover(function func(), recoverFunc RecoveryHandlerFunc) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := RecoverFrom(r, recoverFunc)
				if err != nil {
					stdlog.Println(err)
				}
			}
		}()

		function()

	}()
}

func formatingMsgError(err error) error {
	for _, v := range listFormatingMsgErrors {
		if strings.Contains(err.Error(), v.String()) {
			return v.ToError()
		}
	}
	return err
}

// ServiceDiscovery returns zk/etcdv3 service instancer
func ServiceDiscovery(nodes []string, serviceName string, logger log.Logger) (sd.Instancer, error) {
	switch GetOsEnv(flags.BB_DISCOVERY_ENV_NAME) {
	case flags.BB_DISCOVERY_MODE_ZK:
		client, err := zkClient(nodes, logger)
		if err != nil {
			return nil, err
		}
		path := flags.SERVICE_PATH + serviceName
		instancer, err := zk.NewInstancer(client, path, logger)
		if err != nil {
			return nil, err
		}
		return instancer, nil
	default:
		//default is etcd
		client, err := etcdv3Client(nodes, logger)
		if err != nil {
			return nil, err
		}
		path := flags.SERVICE_PATH + serviceName
		instancer, err := etcdv3.NewInstancer(client, path, logger)
		if err != nil {
			return nil, err
		}
		return instancer, nil
	}
}
