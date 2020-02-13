package share

import config "../config"

type A interface {
	GetConfig(appConfig *config.AppConfig)
}

