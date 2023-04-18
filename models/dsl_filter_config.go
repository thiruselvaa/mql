package models

import (
	"fmt"

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
type BooleanExpressions []BooleanExpression

func (e BooleanExpression) String() string {
	return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, e.Value)
}

type Not struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}
type NotLogicalConditions []Not

type And struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}
type AndLogicalConditions []And

type Or struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}
type OrLogicalConditions []Or

type Comparision struct {
	Operator string `mapstructure:"operator"`
	Value    any    `mapstructure:"value"`
}
type ArrayLengthFunction struct {
	FieldPath   string      `mapstructure:"field_path"`
	Comparision Comparision `mapstructure:"comparision"`
}

type ArrayElementMatchFunction struct {
	FieldPath string          `mapstructure:"field_path"`
	Condition FilterCondition `mapstructure:"condition"`
}

type AllArrayElementMatchFunction ArrayElementMatchFunction
type AnyArrayElementMatchFunction ArrayElementMatchFunction
type OneArrayElementMatchFunction ArrayElementMatchFunction
type NoneArrayElementMatchFunction ArrayElementMatchFunction

type BuiltInFunction struct {
	Len  *ArrayLengthFunction           `mapstructure:"len,omitempty"`
	All  *AllArrayElementMatchFunction  `mapstructure:"all,omitempty"`
	Any  *AnyArrayElementMatchFunction  `mapstructure:"any,omitempty"`
	One  *OneArrayElementMatchFunction  `mapstructure:"one,omitempty"`
	None *NoneArrayElementMatchFunction `mapstructure:"none,omitempty"`
}
type BuiltInFunctions []BuiltInFunction

// type LogicalCondition struct {
// 	Not *NotLogicalConditions `mapstructure:"not,omitempty"`
// 	And *AndLogicalConditions `mapstructure:"and,omitempty"`
// 	Or  *OrLogicalConditions  `mapstructure:"or,omitempty"`
// }

type FilterCondition struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
	// LogicalCondition
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
	Name      string          `mapstructure:"name"`
	Type      string          `mapstructure:"type"`
	Condition FilterCondition `mapstructure:"condition"`
}

func (f Filter) String() string {
	return fmt.Sprintf("Filter Name:(%v) Type:(%v) and Condition:(%+v)", f.Name, f.Type, f.Condition)
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
