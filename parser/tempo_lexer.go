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
		"NUMBER", "END", "WHITESPACE", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"STRUCT", "INTERFACE", "FUNC", "RETURN", "LET", "ASYNC", "AWAIT", "IF",
		"ELSE", "TRUE", "FALSE", "LPAREN", "LSQUARE", "LCURLY", "RPAREN", "RSQUARE",
		"RCURLY", "PLUS", "MINUS", "MULTIPLY", "DIVIDE", "MODULO", "EQUAL",
		"NOT_EQUAL", "LESS", "LESS_EQ", "GREATER", "GREATER_EQ", "AND", "OR",
		"IS", "ROLE_AT", "COMMA", "DOT", "COLON", "COM", "STRING", "ESC", "SAFECODEPOINT",
		"ID", "NUMBER", "END", "WHITESPACE", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 43, 281, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 5, 36, 216, 8,
		36, 10, 36, 12, 36, 219, 9, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 39, 1, 39, 4, 39, 230, 8, 39, 11, 39, 12, 39, 231, 1, 39, 1,
		39, 5, 39, 236, 8, 39, 10, 39, 12, 39, 239, 9, 39, 3, 39, 241, 8, 39, 1,
		40, 4, 40, 244, 8, 40, 11, 40, 12, 40, 245, 1, 41, 1, 41, 1, 42, 4, 42,
		251, 8, 42, 11, 42, 12, 42, 252, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1,
		43, 5, 43, 261, 8, 43, 10, 43, 12, 43, 264, 9, 43, 1, 43, 1, 43, 1, 44,
		1, 44, 1, 44, 1, 44, 5, 44, 272, 8, 44, 10, 44, 12, 44, 275, 9, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 273, 0, 45, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 0, 77, 0, 79, 38, 81, 39, 83,
		40, 85, 41, 87, 42, 89, 43, 1, 0, 7, 5, 0, 34, 34, 92, 92, 110, 110, 114,
		114, 116, 116, 3, 0, 0, 31, 34, 34, 92, 92, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 10, 10, 13, 13, 287, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0,
		0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 91, 1, 0, 0, 0, 3, 98, 1, 0,
		0, 0, 5, 108, 1, 0, 0, 0, 7, 113, 1, 0, 0, 0, 9, 120, 1, 0, 0, 0, 11, 124,
		1, 0, 0, 0, 13, 130, 1, 0, 0, 0, 15, 136, 1, 0, 0, 0, 17, 139, 1, 0, 0,
		0, 19, 144, 1, 0, 0, 0, 21, 149, 1, 0, 0, 0, 23, 155, 1, 0, 0, 0, 25, 157,
		1, 0, 0, 0, 27, 159, 1, 0, 0, 0, 29, 161, 1, 0, 0, 0, 31, 163, 1, 0, 0,
		0, 33, 165, 1, 0, 0, 0, 35, 167, 1, 0, 0, 0, 37, 169, 1, 0, 0, 0, 39, 171,
		1, 0, 0, 0, 41, 173, 1, 0, 0, 0, 43, 175, 1, 0, 0, 0, 45, 177, 1, 0, 0,
		0, 47, 180, 1, 0, 0, 0, 49, 183, 1, 0, 0, 0, 51, 185, 1, 0, 0, 0, 53, 188,
		1, 0, 0, 0, 55, 190, 1, 0, 0, 0, 57, 193, 1, 0, 0, 0, 59, 196, 1, 0, 0,
		0, 61, 199, 1, 0, 0, 0, 63, 201, 1, 0, 0, 0, 65, 203, 1, 0, 0, 0, 67, 205,
		1, 0, 0, 0, 69, 207, 1, 0, 0, 0, 71, 209, 1, 0, 0, 0, 73, 212, 1, 0, 0,
		0, 75, 222, 1, 0, 0, 0, 77, 225, 1, 0, 0, 0, 79, 240, 1, 0, 0, 0, 81, 243,
		1, 0, 0, 0, 83, 247, 1, 0, 0, 0, 85, 250, 1, 0, 0, 0, 87, 256, 1, 0, 0,
		0, 89, 267, 1, 0, 0, 0, 91, 92, 5, 115, 0, 0, 92, 93, 5, 116, 0, 0, 93,
		94, 5, 114, 0, 0, 94, 95, 5, 117, 0, 0, 95, 96, 5, 99, 0, 0, 96, 97, 5,
		116, 0, 0, 97, 2, 1, 0, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100, 5, 110, 0,
		0, 100, 101, 5, 116, 0, 0, 101, 102, 5, 101, 0, 0, 102, 103, 5, 114, 0,
		0, 103, 104, 5, 102, 0, 0, 104, 105, 5, 97, 0, 0, 105, 106, 5, 99, 0, 0,
		106, 107, 5, 101, 0, 0, 107, 4, 1, 0, 0, 0, 108, 109, 5, 102, 0, 0, 109,
		110, 5, 117, 0, 0, 110, 111, 5, 110, 0, 0, 111, 112, 5, 99, 0, 0, 112,
		6, 1, 0, 0, 0, 113, 114, 5, 114, 0, 0, 114, 115, 5, 101, 0, 0, 115, 116,
		5, 116, 0, 0, 116, 117, 5, 117, 0, 0, 117, 118, 5, 114, 0, 0, 118, 119,
		5, 110, 0, 0, 119, 8, 1, 0, 0, 0, 120, 121, 5, 108, 0, 0, 121, 122, 5,
		101, 0, 0, 122, 123, 5, 116, 0, 0, 123, 10, 1, 0, 0, 0, 124, 125, 5, 97,
		0, 0, 125, 126, 5, 115, 0, 0, 126, 127, 5, 121, 0, 0, 127, 128, 5, 110,
		0, 0, 128, 129, 5, 99, 0, 0, 129, 12, 1, 0, 0, 0, 130, 131, 5, 97, 0, 0,
		131, 132, 5, 119, 0, 0, 132, 133, 5, 97, 0, 0, 133, 134, 5, 105, 0, 0,
		134, 135, 5, 116, 0, 0, 135, 14, 1, 0, 0, 0, 136, 137, 5, 105, 0, 0, 137,
		138, 5, 102, 0, 0, 138, 16, 1, 0, 0, 0, 139, 140, 5, 101, 0, 0, 140, 141,
		5, 108, 0, 0, 141, 142, 5, 115, 0, 0, 142, 143, 5, 101, 0, 0, 143, 18,
		1, 0, 0, 0, 144, 145, 5, 116, 0, 0, 145, 146, 5, 114, 0, 0, 146, 147, 5,
		117, 0, 0, 147, 148, 5, 101, 0, 0, 148, 20, 1, 0, 0, 0, 149, 150, 5, 102,
		0, 0, 150, 151, 5, 97, 0, 0, 151, 152, 5, 108, 0, 0, 152, 153, 5, 115,
		0, 0, 153, 154, 5, 101, 0, 0, 154, 22, 1, 0, 0, 0, 155, 156, 5, 40, 0,
		0, 156, 24, 1, 0, 0, 0, 157, 158, 5, 91, 0, 0, 158, 26, 1, 0, 0, 0, 159,
		160, 5, 123, 0, 0, 160, 28, 1, 0, 0, 0, 161, 162, 5, 41, 0, 0, 162, 30,
		1, 0, 0, 0, 163, 164, 5, 93, 0, 0, 164, 32, 1, 0, 0, 0, 165, 166, 5, 125,
		0, 0, 166, 34, 1, 0, 0, 0, 167, 168, 5, 43, 0, 0, 168, 36, 1, 0, 0, 0,
		169, 170, 5, 45, 0, 0, 170, 38, 1, 0, 0, 0, 171, 172, 5, 42, 0, 0, 172,
		40, 1, 0, 0, 0, 173, 174, 5, 47, 0, 0, 174, 42, 1, 0, 0, 0, 175, 176, 5,
		37, 0, 0, 176, 44, 1, 0, 0, 0, 177, 178, 5, 61, 0, 0, 178, 179, 5, 61,
		0, 0, 179, 46, 1, 0, 0, 0, 180, 181, 5, 33, 0, 0, 181, 182, 5, 61, 0, 0,
		182, 48, 1, 0, 0, 0, 183, 184, 5, 60, 0, 0, 184, 50, 1, 0, 0, 0, 185, 186,
		5, 60, 0, 0, 186, 187, 5, 61, 0, 0, 187, 52, 1, 0, 0, 0, 188, 189, 5, 62,
		0, 0, 189, 54, 1, 0, 0, 0, 190, 191, 5, 62, 0, 0, 191, 192, 5, 61, 0, 0,
		192, 56, 1, 0, 0, 0, 193, 194, 5, 38, 0, 0, 194, 195, 5, 38, 0, 0, 195,
		58, 1, 0, 0, 0, 196, 197, 5, 124, 0, 0, 197, 198, 5, 124, 0, 0, 198, 60,
		1, 0, 0, 0, 199, 200, 5, 61, 0, 0, 200, 62, 1, 0, 0, 0, 201, 202, 5, 64,
		0, 0, 202, 64, 1, 0, 0, 0, 203, 204, 5, 44, 0, 0, 204, 66, 1, 0, 0, 0,
		205, 206, 5, 46, 0, 0, 206, 68, 1, 0, 0, 0, 207, 208, 5, 58, 0, 0, 208,
		70, 1, 0, 0, 0, 209, 210, 5, 45, 0, 0, 210, 211, 5, 62, 0, 0, 211, 72,
		1, 0, 0, 0, 212, 217, 5, 34, 0, 0, 213, 216, 3, 75, 37, 0, 214, 216, 3,
		77, 38, 0, 215, 213, 1, 0, 0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0,
		0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 220, 1, 0, 0, 0,
		219, 217, 1, 0, 0, 0, 220, 221, 5, 34, 0, 0, 221, 74, 1, 0, 0, 0, 222,
		223, 5, 92, 0, 0, 223, 224, 7, 0, 0, 0, 224, 76, 1, 0, 0, 0, 225, 226,
		8, 1, 0, 0, 226, 78, 1, 0, 0, 0, 227, 229, 5, 95, 0, 0, 228, 230, 7, 2,
		0, 0, 229, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0,
		231, 232, 1, 0, 0, 0, 232, 241, 1, 0, 0, 0, 233, 237, 7, 3, 0, 0, 234,
		236, 7, 2, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235,
		1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0,
		0, 0, 240, 227, 1, 0, 0, 0, 240, 233, 1, 0, 0, 0, 241, 80, 1, 0, 0, 0,
		242, 244, 7, 4, 0, 0, 243, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245,
		243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 82, 1, 0, 0, 0, 247, 248, 5,
		59, 0, 0, 248, 84, 1, 0, 0, 0, 249, 251, 7, 5, 0, 0, 250, 249, 1, 0, 0,
		0, 251, 252, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253,
		254, 1, 0, 0, 0, 254, 255, 6, 42, 0, 0, 255, 86, 1, 0, 0, 0, 256, 257,
		5, 47, 0, 0, 257, 258, 5, 47, 0, 0, 258, 262, 1, 0, 0, 0, 259, 261, 8,
		6, 0, 0, 260, 259, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 260, 1, 0, 0,
		0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265,
		266, 6, 43, 0, 0, 266, 88, 1, 0, 0, 0, 267, 268, 5, 47, 0, 0, 268, 269,
		5, 42, 0, 0, 269, 273, 1, 0, 0, 0, 270, 272, 9, 0, 0, 0, 271, 270, 1, 0,
		0, 0, 272, 275, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0,
		274, 276, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 277, 5, 42, 0, 0, 277,
		278, 5, 47, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 6, 44, 0, 0, 280, 90,
		1, 0, 0, 0, 10, 0, 215, 217, 231, 237, 240, 245, 252, 262, 273, 1, 6, 0,
		0,
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
	TempoLexerSTRUCT        = 1
	TempoLexerINTERFACE     = 2
	TempoLexerFUNC          = 3
	TempoLexerRETURN        = 4
	TempoLexerLET           = 5
	TempoLexerASYNC         = 6
	TempoLexerAWAIT         = 7
	TempoLexerIF            = 8
	TempoLexerELSE          = 9
	TempoLexerTRUE          = 10
	TempoLexerFALSE         = 11
	TempoLexerLPAREN        = 12
	TempoLexerLSQUARE       = 13
	TempoLexerLCURLY        = 14
	TempoLexerRPAREN        = 15
	TempoLexerRSQUARE       = 16
	TempoLexerRCURLY        = 17
	TempoLexerPLUS          = 18
	TempoLexerMINUS         = 19
	TempoLexerMULTIPLY      = 20
	TempoLexerDIVIDE        = 21
	TempoLexerMODULO        = 22
	TempoLexerEQUAL         = 23
	TempoLexerNOT_EQUAL     = 24
	TempoLexerLESS          = 25
	TempoLexerLESS_EQ       = 26
	TempoLexerGREATER       = 27
	TempoLexerGREATER_EQ    = 28
	TempoLexerAND           = 29
	TempoLexerOR            = 30
	TempoLexerIS            = 31
	TempoLexerROLE_AT       = 32
	TempoLexerCOMMA         = 33
	TempoLexerDOT           = 34
	TempoLexerCOLON         = 35
	TempoLexerCOM           = 36
	TempoLexerSTRING        = 37
	TempoLexerID            = 38
	TempoLexerNUMBER        = 39
	TempoLexerEND           = 40
	TempoLexerWHITESPACE    = 41
	TempoLexerLINE_COMMENT  = 42
	TempoLexerBLOCK_COMMENT = 43
)
