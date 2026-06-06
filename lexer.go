package main

import (
	"fmt"
)
	
var words = map[string]TokenType {
	"if": ifStmt,
	"for": forStmt,
}
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

func (l *Lexer) getWord() Token {
	init := l.pos
	for l.pos < len(l.program) && isAlpha(l.program[l.pos]) {
		l.pos++
	}

	val, ok := words[l.program[init:l.pos]]
	if ok == true {
		return Token {
			tokenType: val,
			value: l.program[init:l.pos],
		}
	}

	words[l.program[init:l.pos]] = id
	fmt.Println(l.program[init:l.pos])
	return Token {
		tokenType: id,
		value: l.program[init:l.pos],
	}
}

func isNumeric(c byte) bool {
	if c < '0' || c > '9' {
		return false
	}
	return true
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}

func (l *Lexer) peek() (byte, bool) {
	if l.pos < len(l.program) {
		return l.program[l.pos], true
	}

	return 0, false
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
		case '>':
			l.pos++
			ch, ok := l.peek()
			if ok {
				switch ch {
					case '=':
						l.pos++
						return Token{ tokenType: greaterEqual, value: ">=" }
					default:
						return Token{ tokenType: greater, value: ">" }
				}
			}
			fmt.Println("erro")
		case '<':
			l.pos++
			ch, ok := l.peek()
			if ok {
				switch ch {
					case '=':
						l.pos++
						return Token{ tokenType: lessEqual, value: ">=" }
					default:
						return Token{ tokenType: less, value: ">" }
				}
			}
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
			if isAlpha(l.program[l.pos]) {
				return l.getWord();
			}
	}
	panic("token not found")
}

func (l *Lexer) printTokens() {
	for _, t := range l.tokens {
		fmt.Println(t)
	}
}
