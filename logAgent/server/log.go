package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func log() {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zapcore.EncoderConfig{},
		OutputPaths:      []string{"/tmp/zap.log"},
		ErrorOutputPaths: []string{"/tmp/zap.log"},
		InitialFields: map[string]interface{}{
			"app": "test",
		},
	}
	var err error
	logger, err := cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
}
