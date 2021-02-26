package ast

// Node is the basic unit of an AST
type Node interface {
	TokenLiteral() string
	String() string
}
