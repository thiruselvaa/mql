package models

import (
	"fmt"
	"strings"

	"github.com/creasty/defaults"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/dump"
	"github.com/thoas/go-funk"
)

type DSLFilterConfig struct {
	Filter Filter `mapstructure:"filter"`
}

type BooleanExpression struct {
	FieldPath string `mapstructure:"field_path"`
	Operator  string `mapstructure:"operator"`
	Value     any    `mapstructure:"value"`
}

func (e BooleanExpression) String() string {
	return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, e.Value)
}

type BooleanExpressions []BooleanExpression

func (e BooleanExpressions) String() string {
	var sb strings.Builder
	for _, v := range e {
		if !funk.IsEmpty(v) {
			fmt.Fprintf(&sb, "%v and ", v)
		}
	}
	return sb.String()
}

type Not struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}

func (n Not) String() string {
	var sb strings.Builder

	if !funk.IsEmpty(n.Expression) {
		fmt.Fprintf(&sb, "%v and ", n.Expression)
	}
	if !funk.IsEmpty(n.And) {
		fmt.Fprintf(&sb, "%v and ", n.And)
	}
	if !funk.IsEmpty(n.Or) {
		fmt.Fprintf(&sb, "%v and ", n.Or)
	}
	if !funk.IsEmpty(n.Function) {
		fmt.Fprintf(&sb, "%v and ", n.Function)
	}

	return fmt.Sprintf("not (%v)", sb.String())
}

type NotLogicalConditions []Not

func (c NotLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == len(c) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, "%v and ", v)
			}
		}
	}
	return sb.String()
}

type And struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}

func (a And) String() string {
	var sb strings.Builder

	if !funk.IsEmpty(a.Expression) {
		fmt.Fprintf(&sb, "%v and ", a.Expression)
	}
	if !funk.IsEmpty(a.Not) {
		fmt.Fprintf(&sb, "%v and ", a.Not)
	}
	if !funk.IsEmpty(a.Or) {
		fmt.Fprintf(&sb, "%v and ", a.Or)
	}
	if !funk.IsEmpty(a.Function) {
		fmt.Fprintf(&sb, "%v and ", a.Function)
	}

	return fmt.Sprintf("and (%v)", sb.String())
}

type AndLogicalConditions []And

func (c AndLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == len(c) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, "%v and ", v)
			}
		}
	}
	return sb.String()
}

type Or struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}

func (o Or) String() string {
	var sb strings.Builder

	if !funk.IsEmpty(o.Expression) {
		fmt.Fprintf(&sb, "%v or ", o.Expression)
	}
	if !funk.IsEmpty(o.Not) {
		fmt.Fprintf(&sb, "%v or ", o.Not)
	}
	if !funk.IsEmpty(o.And) {
		fmt.Fprintf(&sb, "%v or ", o.And)
	}
	if !funk.IsEmpty(o.Function) {
		fmt.Fprintf(&sb, "%v or ", o.Function)
	}

	return fmt.Sprintf("or (%v)", sb.String())
}

type OrLogicalConditions []Or

func (c OrLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == len(c) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, "%v or ", v)
			}
		}
	}
	return sb.String()
}

type Comparision struct {
	Operator string `mapstructure:"operator"`
	Value    any    `mapstructure:"value"`
}

func (c Comparision) String() string {
	return fmt.Sprintf("%v %v)", c.Operator, c.Value)
}

type ArrayLengthFunction struct {
	FieldPath   string      `mapstructure:"field_path"`
	Comparision Comparision `mapstructure:"comparision"`
}

func (f ArrayLengthFunction) String() string {
	return fmt.Sprintf("(len(%v) %v)", f.FieldPath, f.Comparision)
}

type ArrayElementMatchFunction struct {
	FieldPath string          `mapstructure:"field_path"`
	Condition FilterCondition `mapstructure:"condition"`
}

func (f ArrayElementMatchFunction) String() string {
	return fmt.Sprintf("%v, %v)", f.FieldPath, f.Condition)
}

type AllArrayElementMatchFunction ArrayElementMatchFunction

func (f AllArrayElementMatchFunction) String() string {
	return fmt.Sprintf("all(%v, %v))", f.FieldPath, f.Condition)
}

type AnyArrayElementMatchFunction ArrayElementMatchFunction

func (f AnyArrayElementMatchFunction) String() string {
	return fmt.Sprintf("any(%v, %v))", f.FieldPath, f.Condition)
}

type OneArrayElementMatchFunction ArrayElementMatchFunction

func (f OneArrayElementMatchFunction) String() string {
	return fmt.Sprintf("one(%v, %v))", f.FieldPath, f.Condition)
}

type NoneArrayElementMatchFunction ArrayElementMatchFunction

func (f NoneArrayElementMatchFunction) String() string {
	return fmt.Sprintf("none(%v, %v))", f.FieldPath, f.Condition)
}

type BuiltInFunction struct {
	Len  *ArrayLengthFunction           `mapstructure:"len,omitempty"`
	All  *AllArrayElementMatchFunction  `mapstructure:"all,omitempty"`
	Any  *AnyArrayElementMatchFunction  `mapstructure:"any,omitempty"`
	One  *OneArrayElementMatchFunction  `mapstructure:"one,omitempty"`
	None *NoneArrayElementMatchFunction `mapstructure:"none,omitempty"`
}

func (f BuiltInFunction) String() string {
	var sb strings.Builder

	if !funk.IsEmpty(f.Len) {
		fmt.Fprintf(&sb, "%v and ", f.Len)
	}
	if !funk.IsEmpty(f.All) {
		fmt.Fprintf(&sb, "%v and ", f.All)
	}
	if !funk.IsEmpty(f.Any) {
		fmt.Fprintf(&sb, "%v and ", f.Any)
	}
	if !funk.IsEmpty(f.One) {
		fmt.Fprintf(&sb, "%v and ", f.One)
	}
	if !funk.IsEmpty(f.None) {
		fmt.Fprintf(&sb, "%v and ", f.None)
	}

	return sb.String()
}

type BuiltInFunctions []BuiltInFunction

func (f BuiltInFunctions) String() string {
	var sb strings.Builder
	for _, v := range f {
		if !funk.IsEmpty(v) {
			fmt.Fprintf(&sb, "%v and ", v)
		}
	}
	return sb.String()
}

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

func (f FilterCondition) String() string {
	var sb strings.Builder

	if !funk.IsEmpty(f.Expression) {
		fmt.Fprintf(&sb, "%v and ", f.Expression)
	}
	if !funk.IsEmpty(f.Not) {
		fmt.Fprintf(&sb, "%v and ", f.Not)
	}
	if !funk.IsEmpty(f.And) {
		fmt.Fprintf(&sb, "%v and ", f.And)
	}
	if !funk.IsEmpty(f.Or) {
		fmt.Fprintf(&sb, "%v and ", f.Or)
	}
	if !funk.IsEmpty(f.Function) {
		fmt.Fprintf(&sb, "%v and ", f.Function)
	}

	return sb.String()
}

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
	if !funk.IsEmpty(cfg.Filter.Condition.Expression) {
		fmt.Printf("filter condition expressions: %+v", *cfg.Filter.Condition.Expression)
	}
	fmt.Printf("filter condition functions: %+v\n", cfg.Filter.Condition.Function)
	fmt.Printf("filter condition functions - len: %+v\n", (*cfg.Filter.Condition.Function)[0].Len)
	fmt.Printf("filter condition functions - any: %+v\n", (*cfg.Filter.Condition.Function)[1].Any)

	return cfg, nil
}
