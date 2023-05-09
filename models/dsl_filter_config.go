package models

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/antonmedv/expr"
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

const (
	asteriskStringValue = "*"
	nullStringValue     = "null"
)

var FlattenFunc = expr.Function(
	"flatten",
	func(params ...any) (any, error) {
		return Flatten(params...)
	},
	new(func([]interface{}) []interface{}),
	// new(func([]interface{}) bool),
)

// Flatten will take an array of nested array and return
// all nested elements in an array. e.g. [[1,2,[3]],4] -> [1,2,3,4]
// func Flatten(nested ...any) ([]interface{}, error) {
// 	// This function is failing with stack overflow due to recursion - use at own risk
// 	// TODO: handle error
// 	flattened := make([]interface{}, 0)

// 	for _, i := range nested {
// 		switch i := i.(type) {
// 		case []interface{}:
// 			flattenedSubArray, _ := Flatten(i)
// 			flattened = append(flattened, flattenedSubArray...)
// 		case interface{}:
// 			flattened = append(flattened, i)
// 		}
// 	}

// 	return flattened, nil
// }

func Flatten(nested ...any) (any, error) {
	dump.V(nested...)
	return funk.FlattenDeep(nested), nil
}

var SliceToStringsFunc = expr.Function(
	"sliceToStrings",
	func(params ...any) (any, error) {
		return SliceToStrings(params...)
	},
	new(func([]interface{}) interface{}),
)

func SliceToStrings(nested ...any) ([]string, error) {
	return arrutil.SliceToStrings(nested), nil
}

var GroupedExpressionFunc = expr.Function(
	"groupedExpression",
	func(params ...any) (any, error) {
		return groupedExpression(params...)
	},
	new(func([]interface{}, []interface{}, []interface{}, []interface{}) bool),
)

// TODO: refactor the below global variables to use singleton/other struct values
// var dataframe *qframe.QFrame = new(qframe.QFrame)
// var dataframe *qframe.QFrame

var dataframe qframe.QFrame
var msgVal string
var afterDateComparatorFunc = func(colVal *string) bool {
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

	fmt.Printf("\nafterDateComparatorFunc: result type=%T, value=%v\n", result, result)
	switch bresult := result.(type) {
	case bool:
		return bresult
	}
	return false
}

var eq = func(column string, arg interface{}) qframe.FilterClause {

	return qframe.Filter{Column: column, Comparator: filter.Eq, Arg: arg}
}

