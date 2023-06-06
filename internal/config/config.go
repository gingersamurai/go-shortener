package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	StorageType     string           `mapstructure:"storage_type"`
	ShutdownTimeout time.Duration    `mapstructure:"shutdown_timeout"`
	Handler         HandlerConfig    `mapstructure:"handler"`
	HttpServer      HttpServerConfig `mapstructure:"http_server"`
	GrpcServer      GrpcServerConfig `mapstructure:"grpc_server"`
	Postgres        PostgresConfig   `mapstructure:"postgres"`
}

type HandlerConfig struct {
	HostAddr      string        `mapstructure:"host_addr"`
	HandleTimeout time.Duration `mapstructure:"handle_timeout"`
}

type HttpServerConfig struct {
	ListenAddr string `mapstructure:"listen_addr"`
}

type GrpcServerConfig struct {
	ListenAddr string `mapstructure:"listen_addr"`
}

type PostgresConfig struct {
	Host   string `mapstructure:"host"`
	User   string `mapstructure:"user"`
	DBName string `mapstructure:"dbname"`
}

func NewConfig(configFilePath string) (Config, error) {
	absPath, err := filepath.Abs(configFilePath)
	if err != nil {
		return Config{}, err
	}
	configFileDir, configFileName := filepath.Split(absPath)
	id := strings.LastIndex(configFileName, ".")
	configFileName = configFileName[:id]

	myViper := viper.New()
	myViper.AddConfigPath(configFileDir)
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
