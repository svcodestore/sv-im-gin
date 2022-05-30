package config

import (
	"github.com/spf13/viper"
)

var (
	ConfigFile    = "config.yaml"
	ConfigEnvFile = "SAG_CONFIG"
)

type Configurator *viper.Viper

type Config struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []DB   `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Zap    Zap    `mapstructure:"logger" json:"zap" yaml:"logger"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
