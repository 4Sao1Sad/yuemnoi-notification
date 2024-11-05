package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName     string `yaml:"appname"`
	Host        string `mapstructure:"host"`
	Port        int32  `yaml:"port"`
	Db          DB     `mapstructure:"DB"`
	RabbitMQUrl string `yaml:"rabbitmqurl"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func Load() *Config {
	config := Config{}
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error occurs while reading the config. ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("error occurs while unmarshalling the config. ", err)
	}
	return &config
}
