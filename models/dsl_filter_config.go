package models

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/thoas/go-funk"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/config/csv"
	"github.com/tobgu/qframe/filter"
	"github.com/tobgu/qframe/types"
)

type DSLFilterConfig struct {
	Filter Filter `mapstructure:"filter"`
}

type BooleanExpression struct {
	FieldPath string `mapstructure:"field_path"`
	Operator  string `mapstructure:"operator"`
	Value     any    `mapstructure:"value"`
	// Value     *any   `mapstructure:"value"`
}

func (e BooleanExpression) String() string {

	if (!funk.IsEmpty(e.FieldPath) && strings.Contains(e.FieldPath, ",")) &&
		(!funk.IsEmpty(e.Operator) && strings.Contains(e.Operator, ",")) &&
		(!funk.IsEmpty(e.Value)) {

		res := "found group expression"
		groupedFieldNames := strings.Split(e.FieldPath, ",")
		groupedOperators := strings.Split(e.Operator, ",")
		// groupedOperators :=  strings.TrimSpace(strings.Split(e.Operator, ","))
		groupedFieldValues := make([]string, e.valueLength())

		switch val := e.Value.(type) {
		case string:
			groupedFieldValues[0] = val
		case []interface{}:
			switch val[0].(type) {
			case string:
				for i, v := range val {
					groupedFieldValues[i] = v.(string)
				}
			default:
				return fmt.Sprintf("unsupported data-type(%T) in '[]interface{}' value field for group expression", val[0])
			}
		default:
			return fmt.Sprintf("unsupported data-type(%T) in value field for group expression", val)
		}

		dump.V(groupedFieldNames)
		dump.V(groupedOperators)
		// dump.V(groupedFieldValues)

		csvData := make([]string, len(groupedFieldValues)+1)
		for i := 0; i < len(csvData); i++ {
			if i == 0 {
				csvData[i] = e.FieldPath
				// csvData[i] = groupedFieldNames
			} else {
				csvData[i] = groupedFieldValues[i-1]
			}
		}

		dump.V(csvData)

		// colNames := strings.Split(fieldNames, ",")
		colTypes := make(map[string]string, len(groupedFieldNames))
		for _, colName := range groupedFieldNames {
			colTypes[colName] = types.String
		}
		dump.V(colTypes)

		csvReader := strings.NewReader(strings.Join(csvData, "\n"))
		csvDF := qframe.ReadCSV(csvReader, csv.Types(colTypes))
		fmt.Println(csvDF)

		// fmt.Println(
		// 	csvDF.Sort(
		// 		qframe.Order{
		// 			Column: csvDF.ColumnNames()[0],
		// 			// Reverse:  true,
		// 			// NullLast: true,
		// 		},
		// 	),
		// )

		columnOrder := make([]qframe.Order, len(groupedFieldNames))
		// columnOrder := make([]qframe.Order, csvDF.Len())
		for idx, cName := range groupedFieldNames {
			columnOrder[idx] = qframe.Order{Column: cName}
		}
		// fmt.Printf("columnOrder: %#v\n", columnOrder)
		dump.V(columnOrder)

		sortedCsvDF := csvDF.Distinct().Sort(columnOrder...)
		fmt.Println(sortedCsvDF)

		var msgVal string
		afterDateComparatorFunc := func(colVal *string) bool {
			var exprSb strings.Builder
			exprSb.WriteString(msgVal)
			// exprSb.WriteString("=") //this won't work
			// exprSb.WriteString("==")
			// exprSb.WriteString(">=")
			exprSb.WriteString(">")
			// exprSb.WriteString("<=")
			// exprSb.WriteString("<")
			exprSb.WriteString(*colVal)

			expression, err := govaluate.NewEvaluableExpression(exprSb.String())
			if err != nil {
				return false
			}
			result, err := expression.Evaluate(nil)
			if err != nil {
				return false
			}

			fmt.Printf("afterDateComparatorFunc: result type=%T, value=%v\n", result, result)
			switch bresult := result.(type) {
			case bool:
				return bresult
			}
			return false
		}
		fmt.Printf("afterDateComparatorFunc: type is %T\n\n", afterDateComparatorFunc)

		eq := func(column string, arg interface{}) qframe.FilterClause {
			return qframe.Filter{Column: column, Comparator: filter.Eq, Arg: arg}
		}
		fmt.Printf("eq Func: type is %T\n\n", eq)

		after_date := func(column string, arg interface{}) qframe.FilterClause {
			dump.V(arg)
			msgVal = arg.(string)
			return qframe.Filter{Column: column, Comparator: afterDateComparatorFunc, Arg: arg}
		}
		fmt.Printf("after_date Func: type is %T\n\n", after_date)

		filterOperatorMap := map[string]string{
			csvDF.ColumnNames()[0]: "=",
			csvDF.ColumnNames()[1]: "=",
			csvDF.ColumnNames()[2]: "=",
			csvDF.ColumnNames()[3]: "=",
			csvDF.ColumnNames()[4]: "after_date",
		}
		// fmt.Printf("filterOperatorMap: %v\n\n", filterOperatorMap)
		fmt.Println("filterOperatorMap:")
		dump.V(filterOperatorMap)

		searchValuesMap := map[string]string{
			csvDF.ColumnNames()[0]: "H2001",
			csvDF.ColumnNames()[1]: "018",
			csvDF.ColumnNames()[2]: "null",
			csvDF.ColumnNames()[3]: "*",
			csvDF.ColumnNames()[4]: "2021-12-31",
		}
		// fmt.Printf("searchValuesMap: %v\n\n", searchValuesMap)
		fmt.Println("searchValuesMap:")
		dump.V(searchValuesMap)

		filterClauses := make([]qframe.FilterClause, len(groupedFieldNames))
		for idx, cName := range groupedFieldNames {
			switch filterOperatorMap[cName] {
			case filter.Eq:
				filterClauses[idx] = eq(cName, searchValuesMap[cName])
			case "after_date":
				filterClauses[idx] = after_date(cName, searchValuesMap[cName])
			}
		}
		// fmt.Printf("filterClauses: %#v\n", filterClauses)
		dump.V(filterClauses)

		filteredCsvDF := sortedCsvDF.Filter(
			qframe.And(
				// filterClauses[0:3]...,
				filterClauses...,
			),
		)
		fmt.Println(filteredCsvDF)

		result := filteredCsvDF.Len() > 0
		fmt.Printf("filter condition group expression result: %v\n\n", result)
		return res
	}

	switch val := e.Value.(type) {
	case string:
		return fmt.Sprintf("(%v %v '%v')", e.FieldPath, e.Operator, val)
	case int64, uint64, float64:
		return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, val)
	case []interface{}:
		switch val[0].(type) {
		case string:
			var sb strings.Builder
			sb.WriteByte('[')
			for i, v := range val {
				if i > 0 {
					sb.WriteByte(',')
				}
				sb.WriteString("'")
				sb.WriteString(strutil.QuietString(v))
				sb.WriteString("'")
			}
			sb.WriteByte(']')
			return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, sb.String())
		default:
			return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, arrutil.AnyToString(val))
			// return fmt.Sprintf("(%v %v %T %v)", e.FieldPath, e.Operator, val[0], arrutil.AnyToString(val))
		}
	default:
		return "error"
	}
}

func (e BooleanExpression) valueLength() int {
	switch val := e.Value.(type) {
	case string:
		return 1
	case []interface{}:
		return len(val)
	}
	return 0
}

// SliceToString convert []any to string
func SliceToString(arr ...any) string { return ToString(arr) }

// ToString simple and quickly convert []any to string
func ToString(arr []any) string {
	// like fmt.Println([]any(nil))
	if arr == nil {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteByte('[')

	for i, v := range arr {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strutil.QuietString(v))
	}

	sb.WriteByte(']')
	return sb.String()
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
	return fmt.Sprintf("Filter Name:(%v) Type:(%v) and Condition:(%v)", f.Name, f.Type, f.Condition)
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

	// fmt.Printf("\nfilter: %v\n", cfg.Filter)
	fmt.Printf("\nfilter condition: %+v\n", cfg.Filter.Condition)

	return cfg, nil
}
