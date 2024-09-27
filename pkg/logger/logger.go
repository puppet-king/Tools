package logger

import (
	"go.uber.org/zap"
)

var (
	// Logger 全局日志器
	logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

// InitLogger 初始化日志器
func InitLogger() error {
	//logger, _ = zap.NewProduction() // 你也可以用 zap.NewDevelopment() 创建开发环境的日志器
	logger, _ = zap.NewDevelopment()
	Sugar = logger.Sugar()
	return nil
}

// Cleanup 用于在程序退出时清理日志器资源
func Cleanup() {
	if logger != nil {
		_ = logger.Sync()
	}
}
