package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

type Config struct {
	Level      string
	File       string
	MaxSizeMB  int
	MaxBackups int
}
