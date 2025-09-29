package logs

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// ตั้งเวลาเป็น Asia/Bangkok
func utcPlus7TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	t = t.In(loc)
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// init() จะรันอัตโนมัติเมื่อ import package
func init() {
	var	err error

	if os.Getenv("ENV") != "production" {
		log, err = zapLoggerDevelopment()
	} else {
		log, err = zapLoggerProduction()
	}

	if err != nil {
		panic(err)
	}
}

func zapLoggerProduction() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncoderConfig.EncodeTime = utcPlus7TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.StacktraceKey=""
	return cfg.Build(zap.AddCaller(),zap.AddCallerSkip(1))
}

func zapLoggerDevelopment() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = utcPlus7TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.StacktraceKey=""
	return cfg.Build(zap.AddCaller(),zap.AddCallerSkip(1))
}


func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type){
	case error:
		log.Error(v.Error(),fields...)
	case string:
		log.Error(v,fields...)
	}
}

func Sync() {
	_ = log.Sync()
}
