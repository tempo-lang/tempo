package token

type TokenType string

type SourcePos struct {
	Line int
	Col  int
}

type TokenError struct {
	Message string
}

type Token struct {
	Type    TokenType
	Literal string
	Value   any
	Pos     SourcePos
}

func New(kind TokenType, literal string, pos SourcePos, value any) Token {
	return Token{
		Type:    kind,
		Literal: literal,
		Pos:     pos,
		Value:   value,
	}
}

func Error(literal string, pos SourcePos, errorMessage string) Token {
	return Token{
		Type:    ILLEGAL,
		Literal: literal,
		Pos:     pos,
		Value: TokenError{
			Message: errorMessage,
		},
	}
}

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Keywords
	STRUCT     TokenType = "STRUCT"
	INTERFACE  TokenType = "INTERFACE"
	IMPLEMENTS TokenType = "IMPLEMENTS"
	FUNC       TokenType = "FUNC"
	RETURN     TokenType = "RETURN"
	LET        TokenType = "LET"
	ASYNC      TokenType = "ASYNC"
	AWAIT      TokenType = "AWAIT"
	IF         TokenType = "IF"
	ELSE       TokenType = "ELSE"
	WHILE      TokenType = "WHILE"
	TRUE       TokenType = "TRUE"
	FALSE      TokenType = "FALSE"

	// Parenthesis and brackets
	LPAREN  TokenType = "LPAREN"
	RPAREN  TokenType = "RPAREN"
	LSQUARE TokenType = "LSQUARE"
	RSQUARE TokenType = "RSQUARE"
	LCURLY  TokenType = "LCURLY"
	RCURLY  TokenType = "RCURLY"

	// Arithmetic
	PLUS     TokenType = "PLUS"
	MINUS    TokenType = "MINUS"
	MULTIPLY TokenType = "MULTIPLY"
	DIVIDE   TokenType = "DIVIDE"
	MODULO   TokenType = "MODULO"

	// Comparison
	EQUAL      TokenType = "EQUAL"
	NOT_EQUAL  TokenType = "NOT_EQUAL"
	LESS       TokenType = "LESS"
	LESS_EQ    TokenType = "LESS_EQ"
	GREATER    TokenType = "GREATER"
	GREATER_EQ TokenType = "GREATER_EQ"

	// Logical
	AND TokenType = "AND"
	OR  TokenType = "OR"

	// Assignment and symbols
	ASSIGN     TokenType = "ASSIGN"     // =
	ROLE_AT    TokenType = "ROLE_AT"    // @
	COMMA      TokenType = "COMMA"      // ,
	DOT        TokenType = "DOT"        // .
	COLON      TokenType = "COLON"      // :
	SEMICOLON  TokenType = "SEMICOLON"  // ;
	UNDERSCORE TokenType = "UNDERSCORE" // _
	COM        TokenType = "COM"        // ->

	// Literals and identifiers
	STRING TokenType = "STRING"
	IDENT  TokenType = "IDENT"
	FLOAT  TokenType = "FLOAT"
	INT    TokenType = "INT"
)
