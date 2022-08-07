package global

import (
	"github.com/Youngkingman/Kchat/kchat/pkg/logger"
	"github.com/Youngkingman/Kchat/kchat/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
