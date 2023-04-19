package global

import (
	"github.com/Wpc-0601/copilot-genuine/config"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper   *viper.Viper
	Configuration config.Configuration
}

var App = new(Application)
