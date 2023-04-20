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
	for idx, v := range e {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(e) {
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Fprintf(&sb, " and %v", v)
			}
		}
	}
	fmt.Printf("BooleanExpressions.String: %v\n", sb.String())

	return fmt.Sprintf("(%v)", sb.String())
}

type NotLogicalCondition struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}

func (n NotLogicalCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(n.Expression) && len(*n.Expression) > 0 {
		fmt.Println("NotLogicalCondition: add EXPRESSION to tmp array")
		tmp = append(tmp, n.Expression)
	}
	if !funk.IsEmpty(n.And) && len(*n.And) > 0 {
		fmt.Println("NotLogicalCondition: add AND to tmp array")
		tmp = append(tmp, n.And)
	}
	if !funk.IsEmpty(n.Or) && len(*n.Or) > 0 {
		fmt.Println("NotLogicalCondition: add OR to tmp array")
		tmp = append(tmp, n.Or)
	}
	if !funk.IsEmpty(n.Function) && len(*n.Function) > 0 {
		fmt.Println("NotLogicalCondition: add FUNCTION to tmp array")
		tmp = append(tmp, n.Function)
	}

	dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Printf("NotLogicalCondition: if block: %v\n", v)
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Printf("NotLogicalCondition: else block: %v\n", v)
			fmt.Fprintf(&sb, " and %v", v)
		}
	}

	fmt.Printf("NotLogicalCondition: sb.String is %v\n", sb.String())

	// return fmt.Sprintf(" not (%v)", sb.String())
	return sb.String()
}

type NotLogicalConditions []NotLogicalCondition

func (c NotLogicalConditions) String() string {
	var sb strings.Builder
	for idx, v := range c {
		if !funk.IsEmpty(v) {
			if idx == 0 || idx == len(c) {
				fmt.Printf("NotLogicalConditions: if block: %v\n", v)
				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Printf("NotLogicalConditions: else block: %v\n", v)
				fmt.Fprintf(&sb, " and %v", v)
			}
		}
	}
	// return fmt.Sprintf("(%v)", sb.String())
	return fmt.Sprintf(" not (%v)", sb.String())
	// return sb.String()
}

