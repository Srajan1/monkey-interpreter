package lexer

import (
	"github.com/Srajan1/monkey-interpreter/token"
)

type Lexer struct {
	input        string
	position     int // Points to the position which was last read
	readPosition int // Points to the position which will be read next
	ch           byte
}

// The purpose of readChar is to give us the next character and advance our position in the input
// string. The first thing it does is to check whether we have reached the end of input. If that’s
// the case it sets l.ch to 0, which is the ASCII code for the "NUL" character and signifies either
// “we haven’t read anything yet” or “end of file” for us. But if we haven’t reached the end of
// input yet it sets l.ch to the next character by accessing l.input[l.readPosition].
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// peekChar() is really similar to readChar(), except that it doesn’t increment l.position and
// l.readPosition. We only want to “peek” ahead in the input and not move around in it, so we know what a call to readChar() would return.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// readIdentifier() does exactly what its name suggests: it reads in an identifier and advances
// our lexer’s positions until it encounters a non-letter-character.
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// We implified things a lot in readNumber. We only read in
// integers. We ignore floats, hex, octals
// and just say that Monkey doesn’t support this.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}

}

// We look at the current character under
// examination (l.ch) and return a token depending on which character it is. Before returning the
// token we advance our pointers into the input so when we call NextToken() again the l.ch field
// is already updated
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			return token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	// Finding IDENTIFIERS
	// 	We added a default branch to our switch statement, so we can check for identifiers whenever
	// 	the l.ch is not one of the recognized characters. What our lexer needs to do is recognize whether
	// 	the current character is a letter and if so, it needs to read the rest of the identifier/keyword
	// 	until it encounters a non-letter-character. Having read that identifier/keyword, we then need
	// 	to find out if it is a identifier or a keyword, so we can use the correct token.TokenType
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok

}

func newToken(tokenType token.TokenType, tokenLiteral byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(tokenLiteral)}
}

// As you can see, in our case it contains the check ch ==
// '_', which means that we’ll treat _ as a letter and allow it
// in identifiers and keywords. That means we can use variable
// names like foo_bar. Other programming languages even allow ! and ? in identifiers. If we
// want to allow that too, this is the place to sneak it in.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}
