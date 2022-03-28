package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port      int    `yaml:"port"`
	SecretJWT string `yaml:"secret"`
	Database  struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		Username string `yaml:"root"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appconfig *AppConfig

func initConfig() *AppConfig {
	var Configdefault AppConfig
	Configdefault.Port = 8000
	Configdefault.SecretJWT = "SECRET"
	Configdefault.Database.Driver = "mysql"
	Configdefault.Database.Name = "ProjectTest1"
	Configdefault.Database.Address = "127.0.0.1"
	Configdefault.Database.Port = 3306
	Configdefault.Database.Username = "root"
	Configdefault.Database.Password = ""

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Info("failed to open file")
		return &Configdefault
	}
	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract external config, use default value")
		return &Configdefault
	}
	return &finalConfig

}

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appconfig == nil {
		appconfig = initConfig()
	}
	return appconfig
}
