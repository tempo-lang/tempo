// Code generated from Tempo.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Tempo
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TempoParser struct {
	*antlr.BaseParser
}

var TempoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tempoParserInit() {
	staticData := &TempoParserStaticData
	staticData.LiteralNames = []string{
		"", "'struct'", "'interface'", "'func'", "'return'", "'let'", "'async'",
		"'await'", "'if'", "'else'", "'while'", "'true'", "'false'", "'('",
		"'['", "'{'", "')'", "']'", "'}'", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'='",
		"'@'", "','", "'.'", "':'", "'->'", "", "", "", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "STRUCT", "INTERFACE", "FUNC", "RETURN", "LET", "ASYNC", "AWAIT",
		"IF", "ELSE", "WHILE", "TRUE", "FALSE", "LPAREN", "LSQUARE", "LCURLY",
		"RPAREN", "RSQUARE", "RCURLY", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "EQUAL", "NOT_EQUAL", "LESS", "LESS_EQ", "GREATER", "GREATER_EQ",
		"AND", "OR", "IS", "ROLE_AT", "COMMA", "DOT", "COLON", "COM", "STRING",
		"ID", "NUMBER", "END", "WHITESPACE", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"sourceFile", "ident", "valueType", "roleType", "closureParamList",
		"closureSig", "struct", "structFieldList", "structField", "interface",
		"interfaceMethodsList", "interfaceMethod", "func", "funcSig", "funcParamList",
		"funcParam", "funcArgList", "scope", "stmt", "expr", "exprStructField",
		"identAccess", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 330, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 5, 0, 50, 8, 0, 10, 0, 12, 0,
		53, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		76, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 82, 8, 3, 10, 3, 12, 3, 85, 9,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 94, 8, 3, 10, 3, 12,
		3, 97, 9, 3, 1, 3, 1, 3, 3, 3, 101, 8, 3, 3, 3, 103, 8, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 5, 4, 109, 8, 4, 10, 4, 12, 4, 112, 9, 4, 3, 4, 114, 8, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 123, 8, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 135, 8, 7, 10, 7,
		12, 7, 138, 9, 7, 3, 7, 140, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 156, 8, 10, 10,
		10, 12, 10, 159, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 175, 8, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 5, 14, 181, 8, 14, 10, 14, 12, 14, 184, 9, 14,
		3, 14, 186, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 198, 8, 16, 10, 16, 12, 16, 201, 9, 16, 3, 16,
		203, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 209, 8, 17, 10, 17, 12,
		17, 212, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 220, 8,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18,
		231, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 239, 8, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 250,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 256, 8, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 278, 8, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 288, 8, 19,
		10, 19, 12, 19, 291, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 5, 20, 302, 8, 20, 10, 20, 12, 20, 305, 9, 20, 3, 20,
		307, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 3, 21, 314, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 323, 8, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 328, 8, 22, 1, 22, 0, 1, 38, 23, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 2, 1,
		0, 19, 31, 1, 0, 11, 12, 355, 0, 51, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0, 4,
		75, 1, 0, 0, 0, 6, 102, 1, 0, 0, 0, 8, 104, 1, 0, 0, 0, 10, 117, 1, 0,
		0, 0, 12, 124, 1, 0, 0, 0, 14, 130, 1, 0, 0, 0, 16, 143, 1, 0, 0, 0, 18,
		147, 1, 0, 0, 0, 20, 153, 1, 0, 0, 0, 22, 162, 1, 0, 0, 0, 24, 165, 1,
		0, 0, 0, 26, 168, 1, 0, 0, 0, 28, 176, 1, 0, 0, 0, 30, 189, 1, 0, 0, 0,
		32, 193, 1, 0, 0, 0, 34, 206, 1, 0, 0, 0, 36, 249, 1, 0, 0, 0, 38, 277,
		1, 0, 0, 0, 40, 292, 1, 0, 0, 0, 42, 310, 1, 0, 0, 0, 44, 327, 1, 0, 0,
		0, 46, 50, 3, 24, 12, 0, 47, 50, 3, 12, 6, 0, 48, 50, 3, 18, 9, 0, 49,
		46, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0,
		0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51,
		1, 0, 0, 0, 54, 55, 5, 0, 0, 1, 55, 1, 1, 0, 0, 0, 56, 57, 5, 39, 0, 0,
		57, 3, 1, 0, 0, 0, 58, 59, 5, 6, 0, 0, 59, 76, 3, 4, 2, 0, 60, 61, 5, 14,
		0, 0, 61, 62, 3, 4, 2, 0, 62, 63, 5, 17, 0, 0, 63, 76, 1, 0, 0, 0, 64,
		65, 5, 3, 0, 0, 65, 66, 5, 33, 0, 0, 66, 67, 3, 6, 3, 0, 67, 69, 3, 8,
		4, 0, 68, 70, 3, 4, 2, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 76,
		1, 0, 0, 0, 71, 72, 3, 2, 1, 0, 72, 73, 5, 33, 0, 0, 73, 74, 3, 6, 3, 0,
		74, 76, 1, 0, 0, 0, 75, 58, 1, 0, 0, 0, 75, 60, 1, 0, 0, 0, 75, 64, 1,
		0, 0, 0, 75, 71, 1, 0, 0, 0, 76, 5, 1, 0, 0, 0, 77, 78, 5, 14, 0, 0, 78,
		83, 3, 2, 1, 0, 79, 80, 5, 34, 0, 0, 80, 82, 3, 2, 1, 0, 81, 79, 1, 0,
		0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86,
		1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 87, 5, 17, 0, 0, 87, 103, 1, 0, 0,
		0, 88, 101, 3, 2, 1, 0, 89, 90, 5, 13, 0, 0, 90, 95, 3, 2, 1, 0, 91, 92,
		5, 34, 0, 0, 92, 94, 3, 2, 1, 0, 93, 91, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0,
		95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 95, 1,
		0, 0, 0, 98, 99, 5, 16, 0, 0, 99, 101, 1, 0, 0, 0, 100, 88, 1, 0, 0, 0,
		100, 89, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 77, 1, 0, 0, 0, 102, 100,
		1, 0, 0, 0, 103, 7, 1, 0, 0, 0, 104, 113, 5, 13, 0, 0, 105, 110, 3, 4,
		2, 0, 106, 107, 5, 34, 0, 0, 107, 109, 3, 4, 2, 0, 108, 106, 1, 0, 0, 0,
		109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111,
		114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 105, 1, 0, 0, 0, 113, 114,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 5, 16, 0, 0, 116, 9, 1, 0,
		0, 0, 117, 118, 5, 3, 0, 0, 118, 119, 5, 33, 0, 0, 119, 120, 3, 6, 3, 0,
		120, 122, 3, 28, 14, 0, 121, 123, 3, 4, 2, 0, 122, 121, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 11, 1, 0, 0, 0, 124, 125, 5, 1, 0, 0, 125, 126, 5,
		33, 0, 0, 126, 127, 3, 6, 3, 0, 127, 128, 3, 2, 1, 0, 128, 129, 3, 14,
		7, 0, 129, 13, 1, 0, 0, 0, 130, 139, 5, 15, 0, 0, 131, 136, 3, 16, 8, 0,
		132, 133, 5, 34, 0, 0, 133, 135, 3, 16, 8, 0, 134, 132, 1, 0, 0, 0, 135,
		138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 140,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 131, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 5, 18, 0, 0, 142, 15, 1, 0, 0, 0,
		143, 144, 3, 2, 1, 0, 144, 145, 5, 36, 0, 0, 145, 146, 3, 4, 2, 0, 146,
		17, 1, 0, 0, 0, 147, 148, 5, 2, 0, 0, 148, 149, 5, 33, 0, 0, 149, 150,
		3, 6, 3, 0, 150, 151, 3, 2, 1, 0, 151, 152, 3, 20, 10, 0, 152, 19, 1, 0,
		0, 0, 153, 157, 5, 15, 0, 0, 154, 156, 3, 22, 11, 0, 155, 154, 1, 0, 0,
		0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 18, 0, 0, 161, 21,
		1, 0, 0, 0, 162, 163, 3, 26, 13, 0, 163, 164, 5, 41, 0, 0, 164, 23, 1,
		0, 0, 0, 165, 166, 3, 26, 13, 0, 166, 167, 3, 34, 17, 0, 167, 25, 1, 0,
		0, 0, 168, 169, 5, 3, 0, 0, 169, 170, 5, 33, 0, 0, 170, 171, 3, 6, 3, 0,
		171, 172, 3, 2, 1, 0, 172, 174, 3, 28, 14, 0, 173, 175, 3, 4, 2, 0, 174,
		173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 27, 1, 0, 0, 0, 176, 185, 5,
		13, 0, 0, 177, 182, 3, 30, 15, 0, 178, 179, 5, 34, 0, 0, 179, 181, 3, 30,
		15, 0, 180, 178, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0,
		182, 183, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185,
		177, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188,
		5, 16, 0, 0, 188, 29, 1, 0, 0, 0, 189, 190, 3, 2, 1, 0, 190, 191, 5, 36,
		0, 0, 191, 192, 3, 4, 2, 0, 192, 31, 1, 0, 0, 0, 193, 202, 5, 13, 0, 0,
		194, 199, 3, 38, 19, 0, 195, 196, 5, 34, 0, 0, 196, 198, 3, 38, 19, 0,
		197, 195, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 194,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 5, 16,
		0, 0, 205, 33, 1, 0, 0, 0, 206, 210, 5, 15, 0, 0, 207, 209, 3, 36, 18,
		0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210,
		211, 1, 0, 0, 0, 211, 213, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214,
		5, 18, 0, 0, 214, 35, 1, 0, 0, 0, 215, 216, 5, 5, 0, 0, 216, 219, 3, 2,
		1, 0, 217, 218, 5, 36, 0, 0, 218, 220, 3, 4, 2, 0, 219, 217, 1, 0, 0, 0,
		219, 220, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 222, 5, 32, 0, 0, 222,
		223, 3, 38, 19, 0, 223, 224, 5, 41, 0, 0, 224, 250, 1, 0, 0, 0, 225, 226,
		5, 8, 0, 0, 226, 227, 3, 38, 19, 0, 227, 230, 3, 34, 17, 0, 228, 229, 5,
		9, 0, 0, 229, 231, 3, 34, 17, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0,
		0, 0, 231, 250, 1, 0, 0, 0, 232, 233, 5, 10, 0, 0, 233, 234, 3, 38, 19,
		0, 234, 235, 3, 34, 17, 0, 235, 250, 1, 0, 0, 0, 236, 238, 5, 4, 0, 0,
		237, 239, 3, 38, 19, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		240, 1, 0, 0, 0, 240, 250, 5, 41, 0, 0, 241, 242, 3, 2, 1, 0, 242, 243,
		5, 32, 0, 0, 243, 244, 3, 38, 19, 0, 244, 245, 5, 41, 0, 0, 245, 250, 1,
		0, 0, 0, 246, 247, 3, 38, 19, 0, 247, 248, 5, 41, 0, 0, 248, 250, 1, 0,
		0, 0, 249, 215, 1, 0, 0, 0, 249, 225, 1, 0, 0, 0, 249, 232, 1, 0, 0, 0,
		249, 236, 1, 0, 0, 0, 249, 241, 1, 0, 0, 0, 249, 246, 1, 0, 0, 0, 250,
		37, 1, 0, 0, 0, 251, 252, 6, 19, -1, 0, 252, 255, 3, 44, 22, 0, 253, 254,
		5, 33, 0, 0, 254, 256, 3, 6, 3, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0,
		0, 0, 256, 278, 1, 0, 0, 0, 257, 258, 5, 7, 0, 0, 258, 278, 3, 38, 19,
		8, 259, 260, 3, 10, 5, 0, 260, 261, 3, 34, 17, 0, 261, 278, 1, 0, 0, 0,
		262, 263, 3, 2, 1, 0, 263, 264, 5, 33, 0, 0, 264, 265, 3, 6, 3, 0, 265,
		266, 3, 40, 20, 0, 266, 278, 1, 0, 0, 0, 267, 278, 3, 42, 21, 0, 268, 269,
		3, 6, 3, 0, 269, 270, 5, 37, 0, 0, 270, 271, 3, 6, 3, 0, 271, 272, 3, 38,
		19, 2, 272, 278, 1, 0, 0, 0, 273, 274, 5, 13, 0, 0, 274, 275, 3, 38, 19,
		0, 275, 276, 5, 16, 0, 0, 276, 278, 1, 0, 0, 0, 277, 251, 1, 0, 0, 0, 277,
		257, 1, 0, 0, 0, 277, 259, 1, 0, 0, 0, 277, 262, 1, 0, 0, 0, 277, 267,
		1, 0, 0, 0, 277, 268, 1, 0, 0, 0, 277, 273, 1, 0, 0, 0, 278, 289, 1, 0,
		0, 0, 279, 280, 10, 10, 0, 0, 280, 281, 7, 0, 0, 0, 281, 288, 3, 38, 19,
		11, 282, 283, 10, 5, 0, 0, 283, 288, 3, 32, 16, 0, 284, 285, 10, 4, 0,
		0, 285, 286, 5, 35, 0, 0, 286, 288, 3, 2, 1, 0, 287, 279, 1, 0, 0, 0, 287,
		282, 1, 0, 0, 0, 287, 284, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287,
		1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 39, 1, 0, 0, 0, 291, 289, 1, 0,
		0, 0, 292, 306, 5, 15, 0, 0, 293, 294, 3, 2, 1, 0, 294, 295, 5, 36, 0,
		0, 295, 303, 3, 38, 19, 0, 296, 297, 5, 34, 0, 0, 297, 298, 3, 2, 1, 0,
		298, 299, 5, 36, 0, 0, 299, 300, 3, 38, 19, 0, 300, 302, 1, 0, 0, 0, 301,
		296, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304,
		1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 293, 1, 0,
		0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 5, 18, 0, 0,
		309, 41, 1, 0, 0, 0, 310, 313, 3, 2, 1, 0, 311, 312, 5, 33, 0, 0, 312,
		314, 3, 6, 3, 0, 313, 311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 43, 1,
		0, 0, 0, 315, 316, 5, 40, 0, 0, 316, 317, 5, 35, 0, 0, 317, 323, 5, 40,
		0, 0, 318, 319, 5, 40, 0, 0, 319, 323, 5, 35, 0, 0, 320, 321, 5, 35, 0,
		0, 321, 323, 5, 40, 0, 0, 322, 315, 1, 0, 0, 0, 322, 318, 1, 0, 0, 0, 322,
		320, 1, 0, 0, 0, 323, 328, 1, 0, 0, 0, 324, 328, 5, 40, 0, 0, 325, 328,
		5, 38, 0, 0, 326, 328, 7, 1, 0, 0, 327, 322, 1, 0, 0, 0, 327, 324, 1, 0,
		0, 0, 327, 325, 1, 0, 0, 0, 327, 326, 1, 0, 0, 0, 328, 45, 1, 0, 0, 0,
		33, 49, 51, 69, 75, 83, 95, 100, 102, 110, 113, 122, 136, 139, 157, 174,
		182, 185, 199, 202, 210, 219, 230, 238, 249, 255, 277, 287, 289, 303, 306,
		313, 322, 327,
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

// TempoParserInit initializes any static state used to implement TempoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTempoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TempoParserInit() {
	staticData := &TempoParserStaticData
	staticData.once.Do(tempoParserInit)
}

// NewTempoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTempoParser(input antlr.TokenStream) *TempoParser {
	TempoParserInit()
	this := new(TempoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TempoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tempo.g4"

	return this
}

// TempoParser tokens.
const (
	TempoParserEOF           = antlr.TokenEOF
	TempoParserSTRUCT        = 1
	TempoParserINTERFACE     = 2
	TempoParserFUNC          = 3
	TempoParserRETURN        = 4
	TempoParserLET           = 5
	TempoParserASYNC         = 6
	TempoParserAWAIT         = 7
	TempoParserIF            = 8
	TempoParserELSE          = 9
	TempoParserWHILE         = 10
	TempoParserTRUE          = 11
	TempoParserFALSE         = 12
	TempoParserLPAREN        = 13
	TempoParserLSQUARE       = 14
	TempoParserLCURLY        = 15
	TempoParserRPAREN        = 16
	TempoParserRSQUARE       = 17
	TempoParserRCURLY        = 18
	TempoParserPLUS          = 19
	TempoParserMINUS         = 20
	TempoParserMULTIPLY      = 21
	TempoParserDIVIDE        = 22
	TempoParserMODULO        = 23
	TempoParserEQUAL         = 24
	TempoParserNOT_EQUAL     = 25
	TempoParserLESS          = 26
	TempoParserLESS_EQ       = 27
	TempoParserGREATER       = 28
	TempoParserGREATER_EQ    = 29
	TempoParserAND           = 30
	TempoParserOR            = 31
	TempoParserIS            = 32
	TempoParserROLE_AT       = 33
	TempoParserCOMMA         = 34
	TempoParserDOT           = 35
	TempoParserCOLON         = 36
	TempoParserCOM           = 37
	TempoParserSTRING        = 38
	TempoParserID            = 39
	TempoParserNUMBER        = 40
	TempoParserEND           = 41
	TempoParserWHITESPACE    = 42
	TempoParserLINE_COMMENT  = 43
	TempoParserBLOCK_COMMENT = 44
)

// TempoParser rules.
const (
	TempoParserRULE_sourceFile           = 0
	TempoParserRULE_ident                = 1
	TempoParserRULE_valueType            = 2
	TempoParserRULE_roleType             = 3
	TempoParserRULE_closureParamList     = 4
	TempoParserRULE_closureSig           = 5
	TempoParserRULE_struct               = 6
	TempoParserRULE_structFieldList      = 7
	TempoParserRULE_structField          = 8
	TempoParserRULE_interface            = 9
	TempoParserRULE_interfaceMethodsList = 10
	TempoParserRULE_interfaceMethod      = 11
	TempoParserRULE_func                 = 12
	TempoParserRULE_funcSig              = 13
	TempoParserRULE_funcParamList        = 14
	TempoParserRULE_funcParam            = 15
	TempoParserRULE_funcArgList          = 16
	TempoParserRULE_scope                = 17
	TempoParserRULE_stmt                 = 18
	TempoParserRULE_expr                 = 19
	TempoParserRULE_exprStructField      = 20
	TempoParserRULE_identAccess          = 21
	TempoParserRULE_literal              = 22
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFunc_() []IFuncContext
	Func_(i int) IFuncContext
	AllStruct_() []IStructContext
	Struct_(i int) IStructContext
	AllInterface_() []IInterfaceContext
	Interface_(i int) IInterfaceContext

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_sourceFile
	return p
}

func InitEmptySourceFileContext(p *SourceFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_sourceFile
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(TempoParserEOF, 0)
}

func (s *SourceFileContext) AllFunc_() []IFuncContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncContext); ok {
			len++
		}
	}

	tst := make([]IFuncContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncContext); ok {
			tst[i] = t.(IFuncContext)
			i++
		}
	}

	return tst
}

