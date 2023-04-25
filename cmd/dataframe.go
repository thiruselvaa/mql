package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/gookit/goutil/dump"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/config/csv"
	"github.com/tobgu/qframe/types"
)

func main() {
	f := qframe.New(
		map[string]interface{}{
			"c1": []int{1, 2, 3},
			"c2": []string{"a", "b", "c"}},
	)
	args := map[string]interface{}{
		"intVal": 2,
		"strVal": "b",
	}
	newF := f.Filter(qframe.Or(
		qframe.Filter{Column: "c1", Comparator: ">", Arg: args["intVal"]},
		qframe.Filter{Column: "c2", Comparator: "=", Arg: args["strVal"]}))
	fmt.Println(newF.Len())

	newF.ToJSON(os.Stdout)
	newF.ToCSV(os.Stdout)
	// fmt.Println(newF.ToJSON(os.Stdout))

	jsonStr := `[{"c1":2,"c2":"b"},{"c1":3,"c2":"c"}]`
	reader := strings.NewReader(jsonStr)
	jdf := qframe.ReadJSON(reader)
	fmt.Println(jdf)
	// result, reason := jdf.Equals(newF)
	// fmt.Printf("result: %v, reason: %v\n", result, reason)

	fieldNames := "hContractId,packageBenefitPlanCode,segmentId,groupNumber,effectiveDate"
	fieldValues := []string{
		"H0169,003,null,*,2021-12-31",
		"H0251,004,null,*,2021-12-31",
		"H0169,001,null,*,2020-12-31",
		"H0251,002,null,*,2020-12-31",
		"H0169,002,null,*,2020-12-31",
		"H0251,004,null,*,2021-12-31",
		"H0169,004,null,*,2022-12-31",
	}

	// fieldValues := []string{
	// 	"H0169,003,null,*,2021-12-31",
	// 	"H0251,004,null,*,2021-12-31",
	// 	"H0169,001,null,*,2020-12-31",
	// 	"H0251,002,null,12345,2020-12-31",
	// 	"H0169,002,null,*,2020-12-31",
	// 	"H0251,004,null,*,2021-12-31",
	// 	"H0169,004,null,*,2022-12-31",
	// }

	csvData := make([]string, len(fieldValues)+1)
	for i := 0; i < len(csvData); i++ {
		if i == 0 {
			csvData[i] = fieldNames
		} else {
			csvData[i] = fieldValues[i-1]
		}
	}

	dump.V(csvData)

	colNames := strings.Split(fieldNames, ",")
	colTypes := make(map[string]string, len(colNames))
	for _, colName := range colNames {
		colTypes[colName] = types.String

	}

	csvReader := strings.NewReader(strings.Join(csvData, "\n"))
	csvDF := qframe.ReadCSV(csvReader, csv.Types(colTypes))

	// dump.V(csvDF)

	fmt.Println(csvDF)
	fmt.Printf("csvDF.Len(): %v\n", csvDF.Len())
	fmt.Printf("csvDF.ColumnNames(): %v\n", csvDF.ColumnNames())
	fmt.Printf("csvDF.ColumnNames()[0]: %v\n", csvDF.ColumnNames()[0])

	fmt.Println(csvDF.Select(csvDF.ColumnNames()[0], csvDF.ColumnNames()[1]))
	fmt.Println(csvDF.Select(csvDF.ColumnNames()[0]).Distinct())

	fmt.Println(csvDF.WithRowNums("rowNum"))
	fmt.Println(csvDF.Sort(qframe.Order{
		Column: "hContractId",
		// Reverse:  true,
		// NullLast: true,
	}))

	columnOrder := make([]qframe.Order, len(colNames))
	// columnOrder := make([]qframe.Order, csvDF.Len())
	for idx, cName := range colNames {
		columnOrder[idx] = qframe.Order{Column: cName}
	}
	// fmt.Printf("columnOrder: %#v\n", columnOrder)
	dump.V(columnOrder)

	sortedCsvDF := csvDF.Distinct().Sort(columnOrder...)
	fmt.Println(sortedCsvDF)

	// msgVals := []string{
	// 	"2021-12-31",
	// 	// "2021-12-31",
	// 	// "2021-12-31",
	// 	// "2021-12-31",
	// 	// "2021-12-31",
	// 	// "2021-12-31",
	// }
	// msgFieldValue := map[string]interface{}{
	// 	"effectiveDateFromMsg": msgVals,
	// }
	// msgDF := qframe.New(msgFieldValue)
	// fmt.Println(msgDF)

	// // sortedCsvDFWithmsgDF := sortedCsvDF.Append(msgDF)
	// // fmt.Println(sortedCsvDFWithmsgDF)

	// sortedCsvDFWithmsgDF := sortedCsvDF.Copy("effectiveDateFromMsg", "effectiveDate")
	// fmt.Println(sortedCsvDFWithmsgDF)

	// afterDate := func(colVal, msgVal *string) bool {
	// 	*msgVal = "2021-12-31"

	// 	var exprSb strings.Builder
	// 	exprSb.WriteString(*msgVal)
	// 	exprSb.WriteString(">")
	// 	exprSb.WriteString(*colVal)

	// 	// expression, err := govaluate.NewEvaluableExpression("'2014-01-02' > '2014-01-01 23:59:59'")

	// 	expression, err := govaluate.NewEvaluableExpression(exprSb.String())
	// 	if err != nil {
	// 		return false
	// 	}
	// 	result, err := expression.Evaluate(nil)
	// 	if err != nil {
	// 		return false
	// 	}

	// 	fmt.Printf("afterDateComparatorFunc: result=%v\n", result)
	// 	return result.(bool)
	// }

	msgVal := "2021-12-31"
	// sortedCsvDF = sortedCsvDF.Apply(
	// 	qframe.Instruction{
	// 		Fn: func(m *string) *string { return &msgVal },
	// 		// DstCol:  "after_date",
	// 		DstCol:  "effectiveDateFromMsg",
	// 		SrcCol1: "effectiveDate"},
	// )
	// fmt.Println(sortedCsvDF)

	// msgVal = "nil"
	sortedCsvDF = sortedCsvDF.Apply(
		qframe.Instruction{
			Fn:     &msgVal,
			DstCol: "effectiveDateFromMsg",
		},
	)
	fmt.Println(sortedCsvDF)

	afterDate := false
	sortedCsvDF = sortedCsvDF.Apply(
		qframe.Instruction{
			Fn:     afterDate,
			DstCol: "after_date",
		},
	)
	fmt.Println(sortedCsvDF)

	// afterDateComparatorFunc := func(colVal, msgVal *string) *bool {
	// 	var (
	// 		exprSb     strings.Builder
	// 		boolResult bool
	// 	)

	// 	// msgVal := "2021-12-31"
	// 	// exprSb.WriteString(msgVal)
	// 	exprSb.WriteString(*msgVal)
	// 	exprSb.WriteString(">")
	// 	exprSb.WriteString(*colVal)

	// 	expression, err := govaluate.NewEvaluableExpression(exprSb.String())
	// 	if err != nil {
	// 		return &boolResult
	// 	}
	// 	result, err := expression.Evaluate(nil)
	// 	if err != nil {
	// 		return &boolResult
	// 	}

	// 	fmt.Printf("afterDateComparatorFunc: result type=%T, value=%v\n", result, result)
	// 	switch bresult := result.(type) {
	// 	case bool:
	// 		return &bresult
	// 	}
	// 	return &boolResult
	// 	// return &msgVal
	// 	// return msgVal

	// 	// switch result.(type) {
	// 	// case bool:
	// 	// 	*msgVal = strutil.MustString(result)
	// 	// 	return msgVal
	// 	// }
	// 	// return msgVal
	// }

	// sortedCsvDF = sortedCsvDF.Apply(
	// 	qframe.Instruction{
	// 		Fn:      afterDateComparatorFunc,
	// 		DstCol:  "after_date",
	// 		SrcCol1: "effectiveDate",
	// 		// SrcCol2: "effectiveDateFromMsg",
	// 		SrcCol2: "effectiveDateFromMsg",
	// 	},
	// )
	// fmt.Println(sortedCsvDF)

	//
	// afterDateComparatorFunc := func(colVal, msgVal *string) *string {
	// 	var (
	// 		exprSb strings.Builder
	// 		// boolResult bool
	// 	)

	// 	// msgVal := "2021-12-31"
	// 	// exprSb.WriteString(msgVal)
	// 	exprSb.WriteString(*msgVal)
	// 	exprSb.WriteString(">")
	// 	exprSb.WriteString(*colVal)

	// 	expression, err := govaluate.NewEvaluableExpression(exprSb.String())
	// 	if err != nil {
	// 		// return &msgVal
	// 		return msgVal
	// 	}
	// 	result, err := expression.Evaluate(nil)
	// 	if err != nil {
	// 		// return &msgVal
	// 		return msgVal
	// 	}

	// 	fmt.Printf("afterDateComparatorFunc: result type=%T, value=%v\n", result, result)
	// 	// switch bresult := result.(type) {
	// 	// case bool:
	// 	// 	return bresult
	// 	// }
	// 	// return boolResult
	// 	// return &msgVal
	// 	// return msgVal

	// 	switch result.(type) {
	// 	case bool:
	// 		*msgVal = strutil.MustString(result)
	// 		return msgVal
	// 	}
	// 	return msgVal
	// }

	// sortedCsvDF = sortedCsvDF.Apply(
	// 	qframe.Instruction{
	// 		Fn:      afterDateComparatorFunc,
	// 		DstCol:  "after_date",
	// 		SrcCol1: "effectiveDate",
	// 		// SrcCol2: "effectiveDateFromMsg",
	// 		SrcCol2: "effectiveDateFromMsg",
	// 	},
	// )
	// fmt.Println(sortedCsvDF)
	//

	// dump.V(sortedCsvDF.ColumnTypes())
	// dump.V(sortedCsvDF.ColumnTypeMap())

	input := qframe.New(map[string]interface{}{
		"COL1": []string{"2020-12-31", "2020-12-31", "2020-12-31"},
		"COL2": []string{"2021-12-31", "2021-12-31", "2021-12-31"},
	})
	fmt.Println(input)

	// output := input.Apply(qframe.Instruction{Fn: func(x *string) *string { return &msgVal }, DstCol: "IS_LONG", SrcCol1: "COL1"})
	// output := input.Apply(qframe.Instruction{Fn: func(x, y *string) bool { return len(*x) >= len(*y) }, DstCol: "IS_LONG", SrcCol1: "COL1", SrcCol2: "COL2"})
	// output := input.Apply(qframe.Instruction{Fn: func(x *string) bool { return x == &msgVal }, DstCol: "IS_EQUAL", SrcCol1: "COL1"})
	// output := input.Apply(qframe.Instruction{Fn: func(x *string) bool { return len(*x) > 2 }, DstCol: "IS_LONG", SrcCol1: "COL1"})
	// output := input.Apply(
	// 	qframe.Instruction{
	// 		Fn: func(colVal *string) bool {
	// 			var exprSb strings.Builder
	// 			exprSb.WriteString(msgVal)
	// 			exprSb.WriteString(">")
	// 			exprSb.WriteString(*colVal)

	// 			expression, err := govaluate.NewEvaluableExpression(exprSb.String())
	// 			if err != nil {
	// 				return false
	// 			}
	// 			result, err := expression.Evaluate(nil)
	// 			if err != nil {
	// 				return false
	// 			}

	// 			fmt.Printf("afterDateComparatorFunc: result type=%T, value=%v\n", result, result)
	// 			switch bresult := result.(type) {
	// 			case bool:
	// 				return bresult
	// 			}
	// 			return false
	// 			// return colVal == &msgVal
	// 		},
	// 		DstCol:  "IS_EQUAL",
	// 		SrcCol1: "COL1",
	// 	},
	// )
	// fmt.Println(output)

	afterDateComparatorFunc := func(colVal *string) bool {
		var exprSb strings.Builder
		exprSb.WriteString(msgVal)
		exprSb.WriteString(">")
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
		// return colVal == &msgVal
	}

	output := sortedCsvDF.Apply(
		qframe.Instruction{
			Fn:      afterDateComparatorFunc,
			DstCol:  "after_date",
			SrcCol1: "effectiveDate",
			// SrcCol1: "COL1",
		},
	)
	fmt.Println(output)

	// filterOperatorMap := map[string]string{
	// 	csvDF.ColumnNames()[0]: "=",
	// 	csvDF.ColumnNames()[1]: "=",
	// 	csvDF.ColumnNames()[2]: "=",
	// 	csvDF.ColumnNames()[3]: "=",
	// 	csvDF.ColumnNames()[4]: "after_date",
	// }

	// searchValuesMap := map[string]string{
	// 	// "hContractId": "H0251",
	// 	// "  hContractId  ": "H0251",
	// 	// "hContractId  ": "H0251",
	// 	csvDF.ColumnNames()[0]: "H0251",

	// 	csvDF.ColumnNames()[1]: "002",
	// 	// csvDF.ColumnNames()[1]: "003",

	// 	// csvDF.ColumnNames()[2]: "",
	// 	csvDF.ColumnNames()[2]: "null",

	// 	// csvDF.ColumnNames()[3]: "",
	// 	csvDF.ColumnNames()[3]: "*",

	// 	csvDF.ColumnNames()[4]: "2021-12-31",
	// }

	// dump.V(searchValuesMap)

	// eq := func(column string, arg interface{}) qframe.FilterClause {
	// 	return qframe.Filter{Column: column, Comparator: filter.Eq, Arg: arg}
	// }
	// // func(f float64) bool { return f > 1.2 }

	// // afterDateComparatorFunc := func(colVal, msgVal string) bool {
	// // 	var exprSb strings.Builder
	// // 	exprSb.WriteString(msgVal)
	// // 	exprSb.WriteString(">")
	// // 	exprSb.WriteString(colVal)

	// // 	// expression, err := govaluate.NewEvaluableExpression("'2014-01-02' > '2014-01-01 23:59:59'")

	// // 	expression, err := govaluate.NewEvaluableExpression(exprSb.String())
	// // 	if err != nil {
	// // 		return false
	// // 	}
	// // 	result, err := expression.Evaluate(nil)
	// // 	if err != nil {
	// // 		return false
	// // 	}

	// // 	fmt.Printf("afterDateComparatorFunc: result=%v\n", result)
	// // 	return result.(bool)
	// // }

	// // fmt.Printf("afterDateComparatorFunc: %v\n", afterDateComparatorFunc("2020-12-31", "2021-12-31"))

	// // after_date := func(column string, arg interface{}) qframe.FilterClause {
	// // 	return qframe.Filter{Column: column, Comparator: afterDateComparatorFunc, Arg: arg}
	// // }

	// afterDateComparatorFunc := func(colVal, msgVal *string) bool {
	// 	var exprSb strings.Builder
	// 	exprSb.WriteString(*msgVal)
	// 	exprSb.WriteString(">")
	// 	exprSb.WriteString(*colVal)

	// 	// expression, err := govaluate.NewEvaluableExpression("'2014-01-02' > '2014-01-01 23:59:59'")

	// 	expression, err := govaluate.NewEvaluableExpression(exprSb.String())
	// 	if err != nil {
	// 		return false
	// 	}
	// 	result, err := expression.Evaluate(nil)
	// 	if err != nil {
	// 		return false
	// 	}

	// 	fmt.Printf("afterDateComparatorFunc: result=%v\n", result)
	// 	return result.(bool)
	// }

	// colVal := "2020-12-31"
	// msgVal := "2021-12-31"
	// fmt.Printf("afterDateComparatorFunc: %v\n", afterDateComparatorFunc(&colVal, &msgVal))

	// after_date := func(column, arg string) qframe.FilterClause {
	// 	// return qframe.Filter{Column: column, Comparator: afterDateComparatorFunc, Arg: map[string]interface{}{"tmp": []string{arg}}}
	// 	return qframe.Filter{Column: column, Comparator: afterDateComparatorFunc, Arg: types.ColumnName(arg)}
	// 	// return qframe.Filter{Column: column, Comparator: afterDateComparatorFunc, Arg: arg}
	// }

	// // func col(c string) types.ColumnName {
	// // 	return types.ColumnName(c)
	// // }

	// // csvDF.FilteredApply()

	// filterClauses := make([]qframe.FilterClause, len(colNames))
	// for idx, cName := range colNames {
	// 	switch filterOperatorMap[cName] {
	// 	case filter.Eq:
	// 		filterClauses[idx] = eq(cName, searchValuesMap[cName])
	// 	case "after_date":
	// 		filterClauses[idx] = after_date(cName, searchValuesMap[cName])
	// 	}

	// 	// filterClauses[idx] = qframe.Filter{
	// 	// 	Column:     cName,
	// 	// 	Comparator: filterOperatorMap[cName],
	// 	// 	Arg:        searchValuesMap[cName],
	// 	// }
	// }
	// // fmt.Printf("filterClauses: %#v\n", filterClauses)
	// dump.V(filterClauses)

	// filteredCsvDF := sortedCsvDF.Filter(
	// 	qframe.And(
	// 		// filterClauses[0:3]...,
	// 		filterClauses...,
	// 	),
	// )

	// filteredCsvDF.Apply()

	// fmt.Println(filteredCsvDF)
}

