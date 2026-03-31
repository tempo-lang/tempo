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
		"", "'struct'", "'interface'", "'implements'", "'func'", "'return'",
		"'let'", "'async'", "'await'", "'if'", "'else'", "'while'", "'true'",
		"'false'", "'('", "'['", "'{'", "')'", "']'", "'}'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'",
		"'||'", "'='", "'@'", "','", "'.'", "':'", "'->'", "", "'_'", "", "",
		"';'",
	}
	staticData.SymbolicNames = []string{
		"", "STRUCT", "INTERFACE", "IMPLEMENTS", "FUNC", "RETURN", "LET", "ASYNC",
		"AWAIT", "IF", "ELSE", "WHILE", "TRUE", "FALSE", "LPAREN", "LSQUARE",
		"LCURLY", "RPAREN", "RSQUARE", "RCURLY", "PLUS", "MINUS", "MULTIPLY",
		"DIVIDE", "MODULO", "EQUAL", "NOT_EQUAL", "LESS", "LESS_EQ", "GREATER",
		"GREATER_EQ", "AND", "OR", "IS", "ROLE_AT", "COMMA", "DOT", "COLON",
		"COM", "STRING", "UNDERSCORE", "ID", "NUMBER", "END", "WHITESPACE",
		"LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"sourceFile", "ident", "roleIdent", "valueType", "role", "roleType",
		"closureParamList", "closureSig", "struct", "structImplements", "structBody",
		"structField", "interface", "interfaceMethodsList", "interfaceMethod",
		"func", "funcSig", "funcParamList", "funcParam", "funcArgList", "scope",
		"stmt", "assignExpr", "assignSpecifier", "expr", "exprStructField",
		"identAccess", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 392, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 5, 0, 60, 8, 0, 10, 0, 12, 0, 63,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 72, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 85, 8, 3,
		1, 3, 3, 3, 88, 8, 3, 1, 4, 1, 4, 3, 4, 92, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 98, 8, 5, 10, 5, 12, 5, 101, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 110, 8, 5, 10, 5, 12, 5, 113, 9, 5, 1, 5, 1, 5, 3,
		5, 117, 8, 5, 3, 5, 119, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 125, 8, 6,
		10, 6, 12, 6, 128, 9, 6, 3, 6, 130, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 139, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 144, 8, 8, 1, 8, 1,
		8, 3, 8, 148, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 156, 8, 9,
		10, 9, 12, 9, 159, 9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 164, 8, 10, 10, 10,
		12, 10, 167, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 3, 12, 179, 8, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		5, 13, 186, 8, 13, 10, 13, 12, 13, 189, 9, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 202, 8, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 207, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5,
		17, 213, 8, 17, 10, 17, 12, 17, 216, 9, 17, 3, 17, 218, 8, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 230,
		8, 19, 10, 19, 12, 19, 233, 9, 19, 3, 19, 235, 8, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 5, 20, 241, 8, 20, 10, 20, 12, 20, 244, 9, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 252, 8, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 263, 8, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 271, 8, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 282, 8, 21, 1, 22, 1, 22,
		5, 22, 286, 8, 22, 10, 22, 12, 22, 289, 9, 22, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 297, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24,
		303, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 5, 24, 315, 8, 24, 10, 24, 12, 24, 318, 9, 24, 3, 24, 320, 8,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 3, 24, 335, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 350,
		8, 24, 10, 24, 12, 24, 353, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 5, 25, 364, 8, 25, 10, 25, 12, 25, 367, 9, 25,
		3, 25, 369, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 3, 26, 376, 8, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 385, 8, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 390, 8, 27, 1, 27, 0, 1, 48, 28, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 0, 2, 1, 0, 20, 32, 1, 0, 12, 13, 425, 0, 61, 1, 0,
		0, 0, 2, 66, 1, 0, 0, 0, 4, 68, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 91, 1,
		0, 0, 0, 10, 118, 1, 0, 0, 0, 12, 120, 1, 0, 0, 0, 14, 133, 1, 0, 0, 0,
		16, 140, 1, 0, 0, 0, 18, 151, 1, 0, 0, 0, 20, 160, 1, 0, 0, 0, 22, 170,
		1, 0, 0, 0, 24, 175, 1, 0, 0, 0, 26, 183, 1, 0, 0, 0, 28, 192, 1, 0, 0,
		0, 30, 195, 1, 0, 0, 0, 32, 198, 1, 0, 0, 0, 34, 208, 1, 0, 0, 0, 36, 221,
		1, 0, 0, 0, 38, 225, 1, 0, 0, 0, 40, 238, 1, 0, 0, 0, 42, 281, 1, 0, 0,
		0, 44, 283, 1, 0, 0, 0, 46, 296, 1, 0, 0, 0, 48, 334, 1, 0, 0, 0, 50, 354,
		1, 0, 0, 0, 52, 372, 1, 0, 0, 0, 54, 389, 1, 0, 0, 0, 56, 60, 3, 30, 15,
		0, 57, 60, 3, 16, 8, 0, 58, 60, 3, 24, 12, 0, 59, 56, 1, 0, 0, 0, 59, 57,
		1, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0,
		61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 5,
		0, 0, 1, 65, 1, 1, 0, 0, 0, 66, 67, 5, 41, 0, 0, 67, 3, 1, 0, 0, 0, 68,
		71, 3, 2, 1, 0, 69, 70, 5, 34, 0, 0, 70, 72, 3, 10, 5, 0, 71, 69, 1, 0,
		0, 0, 71, 72, 1, 0, 0, 0, 72, 5, 1, 0, 0, 0, 73, 74, 5, 7, 0, 0, 74, 88,
		3, 6, 3, 0, 75, 76, 5, 15, 0, 0, 76, 77, 3, 6, 3, 0, 77, 78, 5, 18, 0,
		0, 78, 88, 1, 0, 0, 0, 79, 80, 5, 4, 0, 0, 80, 81, 5, 34, 0, 0, 81, 82,
		3, 10, 5, 0, 82, 84, 3, 12, 6, 0, 83, 85, 3, 6, 3, 0, 84, 83, 1, 0, 0,
		0, 84, 85, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 88, 3, 4, 2, 0, 87, 73,
		1, 0, 0, 0, 87, 75, 1, 0, 0, 0, 87, 79, 1, 0, 0, 0, 87, 86, 1, 0, 0, 0,
		88, 7, 1, 0, 0, 0, 89, 92, 3, 2, 1, 0, 90, 92, 5, 40, 0, 0, 91, 89, 1,
		0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 9, 1, 0, 0, 0, 93, 94, 5, 15, 0, 0, 94,
		99, 3, 8, 4, 0, 95, 96, 5, 35, 0, 0, 96, 98, 3, 8, 4, 0, 97, 95, 1, 0,
		0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100,
		102, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 103, 5, 18, 0, 0, 103, 119,
		1, 0, 0, 0, 104, 117, 3, 8, 4, 0, 105, 106, 5, 14, 0, 0, 106, 111, 3, 8,
		4, 0, 107, 108, 5, 35, 0, 0, 108, 110, 3, 8, 4, 0, 109, 107, 1, 0, 0, 0,
		110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112,
		114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5, 17, 0, 0, 115, 117,
		1, 0, 0, 0, 116, 104, 1, 0, 0, 0, 116, 105, 1, 0, 0, 0, 117, 119, 1, 0,
		0, 0, 118, 93, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 11, 1, 0, 0, 0, 120,
		129, 5, 14, 0, 0, 121, 126, 3, 6, 3, 0, 122, 123, 5, 35, 0, 0, 123, 125,
		3, 6, 3, 0, 124, 122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0,
		0, 0, 126, 127, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0,
		129, 121, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		132, 5, 17, 0, 0, 132, 13, 1, 0, 0, 0, 133, 134, 5, 4, 0, 0, 134, 135,
		5, 34, 0, 0, 135, 136, 3, 10, 5, 0, 136, 138, 3, 34, 17, 0, 137, 139, 3,
		6, 3, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 15, 1, 0, 0,
		0, 140, 143, 5, 1, 0, 0, 141, 142, 5, 34, 0, 0, 142, 144, 3, 10, 5, 0,
		143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		147, 3, 2, 1, 0, 146, 148, 3, 18, 9, 0, 147, 146, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150, 3, 20, 10, 0, 150, 17, 1, 0,
		0, 0, 151, 152, 5, 3, 0, 0, 152, 157, 3, 4, 2, 0, 153, 154, 5, 35, 0, 0,
		154, 156, 3, 4, 2, 0, 155, 153, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157,
		155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 19, 1, 0, 0, 0, 159, 157, 1,
		0, 0, 0, 160, 165, 5, 16, 0, 0, 161, 164, 3, 22, 11, 0, 162, 164, 3, 30,
		15, 0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0,
		165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 168, 169, 5, 19, 0, 0, 169, 21, 1, 0, 0, 0, 170, 171,
		3, 2, 1, 0, 171, 172, 5, 37, 0, 0, 172, 173, 3, 6, 3, 0, 173, 174, 5, 43,
		0, 0, 174, 23, 1, 0, 0, 0, 175, 178, 5, 2, 0, 0, 176, 177, 5, 34, 0, 0,
		177, 179, 3, 10, 5, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 181, 3, 2, 1, 0, 181, 182, 3, 26, 13, 0, 182, 25,
		1, 0, 0, 0, 183, 187, 5, 16, 0, 0, 184, 186, 3, 28, 14, 0, 185, 184, 1,
		0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0,
		0, 188, 190, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 19, 0, 0, 191,
		27, 1, 0, 0, 0, 192, 193, 3, 32, 16, 0, 193, 194, 5, 43, 0, 0, 194, 29,
		1, 0, 0, 0, 195, 196, 3, 32, 16, 0, 196, 197, 3, 40, 20, 0, 197, 31, 1,
		0, 0, 0, 198, 201, 5, 4, 0, 0, 199, 200, 5, 34, 0, 0, 200, 202, 3, 10,
		5, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0,
		203, 204, 3, 2, 1, 0, 204, 206, 3, 34, 17, 0, 205, 207, 3, 6, 3, 0, 206,
		205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 33, 1, 0, 0, 0, 208, 217, 5,
		14, 0, 0, 209, 214, 3, 36, 18, 0, 210, 211, 5, 35, 0, 0, 211, 213, 3, 36,
		18, 0, 212, 210, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0,
		214, 215, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217,
		209, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220,
		5, 17, 0, 0, 220, 35, 1, 0, 0, 0, 221, 222, 3, 2, 1, 0, 222, 223, 5, 37,
		0, 0, 223, 224, 3, 6, 3, 0, 224, 37, 1, 0, 0, 0, 225, 234, 5, 14, 0, 0,
		226, 231, 3, 48, 24, 0, 227, 228, 5, 35, 0, 0, 228, 230, 3, 48, 24, 0,
		229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231,
		232, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 226,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 5, 17,
		0, 0, 237, 39, 1, 0, 0, 0, 238, 242, 5, 16, 0, 0, 239, 241, 3, 42, 21,
		0, 240, 239, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242,
		243, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 246,
		5, 19, 0, 0, 246, 41, 1, 0, 0, 0, 247, 248, 5, 6, 0, 0, 248, 251, 3, 2,
		1, 0, 249, 250, 5, 37, 0, 0, 250, 252, 3, 6, 3, 0, 251, 249, 1, 0, 0, 0,
		251, 252, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 5, 33, 0, 0, 254,
		255, 3, 48, 24, 0, 255, 256, 5, 43, 0, 0, 256, 282, 1, 0, 0, 0, 257, 258,
		5, 9, 0, 0, 258, 259, 3, 48, 24, 0, 259, 262, 3, 40, 20, 0, 260, 261, 5,
		10, 0, 0, 261, 263, 3, 40, 20, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0,
		0, 0, 263, 282, 1, 0, 0, 0, 264, 265, 5, 11, 0, 0, 265, 266, 3, 48, 24,
		0, 266, 267, 3, 40, 20, 0, 267, 282, 1, 0, 0, 0, 268, 270, 5, 5, 0, 0,
		269, 271, 3, 48, 24, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271,
		272, 1, 0, 0, 0, 272, 282, 5, 43, 0, 0, 273, 274, 3, 44, 22, 0, 274, 275,
		5, 33, 0, 0, 275, 276, 3, 48, 24, 0, 276, 277, 5, 43, 0, 0, 277, 282, 1,
		0, 0, 0, 278, 279, 3, 48, 24, 0, 279, 280, 5, 43, 0, 0, 280, 282, 1, 0,
		0, 0, 281, 247, 1, 0, 0, 0, 281, 257, 1, 0, 0, 0, 281, 264, 1, 0, 0, 0,
		281, 268, 1, 0, 0, 0, 281, 273, 1, 0, 0, 0, 281, 278, 1, 0, 0, 0, 282,
		43, 1, 0, 0, 0, 283, 287, 3, 2, 1, 0, 284, 286, 3, 46, 23, 0, 285, 284,
		1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287, 288, 1, 0,
		0, 0, 288, 45, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 291, 5, 36, 0, 0,
		291, 297, 3, 2, 1, 0, 292, 293, 5, 15, 0, 0, 293, 294, 3, 48, 24, 0, 294,
		295, 5, 18, 0, 0, 295, 297, 1, 0, 0, 0, 296, 290, 1, 0, 0, 0, 296, 292,
		1, 0, 0, 0, 297, 47, 1, 0, 0, 0, 298, 299, 6, 24, -1, 0, 299, 302, 3, 54,
		27, 0, 300, 301, 5, 34, 0, 0, 301, 303, 3, 10, 5, 0, 302, 300, 1, 0, 0,
		0, 302, 303, 1, 0, 0, 0, 303, 335, 1, 0, 0, 0, 304, 305, 3, 14, 7, 0, 305,
		306, 3, 40, 20, 0, 306, 335, 1, 0, 0, 0, 307, 308, 3, 4, 2, 0, 308, 309,
		3, 50, 25, 0, 309, 335, 1, 0, 0, 0, 310, 319, 5, 15, 0, 0, 311, 316, 3,
		48, 24, 0, 312, 313, 5, 35, 0, 0, 313, 315, 3, 48, 24, 0, 314, 312, 1,
		0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0,
		0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 311, 1, 0, 0, 0, 319,
		320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 335, 5, 18, 0, 0, 322, 335,
		3, 52, 26, 0, 323, 324, 3, 10, 5, 0, 324, 325, 5, 38, 0, 0, 325, 326, 3,
		10, 5, 0, 326, 327, 3, 48, 24, 3, 327, 335, 1, 0, 0, 0, 328, 329, 5, 8,
		0, 0, 329, 335, 3, 48, 24, 2, 330, 331, 5, 14, 0, 0, 331, 332, 3, 48, 24,
		0, 332, 333, 5, 17, 0, 0, 333, 335, 1, 0, 0, 0, 334, 298, 1, 0, 0, 0, 334,
		304, 1, 0, 0, 0, 334, 307, 1, 0, 0, 0, 334, 310, 1, 0, 0, 0, 334, 322,
		1, 0, 0, 0, 334, 323, 1, 0, 0, 0, 334, 328, 1, 0, 0, 0, 334, 330, 1, 0,
		0, 0, 335, 351, 1, 0, 0, 0, 336, 337, 10, 12, 0, 0, 337, 338, 7, 0, 0,
		0, 338, 350, 3, 48, 24, 13, 339, 340, 10, 8, 0, 0, 340, 350, 3, 38, 19,
		0, 341, 342, 10, 7, 0, 0, 342, 343, 5, 36, 0, 0, 343, 350, 3, 2, 1, 0,
		344, 345, 10, 6, 0, 0, 345, 346, 5, 15, 0, 0, 346, 347, 3, 48, 24, 0, 347,
		348, 5, 18, 0, 0, 348, 350, 1, 0, 0, 0, 349, 336, 1, 0, 0, 0, 349, 339,
		1, 0, 0, 0, 349, 341, 1, 0, 0, 0, 349, 344, 1, 0, 0, 0, 350, 353, 1, 0,
		0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 49, 1, 0, 0, 0,
		353, 351, 1, 0, 0, 0, 354, 368, 5, 16, 0, 0, 355, 356, 3, 2, 1, 0, 356,
		357, 5, 37, 0, 0, 357, 365, 3, 48, 24, 0, 358, 359, 5, 35, 0, 0, 359, 360,
		3, 2, 1, 0, 360, 361, 5, 37, 0, 0, 361, 362, 3, 48, 24, 0, 362, 364, 1,
		0, 0, 0, 363, 358, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365, 363, 1, 0, 0,
		0, 365, 366, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 368,
		355, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371,
		5, 19, 0, 0, 371, 51, 1, 0, 0, 0, 372, 375, 3, 2, 1, 0, 373, 374, 5, 34,
		0, 0, 374, 376, 3, 10, 5, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0,
		376, 53, 1, 0, 0, 0, 377, 378, 5, 42, 0, 0, 378, 379, 5, 36, 0, 0, 379,
		385, 5, 42, 0, 0, 380, 381, 5, 42, 0, 0, 381, 385, 5, 36, 0, 0, 382, 383,
		5, 36, 0, 0, 383, 385, 5, 42, 0, 0, 384, 377, 1, 0, 0, 0, 384, 380, 1,
		0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 390, 1, 0, 0, 0, 386, 390, 5, 42, 0,
		0, 387, 390, 5, 39, 0, 0, 388, 390, 7, 1, 0, 0, 389, 384, 1, 0, 0, 0, 389,
		386, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 388, 1, 0, 0, 0, 390, 55, 1,
		0, 0, 0, 44, 59, 61, 71, 84, 87, 91, 99, 111, 116, 118, 126, 129, 138,
		143, 147, 157, 163, 165, 178, 187, 201, 206, 214, 217, 231, 234, 242, 251,
		262, 270, 281, 287, 296, 302, 316, 319, 334, 349, 351, 365, 368, 375, 384,
		389,
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
	TempoParserIMPLEMENTS    = 3
	TempoParserFUNC          = 4
	TempoParserRETURN        = 5
	TempoParserLET           = 6
	TempoParserASYNC         = 7
	TempoParserAWAIT         = 8
	TempoParserIF            = 9
	TempoParserELSE          = 10
	TempoParserWHILE         = 11
	TempoParserTRUE          = 12
	TempoParserFALSE         = 13
	TempoParserLPAREN        = 14
	TempoParserLSQUARE       = 15
	TempoParserLCURLY        = 16
	TempoParserRPAREN        = 17
	TempoParserRSQUARE       = 18
	TempoParserRCURLY        = 19
	TempoParserPLUS          = 20
	TempoParserMINUS         = 21
	TempoParserMULTIPLY      = 22
	TempoParserDIVIDE        = 23
	TempoParserMODULO        = 24
	TempoParserEQUAL         = 25
	TempoParserNOT_EQUAL     = 26
	TempoParserLESS          = 27
	TempoParserLESS_EQ       = 28
	TempoParserGREATER       = 29
	TempoParserGREATER_EQ    = 30
	TempoParserAND           = 31
	TempoParserOR            = 32
	TempoParserIS            = 33
	TempoParserROLE_AT       = 34
	TempoParserCOMMA         = 35
	TempoParserDOT           = 36
	TempoParserCOLON         = 37
	TempoParserCOM           = 38
	TempoParserSTRING        = 39
	TempoParserUNDERSCORE    = 40
	TempoParserID            = 41
	TempoParserNUMBER        = 42
	TempoParserEND           = 43
	TempoParserWHITESPACE    = 44
	TempoParserLINE_COMMENT  = 45
	TempoParserBLOCK_COMMENT = 46
)

