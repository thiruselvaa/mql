package main

import (
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/gookit/goutil/dump"
	"github.com/thiruselvaa/mql/models"
	"github.com/thoas/go-funk"
)

type Message struct {
	Data map[string]interface{} `expr:"message"`
}

func main() {
	dslConfigFile := "configs/dsl/solutran/json/solutran-dsl-filter-config.json"
	smfConfig, err := models.NewDSLFilterConfig(dslConfigFile)
	if err != nil {
		fmt.Printf("error parsing smf config file: %v", err)
		return
	}

	//TODO: sample code
	expression := smfConfig.Filter.Condition.String()
	env := Message{
		Data: map[string]interface{}{},
	}
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

	dump.V(env)

	switch e := env.(type) {
	case Message:
		if !funk.IsEmpty(e.Data) {
			result, err = expr.Run(program, env)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unsupported data type: %T for message format: %v", env, env)
	}

	return result, err
}
