package share

import (
	config2 "config"
)

type A interface {
	GetConfig(appConfig *config2.AppConfig)
}
