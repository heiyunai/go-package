package clog

import (
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	DefaultLogOath     = ""
	DefaultLogFileName = "app.log"
)

func NewLog(conf *Config) log.Logger {

	z := zap.New(zapcore.NewCore(
		conf.getEncoder(conf.getEncoderConfig()),
		conf.getWriter(conf.getLogger()),
		conf.getLevel(),
	))

	return kratoszap.NewLogger(z)
}

type Config struct {
	LogEnv            string // 开发环境 Prod Dev
	LogLevel          string // 日志打印级别 debug  info  warning  error
	LogFormat         string // 输出日志格式	logfmt, json
	LogPath           string // 输出日志文件路径
	LogFileName       string // 输出日志文件名称
	LogFileMaxSize    int    // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    // 日志保留时间，单位: 天 (day)
	LogCompress       bool   // 是否压缩日志
	LogStdout         bool   // 是否输出到控制台
}

func (c *Config) getEncoderConfig() zapcore.EncoderConfig {
	switch c.LogEnv {
	case "prod":
		return zap.NewProductionEncoderConfig()
	default:
		return zap.NewDevelopmentEncoderConfig()
	}
}

func (c *Config) getEncoder(config zapcore.EncoderConfig) zapcore.Encoder {
	switch c.LogFormat {
	case "json":
		return zapcore.NewJSONEncoder(config)
	default:
		return zapcore.NewConsoleEncoder(config)
	}
}

func (c *Config) getLevel() zapcore.Level {
	switch c.LogLevel {
	case "debuf":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func (c *Config) getWriter(logger zapcore.WriteSyncer) zapcore.WriteSyncer {
	if c.LogStdout {
		return zapcore.NewMultiWriteSyncer(logger, zapcore.AddSync(os.Stdout))
	} else {
		return zapcore.AddSync(logger)
	}
}

func (c *Config) getLogger() zapcore.WriteSyncer {
	logger := &lumberjack.Logger{
		Filename:   path.Join(c.LogPath, c.LogFileName),
		MaxSize:    c.LogFileMaxSize,
		MaxAge:     c.LogMaxAge,
		MaxBackups: c.LogFileMaxBackups,
		Compress:   c.LogCompress,
	}

	if len(logger.Filename) == 0 {
		logger.Filename = path.Join(DefaultLogOath, DefaultLogFileName)
	}

	return zapcore.AddSync(logger)
}
