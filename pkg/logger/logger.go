package logger

import (
	"time"

	"github.com/caarlos0/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

type LogConfig struct {
	WEBOOK_LOG_LEVEL int `env:"WEBOOK_LOG_LEVEL" envDefault:"0"`
}

func init() {
	cfg := &LogConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return
	}

	logConfig := zap.NewProductionConfig()

	logConfig.EncoderConfig.EncodeTime = customTimeEncoder
	logConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(cfg.WEBOOK_LOG_LEVEL))
	logConfig.Encoding                 = "console"  // 编码
	logConfig.DisableStacktrace        = false      // 打印堆栈
	logConfig.OutputPaths              = append(logConfig.OutputPaths, "./tmp.log")
	logConfig.InitialFields            = map[string]interface{}{"service": "小微书"}

	baseLogger, err := logConfig.Build()
	if err != nil {
		panic("failed to create the default logger: " + err.Error())
	}
	logger = baseLogger.Sugar()
}

func NewSugardLogger() *zap.SugaredLogger {
	return logger
}



// customTimeEncoder 自定义友好的时间戳格式   
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
