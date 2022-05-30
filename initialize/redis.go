package initialize

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/svcodestore/sv-resource-gin/global"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOGGER.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.LOGGER.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
