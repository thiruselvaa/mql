package models

import (
	"github.com/creasty/defaults"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/dump"
)

type SMFConfig struct {
	Query QueryConfig `mapstructure:"query" validate:"required"`
}

type QueryConfig struct {
	Name  string      `mapstructure:"name" validate:"required"`
	Type  string      `mapstructure:"type" default:"native"`
	Where interface{} `mapstructure:"where" validate:"required"`
}

func NewSMFConfig(file string) (cfg *SMFConfig, err error) {
	c := config.New("test").WithOptions(config.ParseEnv).WithDriver(yamlv3.Driver)
	err = c.LoadExistsByFormat(config.Yaml, file)
	if err != nil {
		return nil, err
	}

	cfg = &SMFConfig{}
	err = c.Decode(cfg)
	if err != nil {
		return nil, err
	}
	dump.V(cfg)

	defaults.MustSet(cfg)

	dump.V(cfg)
	return cfg, nil
}
