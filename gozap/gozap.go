package gozap

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

var (
	zapLog *zap.Logger
	once   sync.Once
)

func Info(funStr,event string)  {
	zapLog.Info("info",zap.Any("function",funStr),zap.Any("event",event))
}

func Error(err error,funStr,event string)  {
	zapLog.Error("error",zap.Any("error",err),zap.Any("function",funStr),zap.Any("event",event))
}

// InitLogger loglevel 日志级别
func InitLogger(logpath string, loglevel string){
	once.Do(func() {
		hook := lumberjack.Logger{
			Filename:   logpath, // 日志文件路径
			MaxSize:    1,       // megabytes MB
			MaxBackups: 5,       // 最多保留300个备份
			MaxAge:     7,       // days
			Compress:   false,   // 是否压缩 disabled by default
		}
		w := zapcore.AddSync(&hook)
		// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
		// debug->info->warn->error
		var level zapcore.Level
		switch loglevel {
		case "debug":
			level = zap.DebugLevel
		case "info":
			level = zap.InfoLevel
		case "error":
			level = zap.ErrorLevel
		default:
			level = zap.InfoLevel
		}
		encoderConfig := zap.NewProductionEncoderConfig()
		// 时间格式
		encoderConfig.EncodeTime = ZapTimeEncoder
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			w,
			level,
		)
		zapLog = zap.New(core)
	})
	fmt.Println("logger success")
}

func ZapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}
