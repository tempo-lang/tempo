// Code generated from Tempo.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TempoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TempoLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tempolexerLexerInit() {
	staticData := &TempoLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'struct'", "'interface'", "'func'", "'return'", "'let'", "'async'",
		"'await'", "'if'", "'else'", "'true'", "'false'", "'('", "'['", "'{'",
		"')'", "']'", "'}'", "'+'", "'-'", "'*'", "'/'", "'%'", "'=='", "'!='",
		"'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'='", "'@'", "','", "'.'",
		"':'", "'->'", "", "", "", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "STRUCT", "INTERFACE", "FUNC", "RETURN", "LET", "ASYNC", "AWAIT",
		"IF", "ELSE", "TRUE", "FALSE", "LPAREN", "LSQUARE", "LCURLY", "RPAREN",
		"RSQUARE", "RCURLY", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO",
		"EQUAL", "NOT_EQUAL", "LESS", "LESS_EQ", "GREATER", "GREATER_EQ", "AND",
		"OR", "IS", "ROLE_AT", "COMMA", "DOT", "COLON", "COM", "STRING", "ID",
		"NUMBER", "END", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"STRUCT", "INTERFACE", "FUNC", "RETURN", "LET", "ASYNC", "AWAIT", "IF",
		"ELSE", "TRUE", "FALSE", "LPAREN", "LSQUARE", "LCURLY", "RPAREN", "RSQUARE",
		"RCURLY", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "EQUAL",
		"NOT_EQUAL", "LESS", "LESS_EQ", "GREATER", "GREATER_EQ", "AND", "OR",
		"IS", "ROLE_AT", "COMMA", "DOT", "COLON", "COM", "STRING", "ESC", "SAFECODEPOINT",
		"ID", "NUMBER", "END", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 252, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 36, 5, 36, 212, 8, 36, 10, 36, 12, 36, 215, 9,
		36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 4, 39,
		226, 8, 39, 11, 39, 12, 39, 227, 1, 39, 1, 39, 5, 39, 232, 8, 39, 10, 39,
		12, 39, 235, 9, 39, 3, 39, 237, 8, 39, 1, 40, 4, 40, 240, 8, 40, 11, 40,
		12, 40, 241, 1, 41, 1, 41, 1, 42, 4, 42, 247, 8, 42, 11, 42, 12, 42, 248,
		1, 42, 1, 42, 0, 0, 43, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 0, 77, 0, 79, 38, 81, 39, 83, 40, 85, 41, 1, 0, 6,
		5, 0, 34, 34, 92, 92, 110, 110, 114, 114, 116, 116, 3, 0, 0, 31, 34, 34,
		92, 92, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 65, 90, 97, 122, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 256, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0,
		73, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 1, 87, 1, 0, 0, 0, 3, 94, 1, 0, 0, 0, 5, 104, 1, 0,
		0, 0, 7, 109, 1, 0, 0, 0, 9, 116, 1, 0, 0, 0, 11, 120, 1, 0, 0, 0, 13,
		126, 1, 0, 0, 0, 15, 132, 1, 0, 0, 0, 17, 135, 1, 0, 0, 0, 19, 140, 1,
		0, 0, 0, 21, 145, 1, 0, 0, 0, 23, 151, 1, 0, 0, 0, 25, 153, 1, 0, 0, 0,
		27, 155, 1, 0, 0, 0, 29, 157, 1, 0, 0, 0, 31, 159, 1, 0, 0, 0, 33, 161,
		1, 0, 0, 0, 35, 163, 1, 0, 0, 0, 37, 165, 1, 0, 0, 0, 39, 167, 1, 0, 0,
		0, 41, 169, 1, 0, 0, 0, 43, 171, 1, 0, 0, 0, 45, 173, 1, 0, 0, 0, 47, 176,
		1, 0, 0, 0, 49, 179, 1, 0, 0, 0, 51, 181, 1, 0, 0, 0, 53, 184, 1, 0, 0,
		0, 55, 186, 1, 0, 0, 0, 57, 189, 1, 0, 0, 0, 59, 192, 1, 0, 0, 0, 61, 195,
		1, 0, 0, 0, 63, 197, 1, 0, 0, 0, 65, 199, 1, 0, 0, 0, 67, 201, 1, 0, 0,
		0, 69, 203, 1, 0, 0, 0, 71, 205, 1, 0, 0, 0, 73, 208, 1, 0, 0, 0, 75, 218,
		1, 0, 0, 0, 77, 221, 1, 0, 0, 0, 79, 236, 1, 0, 0, 0, 81, 239, 1, 0, 0,
		0, 83, 243, 1, 0, 0, 0, 85, 246, 1, 0, 0, 0, 87, 88, 5, 115, 0, 0, 88,
		89, 5, 116, 0, 0, 89, 90, 5, 114, 0, 0, 90, 91, 5, 117, 0, 0, 91, 92, 5,
		99, 0, 0, 92, 93, 5, 116, 0, 0, 93, 2, 1, 0, 0, 0, 94, 95, 5, 105, 0, 0,
		95, 96, 5, 110, 0, 0, 96, 97, 5, 116, 0, 0, 97, 98, 5, 101, 0, 0, 98, 99,
		5, 114, 0, 0, 99, 100, 5, 102, 0, 0, 100, 101, 5, 97, 0, 0, 101, 102, 5,
		99, 0, 0, 102, 103, 5, 101, 0, 0, 103, 4, 1, 0, 0, 0, 104, 105, 5, 102,
		0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5, 110, 0, 0, 107, 108, 5, 99,
		0, 0, 108, 6, 1, 0, 0, 0, 109, 110, 5, 114, 0, 0, 110, 111, 5, 101, 0,
		0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 117, 0, 0, 113, 114, 5, 114, 0,
		0, 114, 115, 5, 110, 0, 0, 115, 8, 1, 0, 0, 0, 116, 117, 5, 108, 0, 0,
		117, 118, 5, 101, 0, 0, 118, 119, 5, 116, 0, 0, 119, 10, 1, 0, 0, 0, 120,
		121, 5, 97, 0, 0, 121, 122, 5, 115, 0, 0, 122, 123, 5, 121, 0, 0, 123,
		124, 5, 110, 0, 0, 124, 125, 5, 99, 0, 0, 125, 12, 1, 0, 0, 0, 126, 127,
		5, 97, 0, 0, 127, 128, 5, 119, 0, 0, 128, 129, 5, 97, 0, 0, 129, 130, 5,
		105, 0, 0, 130, 131, 5, 116, 0, 0, 131, 14, 1, 0, 0, 0, 132, 133, 5, 105,
		0, 0, 133, 134, 5, 102, 0, 0, 134, 16, 1, 0, 0, 0, 135, 136, 5, 101, 0,
		0, 136, 137, 5, 108, 0, 0, 137, 138, 5, 115, 0, 0, 138, 139, 5, 101, 0,
		0, 139, 18, 1, 0, 0, 0, 140, 141, 5, 116, 0, 0, 141, 142, 5, 114, 0, 0,
		142, 143, 5, 117, 0, 0, 143, 144, 5, 101, 0, 0, 144, 20, 1, 0, 0, 0, 145,
		146, 5, 102, 0, 0, 146, 147, 5, 97, 0, 0, 147, 148, 5, 108, 0, 0, 148,
		149, 5, 115, 0, 0, 149, 150, 5, 101, 0, 0, 150, 22, 1, 0, 0, 0, 151, 152,
		5, 40, 0, 0, 152, 24, 1, 0, 0, 0, 153, 154, 5, 91, 0, 0, 154, 26, 1, 0,
		0, 0, 155, 156, 5, 123, 0, 0, 156, 28, 1, 0, 0, 0, 157, 158, 5, 41, 0,
		0, 158, 30, 1, 0, 0, 0, 159, 160, 5, 93, 0, 0, 160, 32, 1, 0, 0, 0, 161,
		162, 5, 125, 0, 0, 162, 34, 1, 0, 0, 0, 163, 164, 5, 43, 0, 0, 164, 36,
		1, 0, 0, 0, 165, 166, 5, 45, 0, 0, 166, 38, 1, 0, 0, 0, 167, 168, 5, 42,
		0, 0, 168, 40, 1, 0, 0, 0, 169, 170, 5, 47, 0, 0, 170, 42, 1, 0, 0, 0,
		171, 172, 5, 37, 0, 0, 172, 44, 1, 0, 0, 0, 173, 174, 5, 61, 0, 0, 174,
		175, 5, 61, 0, 0, 175, 46, 1, 0, 0, 0, 176, 177, 5, 33, 0, 0, 177, 178,
		5, 61, 0, 0, 178, 48, 1, 0, 0, 0, 179, 180, 5, 60, 0, 0, 180, 50, 1, 0,
		0, 0, 181, 182, 5, 60, 0, 0, 182, 183, 5, 61, 0, 0, 183, 52, 1, 0, 0, 0,
		184, 185, 5, 62, 0, 0, 185, 54, 1, 0, 0, 0, 186, 187, 5, 62, 0, 0, 187,
		188, 5, 61, 0, 0, 188, 56, 1, 0, 0, 0, 189, 190, 5, 38, 0, 0, 190, 191,
		5, 38, 0, 0, 191, 58, 1, 0, 0, 0, 192, 193, 5, 124, 0, 0, 193, 194, 5,
		124, 0, 0, 194, 60, 1, 0, 0, 0, 195, 196, 5, 61, 0, 0, 196, 62, 1, 0, 0,
		0, 197, 198, 5, 64, 0, 0, 198, 64, 1, 0, 0, 0, 199, 200, 5, 44, 0, 0, 200,
		66, 1, 0, 0, 0, 201, 202, 5, 46, 0, 0, 202, 68, 1, 0, 0, 0, 203, 204, 5,
		58, 0, 0, 204, 70, 1, 0, 0, 0, 205, 206, 5, 45, 0, 0, 206, 207, 5, 62,
		0, 0, 207, 72, 1, 0, 0, 0, 208, 213, 5, 34, 0, 0, 209, 212, 3, 75, 37,
		0, 210, 212, 3, 77, 38, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1, 0, 0, 0,
		212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 34, 0, 0, 217, 74,
		1, 0, 0, 0, 218, 219, 5, 92, 0, 0, 219, 220, 7, 0, 0, 0, 220, 76, 1, 0,
		0, 0, 221, 222, 8, 1, 0, 0, 222, 78, 1, 0, 0, 0, 223, 225, 5, 95, 0, 0,
		224, 226, 7, 2, 0, 0, 225, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 237, 1, 0, 0, 0, 229, 233,
		7, 3, 0, 0, 230, 232, 7, 2, 0, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0,
		0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0,
		235, 233, 1, 0, 0, 0, 236, 223, 1, 0, 0, 0, 236, 229, 1, 0, 0, 0, 237,
		80, 1, 0, 0, 0, 238, 240, 7, 4, 0, 0, 239, 238, 1, 0, 0, 0, 240, 241, 1,
		0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 82, 1, 0, 0,
		0, 243, 244, 5, 59, 0, 0, 244, 84, 1, 0, 0, 0, 245, 247, 7, 5, 0, 0, 246,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 6, 42, 0, 0, 251, 86, 1, 0,
		0, 0, 8, 0, 211, 213, 227, 233, 236, 241, 248, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TempoLexerInit initializes any static state used to implement TempoLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTempoLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TempoLexerInit() {
	staticData := &TempoLexerLexerStaticData
	staticData.once.Do(tempolexerLexerInit)
}

// NewTempoLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTempoLexer(input antlr.CharStream) *TempoLexer {
	TempoLexerInit()
	l := new(TempoLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TempoLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Tempo.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TempoLexer tokens.
const (
	TempoLexerSTRUCT     = 1
	TempoLexerINTERFACE  = 2
	TempoLexerFUNC       = 3
	TempoLexerRETURN     = 4
	TempoLexerLET        = 5
	TempoLexerASYNC      = 6
	TempoLexerAWAIT      = 7
	TempoLexerIF         = 8
	TempoLexerELSE       = 9
	TempoLexerTRUE       = 10
	TempoLexerFALSE      = 11
	TempoLexerLPAREN     = 12
	TempoLexerLSQUARE    = 13
	TempoLexerLCURLY     = 14
	TempoLexerRPAREN     = 15
	TempoLexerRSQUARE    = 16
	TempoLexerRCURLY     = 17
	TempoLexerPLUS       = 18
	TempoLexerMINUS      = 19
	TempoLexerMULTIPLY   = 20
	TempoLexerDIVIDE     = 21
	TempoLexerMODULO     = 22
	TempoLexerEQUAL      = 23
	TempoLexerNOT_EQUAL  = 24
	TempoLexerLESS       = 25
	TempoLexerLESS_EQ    = 26
	TempoLexerGREATER    = 27
	TempoLexerGREATER_EQ = 28
	TempoLexerAND        = 29
	TempoLexerOR         = 30
	TempoLexerIS         = 31
	TempoLexerROLE_AT    = 32
	TempoLexerCOMMA      = 33
	TempoLexerDOT        = 34
	TempoLexerCOLON      = 35
	TempoLexerCOM        = 36
	TempoLexerSTRING     = 37
	TempoLexerID         = 38
	TempoLexerNUMBER     = 39
	TempoLexerEND        = 40
	TempoLexerWHITESPACE = 41
)