func (s *SourceFileContext) Func_(i int) IFuncContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncContext)
}

func (s *SourceFileContext) AllStruct_() []IStructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructContext); ok {
			len++
		}
	}

	tst := make([]IStructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructContext); ok {
			tst[i] = t.(IStructContext)
			i++
		}
	}

	return tst
}

func (s *SourceFileContext) Struct_(i int) IStructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructContext)
}

func (s *SourceFileContext) AllInterface_() []IInterfaceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInterfaceContext); ok {
			len++
		}
	}

	tst := make([]IInterfaceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInterfaceContext); ok {
			tst[i] = t.(IInterfaceContext)
			i++
		}
	}

	return tst
}

func (s *SourceFileContext) Interface_(i int) IInterfaceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TempoParserRULE_sourceFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0 {
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserFUNC:
			{
				p.SetState(46)
				p.Func_()
			}

		case TempoParserSTRUCT:
			{
				p.SetState(47)
				p.Struct_()
			}

		case TempoParserINTERFACE:
			{
				p.SetState(48)
				p.Interface_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(TempoParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) ID() antlr.TerminalNode {
	return s.GetToken(TempoParserID, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TempoParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(TempoParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueTypeContext is an interface to support dynamic dispatch.
type IValueTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueTypeContext differentiates from other interfaces.
	IsValueTypeContext()
}

type ValueTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueTypeContext() *ValueTypeContext {
	var p = new(ValueTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_valueType
	return p
}

func InitEmptyValueTypeContext(p *ValueTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_valueType
}

func (*ValueTypeContext) IsValueTypeContext() {}

func NewValueTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueTypeContext {
	var p = new(ValueTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_valueType

	return p
}

func (s *ValueTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueTypeContext) CopyAll(ctx *ValueTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedTypeContext struct {
	ValueTypeContext
}

func NewNamedTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedTypeContext {
	var p = new(NamedTypeContext)

	InitEmptyValueTypeContext(&p.ValueTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueTypeContext))

	return p
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *NamedTypeContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *NamedTypeContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *NamedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterNamedType(s)
	}
}

func (s *NamedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitNamedType(s)
	}
}

func (s *NamedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitNamedType(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsyncTypeContext struct {
	ValueTypeContext
	inner IValueTypeContext
}

func NewAsyncTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsyncTypeContext {
	var p = new(AsyncTypeContext)

	InitEmptyValueTypeContext(&p.ValueTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueTypeContext))

	return p
}

func (s *AsyncTypeContext) GetInner() IValueTypeContext { return s.inner }

func (s *AsyncTypeContext) SetInner(v IValueTypeContext) { s.inner = v }

func (s *AsyncTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsyncTypeContext) ASYNC() antlr.TerminalNode {
	return s.GetToken(TempoParserASYNC, 0)
}

func (s *AsyncTypeContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *AsyncTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterAsyncType(s)
	}
}

func (s *AsyncTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitAsyncType(s)
	}
}

func (s *AsyncTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitAsyncType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListTypeContext struct {
	ValueTypeContext
	inner IValueTypeContext
}

func NewListTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListTypeContext {
	var p = new(ListTypeContext)

	InitEmptyValueTypeContext(&p.ValueTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueTypeContext))

	return p
}

func (s *ListTypeContext) GetInner() IValueTypeContext { return s.inner }

func (s *ListTypeContext) SetInner(v IValueTypeContext) { s.inner = v }

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserLSQUARE, 0)
}

func (s *ListTypeContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserRSQUARE, 0)
}

