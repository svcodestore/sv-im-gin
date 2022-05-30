package config

type DB struct {
	Disable      bool   `mapstructure:"disable" json:"disable" yaml:"disable"`
	Type         string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName    string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	Mysql
}
