package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/milangrahovac/geekbrainsgo/homework10/pkg/calculator"
)

func operatorChecker(s string) error {
	switch s {
	case
		"+",
		"-",
		"*",
		"/":
		return nil
	}
	fmt.Println("greska")
	return errors.New("error: wrong operator")
}

func main() {
	fmt.Println("homework 10")

	var (
		x  float64
		y  float64
		op string
	)

	flag.Float64Var(&x, "x", 0, "x value")
	flag.Float64Var(&y, "y", 0, "y value")
	flag.StringVar(&op, "op", "", "math operator")
	flag.Parse()

	flagset := make(map[string]bool)
	flag.Visit(
		func(f *flag.Flag) {
			flagset[f.Name] = true
		})

	if len(flagset) != 3 {
		log.Println("You must add x, y and op flag")
	} else {
		err := operatorChecker(op)
		if err != nil {
			log.Fatal(err)
		}

		c := calculator.Calc{X: x, Y: y, Operator: op}
		result := c.Calculate()
		fmt.Printf("%v %s %v = %v\n", c.X, op, c.Y, result)
	}

}