var after_date = func(column string, arg interface{}) qframe.FilterClause {
	fmt.Printf("\nafter_date: afterDateComparatorFunc type is %T\n\n", afterDateComparatorFunc)
	dump.V(arg)
	msgVal = arg.(string)
	return qframe.Filter{Column: column, Comparator: afterDateComparatorFunc, Arg: arg}
}

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

	// searchValues := []string{"H2001", "868", "null", "97004", "2023-01-01"}
	// fmt.Printf("\nsearchValues: %v\n", searchValues)
	// fmt.Println("\nsearchValues:")
	// dump.V(searchValues)

	if (!funk.IsEmpty(e.FieldPath) && strings.Contains(e.FieldPath, ",")) &&
		(!funk.IsEmpty(e.Operator) && strings.Contains(e.Operator, ",")) &&
		(!funk.IsEmpty(e.Value)) {

		fmt.Println("found group expression")
		// shouldReturn, returnValue := e.compositeExpression(searchValues)
		// if shouldReturn {
		// 	return returnValue
		// }
		// return fmt.Sprintf("groupExpression(%v, %v, %v)", e.FieldPath, e.Operator, e.getValueAsString())
		// return fmt.Sprintf(
		// 	"groupExpression(%v, %v, %v, %v)",
		// 	arrutil.AnyToString(strings.Split(e.FieldPath, ",")),
		// 	formatAnyArrToString(strings.Split(e.Operator, ",")),
		// 	e.getValueAsString(),
		// 	e.FieldPath,
		// )
		// return fmt.Sprintf(
		// 	"groupExpression(%v)", "[#.hContractId.string, #.packageBenefitPlanCode.string, #.segmentId.string]",
		// )

		groupedFieldNames := strings.Split(e.FieldPath, ",")
		fmt.Println("\ngroupedFieldNames:")
		dump.V(groupedFieldNames)

		groupedFieldValuesFromConfig := arrutil.AnyToStrings(e.Value)
		fmt.Println("\ngroupedFieldValuesFromConfig:")
		dump.V(groupedFieldValuesFromConfig)

		if funk.IsEmpty(dataframe) {
			fmt.Println("\ncreating dataframe:")

			dataframe = createDataFrame(groupedFieldNames, groupedFieldValuesFromConfig)

			fmt.Println(dataframe)
		}
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, #.membershipGroupData.array[:].groupNumber.string, #.effectiveDate.string]",
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, map(#?.membershipGroupData?.array??[],.groupNumber?.string??nil), #.effectiveDate.string]",
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, #?.membershipGroupData?.array!=nil??map(#?.membershipGroupData?.array??[],.groupNumber?.string), #.effectiveDate.string]",
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, map( filter(#?.membershipGroupData?.array[:], len(#?.membershipGroupData?.array??[]) > 0), .groupNumber?.string??nil), #.effectiveDate.string]",
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, map(filter(#?.membershipGroupData?.array??[], len(#?.membershipGroupData?.array??['']) > 0), .groupNumber?.string), #.effectiveDate.string]",
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, #?.membershipGroupData != nil && (#?.membershipGroupData?.array[0].groupNumber?.string ?? 'null'), #.effectiveDate.string]",
		// "[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, len(#?.membershipGroupData?.array ?? []), #.effectiveDate.string]",

		// exp := "#?.membershipGroupData != nil && (#?.membershipGroupData?.array[0].groupNumber?.string ?? 'null')"
		// env := ValueJsonMap

		return fmt.Sprintf(
			"groupedExpression(%v, %v, %v, %v)",
			formatAnyArrToString(strings.Split(e.FieldPath, ",")),
			formatAnyArrToString(strings.Split(e.Operator, ",")),
			e.getValueAsString(),
			"[#.hContractId.string, #.packageBenefitPlanCode.string,  #.segmentId.string, map(filter(#?.membershipGroupData?.array??[], len(#?.membershipGroupData?.array??['']) > 0), .groupNumber?.string), #.effectiveDate.string]",
		)

		//below not-working due to groupNumber being inside array
		// return fmt.Sprintf(
		// 	"groupExpression(%v)", "[#.hContractId.string, #.packageBenefitPlanCode.string, #.segmentId.string,#.groupNumber.string,#.effectiveDate.string]",
		// )
	}

	return fmt.Sprintf("(%v %v %v)", e.FieldPath, e.Operator, e.getValueAsString())
}

// TODO: handle error
func (e BooleanExpression) getValueAsString() string {
	switch val := e.Value.(type) {
	case string:
		return val
	case int64, uint64, float64:
		return strutil.MustString(val)
	case []interface{}:
		switch val[0].(type) {
		case string:
			return formatAnyArrToString(val)
		default:
			return arrutil.AnyToString(val)
		}
	default:
		return strutil.MustString(val) //panic when unsupported value type is used
	}
}

func formatAnyArrToString(arr any) string {
	var sb strings.Builder
	sb.WriteByte('[')

	switch val := arr.(type) {
	case []interface{}:
		for i, v := range val {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString("'")
			sb.WriteString(strutil.MustString(v))
			sb.WriteString("'")
		}
	case []string:
		for i, v := range val {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString("'")
			sb.WriteString(strutil.MustString(v))
			sb.WriteString("'")
		}
	}
	sb.WriteByte(']')

	return sb.String()
}

// arrutil.AnyToStrings(params[0])
// func groupedExpression(groupedFieldValuesFromMsg, groupedOperators, groupedFieldValuesFromConfig []string) (bool, error) {
func groupedExpression(params ...any) (result bool, err error) {

	fmt.Println("\nparams:")
	dump.V(params)

	//TODO: start - initialize the below values only once and not for every message processing
	// Thats because these values are coming from config and expected to not change for every message
	groupedFieldNames := arrutil.AnyToStrings(params[0])
	// fmt.Println("\ngroupedFieldNames:")
	// dump.V(groupedFieldNames)

	groupedOperators := arrutil.AnyToStrings(params[1])
	// fmt.Println("\ngroupedOperators:")
	// dump.V(groupedOperators)

	groupedFieldValuesFromConfig := arrutil.AnyToStrings(params[2])
	// fmt.Println("\ngroupedFieldValuesFromConfig:")
	// dump.V(groupedFieldValuesFromConfig)
	//TODO: end

	//TODO: avoid using dataframe global variables and instead use struct variables
	if funk.IsEmpty(dataframe) {
		fmt.Println("\ndataframe:")

		dataframe = createDataFrame(groupedFieldNames, groupedFieldValuesFromConfig)

		fmt.Println(dataframe)
	}

	groupedFieldValuesFromMsg := make([]string, len(groupedOperators))

	rawGroupedFieldValuesFromMsg := params[3]
	fmt.Println("\nrawGroupedFieldValuesFromMsg:")
	dump.V(rawGroupedFieldValuesFromMsg)

	var (
		arrLen int
		arrEle interface{}
		sb     strings.Builder
	)
	switch iv := rawGroupedFieldValuesFromMsg.(type) {
	case []interface{}:
		for arrLen, arrEle = range iv {
			// fmt.Printf("message value[%v] %v\n", i, e)
			switch v := arrEle.(type) {
			case string:
				fmt.Printf("%T type of message value[%v] %v\n", v, arrLen, v)
				groupedFieldValuesFromMsg[arrLen] = v
			case []interface{}:
				fmt.Printf("%T type of message value[%v] %v\n", v, arrLen, v)
				for idx, ival := range v {
					switch val := ival.(type) {
					case string:
						fmt.Printf("\t%T type of message value[%v] %v\n", val, idx, val)
						if idx != 0 {
							sb.WriteString("^")
						}
						sb.WriteString(val)
					default:
						return false, fmt.Errorf("array of array: unsupported data type: %T for message element: %v", arrEle, arrEle)
					}
				}
				if len(sb.String()) == 0 {
					sb.WriteString("null")
				}
				groupedFieldValuesFromMsg[arrLen] = sb.String()
			default:
				return false, fmt.Errorf("unsupported data type: %T for message element: %v", arrEle, arrEle)
			}
		}
	default:
		return false, fmt.Errorf("unsupported data type: %T for message element: %v", iv, iv)
	}

	fmt.Printf("\ngroupedFieldValuesFromMsg: %v\n", groupedFieldValuesFromMsg)
	dump.V(groupedFieldValuesFromMsg)

	result = evaluate(groupedFieldValuesFromMsg, groupedOperators)

	return result, nil
}

func evaluate(groupedFieldValuesFromMsg []string, groupedOperators []string) (result bool) {
	var filteredCsvDF qframe.QFrame
	var filterClause qframe.FilterClause
	for i := 0; i < len(groupedOperators); i++ {
		groupedFieldValueFromMsg := groupedFieldValuesFromMsg[i]
		groupedOperator := groupedOperators[i]

		if i == 0 {
			cName := dataframe.ColumnNames()[i]
			filterClause = getFilterClause(dataframe, cName, groupedOperator, groupedFieldValueFromMsg)

			fmt.Printf("\nfilterClause[%v]:\n", i)
			dump.V(filterClause)

			filteredCsvDF = dataframe.Filter(filterClause)

			fmt.Printf("\nfilteredCsvDF[%v]:\n", i)
			fmt.Println(filteredCsvDF)
		} else {
			if filteredCsvDF.Len() == 0 {
				break
			} else {
				cName := filteredCsvDF.ColumnNames()[i]
				if strings.Contains(cName, "[:]") {
					fmt.Printf("\nColumnName[%v]:\n", cName)
					if strings.HasPrefix(groupedFieldValueFromMsg, "[") {
						str := strings.ReplaceAll(groupedFieldValueFromMsg, "[", "")
						str = strings.ReplaceAll(str, "]", "")

						strs := strings.Split(str, ",")
						fmt.Printf("\nElements:%#v:\n", strs)
						filterClauses := make([]qframe.FilterClause, len(str))
						for idx, s := range strs {
							filterClauses[idx] = getFilterClause(filteredCsvDF, cName, groupedOperator, strings.TrimSpace(s))
						}
						fmt.Printf("\nArray filterClauses[%v]:\n", i)
						fmt.Println(filterClauses)
					}
				}

				filterClause = getFilterClause(filteredCsvDF, cName, groupedOperator, groupedFieldValueFromMsg)

				fmt.Printf("\nfilterClause[%v]:\n", i)
				dump.V(filterClause)

				filteredCsvDF = filteredCsvDF.Filter(filterClause)

				fmt.Printf("\nfilteredCsvDF[%v]:\n", i)
				fmt.Println(filteredCsvDF)
			}
		}
	}

	fmt.Printf("\nfinal filteredCsvDF:\n")
	fmt.Println(filteredCsvDF)

	result = filteredCsvDF.Len() > 0
	fmt.Printf("filter condition group expression result: %v\n\n", result)

	return result
}

func getFilterClause(df qframe.QFrame,
	cName, groupedOperator, groupedFieldValueFromMsg string) (filterClause qframe.FilterClause) {
	// switch groupedOperator {
	switch strings.TrimSpace(groupedOperator) {
	// case filter.Eq:
	case "eq":
		fmt.Printf("\neq Func: type is %T\n", eq)

		fmt.Printf("\nfilteredCsvDF.Select(%v).Distinct():\n", cName)
		distinctFilteredCsvDF := dataframe.Select(cName).Distinct()
		fmt.Println(distinctFilteredCsvDF)

		asteriskFilteredCsvDF := distinctFilteredCsvDF.Filter(
			eq(cName, asteriskStringValue),
		)
		fmt.Println("\nasteriskFilteredCsvDF:")
		fmt.Println(asteriskFilteredCsvDF)

		nullFilteredCsvDF := distinctFilteredCsvDF.Filter(
			eq(cName, nullStringValue),
		)
		fmt.Println("\nnullFilteredCsvDF:")
		fmt.Println(nullFilteredCsvDF)

		if groupedFieldValueFromMsg != nullStringValue {
			if asteriskFilteredCsvDF.Len()+nullFilteredCsvDF.Len() == distinctFilteredCsvDF.Len() {
				groupedFieldValueFromMsg = asteriskStringValue
			}
		} else {
			//TODO: handle all other possible values than '*' and 'null', may need looping based lookup
		}

		filterClause = eq(cName, groupedFieldValueFromMsg)
	case "after_date":
		fmt.Printf("\nafter_date Func: type is %T\n", after_date)
		filterClause = after_date(cName, groupedFieldValueFromMsg)
	}
	return filterClause
}

// func createDataFrame(groupedFieldNames, groupedFieldValuesFromConfig []string) *qframe.QFrame {
func createDataFrame(groupedFieldNames, groupedFieldValuesFromConfig []string) qframe.QFrame {
	csvData := make([]string, len(groupedFieldValuesFromConfig)+1)
	for i := 0; i < len(csvData); i++ {
		if i == 0 {
			csvData[i] = strings.Join(groupedFieldNames, ",")
		} else {
			csvData[i] = groupedFieldValuesFromConfig[i-1]
		}
	}

	fmt.Println("\ncsvData:")
	dump.V(csvData)

	//TODO: handle other data types than string
	// colNames := strings.Split(fieldNames, ",")
	colTypes := make(map[string]string, len(groupedFieldNames))
	for _, colName := range groupedFieldNames {
		colTypes[colName] = types.String
	}
	fmt.Println("\ncolTypes:")
	dump.V(colTypes)

	csvReader := strings.NewReader(strings.Join(csvData, "\n"))
	csvDF := qframe.ReadCSV(csvReader, csv.Types(colTypes))

	fmt.Println("\ncsvDF:")
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
	fmt.Println("\ncolumnOrder:")
	dump.V(columnOrder)

	sortedCsvDF := csvDF.Distinct().Sort(columnOrder...)
	fmt.Println("\nsortedCsvDF:")
	fmt.Println(sortedCsvDF)

	// return &sortedCsvDF
	return sortedCsvDF
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
