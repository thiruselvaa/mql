package models

import (
	"fmt"
	"strings"

	"github.com/creasty/defaults"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/dump"
)

type DSLFilterConfig struct {
	Filter Filter `mapstructure:"filter"`
}

type BooleanExpression struct {
	FieldPath string `mapstructure:"field_path,omitempty"`
	Operator  string `mapstructure:"operator,omitempty"`
	Value     any    `mapstructure:"value,omitempty"`
}

func (e BooleanExpression) String() string {
	return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, e.Value)
}

type NotLogicalCondition []Not
type Not struct {
	Expression []BooleanExpression `mapstructure:"expression,omitempty"`
	And        AndLogicalCondition `mapstructure:"and,omitempty"`
	Or         OrLogicalCondition  `mapstructure:"or,omitempty"`
}

type AndLogicalCondition []And
type And struct {
	Expression []BooleanExpression `mapstructure:"expression,omitempty"`
	Not        NotLogicalCondition `mapstructure:"not,omitempty"`
	Or         OrLogicalCondition  `mapstructure:"or,omitempty"`
}

type OrLogicalCondition []Or
type Or struct {
	Expression []BooleanExpression `mapstructure:"expression,omitempty"`
	Not        NotLogicalCondition `mapstructure:"not,omitempty"`
	And        AndLogicalCondition `mapstructure:"and,omitempty"`
}

type Comparision struct {
	Operator string `mapstructure:"operator"`
	Value    any    `mapstructure:"value"`
}
type Len struct {
	FieldPath   string      `mapstructure:"field_path"`
	Comparision Comparision `mapstructure:"comparision"`
}
type All struct {
	FieldPath string            `mapstructure:"field_path"`
	Condition []FilterCondition `mapstructure:"condition"`
}
type Any struct {
	FieldPath string            `mapstructure:"field_path"`
	Condition []FilterCondition `mapstructure:"condition"`
}
type One struct {
	FieldPath string            `mapstructure:"field_path"`
	Condition []FilterCondition `mapstructure:"condition"`
}
type None struct {
	FieldPath string            `mapstructure:"field_path"`
	Condition []FilterCondition `mapstructure:"condition"`
}

type Function struct {
	Len  Len  `mapstructure:"len,omitempty"`
	All  All  `mapstructure:"all,omitempty"`
	Any  Any  `mapstructure:"any,omitempty"`
	One  One  `mapstructure:"one,omitempty"`
	None None `mapstructure:"none,omitempty"`
}

// type LogicalCondition []interface{}
type FilterCondition struct {
	Expression []BooleanExpression `mapstructure:"expression,omitempty"`
	Not        NotLogicalCondition `mapstructure:"not,omitempty"`
	And        AndLogicalCondition `mapstructure:"and,omitempty"`
	Or         OrLogicalCondition  `mapstructure:"or,omitempty"`
	Function   Function            `mapstructure:"function,omitempty"`
}

// func (f FilterCondition) String() string {

// 	// var str StringBuilder
// 	var str string
// 	for idx, c := range f {
// 		str = fmt.Sprintf("Filter Condition(%v):(%+v)", idx, c)
// 	}
// 	return str
// }

type Filter struct {
	Name      string            `mapstructure:"name"`
	Type      string            `mapstructure:"type"`
	Condition []FilterCondition `mapstructure:"condition"`
}

func (f Filter) String() string {
	// var str StringBuilder
	var str []string
	for idx, c := range f.Condition {
		str = append(str, fmt.Sprintf("Filter Condition(%v):%+v", idx, c.Expression))
	}

	return strings.Join(str, "")
	// fmt.Sprintf("Filter Name:(%v) Type:(%v) and Condition:(%+v)", f.Name, f.Type, f.Condition)
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

	fmt.Println(cfg.Filter.String())

	return cfg, nil
}
