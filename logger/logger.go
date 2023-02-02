package logger

import (
	"github.com/jclem/template-gofiber/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a new logger based on the app environment.
func New(cfg *config.Config) (*zap.SugaredLogger, error) {
	var config zap.Config

	if !cfg.IsProd() && !cfg.IsDev() {
		return zap.NewNop().Sugar(), nil
	}

	if cfg.IsProd() {
		config = zap.NewProductionConfig()
	} else if cfg.IsDev() {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
