package main

import (
	"fmt"
)

type Lexer struct {
	program string
	pos int
	tokens []Token
}

func getTokens(program string) Lexer {
	lexer := Lexer { program: program, pos: 0 }

	for {
		if lexer.pos >= len(lexer.program) {
			break
		}
		if program[lexer.pos] == ' ' {
			lexer.pos++
			continue
		}
		lexer.tokens = append(lexer.tokens, lexer.getToken())
	}
	return lexer
}

func (l *Lexer) getNumber() Token {	
	init := l.pos
	for l.pos < len(l.program) && isNumeric(l.program[l.pos]) {
		l.pos++
	}
	return Token {
		tokenType: number,
		value: l.program[init:l.pos],
	}
}

func isNumeric(c byte) bool {
	if c < '0' || c > '9' {
		return false
	}
	return true
}

func (l *Lexer) getToken() Token {
	switch l.program[l.pos] {
		case '+':
			l.pos++
			return Token{tokenType: plus, value: "+"}
		case '-':
			l.pos++
			return Token{tokenType: minus, value: "-"}
		case '*':
			l.pos++
			return Token{tokenType: mul, value: "*"}
		case '/':
			l.pos++
			return Token{tokenType: div, value: "/"}
		case '(':
			l.pos++
			return Token{tokenType: leftParen, value: "("}
		case ')':
			l.pos++
			return Token{tokenType: rightParen, value: ")"}
		default:
			if isNumeric(l.program[l.pos]) {
				return l.getNumber()
			}
	}
	panic("token not found")
}

func (l *Lexer) printTokens() {
	for _, t := range l.tokens {
		fmt.Println(t)
	}
}
