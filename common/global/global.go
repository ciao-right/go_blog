package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB   *gorm.DB
	GlobalViper *viper.Viper
)
