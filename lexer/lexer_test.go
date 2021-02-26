package lexer

import (
	"demo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "5+5*2/(1+2)-1"

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.NUMBER, "5"},
		{token.PLUS, "+"},
		{token.NUMBER, "5"},
		{token.MULTIPLY, "*"},
		{token.NUMBER, "2"},
		{token.DIVIDE, "/"},
		{token.LPAREN, "("},
		{token.NUMBER, "1"},
		{token.PLUS, "+"},
		{token.NUMBER, "2"},
		{token.RPAREN, ")"},
		{token.MINUS, "-"},
		{token.NUMBER, "1"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

	tok := l.NextToken()
	if tok.Type != token.EOF {
		t.Fatalf("tests failed - type wrong. expected=%q, got=%q",
			token.EOF, tok.Type)
	}
	if tok.Literal != "" {
		t.Fatalf("tests failed - literal wrong. expected=%q, got=%q",
			"", tok.Literal)
	}

}
