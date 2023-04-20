package models

import (
	"fmt"
	"strings"

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
	for idx, v := range e {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(e) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " and %v", v)
			}
		}
	}

	return fmt.Sprintf("(%v)", sb.String())
}

type OrBooleanExpressions []BooleanExpression

func (e OrBooleanExpressions) String() string {
	var sb strings.Builder
	for idx, v := range e {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(e) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " or %v", v)
			}
		}
	}

	return fmt.Sprintf("(%v)", sb.String())
}

type OrLogicalCondition FilterCondition

// type OrLogicalCondition struct {
// 	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
// 	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
// 	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
// 	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
// }

func (c OrLogicalCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(c.Expression) && len(*c.Expression) > 0 {
		orBooleanExpressions := make(OrBooleanExpressions, len(*c.Expression))
		copy(orBooleanExpressions, OrBooleanExpressions(*c.Expression))
		tmp = append(tmp, orBooleanExpressions)
	}
	if !funk.IsEmpty(c.Or) && len(*c.Or) > 0 {
		tmp = append(tmp, c.Or)
	}
	if !funk.IsEmpty(c.And) && len(*c.And) > 0 {
		tmp = append(tmp, c.And)
	}
	if !funk.IsEmpty(c.Not) && len(*c.Not) > 0 {
		tmp = append(tmp, c.Not)
	}
	if !funk.IsEmpty(c.Function) && len(*c.Function) > 0 {
		tmp = append(tmp, c.Function)
	}

	// dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Fprintf(&sb, " or %v", v)
		}
	}

	return sb.String()
}

type OrLogicalConditions []OrLogicalCondition

func (c OrLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(c) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " or %v", v)
			}
		}
	}
	return fmt.Sprintf("(%v)", sb.String())
}

type AndLogicalCondition FilterCondition

// type AndLogicalCondition struct {
// 	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
// 	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
// 	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
// 	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
// }

func (c AndLogicalCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(c.Expression) && len(*c.Expression) > 0 {
		tmp = append(tmp, c.Expression)
	}
	if !funk.IsEmpty(c.Or) && len(*c.Or) > 0 {
		tmp = append(tmp, c.Or)
	}
	if !funk.IsEmpty(c.And) && len(*c.And) > 0 {
		tmp = append(tmp, c.And)
	}
	if !funk.IsEmpty(c.Not) && len(*c.Not) > 0 {
		tmp = append(tmp, c.Not)
	}
	if !funk.IsEmpty(c.Function) && len(*c.Function) > 0 {
		tmp = append(tmp, c.Function)
	}

	// dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Fprintf(&sb, " and %v", v)
		}
	}

	return sb.String()
}

type AndLogicalConditions []AndLogicalCondition

func (c AndLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(c) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " and %v", v)
			}
		}
	}
	return fmt.Sprintf("(%v)", sb.String())
}

type NotLogicalCondition FilterCondition

func (c NotLogicalCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(c.Expression) && len(*c.Expression) > 0 {
		tmp = append(tmp, c.Expression)
	}
	if !funk.IsEmpty(c.Or) && len(*c.Or) > 0 {
		tmp = append(tmp, c.Or)
	}
	if !funk.IsEmpty(c.And) && len(*c.And) > 0 {
		tmp = append(tmp, c.And)
	}
	if !funk.IsEmpty(c.Not) && len(*c.Not) > 0 {
		tmp = append(tmp, c.Not)
	}
	if !funk.IsEmpty(c.Function) && len(*c.Function) > 0 {
		tmp = append(tmp, c.Function)
	}

	// dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Fprintf(&sb, " and %v", v)
		}
	}

	return sb.String()
}

type NotLogicalConditions []NotLogicalCondition

func (c NotLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(c) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " and %v", v)
			}
		}
	}
	return fmt.Sprintf("(not (%v)", sb.String())
}

type Comparision struct {
	Operator string `mapstructure:"operator"`
	Value    any    `mapstructure:"value"`
}

func (c Comparision) String() string {
	return fmt.Sprintf("%v %v", c.Operator, c.Value)
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
	return fmt.Sprintf("(%v %v)", f.FieldPath, f.Condition)
}

type AllArrayElementMatchFunction ArrayElementMatchFunction

func (f AllArrayElementMatchFunction) String() string {
	return fmt.Sprintf("(all(%v, %v))", f.FieldPath, f.Condition)
}

