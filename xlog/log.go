package xlog

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var writer *lumberjack.Logger

type Conf struct {
	Interval time.Duration
	Level    string
	Logger   lumberjack.Logger
}

// viper 配置格式
// `{
//     "logger":{
//         "filename":"./log/server.log",
//         "maxsize":20000,
//         "maxage":7,
//         "maxbackups":10,
//         "compress":false
//     },
//     "interval":"24h",
//     "level":"debug"
// }`
func Init(c *viper.Viper) *lumberjack.Logger {
	var conf Conf
	err := c.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	logger = initLogger(&conf)
	return &conf.Logger
}
func initLogger(conf *Conf) *zap.Logger {
	log := &conf.Logger
	fmt.Printf("path: %s \n", log.Filename)
	go func() {
		for {
			<-time.After(conf.Interval)
			log.Rotate()
		}
	}()

	zapLevle := zapcore.InfoLevel
	l := strings.ToLower(conf.Level)
	switch l {
	case "debug":
		zapLevle = zapcore.DebugLevel
	case "info":
		zapLevle = zapcore.InfoLevel
	case "warn":
		zapLevle = zapcore.WarnLevel
	case "error":
		zapLevle = zapcore.ErrorLevel
	case "painc":
		zapLevle = zapcore.PanicLevel
	case "fatal":
		zapLevle = zapcore.FatalLevel
	}

	encoder := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	w := zapcore.AddSync(log)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		w,
		zapLevle,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger
}