/*
Attempts to parse the [candidate] as a Time.
Tries a series of standardized date formats, returns the Time if one applies,
otherwise returns false through the second return.
*/
func tryParseTime(candidate string) (time.Time, bool) {

	var ret time.Time
	var found bool

	timeFormats := [...]string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.Kitchen,
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02",                         // RFC 3339
		"2006-01-02 15:04",                   // RFC 3339 with minutes
		"2006-01-02 15:04:05",                // RFC 3339 with seconds
		"2006-01-02 15:04:05-07:00",          // RFC 3339 with seconds and timezone
		"2006-01-02T15Z0700",                 // ISO8601 with hour
		"2006-01-02T15:04Z0700",              // ISO8601 with minutes
		"2006-01-02T15:04:05Z0700",           // ISO8601 with seconds
		"2006-01-02T15:04:05.999999999Z0700", // ISO8601 with nanoseconds
	}

	for _, format := range timeFormats {

		ret, found = tryParseExactTime(candidate, format)
		if found {
			return ret, true
		}
	}

	return time.Now(), false
}

func tryParseExactTime(candidate string, format string) (time.Time, bool) {

	var ret time.Time
	var err error

	ret, err = time.Parse(format, candidate)
	if err != nil {
		return time.Now(), false
	}

	return ret, true
}