func (s *ListTypeContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitListType(s)
	}
}

func (s *ListTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitListType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClosureTypeContext struct {
	ValueTypeContext
	params     IClosureParamListContext
	returnType IValueTypeContext
}

func NewClosureTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClosureTypeContext {
	var p = new(ClosureTypeContext)

	InitEmptyValueTypeContext(&p.ValueTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueTypeContext))

	return p
}

func (s *ClosureTypeContext) GetParams() IClosureParamListContext { return s.params }

func (s *ClosureTypeContext) GetReturnType() IValueTypeContext { return s.returnType }

func (s *ClosureTypeContext) SetParams(v IClosureParamListContext) { s.params = v }

func (s *ClosureTypeContext) SetReturnType(v IValueTypeContext) { s.returnType = v }

func (s *ClosureTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosureTypeContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TempoParserFUNC, 0)
}

func (s *ClosureTypeContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *ClosureTypeContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *ClosureTypeContext) ClosureParamList() IClosureParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureParamListContext)
}

func (s *ClosureTypeContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ClosureTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterClosureType(s)
	}
}

func (s *ClosureTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitClosureType(s)
	}
}

func (s *ClosureTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitClosureType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) ValueType() (localctx IValueTypeContext) {
	localctx = NewValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TempoParserRULE_valueType)
	var _la int

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserASYNC:
		localctx = NewAsyncTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(TempoParserASYNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _x = p.ValueType()

			localctx.(*AsyncTypeContext).inner = _x
		}

	case TempoParserLSQUARE:
		localctx = NewListTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)

			var _x = p.ValueType()

			localctx.(*ListTypeContext).inner = _x
		}
		{
			p.SetState(62)
			p.Match(TempoParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TempoParserFUNC:
		localctx = NewClosureTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(TempoParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.RoleType()
		}
		{
			p.SetState(67)

			var _x = p.ClosureParamList()

			localctx.(*ClosureTypeContext).params = _x
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755830344) != 0 {
			{
				p.SetState(68)

				var _x = p.ValueType()

				localctx.(*ClosureTypeContext).returnType = _x
			}

		}

	case TempoParserID:
		localctx = NewNamedTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.Ident()
		}
		{
			p.SetState(72)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.RoleType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRoleTypeContext is an interface to support dynamic dispatch.
type IRoleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRoleTypeContext differentiates from other interfaces.
	IsRoleTypeContext()
}

type RoleTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoleTypeContext() *RoleTypeContext {
	var p = new(RoleTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_roleType
	return p
}

func InitEmptyRoleTypeContext(p *RoleTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_roleType
}

func (*RoleTypeContext) IsRoleTypeContext() {}

func NewRoleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoleTypeContext {
	var p = new(RoleTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_roleType

	return p
}

func (s *RoleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RoleTypeContext) CopyAll(ctx *RoleTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RoleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RoleTypeSharedContext struct {
	RoleTypeContext
}

func NewRoleTypeSharedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoleTypeSharedContext {
	var p = new(RoleTypeSharedContext)

	InitEmptyRoleTypeContext(&p.RoleTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RoleTypeContext))

	return p
}

func (s *RoleTypeSharedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoleTypeSharedContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserLSQUARE, 0)
}

func (s *RoleTypeSharedContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *RoleTypeSharedContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RoleTypeSharedContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserRSQUARE, 0)
}

func (s *RoleTypeSharedContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *RoleTypeSharedContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *RoleTypeSharedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterRoleTypeShared(s)
	}
}

