package rotatelog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestRotateLog(t *testing.T) {

	debugLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	infoLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	errorLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	debugWriter, err := NewRotateLog(
		"./log/debug/log.debug.20060102",
		WithRotateTime(time.Hour*24),
		WithMaxFileSize(1024*1024*20),
		WithCurLogLinkPath("./log/log.debug"),
		WithDeleteExpiredFile(time.Hour*24*7, "./log/debug/log.debug.*.*"),
	)
	if err != nil {
		panic(err)
	}

	infoWriter, err := NewRotateLog(
		"./log/info/log.info.20060102",
		WithRotateTime(time.Hour*24),
		WithMaxFileSize(1024*1024*20),
		WithCurLogLinkPath("./log/log.info"),
		WithDeleteExpiredFile(time.Hour*24*7, "./log/info/log.info.*.*"),
	)
	if err != nil {
		panic(err)
	}

	errorWriter, err := NewRotateLog(
		"./log/error/log.error.20060102",
		WithRotateTime(time.Hour*24),
		WithMaxFileSize(1024*1024*20),
		WithCurLogLinkPath("./log/log.error"),
		WithDeleteExpiredFile(time.Hour*24*7, "./log/error/log.error.*.*"),
	)
	if err != nil {
		panic(err)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(debugWriter), debugLvl),
		zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(infoWriter), infoLvl),
		zapcore.NewCore(getJsonEncoder(), zapcore.AddSync(errorWriter), errorLvl),
	)

	log := zap.New(core, zap.AddCaller())
	for {
		log.Debug("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Info("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Warn("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Error("hello", zap.String("time", time.Now().Format(time.DateTime)))
		//log.Fatal("hello", zap.String("time", time.Now().Format(time.DateTime)))
	}
}

func getJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}