// TempoParser rules.
const (
	TempoParserRULE_sourceFile           = 0
	TempoParserRULE_ident                = 1
	TempoParserRULE_roleIdent            = 2
	TempoParserRULE_valueType            = 3
	TempoParserRULE_role                 = 4
	TempoParserRULE_roleType             = 5
	TempoParserRULE_closureParamList     = 6
	TempoParserRULE_closureSig           = 7
	TempoParserRULE_struct               = 8
	TempoParserRULE_structImplements     = 9
	TempoParserRULE_structBody           = 10
	TempoParserRULE_structField          = 11
	TempoParserRULE_interface            = 12
	TempoParserRULE_interfaceMethodsList = 13
	TempoParserRULE_interfaceMethod      = 14
	TempoParserRULE_func                 = 15
	TempoParserRULE_funcSig              = 16
	TempoParserRULE_funcParamList        = 17
	TempoParserRULE_funcParam            = 18
	TempoParserRULE_funcArgList          = 19
	TempoParserRULE_scope                = 20
	TempoParserRULE_stmt                 = 21
	TempoParserRULE_assignExpr           = 22
	TempoParserRULE_assignSpecifier      = 23
	TempoParserRULE_expr                 = 24
	TempoParserRULE_exprStructField      = 25
	TempoParserRULE_identAccess          = 26
	TempoParserRULE_literal              = 27
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&22) != 0 {
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserFUNC:
			{
				p.SetState(56)
				p.Func_()
			}

		case TempoParserSTRUCT:
			{
				p.SetState(57)
				p.Struct_()
			}

		case TempoParserINTERFACE:
			{
				p.SetState(58)
				p.Interface_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
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
		p.SetState(66)
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

// IRoleIdentContext is an interface to support dynamic dispatch.
type IRoleIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext

	// IsRoleIdentContext differentiates from other interfaces.
	IsRoleIdentContext()
}

type RoleIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoleIdentContext() *RoleIdentContext {
	var p = new(RoleIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_roleIdent
	return p
}

func InitEmptyRoleIdentContext(p *RoleIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_roleIdent
}

func (*RoleIdentContext) IsRoleIdentContext() {}

func NewRoleIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoleIdentContext {
	var p = new(RoleIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_roleIdent

	return p
}

func (s *RoleIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *RoleIdentContext) Ident() IIdentContext {
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

func (s *RoleIdentContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *RoleIdentContext) RoleType() IRoleTypeContext {
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

func (s *RoleIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoleIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoleIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterRoleIdent(s)
	}
}

func (s *RoleIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitRoleIdent(s)
	}
}

func (s *RoleIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitRoleIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) RoleIdent() (localctx IRoleIdentContext) {
	localctx = NewRoleIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TempoParserRULE_roleIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Ident()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserROLE_AT {
		{
			p.SetState(69)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.RoleType()
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

func (s *NamedTypeContext) RoleIdent() IRoleIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleIdentContext)
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
	p.EnterRule(localctx, 6, TempoParserRULE_valueType)
	var _la int

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserASYNC:
		localctx = NewAsyncTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(TempoParserASYNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)

			var _x = p.ValueType()

			localctx.(*AsyncTypeContext).inner = _x
		}

	case TempoParserLSQUARE:
		localctx = NewListTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)

			var _x = p.ValueType()

			localctx.(*ListTypeContext).inner = _x
		}
		{
			p.SetState(77)
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
			p.SetState(79)
			p.Match(TempoParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.RoleType()
		}
		{
			p.SetState(82)

			var _x = p.ClosureParamList()

			localctx.(*ClosureTypeContext).params = _x
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023288464) != 0 {
			{
				p.SetState(83)

				var _x = p.ValueType()

				localctx.(*ClosureTypeContext).returnType = _x
			}

		}

	case TempoParserID:
		localctx = NewNamedTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)
			p.RoleIdent()
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

// IRoleContext is an interface to support dynamic dispatch.
type IRoleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	UNDERSCORE() antlr.TerminalNode

	// IsRoleContext differentiates from other interfaces.
	IsRoleContext()
}

type RoleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoleContext() *RoleContext {
	var p = new(RoleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_role
	return p
}

func InitEmptyRoleContext(p *RoleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_role
}

func (*RoleContext) IsRoleContext() {}

func NewRoleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoleContext {
	var p = new(RoleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_role

	return p
}

func (s *RoleContext) GetParser() antlr.Parser { return s.parser }

func (s *RoleContext) Ident() IIdentContext {
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

func (s *RoleContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(TempoParserUNDERSCORE, 0)
}

func (s *RoleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterRole(s)
	}
}

func (s *RoleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitRole(s)
	}
}

func (s *RoleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitRole(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) Role() (localctx IRoleContext) {
	localctx = NewRoleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TempoParserRULE_role)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(89)
			p.Ident()
		}

	case TempoParserUNDERSCORE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(TempoParserUNDERSCORE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (s *RoleTypeSharedContext) AllRole() []IRoleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRoleContext); ok {
			len++
		}
	}

	tst := make([]IRoleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRoleContext); ok {
			tst[i] = t.(IRoleContext)
			i++
		}
	}

	return tst
}