func (s *RoleTypeSharedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitRoleTypeShared(s)
	}
}

func (s *RoleTypeSharedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitRoleTypeShared(s)

	default:
		return t.VisitChildren(s)
	}
}

type RoleTypeNormalContext struct {
	RoleTypeContext
}

func NewRoleTypeNormalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoleTypeNormalContext {
	var p = new(RoleTypeNormalContext)

	InitEmptyRoleTypeContext(&p.RoleTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RoleTypeContext))

	return p
}

func (s *RoleTypeNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoleTypeNormalContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *RoleTypeNormalContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RoleTypeNormalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserLPAREN, 0)
}

func (s *RoleTypeNormalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserRPAREN, 0)
}

func (s *RoleTypeNormalContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *RoleTypeNormalContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *RoleTypeNormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterRoleTypeNormal(s)
	}
}

func (s *RoleTypeNormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitRoleTypeNormal(s)
	}
}

func (s *RoleTypeNormalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitRoleTypeNormal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) RoleType() (localctx IRoleTypeContext) {
	localctx = NewRoleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TempoParserRULE_roleType)
	var _la int

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserLSQUARE:
		localctx = NewRoleTypeSharedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Ident()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(79)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(80)
				p.Ident()
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(86)
			p.Match(TempoParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TempoParserLPAREN, TempoParserID:
		localctx = NewRoleTypeNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserID:
			{
				p.SetState(88)
				p.Ident()
			}

		case TempoParserLPAREN:
			{
				p.SetState(89)
				p.Match(TempoParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(90)
				p.Ident()
			}
			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == TempoParserCOMMA {
				{
					p.SetState(91)
					p.Match(TempoParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.Ident()
				}

				p.SetState(97)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(98)
				p.Match(TempoParserRPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosureParamListContext is an interface to support dynamic dispatch.
type IClosureParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllValueType() []IValueTypeContext
	ValueType(i int) IValueTypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsClosureParamListContext differentiates from other interfaces.
	IsClosureParamListContext()
}

type ClosureParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosureParamListContext() *ClosureParamListContext {
	var p = new(ClosureParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_closureParamList
	return p
}

func InitEmptyClosureParamListContext(p *ClosureParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_closureParamList
}

func (*ClosureParamListContext) IsClosureParamListContext() {}

func NewClosureParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosureParamListContext {
	var p = new(ClosureParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_closureParamList

	return p
}

func (s *ClosureParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosureParamListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserLPAREN, 0)
}

func (s *ClosureParamListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserRPAREN, 0)
}

func (s *ClosureParamListContext) AllValueType() []IValueTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueTypeContext); ok {
			len++
		}
	}

	tst := make([]IValueTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueTypeContext); ok {
			tst[i] = t.(IValueTypeContext)
			i++
		}
	}

	return tst
}

func (s *ClosureParamListContext) ValueType(i int) IValueTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ClosureParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *ClosureParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *ClosureParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosureParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosureParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterClosureParamList(s)
	}
}

func (s *ClosureParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitClosureParamList(s)
	}
}

func (s *ClosureParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitClosureParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) ClosureParamList() (localctx IClosureParamListContext) {
	localctx = NewClosureParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TempoParserRULE_closureParamList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755830344) != 0 {
		{
			p.SetState(105)
			p.ValueType()
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(106)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(107)
				p.ValueType()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(115)
		p.Match(TempoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosureSigContext is an interface to support dynamic dispatch.
type IClosureSigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetParams returns the params rule contexts.
	GetParams() IFuncParamListContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() IValueTypeContext

	// SetParams sets the params rule contexts.
	SetParams(IFuncParamListContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(IValueTypeContext)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
	FuncParamList() IFuncParamListContext
	ValueType() IValueTypeContext

	// IsClosureSigContext differentiates from other interfaces.
	IsClosureSigContext()
}

type ClosureSigContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	params     IFuncParamListContext
	returnType IValueTypeContext
}

func NewEmptyClosureSigContext() *ClosureSigContext {
	var p = new(ClosureSigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_closureSig
	return p
}

func InitEmptyClosureSigContext(p *ClosureSigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_closureSig
}

func (*ClosureSigContext) IsClosureSigContext() {}

func NewClosureSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosureSigContext {
	var p = new(ClosureSigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_closureSig

	return p
}

func (s *ClosureSigContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosureSigContext) GetParams() IFuncParamListContext { return s.params }

func (s *ClosureSigContext) GetReturnType() IValueTypeContext { return s.returnType }

func (s *ClosureSigContext) SetParams(v IFuncParamListContext) { s.params = v }

func (s *ClosureSigContext) SetReturnType(v IValueTypeContext) { s.returnType = v }

func (s *ClosureSigContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TempoParserFUNC, 0)
}

func (s *ClosureSigContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *ClosureSigContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *ClosureSigContext) FuncParamList() IFuncParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamListContext)
}

func (s *ClosureSigContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ClosureSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosureSigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosureSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterClosureSig(s)
	}
}

func (s *ClosureSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitClosureSig(s)
	}
}

