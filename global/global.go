package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/svcodestore/sv-resource-gin/config"
)

var (
	DB                 *gorm.DB
	DBList             map[string]*gorm.DB
	REDIS              *redis.Client
	CONFIGURATOR       config.Configurator
	CONFIG             config.Config
	LOGGER             *zap.Logger
	ConcurrencyControl = &singleflight.Group{}
)
