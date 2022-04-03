package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	Port      string `yaml:"port"`
	SecretJWT string `yaml:"secret"`
	Database  struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Username string `yaml:"root"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appconfig *AppConfig

func initConfig() *AppConfig {
	var Configdefault AppConfig
	Configdefault.Port = os.Getenv("APP_PORT")
	Configdefault.SecretJWT = os.Getenv("SECRET_JWT")
	Configdefault.Database.Driver = os.Getenv("DB_DRIVER")
	Configdefault.Database.Name = os.Getenv("DB_NAME")
	Configdefault.Database.Address = os.Getenv("DB_ADDRESS")
	Configdefault.Database.Port = os.Getenv("DB_PORT")
	Configdefault.Database.Username = os.Getenv("DB_USERNAME")
	Configdefault.Database.Password = os.Getenv("DB_PASSWORD")

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./configs/")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Info("failed to open file")
	// 	return &Configdefault
	// }
	// var finalConfig AppConfig
	// err := viper.Unmarshal(&finalConfig)
	// if err != nil {
	// 	log.Info("failed to extract external config, use default value")
	// 	return &Configdefault
	// }
	return &Configdefault
}

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appconfig == nil {
		appconfig = initConfig()
	}
	return appconfig
}
