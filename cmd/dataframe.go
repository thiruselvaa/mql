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
	fmt.Println(newF)

	newF.ToJSON(os.Stdout)
	newF.ToCSV(os.Stdout)
	// fmt.Println(newF.ToJSON(os.Stdout))

	jsonStr := `[{"c1":2,"c2":"b"},{"c1":3,"c2":"c"}]`
	reader := strings.NewReader(jsonStr)
	jdf := qframe.ReadJSON(reader)
	fmt.Println(jdf)
	// result, reason := jdf.Equals(newF)
	// fmt.Printf("result: %v, reason: %v\n", result, reason)

	// 	csvStr := `hContractId,packageBenefitPlanCode,segmentId,groupNumber,effectiveDate
	// H0169,001,null,*,2020-12-31
	// H0169,"002",null,*,2020-12-31
	// H0169,"003",null,*,2020-12-31
	// H0169,"004",null,*,2020-12-31
	// H0251,"002",null,*,2020-12-31
	// H0251,"004",null,*,2020-12-31
	// `

	// csvReader := strings.NewReader(csvStr)
	// cdf := qframe.ReadCSV(csvReader)
	// fmt.Println(cdf)
	// fmt.Println(cdf.ColumnNames())
	// fmt.Println(cdf.ColumnNames()[0])

	// cdf.Filter(f)

	fieldNames := "hContractId,packageBenefitPlanCode,segmentId,groupNumber,effectiveDate"
	fieldValues := []string{
		"H0169,001,null,*,2020-12-31",
		"H0169,002,null,*,2020-12-31",
		"H0169,003,null,*,2020-12-31",
		"H0169,004,null,*,2020-12-31",
		"H0251,002,null,*,2020-12-31",
		"H0251,004,null,*,2020-12-31",
	}
	// csvData := []string{}
	csvData := make([]string, len(fieldValues)+1)
	// csvData[0] = fieldNames
	for i := 0; i < len(csvData); i++ {
		if i == 0 {
			csvData[i] = fieldNames
		} else {
			csvData[i] = fieldValues[i-1]
		}
	}
	// csvData = fieldValues
	// csvData = append(csvData, fieldValues...)
	dump.V(csvData)

	// csvStr := []string{
	// 	"hContractId,packageBenefitPlanCode,segmentId,groupNumber,effectiveDate",
	// 	"H0169,001,null,*,2020-12-31",
	// 	"H0169,002,null,*,2020-12-31",
	// 	"H0169,003,null,*,2020-12-31",
	// 	"H0169,004,null,*,2020-12-31",
	// 	"H0251,002,null,*,2020-12-31",
	// 	"H0251,004,null,*,2020-12-31",
	// }
	// csvCols := strings.Split(csvStr[0], ",")
	// // // csvRows := strings.Split(csvStr[0], ",")

	// fmt.Printf("len(csvCols): %v, len(csvStr): %#v\n", len(csvCols), len(csvStr))
	// // csvTable := make([][]string, len(csvCols), len(csvStr))
	// // csvTable := make([][]string, len(csvStr))
	// csvTable := [][]string{}

	// dump.V(csvTable)

	// for _, val := range csvStr {
	// 	// csvTable[idx] = strings.Split(val, ",")
	// 	// csvTable = append(csvTable, csvTable[idx])

	// 	csvTable = append(csvTable, strings.Split(val, ","))
	// }
	// fmt.Printf("csvTable: %#v\n", csvTable)

	// dump.V(csvTable)

	// dump.V(strings.Join(csvStr, "\n"))

	// // csvReader := strings.NewReader(csvTable)
	// // csvReader := strings.NewReader(csvStr)
	// csvReader := strings.NewReader(strings.Join(csvStr, "\n"))
	// cdf := qframe.ReadCSV(csvReader)
	// fmt.Println(cdf)

	colNames := strings.Split(fieldNames, ",")
	// colTypes := map[string]string{}
	colTypes := make(map[string]string, len(colNames))
	for _, colName := range colNames {
		colTypes[colName] = types.String

	}
	// colTypes := map[string]string{
	// 	colNames[0]: types.String,
	// 	colNames[1]: types.String,
	// 	colNames[2]: types.String,
	// 	colNames[3]: types.String,
	// 	colNames[4]: types.String,
	// }

	csvReader := strings.NewReader(strings.Join(csvData, "\n"))
	cdf := qframe.ReadCSV(csvReader, csv.Types(colTypes))
	fmt.Println(cdf)
	// fmt.Println(cdf.ColumnNames())
	// fmt.Println(cdf.ColumnNames()[0])

	// 	input := `COL1,COL2
	// a,1.5
	// b,2.25
	// c,3.0
	// `

	// csvDf := qframe.ReadCSV(strings.NewReader(input))
	// fmt.Println(csvDf)
}
