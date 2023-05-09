package main

import (
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/thiruselvaa/mql/models"
)

func main() {
	dslConfigFile := "configs/dsl/solutran/json/solutran-dsl-filter-config.json"
	smfConfig, err := models.NewDSLFilterConfig(dslConfigFile)
	if err != nil {
		fmt.Printf("error parsing smf config file: %v", err)
		return
	}

	//TODO: sample code
	expression := smfConfig.Filter.Condition.String()
	env := map[string]interface{}{}
	result, err := mql(expression, env)
	if err != nil {
		fmt.Printf("error parsing the incoming message: %v", err)
		return
	}
	fmt.Printf("result: %v\n", result)
}

func mql(expression string, env interface{}) (result interface{}, err error) {
	program, err := expr.Compile(expression, models.GroupedExpressionFunc)
	if err != nil {
		return nil, err
	}

	if _, ok := env.(expr.Option); ok {
		return nil, fmt.Errorf("misused expr.Eval: second argument (env) should be passed without expr.Env")
	}

	result, err = expr.Run(program, env)
	if err != nil {
		return nil, err
	}

	return result, err
}
