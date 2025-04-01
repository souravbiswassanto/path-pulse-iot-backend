package lib

import (
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZapLogger(w io.Writer, vLevel zap.AtomicLevel) *zap.Logger {
	enc := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(enc, zapcore.AddSync(w), vLevel)
	return zap.New(core, zap.AddCaller())
}

func NewLogger(verbosity int) logr.Logger {
	vLevel := convertVerbosityToZapLevel(verbosity)
	return zapr.NewLogger(initZapLogger(os.Stdout, vLevel))
}

func convertVerbosityToZapLevel(verbosity int) zap.AtomicLevel {
	if verbosity < 0 {
		verbosity = 0
	}
	if verbosity >= 4 {
		verbosity = 4
	}

	var lvl zapcore.Level
	switch verbosity {
	case 0:
		lvl = zapcore.DPanicLevel
	case 1:
		lvl = zapcore.ErrorLevel
	case 2:
		lvl = zapcore.WarnLevel
	case 4:
		lvl = zapcore.DebugLevel
	default:
		lvl = zapcore.InfoLevel
	}
	return zap.NewAtomicLevelAt(lvl)
}
