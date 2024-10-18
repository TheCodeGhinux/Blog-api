package config

import "go-blog/config"

func Get() config.AppConfig {
	return configurations
}
