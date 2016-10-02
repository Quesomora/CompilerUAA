package main

import (
	"fmt"
	"io/ioutil"
	//	"os"
	//	"strings"
)

/*
func codeGenerator(n node) string {
	switch n.kind {
	case "program":
		var r []string
		for _, no := range n.body {
			r = append(r, codeGenerator(no))
		}
		return strings.Join(r, "\n")
	case "ExpressionStatement":
		return codeGenerator(*n.expression) + ";"
	case "expression":
		var ra []string
		c := codeGenerator(*n.callee)

		for _, no := range *n.arguments {
			ra = append(ra, codeGenerator(no))
		}
		r := strings.Join(ra, ", ")
		return c + "(" + r + ")"
	case "identificador":
		return n.name
	case "factor":
		return n.value
	default :
		fmt.Println("err")
		os.Exit(1)
		return ""
	}
}

func compiler(input string) string {
	tokens := tokenizer(input)
    fmt.Println("\n", tokens[:])
	tree := parser(tokens)
    fmt.Println(tree)
	ntree := transformer(tree)
	out := codeGenerator(node(ntree))
	return out
}*/

func main() {
	program, err := ioutil.ReadFile("C:/Users/queso/Documents/Go/src/github.com/Quesomora/CompilerUAA/simple.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(string(program))
	tokens := tokenizer(string(program))
	fmt.Println(tokens[:])
	//tree := parser(tokens)
	//fmt.Println("\n", tree)
	fmt.Println("--------------------------------------------------------------------------------")
}
