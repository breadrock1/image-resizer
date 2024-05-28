package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Logger  LoggerConfig
	Server  ServerConfig
	Cacher  CacheConfig
	Storage StorageConfig
	Resizer ResizerConfig
}

type LoggerConfig struct {
	Level         string
	FilePath      string
	EnableFileLog bool
}

type ServerConfig struct {
	Host     string
	HostPort int
}

type CacheConfig struct {
	ExpireSeconds  int
	CapacityValues int
}

type StorageConfig struct {
	UseFilesystem   bool
	UploadDirectory string
}

type ResizerConfig struct {
	TargetQuality int
}

func NewConfig(filePath string) (*Config, error) {
	config := &Config{}

	viperInstance := viper.New()
	viperInstance.AutomaticEnv()
	viperInstance.SetConfigFile(filePath)

	viperInstance.SetDefault("cacher.ExpireSeconds", 20)
	viperInstance.SetDefault("cacher.CapacityValues", 5)

	viperInstance.SetDefault("logger.Level", "INFO")
	viperInstance.SetDefault("logger.FilePath", "logs/app.log")
	viperInstance.SetDefault("logger.EnableFileLog", false)

	viperInstance.SetDefault("resizer.TargetQuality", 90)

	viperInstance.SetDefault("server.Host", "0.0.0.0")
	viperInstance.SetDefault("server.HostPort", 2891)

	viperInstance.SetDefault("storage.UseFilesystem", true)
	viperInstance.SetDefault("storage.UploadDirectory", "uploads")

	if err := viperInstance.ReadInConfig(); err != nil {
		confErr := fmt.Errorf("failed while reading config file %s: %w", filePath, err)
		return config, confErr
	}

	if err := viperInstance.Unmarshal(config); err != nil {
		confErr := fmt.Errorf("failed while unmarshaling config file %s: %w", filePath, err)
		return config, confErr
	}

	return config, nil
}
