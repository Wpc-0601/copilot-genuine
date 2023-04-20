package global

import (
	"github.com/Wpc-0601/copilot-genuine/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper   *viper.Viper
	Configuration config.Configuration
	Log           *zap.Logger
}

var App = new(Application)