type AnyArrayElementMatchFunction ArrayElementMatchFunction

func (f AnyArrayElementMatchFunction) String() string {
	return fmt.Sprintf("(any(%v, %v))", f.FieldPath, f.Condition)
}

type OneArrayElementMatchFunction ArrayElementMatchFunction

func (f OneArrayElementMatchFunction) String() string {
	return fmt.Sprintf("(one(%v, %v))", f.FieldPath, f.Condition)
}

type NoneArrayElementMatchFunction ArrayElementMatchFunction

func (f NoneArrayElementMatchFunction) String() string {
	return fmt.Sprintf("(none(%v, %v))", f.FieldPath, f.Condition)
}

type BuiltInFunction struct {
	Len  *ArrayLengthFunction           `mapstructure:"len,omitempty"`
	All  *AllArrayElementMatchFunction  `mapstructure:"all,omitempty"`
	Any  *AnyArrayElementMatchFunction  `mapstructure:"any,omitempty"`
	One  *OneArrayElementMatchFunction  `mapstructure:"one,omitempty"`
	None *NoneArrayElementMatchFunction `mapstructure:"none,omitempty"`
}

func (f BuiltInFunction) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)
	if !funk.IsEmpty(f.Len) {
		fmt.Println("add LEN to tmp array")
		tmp = append(tmp, f.Len)
	}
	if !funk.IsEmpty(f.All) {
		fmt.Println("add ALL to tmp array")
		tmp = append(tmp, f.All)
	}
	if !funk.IsEmpty(f.Any) {
		fmt.Println("add ANY to tmp array")
		tmp = append(tmp, f.Any)
	}
	if !funk.IsEmpty(f.One) {
		fmt.Println("add ONE to tmp array")
		tmp = append(tmp, f.One)
	}
	if !funk.IsEmpty(f.None) {
		fmt.Println("add NONE to tmp array")
		tmp = append(tmp, f.None)
	}

	// dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Fprintf(&sb, " and %v", v)
		}
	}

	return sb.String()
}

type BuiltInFunctions []BuiltInFunction

func (f BuiltInFunctions) String() string {
	var sb strings.Builder
	for idx, v := range f {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(f) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " and %v", v)
			}
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
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
	// LogicalCondition
}

func (f FilterCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(f.Expression) && len(*f.Expression) > 0 {
		tmp = append(tmp, f.Expression)
	}
	if !funk.IsEmpty(f.Or) && len(*f.Or) > 0 {
		tmp = append(tmp, f.Or)
	}
	if !funk.IsEmpty(f.And) && len(*f.And) > 0 {
		tmp = append(tmp, f.And)
	}
	if !funk.IsEmpty(f.Not) && len(*f.Not) > 0 {
		tmp = append(tmp, f.Not)
	}
	if !funk.IsEmpty(f.Function) && len(*f.Function) > 0 {
		tmp = append(tmp, f.Function)
	}

	// dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			// fmt.Fprintf(&sb, "\n\t(\n\t\t%v\n\t)\n\t", v)
			fmt.Fprintf(&sb, "(%v)", v)
		} else {
			fmt.Fprintf(&sb, "and(%v)", v)
		}
	}

	return fmt.Sprintf("(%v)", sb.String())
}

type Filter struct {
	Name      string          `mapstructure:"name"`
	Type      string          `mapstructure:"type" default:"dsl"`
	Condition FilterCondition `mapstructure:"condition"`
}

func (f Filter) String() string {
	return fmt.Sprintf("Filter Name:(%v) Type:(%v) and Condition:(%+v)", f.Name, f.Type, f.Condition)
}

func NewDSLFilterConfig(file string) (cfg *DSLFilterConfig, err error) {
	c := config.New("dsl-config").
		WithOptions(config.ParseEnv).
		WithOptions(config.ParseDefault).
		WithDriver(yamlv3.Driver)
	err = c.LoadExistsByFormat(config.Yaml, file)
	if err != nil {
		return nil, err
	}

	cfg = &DSLFilterConfig{}
	err = c.Decode(cfg)
	if err != nil {
		return nil, err
	}
	// dump.V(cfg)
	// defaults.MustSet(cfg)
	dump.V(cfg)

	fmt.Printf("\nfilter condition: %+v\n", cfg.Filter.Condition)

	return cfg, nil
}
