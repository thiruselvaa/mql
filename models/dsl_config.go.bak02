package models

import (
	"github.com/creasty/defaults"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/dump"
)

// type DSLConfig struct {
// 	Query Query `mapstructure:"query" validate:"required"`
// }
// type Expression struct {
// 	FieldPath string      `mapstructure:"field_path" validate:"required"`
// 	Operator  string      `mapstructure:"operator" validate:"required"`
// 	Value     interface{} `mapstructure:"value" validate:"required"`
// }

// type LogicalCondition struct {
// 	Expression []Expression `mapstructure:"expression,omitempty"`
// 	Not        interface{}  `mapstructure:"not,omitempty"`
// 	And        interface{}  `mapstructure:"and,omitempty"`
// 	Or         interface{}  `mapstructure:"or,omitempty"`
// }

// type Comparision struct {
// 	Operator string      `mapstructure:"operator" validate:"required"`
// 	Value    interface{} `mapstructure:"value" validate:"required"`
// }
// type Len struct {
// 	FieldPath   string      `mapstructure:"field_path" validate:"required"`
// 	Comparision Comparision `mapstructure:"comparision" validate:"required"`
// }

// type FunctionParameter struct {
// 	FieldPath string      `mapstructure:"field_path" validate:"required"`
// 	Predicate interface{} `mapstructure:"where" validate:"required"`
// }

// type Function struct {
// 	Len  Len               `mapstructure:"len,omitempty"`
// 	All  FunctionParameter `mapstructure:"all,omitempty"`
// 	Any  FunctionParameter `mapstructure:"any,omitempty"`
// 	One  FunctionParameter `mapstructure:"one,omitempty"`
// 	None FunctionParameter `mapstructure:"none,omitempty"`
// }
// type Condition struct {
// 	Expression []Expression     `mapstructure:"expression,omitempty"`
// 	Not        LogicalCondition `mapstructure:"not,omitempty"`
// 	And        LogicalCondition `mapstructure:"and,omitempty"`
// 	Or         LogicalCondition `mapstructure:"or,omitempty"`
// 	Function   Function         `mapstructure:"function,omitempty"`
// }
// type Query struct {
// 	Name  string    `mapstructure:"name"`
// 	Type  string    `mapstructure:"type"`
// 	Where Condition `mapstructure:"where"`
// }

type DSLConfig struct {
	Query Query `mapstructure:"query"`
}

type Expression struct {
	FieldPath string `mapstructure:"field_path"`
	Operator  string `mapstructure:"operator"`
	Value     int    `mapstructure:"value"`
}

type Not struct {
	Expression Expression `mapstructure:"expression"`
	And        []And      `mapstructure:"and"`
	Or         []Or       `mapstructure:"or"`
}

type And struct {
	Expression Expression `mapstructure:"expression,omitempty"`
	Not        Not        `mapstructure:"not,omitempty"`
	Or         []Or       `mapstructure:"or,omitempty"`
}
type Or struct {
	Expression Expression `mapstructure:"expression,omitempty"`
	Not        []Not      `mapstructure:"not,omitempty"`
	And        []And      `mapstructure:"and,omitempty"`
}
type LogicalCondition []interface{}
type Condition struct {
	Not LogicalCondition `mapstructure:"not,omitempty"`
	And LogicalCondition `mapstructure:"and,omitempty"`
	Or  LogicalCondition `mapstructure:"or,omitempty"`
}

type Where struct {
	Condition []Condition `mapstructure:"condition"`
}
type Query struct {
	Name  string `mapstructure:"name"`
	Type  string `mapstructure:"type"`
	Where Where  `mapstructure:"where"`
}

func NewDSLConfig(file string) (cfg *DSLConfig, err error) {
	c := config.New("dsl-config").WithOptions(config.ParseEnv).WithDriver(yamlv3.Driver)
	err = c.LoadExistsByFormat(config.Yaml, file)
	if err != nil {
		return nil, err
	}

	cfg = &DSLConfig{}
	err = c.Decode(cfg)
	if err != nil {
		return nil, err
	}
	dump.V(cfg)

	defaults.MustSet(cfg)

	dump.V(cfg)
	return cfg, nil
}
