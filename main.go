package main

import (
	"fmt"

	"log"
	"path"

	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/parser"
)

func main() {
	vm := otto.New()

	program, err := parser.ParseFile(nil, path.Join("dist", "bundle.js"), nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	result, err := vm.Run(program)
	if err != nil {
		log.Fatal(err)
	}

	value, err := result.ToString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result:\n%s\n", value)
}
