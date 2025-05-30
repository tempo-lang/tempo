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
		"':'", "'->'", "", "", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "STRUCT", "INTERFACE", "FUNC", "RETURN", "LET", "ASYNC", "AWAIT",
		"IF", "ELSE", "TRUE", "FALSE", "LPAREN", "LSQUARE", "LCURLY", "RPAREN",
		"RSQUARE", "RCURLY", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO",
		"EQUAL", "NOT_EQUAL", "LESS", "LESS_EQ", "GREATER", "GREATER_EQ", "AND",
		"OR", "IS", "ROLE_AT", "COMMA", "DOT", "COLON", "COM", "ID", "NUMBER",
		"END", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"STRUCT", "INTERFACE", "FUNC", "RETURN", "LET", "ASYNC", "AWAIT", "IF",
		"ELSE", "TRUE", "FALSE", "LPAREN", "LSQUARE", "LCURLY", "RPAREN", "RSQUARE",
		"RCURLY", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "EQUAL",
		"NOT_EQUAL", "LESS", "LESS_EQ", "GREATER", "GREATER_EQ", "AND", "OR",
		"IS", "ROLE_AT", "COMMA", "DOT", "COLON", "COM", "ID", "NUMBER", "END",
		"WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 40, 231, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 4, 36, 205, 8, 36,
		11, 36, 12, 36, 206, 1, 36, 1, 36, 5, 36, 211, 8, 36, 10, 36, 12, 36, 214,
		9, 36, 3, 36, 216, 8, 36, 1, 37, 4, 37, 219, 8, 37, 11, 37, 12, 37, 220,
		1, 38, 1, 38, 1, 39, 4, 39, 226, 8, 39, 11, 39, 12, 39, 227, 1, 39, 1,
		39, 0, 0, 40, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27,
		55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36,
		73, 37, 75, 38, 77, 39, 79, 40, 1, 0, 4, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32,
		32, 235, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 81, 1, 0, 0, 0, 3, 88, 1, 0, 0,
		0, 5, 98, 1, 0, 0, 0, 7, 103, 1, 0, 0, 0, 9, 110, 1, 0, 0, 0, 11, 114,
		1, 0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 126, 1, 0, 0, 0, 17, 129, 1, 0, 0,
		0, 19, 134, 1, 0, 0, 0, 21, 139, 1, 0, 0, 0, 23, 145, 1, 0, 0, 0, 25, 147,
		1, 0, 0, 0, 27, 149, 1, 0, 0, 0, 29, 151, 1, 0, 0, 0, 31, 153, 1, 0, 0,
		0, 33, 155, 1, 0, 0, 0, 35, 157, 1, 0, 0, 0, 37, 159, 1, 0, 0, 0, 39, 161,
		1, 0, 0, 0, 41, 163, 1, 0, 0, 0, 43, 165, 1, 0, 0, 0, 45, 167, 1, 0, 0,
		0, 47, 170, 1, 0, 0, 0, 49, 173, 1, 0, 0, 0, 51, 175, 1, 0, 0, 0, 53, 178,
		1, 0, 0, 0, 55, 180, 1, 0, 0, 0, 57, 183, 1, 0, 0, 0, 59, 186, 1, 0, 0,
		0, 61, 189, 1, 0, 0, 0, 63, 191, 1, 0, 0, 0, 65, 193, 1, 0, 0, 0, 67, 195,
		1, 0, 0, 0, 69, 197, 1, 0, 0, 0, 71, 199, 1, 0, 0, 0, 73, 215, 1, 0, 0,
		0, 75, 218, 1, 0, 0, 0, 77, 222, 1, 0, 0, 0, 79, 225, 1, 0, 0, 0, 81, 82,
		5, 115, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5, 114, 0, 0, 84, 85, 5, 117,
		0, 0, 85, 86, 5, 99, 0, 0, 86, 87, 5, 116, 0, 0, 87, 2, 1, 0, 0, 0, 88,
		89, 5, 105, 0, 0, 89, 90, 5, 110, 0, 0, 90, 91, 5, 116, 0, 0, 91, 92, 5,
		101, 0, 0, 92, 93, 5, 114, 0, 0, 93, 94, 5, 102, 0, 0, 94, 95, 5, 97, 0,
		0, 95, 96, 5, 99, 0, 0, 96, 97, 5, 101, 0, 0, 97, 4, 1, 0, 0, 0, 98, 99,
		5, 102, 0, 0, 99, 100, 5, 117, 0, 0, 100, 101, 5, 110, 0, 0, 101, 102,
		5, 99, 0, 0, 102, 6, 1, 0, 0, 0, 103, 104, 5, 114, 0, 0, 104, 105, 5, 101,
		0, 0, 105, 106, 5, 116, 0, 0, 106, 107, 5, 117, 0, 0, 107, 108, 5, 114,
		0, 0, 108, 109, 5, 110, 0, 0, 109, 8, 1, 0, 0, 0, 110, 111, 5, 108, 0,
		0, 111, 112, 5, 101, 0, 0, 112, 113, 5, 116, 0, 0, 113, 10, 1, 0, 0, 0,
		114, 115, 5, 97, 0, 0, 115, 116, 5, 115, 0, 0, 116, 117, 5, 121, 0, 0,
		117, 118, 5, 110, 0, 0, 118, 119, 5, 99, 0, 0, 119, 12, 1, 0, 0, 0, 120,
		121, 5, 97, 0, 0, 121, 122, 5, 119, 0, 0, 122, 123, 5, 97, 0, 0, 123, 124,
		5, 105, 0, 0, 124, 125, 5, 116, 0, 0, 125, 14, 1, 0, 0, 0, 126, 127, 5,
		105, 0, 0, 127, 128, 5, 102, 0, 0, 128, 16, 1, 0, 0, 0, 129, 130, 5, 101,
		0, 0, 130, 131, 5, 108, 0, 0, 131, 132, 5, 115, 0, 0, 132, 133, 5, 101,
		0, 0, 133, 18, 1, 0, 0, 0, 134, 135, 5, 116, 0, 0, 135, 136, 5, 114, 0,
		0, 136, 137, 5, 117, 0, 0, 137, 138, 5, 101, 0, 0, 138, 20, 1, 0, 0, 0,
		139, 140, 5, 102, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 108, 0, 0,
		142, 143, 5, 115, 0, 0, 143, 144, 5, 101, 0, 0, 144, 22, 1, 0, 0, 0, 145,
		146, 5, 40, 0, 0, 146, 24, 1, 0, 0, 0, 147, 148, 5, 91, 0, 0, 148, 26,
		1, 0, 0, 0, 149, 150, 5, 123, 0, 0, 150, 28, 1, 0, 0, 0, 151, 152, 5, 41,
		0, 0, 152, 30, 1, 0, 0, 0, 153, 154, 5, 93, 0, 0, 154, 32, 1, 0, 0, 0,
		155, 156, 5, 125, 0, 0, 156, 34, 1, 0, 0, 0, 157, 158, 5, 43, 0, 0, 158,
		36, 1, 0, 0, 0, 159, 160, 5, 45, 0, 0, 160, 38, 1, 0, 0, 0, 161, 162, 5,
		42, 0, 0, 162, 40, 1, 0, 0, 0, 163, 164, 5, 47, 0, 0, 164, 42, 1, 0, 0,
		0, 165, 166, 5, 37, 0, 0, 166, 44, 1, 0, 0, 0, 167, 168, 5, 61, 0, 0, 168,
		169, 5, 61, 0, 0, 169, 46, 1, 0, 0, 0, 170, 171, 5, 33, 0, 0, 171, 172,
		5, 61, 0, 0, 172, 48, 1, 0, 0, 0, 173, 174, 5, 60, 0, 0, 174, 50, 1, 0,
		0, 0, 175, 176, 5, 60, 0, 0, 176, 177, 5, 61, 0, 0, 177, 52, 1, 0, 0, 0,
		178, 179, 5, 62, 0, 0, 179, 54, 1, 0, 0, 0, 180, 181, 5, 62, 0, 0, 181,
		182, 5, 61, 0, 0, 182, 56, 1, 0, 0, 0, 183, 184, 5, 38, 0, 0, 184, 185,
		5, 38, 0, 0, 185, 58, 1, 0, 0, 0, 186, 187, 5, 124, 0, 0, 187, 188, 5,
		124, 0, 0, 188, 60, 1, 0, 0, 0, 189, 190, 5, 61, 0, 0, 190, 62, 1, 0, 0,
		0, 191, 192, 5, 64, 0, 0, 192, 64, 1, 0, 0, 0, 193, 194, 5, 44, 0, 0, 194,
		66, 1, 0, 0, 0, 195, 196, 5, 46, 0, 0, 196, 68, 1, 0, 0, 0, 197, 198, 5,
		58, 0, 0, 198, 70, 1, 0, 0, 0, 199, 200, 5, 45, 0, 0, 200, 201, 5, 62,
		0, 0, 201, 72, 1, 0, 0, 0, 202, 204, 5, 95, 0, 0, 203, 205, 7, 0, 0, 0,
		204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 216, 1, 0, 0, 0, 208, 212, 7, 1, 0, 0, 209, 211,
		7, 0, 0, 0, 210, 209, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0,
		0, 0, 212, 213, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0,
		215, 202, 1, 0, 0, 0, 215, 208, 1, 0, 0, 0, 216, 74, 1, 0, 0, 0, 217, 219,
		7, 2, 0, 0, 218, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0,
		0, 0, 220, 221, 1, 0, 0, 0, 221, 76, 1, 0, 0, 0, 222, 223, 5, 59, 0, 0,
		223, 78, 1, 0, 0, 0, 224, 226, 7, 3, 0, 0, 225, 224, 1, 0, 0, 0, 226, 227,
		1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 1, 0,
		0, 0, 229, 230, 6, 39, 0, 0, 230, 80, 1, 0, 0, 0, 6, 0, 206, 212, 215,
		220, 227, 1, 6, 0, 0,
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
	TempoLexerID         = 37
	TempoLexerNUMBER     = 38
	TempoLexerEND        = 39
	TempoLexerWHITESPACE = 40
)
