package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Monitor  MonitorConfig  `mapstructure:"monitor"`
	Log      LogConfig     `mapstructure:"log"`
	Security SecurityConfig `mapstructure:"security"`
}

type ServerConfig struct {
	Port     int      `mapstructure:"port"`
	Host     string   `mapstructure:"host"`
	Debug    bool     `mapstructure:"debug"`
	AllowIPs []string `mapstructure:"allow_ips"`
}

type DatabaseConfig struct {
	Driver       string `mapstructure:"driver"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/chaos-platform/")
	
	viper.SetEnvPrefix("CHAOS")
	viper.AutomaticEnv()
	
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	
	return &config, validateConfig(&config)
}

func validateConfig(cfg *Config) error {
	// 添加配置验证逻辑
	return nil
} 