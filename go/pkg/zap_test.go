package pkg

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestZap(t *testing.T) {
	zapLogger, err := getZapLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	fields := []zap.Field{zap.Namespace("context"), zap.String("str", "value"), zap.Int("int", 1)}
	zapLogger.Debug("debug", fields...)
}

func getZapLogger() (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	fileName := "./zap.log"
	format := "-%Y%m%d"
	maxAge := 7 * 24 * time.Hour
	rotationTime := 24 * time.Hour
	rotateLogs, err := rotatelogs.New(fileName+format+".log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return nil, err
	}

	level := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(rotateLogs), level),
	)

	opts := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1), zap.Development()}
	return zap.New(core, opts...), nil
}
