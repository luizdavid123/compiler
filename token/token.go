package token

// Type is the type of a token
type Type string

// Token represents a token used to construct a AST
type Token struct {
	Type    Type
	Literal string
}

// New return a token
func New(t Type, l string) Token {
	return Token{Type: t, Literal: l}
}

// Predefined type of tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Literal
	NUMBER = "NUMBER"

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	LPAREN = "("
	RPAREN = ")"
)