func (s *RoleTypeSharedContext) Role(i int) IRoleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleContext); ok {
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

	return t.(IRoleContext)
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

func (s *RoleTypeNormalContext) AllRole() []IRoleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRoleContext); ok {
			len++
		}
	}

	tst := make([]IRoleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRoleContext); ok {
			tst[i] = t.(IRoleContext)
			i++
		}
	}

	return tst
}

func (s *RoleTypeNormalContext) Role(i int) IRoleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleContext); ok {
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

	return t.(IRoleContext)
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
	p.EnterRule(localctx, 10, TempoParserRULE_roleType)
	var _la int

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserLSQUARE:
		localctx = NewRoleTypeSharedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Role()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(95)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(96)
				p.Role()
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(102)
			p.Match(TempoParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TempoParserLPAREN, TempoParserUNDERSCORE, TempoParserID:
		localctx = NewRoleTypeNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserUNDERSCORE, TempoParserID:
			{
				p.SetState(104)
				p.Role()
			}

		case TempoParserLPAREN:
			{
				p.SetState(105)
				p.Match(TempoParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(106)
				p.Role()
			}
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == TempoParserCOMMA {
				{
					p.SetState(107)
					p.Match(TempoParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(108)
					p.Role()
				}

				p.SetState(113)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(114)
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
	p.EnterRule(localctx, 12, TempoParserRULE_closureParamList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023288464) != 0 {
		{
			p.SetState(121)
			p.ValueType()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(122)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(123)
				p.ValueType()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(131)
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
	p.EnterRule(localctx, 14, TempoParserRULE_closureSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(TempoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.RoleType()
	}
	{
		p.SetState(136)

		var _x = p.FuncParamList()

		localctx.(*ClosureSigContext).params = _x
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023288464) != 0 {
		{
			p.SetState(137)

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

	// GetBody returns the body rule contexts.
	GetBody() IStructBodyContext

	// SetBody sets the body rule contexts.
	SetBody(IStructBodyContext)

	// Getter signatures
	STRUCT() antlr.TerminalNode
	Ident() IIdentContext
	StructBody() IStructBodyContext
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
	StructImplements() IStructImplementsContext

	// IsStructContext differentiates from other interfaces.
	IsStructContext()
}

type StructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	body   IStructBodyContext
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

func (s *StructContext) GetBody() IStructBodyContext { return s.body }

func (s *StructContext) SetBody(v IStructBodyContext) { s.body = v }

func (s *StructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(TempoParserSTRUCT, 0)
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

func (s *StructContext) StructBody() IStructBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructBodyContext)
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

func (s *StructContext) StructImplements() IStructImplementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructImplementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructImplementsContext)
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
	p.EnterRule(localctx, 16, TempoParserRULE_struct)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(TempoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserROLE_AT {
		{
			p.SetState(141)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.RoleType()
		}

	}
	{
		p.SetState(145)
		p.Ident()
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserIMPLEMENTS {
		{
			p.SetState(146)
			p.StructImplements()
		}

	}
	{
		p.SetState(149)

		var _x = p.StructBody()

		localctx.(*StructContext).body = _x
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

// IStructImplementsContext is an interface to support dynamic dispatch.
type IStructImplementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPLEMENTS() antlr.TerminalNode
	AllRoleIdent() []IRoleIdentContext
	RoleIdent(i int) IRoleIdentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructImplementsContext differentiates from other interfaces.
	IsStructImplementsContext()
}

type StructImplementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructImplementsContext() *StructImplementsContext {
	var p = new(StructImplementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structImplements
	return p
}

func InitEmptyStructImplementsContext(p *StructImplementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structImplements
}

func (*StructImplementsContext) IsStructImplementsContext() {}

func NewStructImplementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructImplementsContext {
	var p = new(StructImplementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_structImplements

	return p
}

func (s *StructImplementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructImplementsContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(TempoParserIMPLEMENTS, 0)
}

func (s *StructImplementsContext) AllRoleIdent() []IRoleIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRoleIdentContext); ok {
			len++
		}
	}

	tst := make([]IRoleIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRoleIdentContext); ok {
			tst[i] = t.(IRoleIdentContext)
			i++
		}
	}

	return tst
}

func (s *StructImplementsContext) RoleIdent(i int) IRoleIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleIdentContext); ok {
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

	return t.(IRoleIdentContext)
}

func (s *StructImplementsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *StructImplementsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *StructImplementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructImplementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructImplementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStructImplements(s)
	}
}

func (s *StructImplementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStructImplements(s)
	}
}

func (s *StructImplementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStructImplements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) StructImplements() (localctx IStructImplementsContext) {
	localctx = NewStructImplementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TempoParserRULE_structImplements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(TempoParserIMPLEMENTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.RoleIdent()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TempoParserCOMMA {
		{
			p.SetState(153)
			p.Match(TempoParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.RoleIdent()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IStructBodyContext is an interface to support dynamic dispatch.
type IStructBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCURLY() antlr.TerminalNode
	RCURLY() antlr.TerminalNode
	AllStructField() []IStructFieldContext
	StructField(i int) IStructFieldContext
	AllFunc_() []IFuncContext
	Func_(i int) IFuncContext

	// IsStructBodyContext differentiates from other interfaces.
	IsStructBodyContext()
}

type StructBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructBodyContext() *StructBodyContext {
	var p = new(StructBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structBody
	return p
}

func InitEmptyStructBodyContext(p *StructBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_structBody
}

func (*StructBodyContext) IsStructBodyContext() {}

func NewStructBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructBodyContext {
	var p = new(StructBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_structBody

	return p
}

func (s *StructBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructBodyContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserLCURLY, 0)
}

func (s *StructBodyContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(TempoParserRCURLY, 0)
}

func (s *StructBodyContext) AllStructField() []IStructFieldContext {
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

func (s *StructBodyContext) StructField(i int) IStructFieldContext {
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

func (s *StructBodyContext) AllFunc_() []IFuncContext {
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

func (s *StructBodyContext) Func_(i int) IFuncContext {
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

func (s *StructBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterStructBody(s)
	}
}

func (s *StructBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitStructBody(s)
	}
}

func (s *StructBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitStructBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) StructBody() (localctx IStructBodyContext) {
	localctx = NewStructBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TempoParserRULE_structBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TempoParserFUNC || _la == TempoParserID {
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserID:
			{
				p.SetState(161)
				p.StructField()
			}

		case TempoParserFUNC:
			{
				p.SetState(162)
				p.Func_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
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
	END() antlr.TerminalNode

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

func (s *StructFieldContext) END() antlr.TerminalNode {
	return s.GetToken(TempoParserEND, 0)
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
	p.EnterRule(localctx, 22, TempoParserRULE_structField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Ident()
	}
	{
		p.SetState(171)
		p.Match(TempoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.ValueType()
	}
	{
		p.SetState(173)
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

// IInterfaceContext is an interface to support dynamic dispatch.
type IInterfaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	Ident() IIdentContext
	InterfaceMethodsList() IInterfaceMethodsListContext
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext

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
	p.EnterRule(localctx, 24, TempoParserRULE_interface)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(TempoParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserROLE_AT {
		{
			p.SetState(176)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.RoleType()
		}

	}
	{
		p.SetState(180)
		p.Ident()
	}
	{
		p.SetState(181)
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
	p.EnterRule(localctx, 26, TempoParserRULE_interfaceMethodsList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TempoParserFUNC {
		{
			p.SetState(184)
			p.InterfaceMethod()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
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
	p.EnterRule(localctx, 28, TempoParserRULE_interfaceMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.FuncSig()
	}
	{
		p.SetState(193)
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
	p.EnterRule(localctx, 30, TempoParserRULE_func)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.FuncSig()
	}
	{
		p.SetState(196)
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
	Ident() IIdentContext
	FuncParamList() IFuncParamListContext
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
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
	p.EnterRule(localctx, 32, TempoParserRULE_funcSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(TempoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserROLE_AT {
		{
			p.SetState(199)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.RoleType()
		}

	}
	{
		p.SetState(203)

		var _x = p.Ident()

		localctx.(*FuncSigContext).name = _x
	}
	{
		p.SetState(204)

		var _x = p.FuncParamList()

		localctx.(*FuncSigContext).params = _x
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023288464) != 0 {
		{
			p.SetState(205)

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
	p.EnterRule(localctx, 34, TempoParserRULE_funcParamList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(209)
			p.FuncParam()
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(210)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(211)
				p.FuncParam()
			}

			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(219)
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
	p.EnterRule(localctx, 36, TempoParserRULE_funcParam)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Ident()
	}
	{
		p.SetState(222)
		p.Match(TempoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
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
	p.EnterRule(localctx, 38, TempoParserRULE_funcArgList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8315056746768) != 0 {
		{
			p.SetState(226)
			p.expr(0)
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(227)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(228)
				p.expr(0)
			}

			p.SetState(233)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 40, TempoParserRULE_scope)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8315056749424) != 0 {
		{
			p.SetState(239)
			p.Stmt()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(245)
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

func (s *StmtAssignContext) AssignExpr() IAssignExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignExprContext)
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
	p.EnterRule(localctx, 42, TempoParserRULE_stmt)
	var _la int

	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmtVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Match(TempoParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Ident()
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TempoParserCOLON {
			{
				p.SetState(249)
				p.Match(TempoParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(250)
				p.ValueType()
			}

		}
		{
			p.SetState(253)
			p.Match(TempoParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.expr(0)
		}
		{
			p.SetState(255)
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
			p.SetState(257)
			p.Match(TempoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.expr(0)
		}
		{
			p.SetState(259)

			var _x = p.Scope()

			localctx.(*StmtIfContext).thenScope = _x
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TempoParserELSE {
			{
				p.SetState(260)
				p.Match(TempoParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(261)

				var _x = p.Scope()

				localctx.(*StmtIfContext).elseScope = _x
			}

		}

	case 3:
		localctx = NewStmtWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(TempoParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.expr(0)
		}
		{
			p.SetState(266)
			p.Scope()
		}

	case 4:
		localctx = NewStmtReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(268)
			p.Match(TempoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8315056746768) != 0 {
			{
				p.SetState(269)
				p.expr(0)
			}

		}
		{
			p.SetState(272)
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
			p.SetState(273)
			p.AssignExpr()
		}
		{
			p.SetState(274)
			p.Match(TempoParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.expr(0)
		}
		{
			p.SetState(276)
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
			p.SetState(278)
			p.expr(0)
		}
		{
			p.SetState(279)
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

// IAssignExprContext is an interface to support dynamic dispatch.
type IAssignExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	AllAssignSpecifier() []IAssignSpecifierContext
	AssignSpecifier(i int) IAssignSpecifierContext

	// IsAssignExprContext differentiates from other interfaces.
	IsAssignExprContext()
}

type AssignExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignExprContext() *AssignExprContext {
	var p = new(AssignExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_assignExpr
	return p
}

func InitEmptyAssignExprContext(p *AssignExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_assignExpr
}

func (*AssignExprContext) IsAssignExprContext() {}

func NewAssignExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExprContext {
	var p = new(AssignExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_assignExpr

	return p
}

func (s *AssignExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExprContext) Ident() IIdentContext {
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

func (s *AssignExprContext) AllAssignSpecifier() []IAssignSpecifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignSpecifierContext); ok {
			len++
		}
	}

	tst := make([]IAssignSpecifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignSpecifierContext); ok {
			tst[i] = t.(IAssignSpecifierContext)
			i++
		}
	}

	return tst
}

func (s *AssignExprContext) AssignSpecifier(i int) IAssignSpecifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignSpecifierContext); ok {
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

	return t.(IAssignSpecifierContext)
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) AssignExpr() (localctx IAssignExprContext) {
	localctx = NewAssignExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TempoParserRULE_assignExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Ident()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TempoParserLSQUARE || _la == TempoParserDOT {
		{
			p.SetState(284)
			p.AssignSpecifier()
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAssignSpecifierContext is an interface to support dynamic dispatch.
type IAssignSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignSpecifierContext differentiates from other interfaces.
	IsAssignSpecifierContext()
}

type AssignSpecifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignSpecifierContext() *AssignSpecifierContext {
	var p = new(AssignSpecifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_assignSpecifier
	return p
}

func InitEmptyAssignSpecifierContext(p *AssignSpecifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TempoParserRULE_assignSpecifier
}

func (*AssignSpecifierContext) IsAssignSpecifierContext() {}

func NewAssignSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignSpecifierContext {
	var p = new(AssignSpecifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_assignSpecifier

	return p
}

func (s *AssignSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignSpecifierContext) CopyAll(ctx *AssignSpecifierContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignFieldContext struct {
	AssignSpecifierContext
}

func NewAssignFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignFieldContext {
	var p = new(AssignFieldContext)

	InitEmptyAssignSpecifierContext(&p.AssignSpecifierContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignSpecifierContext))

	return p
}

func (s *AssignFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignFieldContext) DOT() antlr.TerminalNode {
	return s.GetToken(TempoParserDOT, 0)
}

func (s *AssignFieldContext) Ident() IIdentContext {
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

func (s *AssignFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterAssignField(s)
	}
}

func (s *AssignFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitAssignField(s)
	}
}

func (s *AssignFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitAssignField(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignIndexContext struct {
	AssignSpecifierContext
}

func NewAssignIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignIndexContext {
	var p = new(AssignIndexContext)

	InitEmptyAssignSpecifierContext(&p.AssignSpecifierContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignSpecifierContext))

	return p
}

func (s *AssignIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignIndexContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserLSQUARE, 0)
}

func (s *AssignIndexContext) Expr() IExprContext {
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

func (s *AssignIndexContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserRSQUARE, 0)
}

func (s *AssignIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterAssignIndex(s)
	}
}

func (s *AssignIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitAssignIndex(s)
	}
}

func (s *AssignIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitAssignIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) AssignSpecifier() (localctx IAssignSpecifierContext) {
	localctx = NewAssignSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TempoParserRULE_assignSpecifier)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserDOT:
		localctx = NewAssignFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Match(TempoParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Ident()
		}

	case TempoParserLSQUARE:
		localctx = NewAssignIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.expr(0)
		}
		{
			p.SetState(294)
			p.Match(TempoParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (s *ExprStructContext) RoleIdent() IRoleIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoleIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoleIdentContext)
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

type ExprListContext struct {
	ExprContext
}

func NewExprListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprListContext {
	var p = new(ExprListContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserLSQUARE, 0)
}

func (s *ExprListContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserRSQUARE, 0)
}

func (s *ExprListContext) AllExpr() []IExprContext {
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

func (s *ExprListContext) Expr(i int) IExprContext {
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

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TempoParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TempoParserCOMMA, i)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (s *ExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprList(s)

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

type ExprIndexContext struct {
	ExprContext
	baseExpr  IExprContext
	indexExpr IExprContext
}

func NewExprIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIndexContext {
	var p = new(ExprIndexContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprIndexContext) GetBaseExpr() IExprContext { return s.baseExpr }

func (s *ExprIndexContext) GetIndexExpr() IExprContext { return s.indexExpr }

func (s *ExprIndexContext) SetBaseExpr(v IExprContext) { s.baseExpr = v }

func (s *ExprIndexContext) SetIndexExpr(v IExprContext) { s.indexExpr = v }

func (s *ExprIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIndexContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserLSQUARE, 0)
}

func (s *ExprIndexContext) RSQUARE() antlr.TerminalNode {
	return s.GetToken(TempoParserRSQUARE, 0)
}

func (s *ExprIndexContext) AllExpr() []IExprContext {
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

func (s *ExprIndexContext) Expr(i int) IExprContext {
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

func (s *ExprIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprIndex(s)
	}
}

func (s *ExprIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprIndex(s)
	}
}

func (s *ExprIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprIndex(s)

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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, TempoParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprPrimitiveContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(299)
			p.Literal()
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(300)
				p.Match(TempoParserROLE_AT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(301)
				p.RoleType()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		localctx = NewExprClosureContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(304)
			p.ClosureSig()
		}
		{
			p.SetState(305)
			p.Scope()
		}

	case 3:
		localctx = NewExprStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)
			p.RoleIdent()
		}
		{
			p.SetState(308)
			p.ExprStructField()
		}

	case 4:
		localctx = NewExprListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(310)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8315056746768) != 0 {
			{
				p.SetState(311)
				p.expr(0)
			}
			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == TempoParserCOMMA {
				{
					p.SetState(312)
					p.Match(TempoParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(313)
					p.expr(0)
				}

				p.SetState(318)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(321)
			p.Match(TempoParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExprIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(322)
			p.IdentAccess()
		}

	case 6:
		localctx = NewExprComContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(323)

			var _x = p.RoleType()

			localctx.(*ExprComContext).sender = _x
		}
		{
			p.SetState(324)
			p.Match(TempoParserCOM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)

			var _x = p.RoleType()

			localctx.(*ExprComContext).receiver = _x
		}
		{
			p.SetState(326)
			p.expr(3)
		}

	case 7:
		localctx = NewExprAwaitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(328)
			p.Match(TempoParserAWAIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.expr(2)
		}

	case 8:
		localctx = NewExprGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(330)
			p.Match(TempoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.expr(0)
		}
		{
			p.SetState(332)
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
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(349)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprBinOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprBinOpContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(337)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8588886016) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(338)

					var _x = p.expr(13)

					localctx.(*ExprBinOpContext).rhs = _x
				}

			case 2:
				localctx = NewExprCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(340)
					p.FuncArgList()
				}

			case 3:
				localctx = NewExprFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(341)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(342)
					p.Match(TempoParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(343)
					p.Ident()
				}

			case 4:
				localctx = NewExprIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprIndexContext).baseExpr = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(345)
					p.Match(TempoParserLSQUARE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(346)

					var _x = p.expr(0)

					localctx.(*ExprIndexContext).indexExpr = _x
				}
				{
					p.SetState(347)
					p.Match(TempoParserRSQUARE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 50, TempoParserRULE_exprStructField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(355)
			p.Ident()
		}
		{
			p.SetState(356)
			p.Match(TempoParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.expr(0)
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(358)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(359)
				p.Ident()
			}
			{
				p.SetState(360)
				p.Match(TempoParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(361)
				p.expr(0)
			}

			p.SetState(367)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(370)
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
	p.EnterRule(localctx, 52, TempoParserRULE_identAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Ident()
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(373)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
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
	p.EnterRule(localctx, 54, TempoParserRULE_literal)
	var _la int

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(377)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(378)
				p.Match(TempoParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(379)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(380)
				p.Match(TempoParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(381)
				p.Match(TempoParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 3:
			{
				p.SetState(382)
				p.Match(TempoParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(383)
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
			p.SetState(386)
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
			p.SetState(387)
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
			p.SetState(388)
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
	case 24:
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
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
