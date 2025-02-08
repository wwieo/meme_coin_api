package config

import (
	"go.uber.org/dig"
)

type MemeCoin struct {
	dig.Out

	ServiceAddress ServiceAddress           `mapstructure:"service_address"`
	DBMS           DatabaseManagementSystem `mapstructure:"database_management_system"`
}

type ServiceAddress struct {
	MemeCoin string `mapstructure:"meme_coin"`
}

type DatabaseManagementSystem struct {
	MariaDBSystems map[string]MySQL `mapstructure:"mysql"`
	RedisSystems   map[string]Redis `mapstructure:"redis"`
}

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Account  string `mapstructure:"account"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
