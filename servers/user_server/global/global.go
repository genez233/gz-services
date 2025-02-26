package global

import (
	"github.com/genez233/go-utils/glog"
	"gorm.io/gorm"
	"gz-services/servers/user_server/pkg/sett"
)

var (
	DB  *gorm.DB
	Log *glog.GLog
	// Server 服务配置
	Server *sett.ServerSetting
	// App 应用配置
	App *sett.AppSetting
	// Database 数据库配置
	Database *sett.DatabaseSetting
	// JWT 鉴权
	JWT *sett.JWTSetting
)
