package main

import (
	"fmt"

	"github.com/antonmedv/expr"
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
}

func mql(expression string, env interface{}) (result interface{}, err error) {
	return expr.Eval(expression, env)
}