func (s *ClosureSigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitClosureSig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) ClosureSig() (localctx IClosureSigContext) {
	localctx = NewClosureSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TempoParserRULE_closureSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(TempoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.RoleType()
	}
	{
		p.SetState(120)

		var _x = p.FuncParamList()

		localctx.(*ClosureSigContext).params = _x
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755830344) != 0 {
		{
			p.SetState(121)

			var _x = p.ValueType()

			localctx.(*ClosureSigContext).returnType = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructContext is an interface to support dynamic dispatch.
type IStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
	Ident() IIdentContext
	StructFieldList() IStructFieldListContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructContext() *StructContext {
	var p = new(StructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_struct
	return p
}

func InitEmptyStructContext(p *StructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_struct
}

func (*StructContext) IsStructContext() {}

func NewStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructContext {
	var p = new(StructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_struct

	return p
}

func (s *StructContext) GetParser() antlr.Parser { return s.parser }

func (s *StructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(TempoParserSTRUCT, 0)
}

func (s *StructContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *StructContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *StructContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *StructContext) StructFieldList() IStructFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldListContext)
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Struct_() (localctx IStructContext) {
	localctx = NewStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TempoParserRULE_struct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(TempoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.RoleType()
	}
	{
		p.SetState(127)
		p.Ident()
	}
	{
		p.SetState(128)
		p.StructFieldList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldListContext is an interface to support dynamic dispatch.
type IStructFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructFieldListContext differentiates from other interfaces.
	IsStructFieldListContext()
}

type StructFieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldListContext() *StructFieldListContext {
	var p = new(StructFieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structFieldList
	return p
}

func InitEmptyStructFieldListContext(p *StructFieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structFieldList
}

func (*StructFieldListContext) IsStructFieldListContext() {}

func NewStructFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldListContext {
	var p = new(StructFieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_structFieldList

	return p
}

func (s *StructFieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldListContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserLCURLY, 0)
}

func (s *StructFieldListContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserRCURLY, 0)
}

func (s *StructFieldListContext) AllStructField() []IStructFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFieldContext); ok {
			len++
		}
	}

	tst := make([]IStructFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFieldContext); ok {
			tst[i] = t.(IStructFieldContext)
			i++
		}
	}

	return tst
}

func (s *StructFieldListContext) StructField(i int) IStructFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFieldContext)
}

func (s *StructFieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *StructFieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *StructFieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStructFieldList(s)
	}
}

func (s *StructFieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStructFieldList(s)
	}
}

func (s *StructFieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStructFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) StructFieldList() (localctx IStructFieldListContext) {
	localctx = NewStructFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TempoParserRULE_structFieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(131)
			p.StructField()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(132)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(133)
				p.StructField()
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(141)
		p.Match(TempoParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFieldContext is an interface to support dynamic dispatch.
type IStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	COLON() antlr.TerminalNode
	ValueType() IValueTypeContext

	// IsStructFieldContext differentiates from other interfaces.
	IsStructFieldContext()
}

type StructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFieldContext() *StructFieldContext {
	var p = new(StructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structField
	return p
}

func InitEmptyStructFieldContext(p *StructFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structField
}

func (*StructFieldContext) IsStructFieldContext() {}

func NewStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFieldContext {
	var p = new(StructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_structField

	return p
}

func (s *StructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFieldContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *StructFieldContext) COLON() antlr.TerminalNode {
	return s.GetToken(TempoParserCOLON, 0)
}

func (s *StructFieldContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) StructField() (localctx IStructFieldContext) {
	localctx = NewStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TempoParserRULE_structField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Ident()
	}
	{
		p.SetState(144)
		p.Match(TempoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.ValueType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceContext is an interface to support dynamic dispatch.
type IInterfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
	Ident() IIdentContext
	InterfaceMethodsList() IInterfaceMethodsListContext

	// IsInterfaceContext differentiates from other interfaces.
	IsInterfaceContext()
}

type InterfaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceContext() *InterfaceContext {
	var p = new(InterfaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_interface
	return p
}

func InitEmptyInterfaceContext(p *InterfaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_interface
}

func (*InterfaceContext) IsInterfaceContext() {}

func NewInterfaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceContext {
	var p = new(InterfaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_interface

	return p
}

func (s *InterfaceContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(TempoParserINTERFACE, 0)
}

func (s *InterfaceContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *InterfaceContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *InterfaceContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *InterfaceContext) InterfaceMethodsList() IInterfaceMethodsListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceMethodsListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceMethodsListContext)
}

func (s *InterfaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterInterface(s)
	}
}

func (s *InterfaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitInterface(s)
	}
}

func (s *InterfaceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitInterface(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Interface_() (localctx IInterfaceContext) {
	localctx = NewInterfaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TempoParserRULE_interface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(TempoParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.RoleType()
	}
	{
		p.SetState(150)
		p.Ident()
	}
	{
		p.SetState(151)
		p.InterfaceMethodsList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceMethodsListContext is an interface to support dynamic dispatch.
type IInterfaceMethodsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllInterfaceMethod() []IInterfaceMethodContext
	InterfaceMethod(i int) IInterfaceMethodContext

	// IsInterfaceMethodsListContext differentiates from other interfaces.
	IsInterfaceMethodsListContext()
}

type InterfaceMethodsListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceMethodsListContext() *InterfaceMethodsListContext {
	var p = new(InterfaceMethodsListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_interfaceMethodsList
	return p
}

func InitEmptyInterfaceMethodsListContext(p *InterfaceMethodsListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_interfaceMethodsList
}

func (*InterfaceMethodsListContext) IsInterfaceMethodsListContext() {}

func NewInterfaceMethodsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceMethodsListContext {
	var p = new(InterfaceMethodsListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_interfaceMethodsList

	return p
}

func (s *InterfaceMethodsListContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceMethodsListContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserLCURLY, 0)
}

func (s *InterfaceMethodsListContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserRCURLY, 0)
}

func (s *InterfaceMethodsListContext) AllInterfaceMethod() []IInterfaceMethodContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInterfaceMethodContext); ok {
			len++
		}
	}

	tst := make([]IInterfaceMethodContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInterfaceMethodContext); ok {
			tst[i] = t.(IInterfaceMethodContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceMethodsListContext) InterfaceMethod(i int) IInterfaceMethodContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceMethodContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceMethodContext)
}

func (s *InterfaceMethodsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceMethodsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceMethodsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterInterfaceMethodsList(s)
	}
}

func (s *InterfaceMethodsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitInterfaceMethodsList(s)
	}
}

func (s *InterfaceMethodsListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitInterfaceMethodsList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) InterfaceMethodsList() (localctx IInterfaceMethodsListContext) {
	localctx = NewInterfaceMethodsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TempoParserRULE_interfaceMethodsList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TempoParserFUNC {
		{
			p.SetState(154)
			p.InterfaceMethod()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
		p.Match(TempoParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceMethodContext is an interface to support dynamic dispatch.
type IInterfaceMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncSig() IFuncSigContext
	END() antlr.TerminalNode

	// IsInterfaceMethodContext differentiates from other interfaces.
	IsInterfaceMethodContext()
}

type InterfaceMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceMethodContext() *InterfaceMethodContext {
	var p = new(InterfaceMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_interfaceMethod
	return p
}

func InitEmptyInterfaceMethodContext(p *InterfaceMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_interfaceMethod
}

func (*InterfaceMethodContext) IsInterfaceMethodContext() {}

func NewInterfaceMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceMethodContext {
	var p = new(InterfaceMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_interfaceMethod

	return p
}

func (s *InterfaceMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceMethodContext) FuncSig() IFuncSigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSigContext)
}

func (s *InterfaceMethodContext) END() antlr.TerminalNode {
	return s.GetToken(TempoParserEND, 0)
}

func (s *InterfaceMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterInterfaceMethod(s)
	}
}

func (s *InterfaceMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitInterfaceMethod(s)
	}
}

func (s *InterfaceMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitInterfaceMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) InterfaceMethod() (localctx IInterfaceMethodContext) {
	localctx = NewInterfaceMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TempoParserRULE_interfaceMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.FuncSig()
	}
	{
		p.SetState(163)
		p.Match(TempoParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncContext is an interface to support dynamic dispatch.
type IFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncSig() IFuncSigContext
	Scope() IScopeContext

	// IsFuncContext differentiates from other interfaces.
	IsFuncContext()
}

type FuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncContext() *FuncContext {
	var p = new(FuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_func
	return p
}

func InitEmptyFuncContext(p *FuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_func
}

func (*FuncContext) IsFuncContext() {}

func NewFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncContext {
	var p = new(FuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_func

	return p
}

func (s *FuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncContext) FuncSig() IFuncSigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncSigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncSigContext)
}

func (s *FuncContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterFunc(s)
	}
}

func (s *FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitFunc(s)
	}
}

