package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	Application     applicationConfig
	Datasource      datasourceConfig
	Route           routeConfig
	InternalService internalServiceConfig
	ExternalService externalServiceConfig
	Module          moduleConfig
	Monitoring      monitoringConfig
)

func Init() {
	var cfg config
	if err := loadConfig(&cfg); err != nil {
		log.Err(err)
	}
}

func loadConfig(cfg *config) error {
	viper.SetConfigName("config") // Filename without extension
	viper.SetConfigType("yml")    // Explicitly set config type
	viper.AddConfigPath(".") // Look in the specified path
	viper.AutomaticEnv() // Override with ENV vars if available

	if err := viper.ReadInConfig(); err != nil {
		return err
	} else if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	syncConfig(cfg)

	return nil
}

func syncConfig(cfg *config) {
	Application = cfg.Application
	Datasource = cfg.Datasource
	Route = cfg.Route
	InternalService = cfg.InternalService
	ExternalService = cfg.ExternalService
	Module = cfg.Module
	Monitoring = cfg.Monitoring
}
