package config

type Configuration struct {
	App App
	Log Log
}

type App struct {
	Env     string
	Port    string
	AppName string `mapstructure:"app_name"`
	AppUrl  string `mapstructure:"app_url"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	RootDir    string `mapstructure:"root_dir"`
	FileName   string `mapstructure:"file_name"`
	Format     string `mapstructure:"format"`
	ShowLine   bool   `mapstructure:"show_line"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}
