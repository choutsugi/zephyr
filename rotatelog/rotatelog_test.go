package rotatelog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
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

	debugWriter, err := newWriter("./log/log.debug.20060102", "./log/log.debug", time.Hour*24, time.Hour*24*7, "./log/log.debug.*")
	if err != nil {
		panic(err)
	}
	infoWriter, err := newWriter("./log/log.info.20060102", "./log/log.info", time.Hour*24, time.Hour*24*7, "./log/log.info.*")
	if err != nil {
		panic(err)
	}
	errorWriter, err := newWriter("./log/log.error.20060102", "./log/log.error", time.Hour*24, time.Hour*24*7, "./log/log.error.*")
	if err != nil {
		panic(err)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(debugWriter), debugLvl),
		zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(infoWriter), infoLvl),
		zapcore.NewCore(getJsonEncoder(), zapcore.AddSync(errorWriter), errorLvl),
	)

	log := zap.New(core, zap.AddCaller())
	{
		log.Debug("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Info("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Warn("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Error("hello", zap.String("time", time.Now().Format(time.DateTime)))
		log.Fatal("hello", zap.String("time", time.Now().Format(time.DateTime)))
	}
}

func newWriter(logPath, linkPath string, rotateTime time.Duration, maxAge time.Duration, fileWildcard string) (io.Writer, error) {
	writer, err := NewRotateLog(
		logPath,
		WithRotateTime(rotateTime),
		WithCurLogLinkPath(linkPath),
		WithDeleteExpiredFile(maxAge, fileWildcard),
	)
	if err != nil {
		return nil, err
	}

	return writer, nil
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
