package main

import (
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/gookit/goutil/dump"
	"github.com/thiruselvaa/mql/models"
)

type Tweet struct {
	Len int
}

type Env struct {
	Tweets []Tweet
}

func main() {
	expression := `any(Tweets, {.Len in [0, 1, 2, 3]})`
	env := Env{
		Tweets: []Tweet{{1}, {10}, {11}},
	}

	result, err := mql(expression, env)
	if err != nil {
		panic(err)
	}

	fmt.Printf("result: %v\n", result)

	configFile := "configs/test-filter-query.json"
	// configFile := "configs/test-filter-query.yaml"
	models.NewSMFConfig(configFile)
}

func mql(expression string, env interface{}) (interface{}, error) {

	program, err := expr.Compile(expression)
	if err != nil {
		return nil, err
	}

	dump.V(program.Disassemble())
	dump.V(program.Node)

	if _, ok := env.(expr.Option); ok {
		return nil, fmt.Errorf("misused expr.Eval: second argument (env) should be passed without expr.Env")
	}

	result, err := expr.Run(program, env)
	if err != nil {
		return nil, err
	}

	return result, err
}
