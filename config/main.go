package config

import "github.com/jinzhu/configor"

type Cfg struct {
	Mysql struct {
		Dsn string
	}

	Redis struct {
		Url string
	}
}

func Get() *Cfg {
	var Config *Cfg
	configor.Load(&Config, "config.yaml")
	return Config
}
