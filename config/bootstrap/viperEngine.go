package bootstrap

import (
	"fmt"
	"github.com/Wpc-0601/copilot-genuine/config/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitializeConfig() *viper.Viper {
	file := "resources/application-dev.yml"
	// fetch the env variable in production
	if env := os.Getenv("VIPER-CONFIG"); env != "" {
		file = env
	}
	instance := viper.New()
	instance.SetConfigFile(file)
	instance.SetConfigType("yml")
	if err := instance.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read file failed: %s \n", err))
	}
	instance.WatchConfig()
	instance.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := instance.Unmarshal(&global.App.Configuration); err != nil {
			fmt.Println(err)
		}
	})
	if err := instance.Unmarshal(&global.App.Configuration); err != nil {
		fmt.Println(err)
	}
	return instance
}
