package config

type Configuration struct {
	App App
}

type App struct {
	Env     string
	Port    string
	AppName string `mapstructure:"app_name"`
	AppUrl  string `mapstructure:"app_url"`
}
