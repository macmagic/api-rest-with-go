package service

import "config"

var BaseConfig *config.AppConfig

func GetConfig(c *config.AppConfig) {
	BaseConfig = c
}
