package util

import (
	"github.com/spf13/viper"
)

var Configuration *Config

type Meta struct {
	Application string `mapstructure:"application"`
	Environment string `mapstructure:"environment"`
	Version     string `mapstructure:"version"`
}

type Otel struct {
	Address string `mapstructure:"address"`
}

type Log struct {
	Level string `mapstructure:"level"`
}

type HTTPServer struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Schema   string `mapstructure:"schema"`
}

type Jwt struct {
	Token string `mapstructure:"token"`
}

type Config struct {
	Meta       Meta       `mapstructure:"meta"`
	Log        Log        `mapstructure:"log"`
	HTTPServer HTTPServer `mapstructure:"http"`
	Database   Database   `mapstructure:"database"`
	Jwt        Jwt        `mapstructure:"jwt"`
	Otel       Otel       `mapstructure:"otel"`
}

func init() {
	logger := GetLogger()

	config, err := LoadConfig()
	if err != nil {
		logger.Errorf("%v", err.Error())
	}
	Configuration = &config
}

// Function to load the configuration from yaml using viper
func LoadConfig() (config Config, err error) {
	for _, location := range CONFIGDIRS {
		viper.AddConfigPath(location)
	}
	viper.SetConfigName(CONFIGNAME)
	viper.SetConfigType(CONFIGTYPE)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
