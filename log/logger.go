package log

import (
	"log"
	"os"

	logkit "github.com/go-kit/kit/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	//LogTime is log key for timestamp
	LogTime = "ts"
	//LogCaller is log key for source file name
	LogCaller = "caller"
	//LogMethod is log key for method name
	LogMethod = "method"
	//LogUser is log key for user
	LogUser = "user"
	//LogEmail is log key for email
	LogEmail = "email"
	//LogMobile is log key for mobile no
	LogMobile = "mobile"
	//LogRole is log key for role
	LogRole = "role"
	//LogTook is log key for call duration
	LogTook = "took"
	//LogInfo is log key for info
	LogInfo = "[INFO]"
	//LogDebug is log key for debug
	LogDebug = "[DEBUG]"
	//LogCritical is log key for critical
	LogCritical = "[CRITICAL]"
	//LogError is log key for error
	LogError = "[ERROR]"
	//LogBasic is log key for basic log
	LogBasic = "[BASIC]"
	//LogWarning is log key for warning log
	LogWarning = "[WARNING]"
	//LogReq is log key for request log
	LogReq = "[REQUEST]"
	//LogResp is log key for response log
	LogResp = "[RESPONSE]"
	//LogData is log key for data log
	LogData = "[DATA]"
	//LogService is log key for service name
	LogService = "service"
	//LogToken is log key for token
	LogToken = "token"
	//LogExit is log key for exit
	LogExit = "exit"
	//default file logger
	logFile = "service.log"
)

type ConfigLog struct {
	Caller int
}

// File set default log to file
func File(file string) {
	logFile := &lumberjack.Logger{
		Filename:  file,
		MaxSize:   1, // megabytes
		LocalTime: true,
		Compress:  true, // disabled by default
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)
}

// Logger returns default logger
func Logger() logkit.Logger {
	File(logFile)
	logger := logkit.NewLogfmtLogger(NewDefaultLogWriter())
	logger = logkit.With(logger, LogCaller, logkit.DefaultCaller)

	return logger
}

// StdLogger returns logger to stderr
func StdLogger() logkit.Logger {
	logger := logkit.NewLogfmtLogger(os.Stderr)
	logger = logkit.With(logger, LogTime, logkit.DefaultTimestampUTC, LogCaller, logkit.DefaultCaller)

	return logger
}

// StdLoggerConf returns logger to stderr with config
func StdLoggerConf(conf ConfigLog) logkit.Logger {
	logger := logkit.NewLogfmtLogger(os.Stderr)
	logger = logkit.With(logger, LogTime, logkit.DefaultTimestampUTC, LogCaller, logkit.Caller(conf.Caller))

	return logger
}

// FileLogger returns file logger
func FileLogger(file string) logkit.Logger {
	File(file)
	logger := logkit.NewLogfmtLogger(NewDefaultLogWriter())
	logger = logkit.With(logger, LogCaller, logkit.DefaultCaller)

	return logger
}

func StackDriverLogger() logkit.Logger {
	logger := NewSDLogger(os.Stdout)
	logger = logkit.With(logger, LogTime, logkit.DefaultTimestampUTC, LogCaller, logkit.DefaultCaller)

	return logger
}

// // ConsoleLog is console log format in service. this for helping logging in service
// type ConsoleLog struct {
// 	OrderID           int64
// 	RequestID         string
// 	VehicleNo         string
// 	DriverID          string
// 	DeviceID          int64
// 	DispatchID        uuid.UUID
// 	SioID             int64
// 	SessionID         int64
// 	VehicleStatus     []byte
// 	LastVehicleStatus []byte
// 	OrderStatus       []byte
// 	LastOrderStatus   []byte
// 	Log               string
// 	TimeStart         time.Time
// 	UserID            string
// }

// // GenerateConsoleLog is genereate log
// func (cslog *ConsoleLog) GenerateConsoleLog() {
// 	var logs []string
// 	if cslog.OrderID != 0 {
// 		logs = append(logs, fmt.Sprintf("order_id = '%d'", cslog.OrderID))
// 	}
// 	if cslog.RequestID == "" {
// 		reqUUID := uuid.New()
// 		cslog.RequestID = reqUUID.String()
// 		logs = append(logs, fmt.Sprintf("request_id = '%s'", cslog.RequestID))
// 	} else {
// 		logs = append(logs, fmt.Sprintf("request_id = '%s'", cslog.RequestID))
// 	}
// 	if cslog.VehicleNo != "" {
// 		logs = append(logs, fmt.Sprintf("vehicle_no = '%s'", cslog.VehicleNo))
// 	}
// 	if cslog.DriverID != "" {
// 		logs = append(logs, fmt.Sprintf("driver_id = '%s'", cslog.DriverID))
// 	}
// 	if cslog.DeviceID != 0 {
// 		logs = append(logs, fmt.Sprintf("device_id = '%d'", cslog.DeviceID))
// 	}
// 	if cslog.DispatchID != uuid.Nil {
// 		logs = append(logs, fmt.Sprintf("dispatch_id = '[%d:%d]'", cslog.DispatchID.MSB, cslog.DispatchID.LSB))
// 	}
// 	if cslog.SioID != 0 {
// 		logs = append(logs, fmt.Sprintf("sio_id = '%d'", cslog.SioID))
// 	}
// 	if cslog.SessionID != 0 {
// 		logs = append(logs, fmt.Sprintf("session_id = '%d'", cslog.SessionID))
// 	}
// 	if cslog.VehicleStatus != nil {
// 		vehst, _ := strconv.Atoi(string(cslog.VehicleStatus))
// 		logs = append(logs, fmt.Sprintf("vehicle_status = '%s', ", cs.GetVehicleStatusInString(vehst)))
// 	}
// 	if cslog.OrderStatus != nil {
// 		ordst, _ := strconv.Atoi(string(cslog.OrderStatus))
// 		logs = append(logs, fmt.Sprintf("order_status = '%s'", cs.GetOrderStatusInString(ordst)))
// 	}
// 	if cslog.LastVehicleStatus != nil {
// 		vehst, _ := strconv.Atoi(string(cslog.LastVehicleStatus))
// 		logs = append(logs, fmt.Sprintf("last_vehicle_status = '%s', ", cs.GetVehicleStatusInString(vehst)))
// 	}
// 	if cslog.LastOrderStatus != nil {
// 		ordst, _ := strconv.Atoi(string(cslog.LastOrderStatus))
// 		logs = append(logs, fmt.Sprintf("last_order_status = '%s'", cs.GetOrderStatusInString(ordst)))
// 	}
// 	if cslog.UserID != "" {
// 		logs = append(logs, fmt.Sprintf("user_id = '%s'", cslog.UserID))
// 	}
// 	cslog.Log = strings.Join(logs, " , ")
// }

// // GetTimeSince is get duration service running.
// func (cslog *ConsoleLog) GetTimeSince() float64 {
// 	return time.Since(cslog.TimeStart).Seconds()
// }

// func ConvertStatus(status int32) []byte {
// 	return []byte(fmt.Sprintf("%d", status))
// }
