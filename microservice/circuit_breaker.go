package microservice

import (
	"fmt"
	"strings"
	"time"

	"github.com/sony/gobreaker"
)

type BreakerConfig struct {
	Name                     string
	MaxRequests              uint32
	Interval                 time.Duration
	Timeout                  time.Duration
	TotalRequestCheckpoint   uint32
	MaxRatioRequestToFailure float64
}

type Breaker struct {
	*gobreaker.CircuitBreaker
}

func (b *Breaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	return b.CircuitBreaker.Execute(req)
}

func NewCircuitBreaker(config *BreakerConfig) *Breaker {
	defConfig := DefaultBreakerSetting("default-breaker", 30*time.Second)
	if config == nil {
		config = defConfig
	}
	if config.Name == "" {
		config.Name = defConfig.Name
	}
	if config.MaxRequests <= 0 {
		config.MaxRequests = defConfig.MaxRequests
	}
	if config.Timeout <= 0 {
		config.Timeout = defConfig.Timeout
	}
	if config.Interval <= 0 {
		config.Interval = defConfig.Interval
	}
	if config.TotalRequestCheckpoint <= 0 {
		config.TotalRequestCheckpoint = defConfig.TotalRequestCheckpoint
	}
	if config.MaxRatioRequestToFailure <= 0 {
		config.MaxRatioRequestToFailure = defConfig.MaxRatioRequestToFailure
	}

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:          config.Name,
		MaxRequests:   config.MaxRequests,
		Interval:      config.Interval,
		Timeout:       config.Timeout,
		ReadyToTrip:   readyToTripWithConfig(config.TotalRequestCheckpoint, config.MaxRatioRequestToFailure),
		OnStateChange: defaultOnStateChange,
		IsSuccessful:  defaultIsSuccessful,
	})

	return &Breaker{
		CircuitBreaker: cb,
	}
}

func DefaultBreakerSetting(name string, timeout time.Duration) *BreakerConfig {
	return &BreakerConfig{
		Name:                     name,
		MaxRequests:              10,
		Interval:                 2 * timeout,
		Timeout:                  timeout,
		TotalRequestCheckpoint:   100,
		MaxRatioRequestToFailure: 0.6,
	}
}

func defaultIsSuccessful(err error) bool {
	if err == nil || strings.Contains(err.Error(), "desc = OK: HTTP status code 200") {
		return true
	}
	return false
}

func defaultOnStateChange(name string, from gobreaker.State, to gobreaker.State) {
	fmt.Println(name, "breaker state changed from", from, "to", to)
}

func readyToTripWithConfig(totalReq uint32, ratioReqFail float64) func(gobreaker.Counts) bool {
	return func(c gobreaker.Counts) bool {
		failureRatio := float64(c.TotalFailures) / float64(c.Requests)
		return c.Requests >= totalReq && failureRatio >= ratioReqFail
	}
}
