package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"kratos-sms/internal/conf"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *zap.Logger
}

func (l Logger) Reset(log *zap.Logger) {
	l.log = log
}

func (l Logger) Log(level log.Level, pairs ...interface{}) error {
	if len(pairs) == 0 || len(pairs)&0x1 != 0 {
		l.log.Warn(fmt.Sprint("log fields must appear in pairs: ", pairs))
		return nil
	}
	// Zap.Field is used when key-val pairs appear
	var data []zap.Field
	var key, value string
	var ok bool
	for i := 0; i < len(pairs); i += 2 {
		key, ok = pairs[i].(string)
		if !ok {
			key = fmt.Sprint(pairs[i])
		}
		value, ok = pairs[i+1].(string)
		if !ok {
			value = fmt.Sprint(pairs[i+1])
		}
		data = append(data, zap.String(key, value))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug("-", data...)
	case log.LevelInfo:
		l.log.Info("-", data...)
	case log.LevelWarn:
		l.log.Warn("-", data...)
	case log.LevelError:
		l.log.Error("-", data...)
	case log.LevelFatal:
		l.log.Fatal("-", data...)
	}
	return nil
}

func Default() *Logger {
	return &Logger{std}
}

// 使用注意事项：
// 1. 环境变量 CONF_LOG_TIME_FORMAT 用于设置日期格式，默认为: 2006-01-02T15:04:05.000
// 2. 生产环境日志策略需调用 ProductionDefault 来设置，或者参照此方法根据需要自己修改合适的日志参数
// 3. 使用 ProductionDefault 进行生产环境日志设置时，环境变量 CONF_LOG_PATH 用于设置日志路径，默认为执行程序的当前目录下的logs目录

type Level = zapcore.Level

type Field = zap.Field

var std = New(os.Stdout, DebugLevel, WithCaller(false), zap.AddStacktrace(ErrorLevel))

// ProductionDefault 设置默认生产日志策略
// 参照此方法根据需要自己修改合适的日志参数, 编写自己的初始化方法
func ProductionDefault(bs *conf.Log, opts ...Option) {
	if bs.AppDebug {
		ResetDefault(New(os.Stdout, DebugLevel, WithCaller(false), zap.AddStacktrace(ErrorLevel)))
		return
	}
	var defaultLog = bs.Default
	var errorLog = bs.Error
	var tops = []TeeOption{
		// 默认JSON格式
		{
			Filename:      BasePath() + defaultLog.Filename,
			TextFormat:    defaultLog.TextFormat,
			TimePrecision: defaultLog.TimePrecision,
			Ropt: RotateOptions{
				MaxSize:    defaultLog.MaxSize,
				MaxAge:     defaultLog.MaxAge,
				MaxBackups: defaultLog.MaxBackups,
				Compress:   defaultLog.Compress,
			},
			Level: Level(defaultLog.Level - 1),
		},
		// 设置为console格式
		{
			Filename:      BasePath() + errorLog.Filename,
			TextFormat:    errorLog.TextFormat,
			TimePrecision: errorLog.TimePrecision,
			Ropt: RotateOptions{
				MaxSize:    errorLog.MaxSize,
				MaxAge:     errorLog.MaxAge,
				MaxBackups: errorLog.MaxBackups,
				Compress:   errorLog.Compress,
			},
			Level: Level(errorLog.Level - 1),
		},
	}

	logger := NewTeeWithRotate(tops, opts...)
	ResetDefault(logger)
}

type Option = zap.Option

type RotateOptions struct {
	MaxSize    uint32 // 单个文件最大大小, 单位MB
	MaxAge     uint32 // 文件最长保存天数
	MaxBackups uint32 // 最大文件个数
	Compress   bool   // 是否开启压缩
}

type TeeOption struct {
	Filename      string                 // 日志文件名
	TimePrecision conf.Log_TimePrecision // 记录日志时，相关的时间精度，该参数选项：SECOND、MILLISECOND，分别表示 秒 和 毫秒 ,默认为毫秒级别
	TextFormat    conf.Log_TextFormat    // 日志文本格式 console or json
	Ropt          RotateOptions          // 日志分隔轮转配置
	Level         zapcore.LevelEnabler   // 日志级别生效级别
}

func NewTeeWithRotate(tops []TeeOption, opts ...Option) *zap.Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.TimeKey = "created_at"
	cfg.EncoderConfig.MessageKey = "zap"

	for _, top := range tops {
		top := top
		// TimePrecision
		cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			timeFormat(top.TimePrecision, &t, enc)
		}
		// TextFormat
		encoder := zapcore.NewJSONEncoder(cfg.EncoderConfig)
		if top.TextFormat == conf.Log_CONSOLE {
			encoder = zapcore.NewConsoleEncoder(cfg.EncoderConfig)
		}
		// 日志分隔轮转配置
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   top.Filename,
			MaxSize:    int(top.Ropt.MaxSize),
			MaxBackups: int(top.Ropt.MaxBackups),
			MaxAge:     int(top.Ropt.MaxAge),
			Compress:   top.Ropt.Compress,
		})

		core := zapcore.NewCore(encoder, zapcore.AddSync(w), top.Level)
		cores = append(cores, core)
	}

	return zap.New(zapcore.NewTee(cores...), opts...)
}

// New create a new logger (not support log rotating).
func New(writer io.Writer, level Level, opts ...Option) *zap.Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		timeFormat(conf.Log_MILLISECOND, &t, enc)
	}
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.TimeKey = "created_at"
	cfg.EncoderConfig.MessageKey = "zap"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)
	return zap.New(core, opts...)
}

func BasePath() (path string) {
	path = os.Getenv("CONF_LOG_PATH")
	if len(path) == 0 {
		path = "logs/"
		return
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return
}

func Sync() {
	if std != nil {
		_ = std.Sync()
	}
}

// ResetDefault not safe for concurrent use
func ResetDefault(l *zap.Logger) {
	std = l

	Info = std.Info
	Warn = std.Warn
	Error = std.Error
	DPanic = std.DPanic
	Panic = std.Panic
	Fatal = std.Fatal
	Debug = std.Debug

	Infof = std.Sugar().Infof
	Warnf = std.Sugar().Warnf
	Errorf = std.Sugar().Errorf
	DPanicf = std.Sugar().DPanicf
	Panicf = std.Sugar().Panicf
	Fatalf = std.Sugar().Fatalf
	Debugf = std.Sugar().Debugf
}

// 根据 TextFormat 参数 或 环境变量 LOG_TIME_FORMAT 的值来设置日期格式
func timeFormat(tf conf.Log_TimePrecision, t *time.Time, enc zapcore.PrimitiveArrayEncoder) {
	if tf > 0 {
		if tf == conf.Log_SECOND {
			enc.AppendString(t.Format("2006-01-02T15:04:05"))
		} else {
			// default
			enc.AppendString(t.Format("2006-01-02T15:04:05.000"))
		}
		// 只要该参数不为空，就采用上述两种格式之一
		return
	}
	str := os.Getenv("CONF_LOG_TIME_FORMAT")
	if len(str) == 0 {
		// default
		enc.AppendString(t.Format("2006-01-02T15:04:05.000"))
	} else {
		// custom
		enc.AppendString(t.Format(str))
	}
}