type AndLogicalCondition struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	Or         *OrLogicalConditions  `mapstructure:"or,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}

func (a AndLogicalCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(a.Expression) && len(*a.Expression) > 0 {
		fmt.Println("add EXPRESSION to tmp array")
		tmp = append(tmp, a.Expression)
	}
	if !funk.IsEmpty(a.Not) && len(*a.Not) > 0 {
		fmt.Println("add NOT to tmp array")
		tmp = append(tmp, a.Not)
	}
	if !funk.IsEmpty(a.Or) && len(*a.Or) > 0 {
		fmt.Println("add OR to tmp array")
		tmp = append(tmp, a.Or)
	}
	if !funk.IsEmpty(a.Function) && len(*a.Function) > 0 {
		fmt.Println("add FUNCTION to tmp array")
		tmp = append(tmp, a.Function)
	}

	dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Fprintf(&sb, " and %v", v)
		}
	}

	// return fmt.Sprintf(" and (%v)", sb.String())
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

type OrLogicalCondition struct {
	Expression *BooleanExpressions   `mapstructure:"expression,omitempty"`
	Not        *NotLogicalConditions `mapstructure:"not,omitempty"`
	And        *AndLogicalConditions `mapstructure:"and,omitempty"`
	Function   *BuiltInFunctions     `mapstructure:"function,omitempty"`
}

func (o OrLogicalCondition) String() string {
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(o.Expression) && len(*o.Expression) > 0 {
		// var orBexprs BooleanExpressions
		fmt.Println("OrLogicalCondition: add EXPRESSION to tmp array")
		for idx, e := range *o.Expression {
			if idx == 0 || idx == len(*o.Expression) {
				fmt.Fprintf(&sb, "%v", e)
			} else {
				fmt.Fprintf(&sb, " or %v", e)
			}
		}
		// tmp = append(tmp, o.Expression)
	}
	if !funk.IsEmpty(o.Not) && len(*o.Not) > 0 {
		fmt.Println("OrLogicalCondition: add NOT to tmp array")
		tmp = append(tmp, o.Not)
	}
	if !funk.IsEmpty(o.And) && len(*o.And) > 0 {
		fmt.Println("OrLogicalCondition: add AND to tmp array")
		tmp = append(tmp, o.And)
	}
	if !funk.IsEmpty(o.Function) && len(*o.Function) > 0 {
		fmt.Println("OrLogicalCondition: add FUNCTION to tmp array")
		tmp = append(tmp, o.Function)
	}

	dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "%v", v)
		} else {
			fmt.Fprintf(&sb, " and %v", v)
		}
	}

	// return fmt.Sprintf(" or (%v)", sb.String())
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
	return fmt.Sprintf("all(%v, %v)", f.FieldPath, f.Condition)
}

type AnyArrayElementMatchFunction ArrayElementMatchFunction

func (f AnyArrayElementMatchFunction) String() string {
	return fmt.Sprintf("any(%v, %v)", f.FieldPath, f.Condition)
}

type OneArrayElementMatchFunction ArrayElementMatchFunction

func (f OneArrayElementMatchFunction) String() string {
	return fmt.Sprintf("one(%v, %v)", f.FieldPath, f.Condition)
}

type NoneArrayElementMatchFunction ArrayElementMatchFunction

func (f NoneArrayElementMatchFunction) String() string {
	return fmt.Sprintf("none(%v, %v)", f.FieldPath, f.Condition)
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

	dump.V(tmp)

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
			fmt.Printf("IsEmpty: length of built-in functions: %v\n", len(f))

			if idx == 0 || idx == len(f) {
				fmt.Printf("idx: length of built-in functions: %v\n", len(f))

				fmt.Fprintf(&sb, "%v", v)
			} else {
				fmt.Printf("else: length of built-in functions: %v\n", len(f))

				fmt.Fprintf(&sb, " and %v", v)
			}
		}
	}
	fmt.Printf("BuiltInFunctions.String: %v\n", sb.String())
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
	var (
		sb  strings.Builder
		tmp []interface{}
	)

	if !funk.IsEmpty(f.Expression) && len(*f.Expression) > 0 {
		fmt.Println("add EXPRESSION to tmp array")
		tmp = append(tmp, f.Expression)
	}
	if !funk.IsEmpty(f.Not) && len(*f.Not) > 0 {
		// for _, v := range *f.Not {
		// 	if !funk.IsEmpty(v) {
		// 		continue
		// 	}
		// 	break
		// }
		fmt.Println("add NOT to tmp array")
		tmp = append(tmp, f.Not)
	}
	if !funk.IsEmpty(f.And) && len(*f.And) > 0 {
		fmt.Println("add AND to tmp array")
		tmp = append(tmp, f.And)
	}
	if !funk.IsEmpty(f.Or) && len(*f.Or) > 0 {
		fmt.Println("add OR to tmp array")
		tmp = append(tmp, f.Or)
	}
	if !funk.IsEmpty(f.Function) && len(*f.Function) > 0 {
		fmt.Println("add FUNCTION to tmp array")
		tmp = append(tmp, f.Function)
	}

	dump.V(tmp)

	for idx, v := range tmp {
		if idx == 0 || idx == len(tmp) {
			fmt.Fprintf(&sb, "\n\t(\n\t\t%v\n\t)\n\t", v)
			// fmt.Fprintf(&sb, "\n\t(\n\t\t%v\n\t)\n\t", v)
		} else {
			fmt.Fprintf(&sb, "and\n\t(\n\t\t%v\n\t)\n", v)
			// fmt.Fprintf(&sb, "and\n\t(\n\t\t%v\n\t)\n", v)
		}
	}

	return fmt.Sprintf("\n(%v)", sb.String())
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

	// var m map[string]interface{}
	// err = c.Decode(&m)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	return nil, err
	// }
	// dump.V(m)

	// fmt.Printf("c.ToJSON(): %v\n", c.ToJSON())
	// dump.V(c.ToJSON())

	// var value []byte
	// value, err = jsonutil.EncodePretty(m)
	// if err != nil {
	// 	fmt.Printf("unable to decode the json string: %v\n", err)
	// }
	// dump.V(string(value))

	cfg = &DSLFilterConfig{}
	err = c.Decode(cfg)
	if err != nil {
		return nil, err
	}
	dump.V(cfg)

	defaults.MustSet(cfg)
	dump.V(cfg)

	// var value []byte
	// value, err = jsonutil.EncodePretty(cfg)
	// if err != nil {
	// 	fmt.Printf("unable to decode the json string: %v\n", err)
	// }
	// dump.V(string(value))

	fmt.Printf("\ncfg.Filter.String: %v\n\n", cfg.Filter.String())
	// if !funk.IsEmpty(cfg.Filter.Condition.Expression) {
	// 	fmt.Printf("filter condition expressions: %+v", *cfg.Filter.Condition.Expression)
	// }
	fmt.Printf("\nfilter condition: %+v\n", cfg.Filter.Condition)
	// fmt.Printf("\nfilter condition functions: %+v\n", cfg.Filter.Condition.Function)
	// fmt.Printf("filter condition functions - len: %v\n", (*cfg.Filter.Condition.Function)[0].Len)
	// // fmt.Printf("filter condition functions - any: %+v\n", (*cfg.Filter.Condition.Function)[1].Any)

	return cfg, nil
}
