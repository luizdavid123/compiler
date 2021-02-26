package lexer

import (
	"demo/token"
)

// Lexer transforms a string int tokens
type Lexer struct {
	input string
	pos   int // current position in input (points to current char)
	rpos  int // current reading position in input (after current char)
	ch    byte
}

// New create a lexer from input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken return next token read
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	s := string(l.ch)
	switch l.ch {
	case '+':
		t = token.New(token.PLUS, s)
	case '-':
		t = token.New(token.MINUS, s)
	case '*':
		t = token.New(token.MULTIPLY, s)
	case '/':
		t = token.New(token.DIVIDE, s)
	case '(':
		t = token.New(token.LPAREN, s)
	case ')':
		t = token.New(token.RPAREN, s)
	case 0:
		t = token.New(token.EOF, "")
	default:
		if isDigit(l.ch) {
			s = l.readNumber()
			t = token.New(token.NUMBER, s)
			return t
		}
		t = token.New(token.ILLEGAL, s)
		return t
	}

	l.readChar()
	return t
}

// HasNextToken check whether curr char is EOF
func (l *Lexer) HasNextToken() bool {
	return l.ch != 0
}

func (l *Lexer) peekChar() byte {
	if l.rpos >= len(l.input) {
		return 0
	}
	return l.input[l.rpos]
}

func (l *Lexer) readChar() {
	if l.rpos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.rpos]
	}
	l.pos = l.rpos
	l.rpos++
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
