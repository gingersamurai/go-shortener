package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	ConfigFilePath = "./"
	ConfigFileName = "config"
)

type Config struct {
	StorageType     string           `mapstructure:"storage_type"`
	ShutdownTimeout time.Duration    `mapstructure:"shutdown_timeout"`
	Handler         HandlerConfig    `mapstructure:"handler"`
	HttpServer      HttpServerConfig `mapstructure:"http_server"`
	Postgres        PostgresConfig   `mapstructure:"postgres"`
}

type HandlerConfig struct {
	HostAddr      string        `mapstructure:"host_addr"`
	HandleTimeout time.Duration `mapstructure:"handle_timeout"`
}

type HttpServerConfig struct {
	ListenAddr string `mapstructure:"listen_addr"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func NewConfig(configFilePath, configFileName string) (Config, error) {
	myViper := viper.New()
	myViper.AddConfigPath(configFilePath)
	myViper.SetConfigName(configFileName)

	if err := myViper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var result Config
	if err := myViper.Unmarshal(&result); err != nil {
		return Config{}, err
	}

	return result, nil
}
