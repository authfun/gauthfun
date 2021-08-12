package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Service ServiceConfig `mapstructure:"service"`
	MySql   MysqlConfig   `mapstructure:"mysql"`
}

type ServiceConfig struct {
	Port int
}

type MysqlConfig struct {
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	DbName     string `mapstructure:"db_name"`
	Parameters string `mapstructure:"parameters"`
}

var AppConfig Config

func init() {

	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}
	parseConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		parseConfig()
	})
}

func parseConfig() {
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error parse config file: %s \n", err))
	}
}
