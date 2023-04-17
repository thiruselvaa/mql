package models

import (
	"github.com/creasty/defaults"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/dump"
)

type DSLFilterConfig struct {
	Filter Filter `json:"filter"`
}

type BooleanExpression struct {
	FieldPath string `mapstructure:"field_path,omitempty"`
	Operator  string `mapstructure:"operator,omitempty"`
	Value     int    `mapstructure:"value,omitempty"`
}

type NotLogicalCondition []Not
type Not struct {
	Expression BooleanExpression   `mapstructure:"expression,omitempty"`
	And        AndLogicalCondition `mapstructure:"and,omitempty"`
	Or         OrLogicalCondition  `mapstructure:"or,omitempty"`
}

type AndLogicalCondition []And
type And struct {
	Expression BooleanExpression   `mapstructure:"expression,omitempty"`
	Not        NotLogicalCondition `mapstructure:"not,omitempty"`
	Or         OrLogicalCondition  `mapstructure:"or,omitempty"`
}

type OrLogicalCondition []Or
type Or struct {
	Expression BooleanExpression   `mapstructure:"expression,omitempty"`
	Not        NotLogicalCondition `mapstructure:"not,omitempty"`
	And        AndLogicalCondition `mapstructure:"and,omitempty"`
}

type LogicalCondition []interface{}
type FilterCondition struct {
	Not NotLogicalCondition `mapstructure:"not,omitempty"`
	And AndLogicalCondition `mapstructure:"and,omitempty"`
	Or  OrLogicalCondition  `mapstructure:"or,omitempty"`
}

type Filter struct {
	Name      string            `mapstructure:"name"`
	Type      string            `mapstructure:"type"`
	Condition []FilterCondition `mapstructure:"condition"`
}

func NewDSLFilterConfig(file string) (cfg *DSLFilterConfig, err error) {
	c := config.New("dsl-config").WithOptions(config.ParseEnv).WithDriver(yamlv3.Driver)
	err = c.LoadExistsByFormat(config.Yaml, file)
	if err != nil {
		return nil, err
	}

	cfg = &DSLFilterConfig{}
	err = c.Decode(cfg)
	if err != nil {
		return nil, err
	}
	dump.V(cfg)

	defaults.MustSet(cfg)

	dump.V(cfg)
	return cfg, nil
}
