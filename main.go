package main

import (
	"demo/lexer"
	"fmt"
)

func main() {
	input := "1+2"
	lexer := lexer.New(input)
	for lexer.HasNextToken() {
		token := lexer.NextToken()
		fmt.Printf("%+v\n", token)
	}
}
