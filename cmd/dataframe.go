package main

import (
	"fmt"
	"os"
	"strings"

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

	// columnOrder := []qframe.Order{
	// 	{Column: csvDF.ColumnNames()[0]},
	// 	{Column: csvDF.ColumnNames()[1]},
	// 	{Column: csvDF.ColumnNames()[2]},
	// 	{Column: csvDF.ColumnNames()[3]},
	// 	{Column: csvDF.ColumnNames()[4]},
	// }
	// fmt.Println(csvDF.Sort(columnOrder...))

	columnOrder := make([]qframe.Order, len(colNames))
	// columnOrder := make([]qframe.Order, csvDF.Len())
	for idx, cName := range colNames {
		columnOrder[idx] = qframe.Order{Column: cName}
	}
	// fmt.Printf("columnOrder: %#v\n", columnOrder)
	dump.V(columnOrder)

	sortedCsvDF := csvDF.Distinct().Sort(columnOrder...)
	fmt.Println(sortedCsvDF)

	// fmt.Println(sortedCsvDF.Distinct())
	fmt.Println(csvDF.Distinct())

	// sortedCsvDF.Filter(qframe.Filter{})
}
