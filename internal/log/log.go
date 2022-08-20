package log

import (
  "fmt"
  "io"
  "os"
  "strings"
  "sync"
  "time"

  "github.com/go-kratos/kratos/v2/log"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "gopkg.in/natefinch/lumberjack.v2"

  "kratos-sms/internal/conf"
)

var _ log.Logger = (*Logger)(nil)
var once sync.Once
var bootstrap *conf.Bootstrap

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
  // Struct format
  // Zap.Field is used when key-val pairs appear
  var data []zap.Field
  var key, value string
  var ok bool

  // Console format
  var sb strings.Builder
  if bootstrap.AppDebug {
    sb.Grow(256)
  }

  for i := 0; i < len(pairs); i += 2 {
    key, ok = pairs[i].(string)
    if !ok {
      key = fmt.Sprint(pairs[i])
    }
    value, ok = pairs[i+1].(string)
    if bootstrap.AppDebug {
      if !ok {
        value = fmt.Sprint(pairs[i+1])
      }
      sb.WriteString(key)
      sb.WriteString("=\"")
      sb.WriteString(value)
      sb.WriteString("\"\t")
    } else {
      if ok {
        data = append(data, zap.String(key, value))
      } else {
        data = append(data, zap.Any(key, pairs[i+1]))
      }
    }
  }

  switch level {
  case log.LevelDebug:
    if bootstrap.AppDebug {
      l.log.Debug(sb.String())
    } else {
      l.log.Debug("-", data...)
    }
  case log.LevelInfo:
    if bootstrap.AppDebug {
      l.log.Info(sb.String())
    } else {
      l.log.Info("-", data...)
    }
  case log.LevelWarn:
    if bootstrap.AppDebug {
      l.log.Warn(sb.String())
    } else {
      l.log.Warn("-", data...)
    }
  case log.LevelError:
    if bootstrap.AppDebug {
      l.log.Error(sb.String())
    } else {
      l.log.Error("-", data...)
    }
  case log.LevelFatal:
    if bootstrap.AppDebug {
      l.log.Fatal(sb.String())
    } else {
      l.log.Fatal("-", data...)
    }
  }
  return nil
}

var std = New(os.Stdout, DebugLevel, conf.Log_RFC3339, WithCaller(false), zap.AddStacktrace(ErrorLevel))
var defaultLogger = &Logger{std}

func Default() *Logger {
  return defaultLogger
}

// ProductionDefault 设置默认生产日志策略
func ProductionDefault(bs *conf.Bootstrap, opts ...Option) {
  // 日志仅初始化一次
  once.Do(func() {
    bootstrap = bs
    // Debug 模式如果开启，不接受自定义配置
    if bs.AppDebug {
      ResetDefault(New(os.Stdout, DebugLevel, conf.Log_RFC3339, WithCaller(false), zap.AddStacktrace(ErrorLevel)))
      return
    }
    var defaultLog = bs.Log.Default
    var errorLog = bs.Log.Error
    var tops = []TeeOption{
      {
        Filename:   bs.Log.Path + defaultLog.Filename,
        TimeFormat: defaultLog.TimeFormat,
        TextFormat: defaultLog.TextFormat,
        Ropt: RotateOptions{
          MaxSize:    defaultLog.MaxSize,
          MaxAge:     defaultLog.MaxAge,
          MaxBackups: defaultLog.MaxBackups,
          Compress:   defaultLog.Compress,
        },
        Level: level(defaultLog.Level),
      },
      {
        Filename:   bs.Log.Path + errorLog.Filename,
        TimeFormat: errorLog.TimeFormat,
        TextFormat: errorLog.TextFormat,
        Ropt: RotateOptions{
          MaxSize:    errorLog.MaxSize,
          MaxAge:     errorLog.MaxAge,
          MaxBackups: errorLog.MaxBackups,
          Compress:   errorLog.Compress,
        },
        Level: level(errorLog.Level),
      },
    }
    logger := NewTeeWithRotate(tops, opts...)
    ResetDefault(logger)
  })
}

type RotateOptions struct {
  MaxSize    uint32 // 单个文件最大大小, 单位MB
  MaxAge     uint32 // 文件最长保存天数
  MaxBackups uint32 // 最大文件个数
  Compress   bool   // 是否开启压缩
}

type TeeOption struct {
  Filename   string               // 日志文件名
  TextFormat conf.Log_TextFormat  // 日志文本格式 console or json
  TimeFormat conf.Log_TimeFormat  // 记录日志时，时间戳格式
  Ropt       RotateOptions        // 日志分隔轮转配置
  Level      zapcore.LevelEnabler // 日志级别生效级别
}

func NewTeeWithRotate(tops []TeeOption, opts ...Option) *zap.Logger {
  var cores []zapcore.Core
  cfg := zap.NewProductionConfig()
  cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
  cfg.EncoderConfig.TimeKey = "ts"
  cfg.EncoderConfig.MessageKey = "_zap"

  for _, top := range tops {
    top := top
    // TimePrecision
    cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
      timeFormat(top.TimeFormat, &t, enc)
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
func New(writer io.Writer, level Level, tf conf.Log_TimeFormat, opts ...Option) *zap.Logger {
  if writer == nil {
    panic("the writer is nil")
  }
  cfg := zap.NewProductionConfig()
  cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    timeFormat(tf, &t, enc)
  }
  cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
  cfg.EncoderConfig.TimeKey = "ts"
  cfg.EncoderConfig.MessageKey = "_zap"

  core := zapcore.NewCore(
    zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    zapcore.AddSync(writer),
    level,
  )
  return zap.New(core, opts...)
}

// ResetDefault not safe for concurrent use
func ResetDefault(l *zap.Logger) {
  Sync()
  std = l
  defaultLogger.log = l

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

func Sync() {
  if std != nil {
    _ = std.Sync()
  }
}

const (
  TfMillis  = "2006-01-02T15:04:05.000"
  TfSeconds = "2006-01-02T15:04:05.000"
)

func timeFormat(tf conf.Log_TimeFormat, t *time.Time, enc zapcore.PrimitiveArrayEncoder) {
  var format = TfMillis
  switch tf {
  case conf.Log_MILLIS:
    format = TfMillis
  case conf.Log_SECONDS:
    format = TfSeconds
  case conf.Log_RFC3339:
    format = time.RFC3339
  case conf.Log_RFC3339_NANO:
    format = time.RFC3339Nano
  default:
  }
  enc.AppendString(t.Format(format))
}

// map config level to zap level
func level(l conf.Log_Level) Level {
  switch l {
  case conf.Log_DEBUG:
    return zapcore.DebugLevel
  case conf.Log_INFO:
    return zapcore.InfoLevel
  case conf.Log_WARN:
    return zapcore.WarnLevel
  case conf.Log_ERROR:
    return zapcore.ErrorLevel
  case conf.Log_FATAL:
    return zapcore.FatalLevel
  }
  return zapcore.DebugLevel
}
