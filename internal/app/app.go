package app

import (
	"context"
	"fx-web/internal/conf"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func StartApp(lifecycle fx.Lifecycle, r *gin.Engine, config *conf.Config, logger *zap.Logger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// 单独启动一个协程提供服务器
			go func() {
				err := r.Run(config.HttpPort)
				if err != nil {
					logger.Error("start server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