func (s *FuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Func_() (localctx IFuncContext) {
	localctx = NewFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TempoParserRULE_func)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.FuncSig()
	}
	{
		p.SetState(166)
		p.Scope()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncSigContext is an interface to support dynamic dispatch.
type IFuncSigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IIdentContext

	// GetParams returns the params rule contexts.
	GetParams() IFuncParamListContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() IValueTypeContext

	// SetName sets the name rule contexts.
	SetName(IIdentContext)

	// SetParams sets the params rule contexts.
	SetParams(IFuncParamListContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(IValueTypeContext)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
	Ident() IIdentContext
	FuncParamList() IFuncParamListContext
	ValueType() IValueTypeContext

	// IsFuncSigContext differentiates from other interfaces.
	IsFuncSigContext()
}

type FuncSigContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	name       IIdentContext
	params     IFuncParamListContext
	returnType IValueTypeContext
}

func NewEmptyFuncSigContext() *FuncSigContext {
	var p = new(FuncSigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcSig
	return p
}

func InitEmptyFuncSigContext(p *FuncSigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcSig
}

func (*FuncSigContext) IsFuncSigContext() {}

func NewFuncSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSigContext {
	var p = new(FuncSigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_funcSig

	return p
}

func (s *FuncSigContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSigContext) GetName() IIdentContext { return s.name }

func (s *FuncSigContext) GetParams() IFuncParamListContext { return s.params }

func (s *FuncSigContext) GetReturnType() IValueTypeContext { return s.returnType }

func (s *FuncSigContext) SetName(v IIdentContext) { s.name = v }

func (s *FuncSigContext) SetParams(v IFuncParamListContext) { s.params = v }

func (s *FuncSigContext) SetReturnType(v IValueTypeContext) { s.returnType = v }

func (s *FuncSigContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TempoParserFUNC, 0)
}

func (s *FuncSigContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *FuncSigContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *FuncSigContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FuncSigContext) FuncParamList() IFuncParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamListContext)
}

func (s *FuncSigContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *FuncSigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterFuncSig(s)
	}
}

func (s *FuncSigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitFuncSig(s)
	}
}

func (s *FuncSigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitFuncSig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) FuncSig() (localctx IFuncSigContext) {
	localctx = NewFuncSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TempoParserRULE_funcSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(TempoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.RoleType()
	}
	{
		p.SetState(171)

		var _x = p.Ident()

		localctx.(*FuncSigContext).name = _x
	}
	{
		p.SetState(172)

		var _x = p.FuncParamList()

		localctx.(*FuncSigContext).params = _x
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549755830344) != 0 {
		{
			p.SetState(173)

			var _x = p.ValueType()

			localctx.(*FuncSigContext).returnType = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncParamListContext is an interface to support dynamic dispatch.
type IFuncParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllFuncParam() []IFuncParamContext
	FuncParam(i int) IFuncParamContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncParamListContext differentiates from other interfaces.
	IsFuncParamListContext()
}

type FuncParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamListContext() *FuncParamListContext {
	var p = new(FuncParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcParamList
	return p
}

func InitEmptyFuncParamListContext(p *FuncParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcParamList
}

func (*FuncParamListContext) IsFuncParamListContext() {}

func NewFuncParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamListContext {
	var p = new(FuncParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_funcParamList

	return p
}

func (s *FuncParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserLPAREN, 0)
}

func (s *FuncParamListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserRPAREN, 0)
}

func (s *FuncParamListContext) AllFuncParam() []IFuncParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncParamContext); ok {
			len++
		}
	}

	tst := make([]IFuncParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncParamContext); ok {
			tst[i] = t.(IFuncParamContext)
			i++
		}
	}

	return tst
}

func (s *FuncParamListContext) FuncParam(i int) IFuncParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncParamContext)
}

func (s *FuncParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *FuncParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *FuncParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterFuncParamList(s)
	}
}

func (s *FuncParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitFuncParamList(s)
	}
}

func (s *FuncParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitFuncParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) FuncParamList() (localctx IFuncParamListContext) {
	localctx = NewFuncParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TempoParserRULE_funcParamList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(177)
			p.FuncParam()
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(178)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(179)
				p.FuncParam()
			}

			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(187)
		p.Match(TempoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncParamContext is an interface to support dynamic dispatch.
type IFuncParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	COLON() antlr.TerminalNode
	ValueType() IValueTypeContext

	// IsFuncParamContext differentiates from other interfaces.
	IsFuncParamContext()
}

type FuncParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamContext() *FuncParamContext {
	var p = new(FuncParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcParam
	return p
}

func InitEmptyFuncParamContext(p *FuncParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcParam
}

func (*FuncParamContext) IsFuncParamContext() {}

func NewFuncParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamContext {
	var p = new(FuncParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_funcParam

	return p
}

func (s *FuncParamContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FuncParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(TempoParserCOLON, 0)
}

func (s *FuncParamContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) FuncParam() (localctx IFuncParamContext) {
	localctx = NewFuncParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TempoParserRULE_funcParam)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Ident()
	}
	{
		p.SetState(190)
		p.Match(TempoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.ValueType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgListContext is an interface to support dynamic dispatch.
type IFuncArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncArgListContext differentiates from other interfaces.
	IsFuncArgListContext()
}

type FuncArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgListContext() *FuncArgListContext {
	var p = new(FuncArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcArgList
	return p
}

func InitEmptyFuncArgListContext(p *FuncArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_funcArgList
}

func (*FuncArgListContext) IsFuncArgListContext() {}

func NewFuncArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgListContext {
	var p = new(FuncArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_funcArgList

	return p
}

func (s *FuncArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserLPAREN, 0)
}

func (s *FuncArgListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserRPAREN, 0)
}

func (s *FuncArgListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *FuncArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *FuncArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterFuncArgList(s)
	}
}

func (s *FuncArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitFuncArgList(s)
	}
}

func (s *FuncArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitFuncArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) FuncArgList() (localctx IFuncArgListContext) {
	localctx = NewFuncArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TempoParserRULE_funcArgList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1958505117832) != 0 {
		{
			p.SetState(194)
			p.expr(0)
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(195)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(196)
				p.expr(0)
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(204)
		p.Match(TempoParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_scope
	return p
}

func InitEmptyScopeContext(p *ScopeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_scope
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserLCURLY, 0)
}

func (s *ScopeContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserRCURLY, 0)
}

func (s *ScopeContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitScope(s)
	}
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Scope() (localctx IScopeContext) {
	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TempoParserRULE_scope)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1958505119160) != 0 {
		{
			p.SetState(207)
			p.Stmt()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
		p.Match(TempoParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StmtWhileContext struct {
	StmtContext
}

func NewStmtWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtWhileContext {
	var p = new(StmtWhileContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(TempoParserWHILE, 0)
}

func (s *StmtWhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtWhileContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *StmtWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStmtWhile(s)
	}
}

func (s *StmtWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStmtWhile(s)
	}
}

func (s *StmtWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStmtWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtExprContext struct {
	StmtContext
}

func NewStmtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtExprContext {
	var p = new(StmtExprContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtExprContext) END() antlr.TerminalNode {
	return s.GetToken(TempoParserEND, 0)
}

func (s *StmtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStmtExpr(s)
	}
}

func (s *StmtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStmtExpr(s)
	}
}

func (s *StmtExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStmtExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtIfContext struct {
	StmtContext
	thenScope IScopeContext
	elseScope IScopeContext
}

func NewStmtIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtIfContext {
	var p = new(StmtIfContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtIfContext) GetThenScope() IScopeContext { return s.thenScope }

func (s *StmtIfContext) GetElseScope() IScopeContext { return s.elseScope }

func (s *StmtIfContext) SetThenScope(v IScopeContext) { s.thenScope = v }

func (s *StmtIfContext) SetElseScope(v IScopeContext) { s.elseScope = v }

func (s *StmtIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtIfContext) IF() antlr.TerminalNode {
	return s.GetToken(TempoParserIF, 0)
}

func (s *StmtIfContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtIfContext) AllScope() []IScopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScopeContext); ok {
			len++
		}
	}

	tst := make([]IScopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScopeContext); ok {
			tst[i] = t.(IScopeContext)
			i++
		}
	}

	return tst
}

