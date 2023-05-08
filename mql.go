package main

import (
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/thiruselvaa/mql/models"
)

func main() {
	// // configFile := "configs/native-filter-query.json"
	// configFile := "configs/native-filter-query.yaml"
	// models.NewSMFConfig(configFile)

	// qConfigFile := "configs/dsl-filter-query.json"
	// // qConfigFile := "configs/dsl-filter-query.yaml"
	// models.NewDSLQueryConfig(qConfigFile)

	// dslConfigFile := "configs/dsl-filter-query.json"
	// dslConfig, err := models.NewDSLConfig(dslConfigFile)
	// dslConfigFile := "configs/dsl-filter-config.json"
	dslConfigFile := "configs/dsl/solutran/json/solutran-dsl-filter-config.json"
	// dslConfig, err := models.NewDSLFilterConfig(dslConfigFile)
	_, err := models.NewDSLFilterConfig(dslConfigFile)
	if err != nil {
		fmt.Printf("error parsing smf config file: %v", err)
		return
	}
	// fmt.Println(dslConfig)
	// dump.V((*dslConfig.Filter.Condition.Function)[0].Any.Condition.Expression)

	// var value []byte
	// value, err = jsonutil.EncodePretty(dslConfig)
	// // value, err = jsonutil.Encode(dslConfig)
	// if err != nil {
	// 	fmt.Printf("unable to decode the json string: %v\n", err)
	// }
	// dump.V(string(value))
}

func mql(expression string, env interface{}) (interface{}, error) {

	program, err := expr.Compile(expression, models.GroupExpression)
	if err != nil {
		return nil, err
	}

	// dump.V(program.Disassemble())

	// dump.V(program.Node)
	// var value []byte
	// value, err = jsonutil.EncodePretty(program.Node)
	// if err != nil {
	// 	fmt.Printf("unable to decode the json string: %v\n", err)
	// }
	// dump.V(string(value))

	if _, ok := env.(expr.Option); ok {
		return nil, fmt.Errorf("misused expr.Eval: second argument (env) should be passed without expr.Env")
	}

	result, err := expr.Run(program, env)
	if err != nil {
		return nil, err
	}

	return result, err
}
