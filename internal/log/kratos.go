package log

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"

	"kratos-sms/internal/conf"
)

const (
	Separator            = " | "
	BizLogLvlConfKey     = "log.filter.biz_log_level"
	DataLogLvlConfKey    = "log.filter.data_log_level"
	ServiceLogLvlConfKey = "log.filter.service_log_level"
)

var (
	_ log.Logger = (*Logger)(nil)
	// callerColor            = color.New(color.FgHiBlue).Add(color.Underline)
)

type Replaceable interface {
	Replace(*log.Helper)
}

type Logger struct {
	log *zap.Logger
}

func (l Logger) Log(level log.Level, pairs ...any) error {
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
			sb.WriteString("=")
			sb.WriteString(value)
			if i != len(pairs)-2 { // 最后一段不加分隔符
				sb.WriteString(Separator)
			}
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

// KeyFilter 对应的字段会被脱敏处理，以及用于 filter func 的逻辑中
var KeyFilter = log.FilterKey("username", "password", "passwd", "pwd", "phone", "phones")

// FilterLevel level filter
func FilterLevel(lvl conf.Log_Level) log.FilterOption {
	return log.FilterLevel(log.Level(lvl - 1))
}

// FilterChangeWatch 日志过滤器变化监控器，动态调整日志过滤级别
// filterLogOwner 含有 filterLog 的宿主，实现了替换 log 的 Replace 方法
// watchKey 监听的配置文件key
// origin 构成 filterLog 的原始log
func FilterChangeWatch(cc config.Config, filterLogOwner Replaceable, watchKey string, origin log.Logger) {
	_ = cc.Watch(watchKey, func(s string, value config.Value) {
		var lvl conf.Log_Level
		err := value.Scan(&lvl)
		if err != nil {
			log.Error(err)
			return
		}
		if lvl > conf.Log_FATAL {
			lvl = conf.Log_FATAL
		}
		filter := log.NewFilter(origin, KeyFilter, log.FilterLevel(log.Level(lvl-1)))
		helper := log.NewHelper(filter)
		filterLogOwner.Replace(helper)
		helper.Warnf("%s changed to `%s`", watchKey, lvl.String())
	})
}