func (s *StmtIfContext) Scope(i int) IScopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *StmtIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TempoParserELSE, 0)
}

func (s *StmtIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStmtIf(s)
	}
}

func (s *StmtIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStmtIf(s)
	}
}

func (s *StmtIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStmtIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtVarDeclContext struct {
	StmtContext
}

func NewStmtVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtVarDeclContext {
	var p = new(StmtVarDeclContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtVarDeclContext) LET() antlr.TerminalNode {
	return s.GetToken(TempoParserLET, 0)
}

func (s *StmtVarDeclContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *StmtVarDeclContext) IS() antlr.TerminalNode {
	return s.GetToken(TempoParserIS, 0)
}

func (s *StmtVarDeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtVarDeclContext) END() antlr.TerminalNode {
	return s.GetToken(TempoParserEND, 0)
}

func (s *StmtVarDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(TempoParserCOLON, 0)
}

func (s *StmtVarDeclContext) ValueType() IValueTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *StmtVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStmtVarDecl(s)
	}
}

func (s *StmtVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStmtVarDecl(s)
	}
}

func (s *StmtVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStmtVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtReturnContext struct {
	StmtContext
}

func NewStmtReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtReturnContext {
	var p = new(StmtReturnContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(TempoParserRETURN, 0)
}

func (s *StmtReturnContext) END() antlr.TerminalNode {
	return s.GetToken(TempoParserEND, 0)
}

func (s *StmtReturnContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStmtReturn(s)
	}
}

func (s *StmtReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStmtReturn(s)
	}
}

func (s *StmtReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStmtReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtAssignContext struct {
	StmtContext
}

func NewStmtAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtAssignContext {
	var p = new(StmtAssignContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtAssignContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *StmtAssignContext) IS() antlr.TerminalNode {
	return s.GetToken(TempoParserIS, 0)
}

func (s *StmtAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtAssignContext) END() antlr.TerminalNode {
	return s.GetToken(TempoParserEND, 0)
}

func (s *StmtAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStmtAssign(s)
	}
}

func (s *StmtAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStmtAssign(s)
	}
}

func (s *StmtAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStmtAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TempoParserRULE_stmt)
	var _la int

	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmtVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.Match(TempoParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Ident()
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TempoParserCOLON {
			{
				p.SetState(217)
				p.Match(TempoParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(218)
				p.ValueType()
			}

		}
		{
			p.SetState(221)
			p.Match(TempoParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.expr(0)
		}
		{
			p.SetState(223)
			p.Match(TempoParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStmtIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.Match(TempoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.expr(0)
		}
		{
			p.SetState(227)

			var _x = p.Scope()

			localctx.(*StmtIfContext).thenScope = _x
		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TempoParserELSE {
			{
				p.SetState(228)
				p.Match(TempoParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(229)

				var _x = p.Scope()

				localctx.(*StmtIfContext).elseScope = _x
			}

		}

	case 3:
		localctx = NewStmtWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.Match(TempoParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.expr(0)
		}
		{
			p.SetState(234)
			p.Scope()
		}

	case 4:
		localctx = NewStmtReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(236)
			p.Match(TempoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1958505117832) != 0 {
			{
				p.SetState(237)
				p.expr(0)
			}

		}
		{
			p.SetState(240)
			p.Match(TempoParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStmtAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(241)
			p.Ident()
		}
		{
			p.SetState(242)
			p.Match(TempoParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.expr(0)
		}
		{
			p.SetState(244)
			p.Match(TempoParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStmtExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(246)
			p.expr(0)
		}
		{
			p.SetState(247)
			p.Match(TempoParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprBinOpContext struct {
	ExprContext
	lhs IExprContext
	rhs IExprContext
}

func NewExprBinOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBinOpContext {
	var p = new(ExprBinOpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprBinOpContext) GetLhs() IExprContext { return s.lhs }

func (s *ExprBinOpContext) GetRhs() IExprContext { return s.rhs }

func (s *ExprBinOpContext) SetLhs(v IExprContext) { s.lhs = v }

func (s *ExprBinOpContext) SetRhs(v IExprContext) { s.rhs = v }

func (s *ExprBinOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBinOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprBinOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprBinOpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TempoParserPLUS, 0)
}

func (s *ExprBinOpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TempoParserMINUS, 0)
}

func (s *ExprBinOpContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(TempoParserMULTIPLY, 0)
}

func (s *ExprBinOpContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(TempoParserDIVIDE, 0)
}

func (s *ExprBinOpContext) MODULO() antlr.TerminalNode {
	return s.GetToken(TempoParserMODULO, 0)
}

func (s *ExprBinOpContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TempoParserEQUAL, 0)
}

func (s *ExprBinOpContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(TempoParserNOT_EQUAL, 0)
}

func (s *ExprBinOpContext) LESS() antlr.TerminalNode {
	return s.GetToken(TempoParserLESS, 0)
}

func (s *ExprBinOpContext) LESS_EQ() antlr.TerminalNode {
	return s.GetToken(TempoParserLESS_EQ, 0)
}

func (s *ExprBinOpContext) GREATER() antlr.TerminalNode {
	return s.GetToken(TempoParserGREATER, 0)
}

func (s *ExprBinOpContext) GREATER_EQ() antlr.TerminalNode {
	return s.GetToken(TempoParserGREATER_EQ, 0)
}

func (s *ExprBinOpContext) AND() antlr.TerminalNode {
	return s.GetToken(TempoParserAND, 0)
}

func (s *ExprBinOpContext) OR() antlr.TerminalNode {
	return s.GetToken(TempoParserOR, 0)
}

func (s *ExprBinOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprBinOp(s)
	}
}

func (s *ExprBinOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprBinOp(s)
	}
}

func (s *ExprBinOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprBinOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprFieldAccessContext struct {
	ExprContext
}

func NewExprFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldAccessContext {
	var p = new(ExprFieldAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprFieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprFieldAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(TempoParserDOT, 0)
}

func (s *ExprFieldAccessContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ExprFieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprFieldAccess(s)
	}
}

func (s *ExprFieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprFieldAccess(s)
	}
}

func (s *ExprFieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprFieldAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprCallContext struct {
	ExprContext
}

func NewExprCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprCallContext {
	var p = new(ExprCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprCallContext) FuncArgList() IFuncArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgListContext)
}

func (s *ExprCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprCall(s)
	}
}

func (s *ExprCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprCall(s)
	}
}

func (s *ExprCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprComContext struct {
	ExprContext
	sender   IRoleTypeContext
	receiver IRoleTypeContext
}

func NewExprComContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprComContext {
	var p = new(ExprComContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprComContext) GetSender() IRoleTypeContext { return s.sender }

func (s *ExprComContext) GetReceiver() IRoleTypeContext { return s.receiver }

func (s *ExprComContext) SetSender(v IRoleTypeContext) { s.sender = v }

func (s *ExprComContext) SetReceiver(v IRoleTypeContext) { s.receiver = v }

func (s *ExprComContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprComContext) COM() antlr.TerminalNode {
	return s.GetToken(TempoParserCOM, 0)
}

func (s *ExprComContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprComContext) AllRoleType() []IRoleTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRoleTypeContext); ok {
			len++
		}
	}

	tst := make([]IRoleTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRoleTypeContext); ok {
			tst[i] = t.(IRoleTypeContext)
			i++
		}
	}

	return tst
}

func (s *ExprComContext) RoleType(i int) IRoleTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *ExprComContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprCom(s)
	}
}

func (s *ExprComContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprCom(s)
	}
}

func (s *ExprComContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprCom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprGroupContext struct {
	ExprContext
}

func NewExprGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprGroupContext {
	var p = new(ExprGroupContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprGroupContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserLPAREN, 0)
}

func (s *ExprGroupContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprGroupContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TempoParserRPAREN, 0)
}

func (s *ExprGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprGroup(s)
	}
}

