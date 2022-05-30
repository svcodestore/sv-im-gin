package initialize

import (
	"github.com/svcodestore/sv-resource-gin/global"
	"gorm.io/gorm"
)

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.Dbname] = GormMysqlByConfig(info)
		default:
			continue
		}
	}

	global.DBList = dbMap
}
