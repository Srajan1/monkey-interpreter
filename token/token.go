package token

type TokenType string

// Note that it'll be a lot more perf friendly to use int/bytes for TokenType, but strings are easier to work with.
type Token struct {
	Type    TokenType // This type attribute of a token helps us diff between different types of tokens, whether it's an Identifier, or a variable, etc.
	Literal string    // The literal itself.
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENTIFIERS = "IDENTIFIERS"
	INT         = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimitors
	COMMA     = ","
	SEMICOLON = ";"

	// Special symbols
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// This will be used to check if a certain identifier is actually a language's keyword
var keyword = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
// If it is, it returns the keyword’s TokenType constant. If it isn’t, we just get back token.IDENT,
// which is the TokenType for all user-defined identifiers.
func LookupIdent(ident string) TokenType {
	tok, ok := keyword[ident]

	if ok {
		return tok
	}
	return IDENTIFIERS
}