func (s *ExprGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprGroup(s)
	}
}

func (s *ExprGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStructContext struct {
	ExprContext
}

func NewExprStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStructContext {
	var p = new(ExprStructContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStructContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ExprStructContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *ExprStructContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *ExprStructContext) ExprStructField() IExprStructFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStructFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStructFieldContext)
}

func (s *ExprStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprStruct(s)
	}
}

func (s *ExprStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprStruct(s)
	}
}

func (s *ExprStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprIdentContext struct {
	ExprContext
}

func NewExprIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIdentContext {
	var p = new(ExprIdentContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIdentContext) IdentAccess() IIdentAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentAccessContext)
}

func (s *ExprIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprIdent(s)
	}
}

func (s *ExprIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprIdent(s)
	}
}

func (s *ExprIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAwaitContext struct {
	ExprContext
}

func NewExprAwaitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAwaitContext {
	var p = new(ExprAwaitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprAwaitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAwaitContext) AWAIT() antlr.TerminalNode {
	return s.GetToken(TempoParserAWAIT, 0)
}

func (s *ExprAwaitContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprAwaitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprAwait(s)
	}
}

func (s *ExprAwaitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprAwait(s)
	}
}

func (s *ExprAwaitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprAwait(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprClosureContext struct {
	ExprContext
}

func NewExprClosureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprClosureContext {
	var p = new(ExprClosureContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprClosureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprClosureContext) ClosureSig() IClosureSigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosureSigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosureSigContext)
}

func (s *ExprClosureContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *ExprClosureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprClosure(s)
	}
}

func (s *ExprClosureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprClosure(s)
	}
}

func (s *ExprClosureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprClosure(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprPrimitiveContext struct {
	ExprContext
}

func NewExprPrimitiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprPrimitiveContext {
	var p = new(ExprPrimitiveContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprPrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPrimitiveContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprPrimitiveContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *ExprPrimitiveContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *ExprPrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprPrimitive(s)
	}
}

func (s *ExprPrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprPrimitive(s)
	}
}

func (s *ExprPrimitiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprPrimitive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TempoParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, TempoParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprPrimitiveContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(252)
			p.Literal()
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(253)
				p.Match(TempoParserROLE_AT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(254)
				p.RoleType()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		localctx = NewExprAwaitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(257)
			p.Match(TempoParserAWAIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.expr(8)
		}

	case 3:
		localctx = NewExprClosureContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.ClosureSig()
		}
		{
			p.SetState(260)
			p.Scope()
		}

	case 4:
		localctx = NewExprStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Ident()
		}
		{
			p.SetState(263)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.RoleType()
		}
		{
			p.SetState(265)
			p.ExprStructField()
		}

	case 5:
		localctx = NewExprIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(267)
			p.IdentAccess()
		}

	case 6:
		localctx = NewExprComContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(268)

			var _x = p.RoleType()

			localctx.(*ExprComContext).sender = _x
		}
		{
			p.SetState(269)
			p.Match(TempoParserCOM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)

			var _x = p.RoleType()

			localctx.(*ExprComContext).receiver = _x
		}
		{
			p.SetState(271)
			p.expr(2)
		}

	case 7:
		localctx = NewExprGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)
			p.Match(TempoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.expr(0)
		}
		{
			p.SetState(275)
			p.Match(TempoParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(287)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprBinOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprBinOpContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(279)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(280)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4294443008) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(281)

					var _x = p.expr(11)

					localctx.(*ExprBinOpContext).rhs = _x
				}

			case 2:
				localctx = NewExprCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(283)
					p.FuncArgList()
				}

			case 3:
				localctx = NewExprFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(285)
					p.Match(TempoParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(286)
					p.Ident()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprStructFieldContext is an interface to support dynamic dispatch.
type IExprStructFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprStructFieldContext differentiates from other interfaces.
	IsExprStructFieldContext()
}

type ExprStructFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStructFieldContext() *ExprStructFieldContext {
	var p = new(ExprStructFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_exprStructField
	return p
}

func InitEmptyExprStructFieldContext(p *ExprStructFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_exprStructField
}

func (*ExprStructFieldContext) IsExprStructFieldContext() {}

func NewExprStructFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStructFieldContext {
	var p = new(ExprStructFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_exprStructField

	return p
}

func (s *ExprStructFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStructFieldContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserLCURLY, 0)
}

func (s *ExprStructFieldContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserRCURLY, 0)
}

func (s *ExprStructFieldContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *ExprStructFieldContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ExprStructFieldContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOLON)
}

func (s *ExprStructFieldContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOLON, i)
}

func (s *ExprStructFieldContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprStructFieldContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStructFieldContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *ExprStructFieldContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *ExprStructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStructFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprStructField(s)
	}
}

func (s *ExprStructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprStructField(s)
	}
}

func (s *ExprStructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) ExprStructField() (localctx IExprStructFieldContext) {
	localctx = NewExprStructFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TempoParserRULE_exprStructField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(293)
			p.Ident()
		}
		{
			p.SetState(294)
			p.Match(TempoParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.expr(0)
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(296)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(297)
				p.Ident()
			}
			{
				p.SetState(298)
				p.Match(TempoParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(299)
				p.expr(0)
			}

			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(308)
		p.Match(TempoParserRCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentAccessContext is an interface to support dynamic dispatch.
type IIdentAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext

	// IsIdentAccessContext differentiates from other interfaces.
	IsIdentAccessContext()
}

type IdentAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentAccessContext() *IdentAccessContext {
	var p = new(IdentAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_identAccess
	return p
}

func InitEmptyIdentAccessContext(p *IdentAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_identAccess
}

func (*IdentAccessContext) IsIdentAccessContext() {}

func NewIdentAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentAccessContext {
	var p = new(IdentAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_identAccess

	return p
}

func (s *IdentAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentAccessContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *IdentAccessContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *IdentAccessContext) RoleType() IRoleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleTypeContext)
}

func (s *IdentAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterIdentAccess(s)
	}
}

func (s *IdentAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitIdentAccess(s)
	}
}

func (s *IdentAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitIdentAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) IdentAccess() (localctx IIdentAccessContext) {
	localctx = NewIdentAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TempoParserRULE_identAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Ident()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(311)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.RoleType()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringContext struct {
	LiteralContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(TempoParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	LiteralContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TempoParserTRUE, 0)
}

func (s *BoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TempoParserFALSE, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatContext struct {
	LiteralContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(TempoParserNUMBER)
}

func (s *FloatContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserNUMBER, i)
}

func (s *FloatContext) DOT() antlr.TerminalNode {
	return s.GetToken(TempoParserDOT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (s *FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	LiteralContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TempoParserNUMBER, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TempoParserRULE_literal)
	var _la int

	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(315)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(316)
				p.Match(TempoParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(317)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(318)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(319)
				p.Match(TempoParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 3:
			{
				p.SetState(320)
				p.Match(TempoParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(321)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(TempoParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(325)
			p.Match(TempoParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(326)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TempoParserTRUE || _la == TempoParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TempoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TempoParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
