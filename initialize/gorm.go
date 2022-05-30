package initialize

import (
	"gorm.io/gorm"

	"github.com/svcodestore/sv-resource-gin/config"
	"github.com/svcodestore/sv-resource-gin/utils"
)

func Gorm() *gorm.DB {
	return utils.Gorm()
}
func GormMysqlByConfig(m config.DB) *gorm.DB {
	return utils.GormMysqlByConfig(m)
}
