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
		"sourceFile", "ident", "valueType", "roleType", "struct", "structFieldList",
		"structField", "interface", "interfaceMethodsList", "interfaceMethod",
		"func", "funcSig", "funcParamList", "funcParam", "funcArgList", "scope",
		"stmt", "expr", "exprStructField", "identAccess",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 266, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2,
		3, 2, 54, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 64,
		8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 76, 8, 3, 10, 3, 12, 3, 79, 9, 3, 1, 3, 1, 3, 3, 3, 83, 8, 3, 3,
		3, 85, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 97, 8, 5, 10, 5, 12, 5, 100, 9, 5, 3, 5, 102, 8, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 5,
		8, 118, 8, 8, 10, 8, 12, 8, 121, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 137,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 143, 8, 12, 10, 12, 12, 12, 146,
		9, 12, 3, 12, 148, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 5, 14, 160, 8, 14, 10, 14, 12, 14, 163, 9, 14,
		3, 14, 165, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 171, 8, 15, 10, 15,
		12, 15, 174, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 191, 8, 16,
		1, 16, 1, 16, 3, 16, 195, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 206, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 228, 8, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 238, 8, 17, 10, 17, 12,
		17, 241, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 5, 18, 252, 8, 18, 10, 18, 12, 18, 255, 9, 18, 3, 18, 257, 8, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 3, 19, 264, 8, 19, 1, 19, 0, 1, 34,
		20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 0, 2, 1, 0, 10, 11, 1, 0, 18, 30, 280, 0, 45, 1, 0, 0, 0, 2, 50,
		1, 0, 0, 0, 4, 53, 1, 0, 0, 0, 6, 84, 1, 0, 0, 0, 8, 86, 1, 0, 0, 0, 10,
		92, 1, 0, 0, 0, 12, 105, 1, 0, 0, 0, 14, 109, 1, 0, 0, 0, 16, 115, 1, 0,
		0, 0, 18, 124, 1, 0, 0, 0, 20, 127, 1, 0, 0, 0, 22, 130, 1, 0, 0, 0, 24,
		138, 1, 0, 0, 0, 26, 151, 1, 0, 0, 0, 28, 155, 1, 0, 0, 0, 30, 168, 1,
		0, 0, 0, 32, 205, 1, 0, 0, 0, 34, 227, 1, 0, 0, 0, 36, 242, 1, 0, 0, 0,
		38, 260, 1, 0, 0, 0, 40, 44, 3, 20, 10, 0, 41, 44, 3, 8, 4, 0, 42, 44,
		3, 14, 7, 0, 43, 40, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0,
		44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 48, 1,
		0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 5, 0, 0, 1, 49, 1, 1, 0, 0, 0, 50,
		51, 5, 37, 0, 0, 51, 3, 1, 0, 0, 0, 52, 54, 5, 6, 0, 0, 53, 52, 1, 0, 0,
		0, 53, 54, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 3, 2, 1, 0, 56, 57,
		5, 32, 0, 0, 57, 58, 3, 6, 3, 0, 58, 5, 1, 0, 0, 0, 59, 60, 5, 13, 0, 0,
		60, 65, 3, 2, 1, 0, 61, 62, 5, 33, 0, 0, 62, 64, 3, 2, 1, 0, 63, 61, 1,
		0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66,
		68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 5, 16, 0, 0, 69, 85, 1, 0,
		0, 0, 70, 83, 3, 2, 1, 0, 71, 72, 5, 12, 0, 0, 72, 77, 3, 2, 1, 0, 73,
		74, 5, 33, 0, 0, 74, 76, 3, 2, 1, 0, 75, 73, 1, 0, 0, 0, 76, 79, 1, 0,
		0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 80, 81, 5, 15, 0, 0, 81, 83, 1, 0, 0, 0, 82, 70, 1, 0, 0, 0,
		82, 71, 1, 0, 0, 0, 83, 85, 1, 0, 0, 0, 84, 59, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 85, 7, 1, 0, 0, 0, 86, 87, 5, 1, 0, 0, 87, 88, 5, 32, 0, 0, 88,
		89, 3, 6, 3, 0, 89, 90, 3, 2, 1, 0, 90, 91, 3, 10, 5, 0, 91, 9, 1, 0, 0,
		0, 92, 101, 5, 14, 0, 0, 93, 98, 3, 12, 6, 0, 94, 95, 5, 33, 0, 0, 95,
		97, 3, 12, 6, 0, 96, 94, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0,
		0, 0, 98, 99, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101,
		93, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 5,
		17, 0, 0, 104, 11, 1, 0, 0, 0, 105, 106, 3, 2, 1, 0, 106, 107, 5, 35, 0,
		0, 107, 108, 3, 4, 2, 0, 108, 13, 1, 0, 0, 0, 109, 110, 5, 2, 0, 0, 110,
		111, 5, 32, 0, 0, 111, 112, 3, 6, 3, 0, 112, 113, 3, 2, 1, 0, 113, 114,
		3, 16, 8, 0, 114, 15, 1, 0, 0, 0, 115, 119, 5, 14, 0, 0, 116, 118, 3, 18,
		9, 0, 117, 116, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0,
		119, 120, 1, 0, 0, 0, 120, 122, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122,
		123, 5, 17, 0, 0, 123, 17, 1, 0, 0, 0, 124, 125, 3, 22, 11, 0, 125, 126,
		5, 39, 0, 0, 126, 19, 1, 0, 0, 0, 127, 128, 3, 22, 11, 0, 128, 129, 3,
		30, 15, 0, 129, 21, 1, 0, 0, 0, 130, 131, 5, 3, 0, 0, 131, 132, 5, 32,
		0, 0, 132, 133, 3, 6, 3, 0, 133, 134, 3, 2, 1, 0, 134, 136, 3, 24, 12,
		0, 135, 137, 3, 4, 2, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		23, 1, 0, 0, 0, 138, 147, 5, 12, 0, 0, 139, 144, 3, 26, 13, 0, 140, 141,
		5, 33, 0, 0, 141, 143, 3, 26, 13, 0, 142, 140, 1, 0, 0, 0, 143, 146, 1,
		0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 148, 1, 0, 0,
		0, 146, 144, 1, 0, 0, 0, 147, 139, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		149, 1, 0, 0, 0, 149, 150, 5, 15, 0, 0, 150, 25, 1, 0, 0, 0, 151, 152,
		3, 2, 1, 0, 152, 153, 5, 35, 0, 0, 153, 154, 3, 4, 2, 0, 154, 27, 1, 0,
		0, 0, 155, 164, 5, 12, 0, 0, 156, 161, 3, 34, 17, 0, 157, 158, 5, 33, 0,
		0, 158, 160, 3, 34, 17, 0, 159, 157, 1, 0, 0, 0, 160, 163, 1, 0, 0, 0,
		161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163,
		161, 1, 0, 0, 0, 164, 156, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166,
		1, 0, 0, 0, 166, 167, 5, 15, 0, 0, 167, 29, 1, 0, 0, 0, 168, 172, 5, 14,
		0, 0, 169, 171, 3, 32, 16, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0, 0,
		0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 175, 1, 0, 0, 0, 174,
		172, 1, 0, 0, 0, 175, 176, 5, 17, 0, 0, 176, 31, 1, 0, 0, 0, 177, 178,
		5, 5, 0, 0, 178, 179, 3, 2, 1, 0, 179, 180, 5, 35, 0, 0, 180, 181, 3, 4,
		2, 0, 181, 182, 5, 31, 0, 0, 182, 183, 3, 34, 17, 0, 183, 184, 5, 39, 0,
		0, 184, 206, 1, 0, 0, 0, 185, 186, 5, 8, 0, 0, 186, 187, 3, 34, 17, 0,
		187, 190, 3, 30, 15, 0, 188, 189, 5, 9, 0, 0, 189, 191, 3, 30, 15, 0, 190,
		188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 206, 1, 0, 0, 0, 192, 194,
		5, 4, 0, 0, 193, 195, 3, 34, 17, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1,
		0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 206, 5, 39, 0, 0, 197, 198, 3, 2, 1,
		0, 198, 199, 5, 31, 0, 0, 199, 200, 3, 34, 17, 0, 200, 201, 5, 39, 0, 0,
		201, 206, 1, 0, 0, 0, 202, 203, 3, 34, 17, 0, 203, 204, 5, 39, 0, 0, 204,
		206, 1, 0, 0, 0, 205, 177, 1, 0, 0, 0, 205, 185, 1, 0, 0, 0, 205, 192,
		1, 0, 0, 0, 205, 197, 1, 0, 0, 0, 205, 202, 1, 0, 0, 0, 206, 33, 1, 0,
		0, 0, 207, 208, 6, 17, -1, 0, 208, 228, 5, 38, 0, 0, 209, 228, 7, 0, 0,
		0, 210, 211, 5, 7, 0, 0, 211, 228, 3, 34, 17, 7, 212, 213, 3, 2, 1, 0,
		213, 214, 5, 32, 0, 0, 214, 215, 3, 6, 3, 0, 215, 216, 3, 36, 18, 0, 216,
		228, 1, 0, 0, 0, 217, 228, 3, 38, 19, 0, 218, 219, 3, 6, 3, 0, 219, 220,
		5, 36, 0, 0, 220, 221, 3, 6, 3, 0, 221, 222, 3, 34, 17, 2, 222, 228, 1,
		0, 0, 0, 223, 224, 5, 12, 0, 0, 224, 225, 3, 34, 17, 0, 225, 226, 5, 15,
		0, 0, 226, 228, 1, 0, 0, 0, 227, 207, 1, 0, 0, 0, 227, 209, 1, 0, 0, 0,
		227, 210, 1, 0, 0, 0, 227, 212, 1, 0, 0, 0, 227, 217, 1, 0, 0, 0, 227,
		218, 1, 0, 0, 0, 227, 223, 1, 0, 0, 0, 228, 239, 1, 0, 0, 0, 229, 230,
		10, 10, 0, 0, 230, 231, 7, 1, 0, 0, 231, 238, 3, 34, 17, 11, 232, 233,
		10, 5, 0, 0, 233, 238, 3, 28, 14, 0, 234, 235, 10, 4, 0, 0, 235, 236, 5,
		34, 0, 0, 236, 238, 3, 2, 1, 0, 237, 229, 1, 0, 0, 0, 237, 232, 1, 0, 0,
		0, 237, 234, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239,
		240, 1, 0, 0, 0, 240, 35, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 256, 5,
		14, 0, 0, 243, 244, 3, 2, 1, 0, 244, 245, 5, 35, 0, 0, 245, 253, 3, 34,
		17, 0, 246, 247, 5, 33, 0, 0, 247, 248, 3, 2, 1, 0, 248, 249, 5, 35, 0,
		0, 249, 250, 3, 34, 17, 0, 250, 252, 1, 0, 0, 0, 251, 246, 1, 0, 0, 0,
		252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 243, 1, 0, 0, 0, 256, 257,
		1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 5, 17, 0, 0, 259, 37, 1, 0,
		0, 0, 260, 263, 3, 2, 1, 0, 261, 262, 5, 32, 0, 0, 262, 264, 3, 6, 3, 0,
		263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 39, 1, 0, 0, 0, 25, 43,
		45, 53, 65, 77, 82, 84, 98, 101, 119, 136, 144, 147, 161, 164, 172, 190,
		194, 205, 227, 237, 239, 253, 256, 263,
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
	TempoParserEOF        = antlr.TokenEOF
	TempoParserSTRUCT     = 1
	TempoParserINTERFACE  = 2
	TempoParserFUNC       = 3
	TempoParserRETURN     = 4
	TempoParserLET        = 5
	TempoParserASYNC      = 6
	TempoParserAWAIT      = 7
	TempoParserIF         = 8
	TempoParserELSE       = 9
	TempoParserTRUE       = 10
	TempoParserFALSE      = 11
	TempoParserLPAREN     = 12
	TempoParserLSQUARE    = 13
	TempoParserLCURLY     = 14
	TempoParserRPAREN     = 15
	TempoParserRSQUARE    = 16
	TempoParserRCURLY     = 17
	TempoParserPLUS       = 18
	TempoParserMINUS      = 19
	TempoParserMULTIPLY   = 20
	TempoParserDIVIDE     = 21
	TempoParserMODULO     = 22
	TempoParserEQUAL      = 23
	TempoParserNOT_EQUAL  = 24
	TempoParserLESS       = 25
	TempoParserLESS_EQ    = 26
	TempoParserGREATER    = 27
	TempoParserGREATER_EQ = 28
	TempoParserAND        = 29
	TempoParserOR         = 30
	TempoParserIS         = 31
	TempoParserROLE_AT    = 32
	TempoParserCOMMA      = 33
	TempoParserDOT        = 34
	TempoParserCOLON      = 35
	TempoParserCOM        = 36
	TempoParserID         = 37
	TempoParserNUMBER     = 38
	TempoParserEND        = 39
	TempoParserWHITESPACE = 40
)

// TempoParser rules.
const (
	TempoParserRULE_sourceFile           = 0
	TempoParserRULE_ident                = 1
	TempoParserRULE_valueType            = 2
	TempoParserRULE_roleType             = 3
	TempoParserRULE_struct               = 4
	TempoParserRULE_structFieldList      = 5
	TempoParserRULE_structField          = 6
	TempoParserRULE_interface            = 7
	TempoParserRULE_interfaceMethodsList = 8
	TempoParserRULE_interfaceMethod      = 9
	TempoParserRULE_func                 = 10
	TempoParserRULE_funcSig              = 11
	TempoParserRULE_funcParamList        = 12
	TempoParserRULE_funcParam            = 13
	TempoParserRULE_funcArgList          = 14
	TempoParserRULE_scope                = 15
	TempoParserRULE_stmt                 = 16
	TempoParserRULE_expr                 = 17
	TempoParserRULE_exprStructField      = 18
	TempoParserRULE_identAccess          = 19
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0 {
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserFUNC:
			{
				p.SetState(40)
				p.Func_()
			}

		case TempoParserSTRUCT:
			{
				p.SetState(41)
				p.Struct_()
			}

		case TempoParserINTERFACE:
			{
				p.SetState(42)
				p.Interface_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
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
		p.SetState(50)
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

	// Getter signatures
	Ident() IIdentContext
	ROLE_AT() antlr.TerminalNode
	RoleType() IRoleTypeContext
	ASYNC() antlr.TerminalNode

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

func (s *ValueTypeContext) Ident() IIdentContext {
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

func (s *ValueTypeContext) ROLE_AT() antlr.TerminalNode {
	return s.GetToken(TempoParserROLE_AT, 0)
}

func (s *ValueTypeContext) RoleType() IRoleTypeContext {
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

func (s *ValueTypeContext) ASYNC() antlr.TerminalNode {
	return s.GetToken(TempoParserASYNC, 0)
}

func (s *ValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterValueType(s)
	}
}

func (s *ValueTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitValueType(s)
	}
}

func (s *ValueTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitValueType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TempoParser) ValueType() (localctx IValueTypeContext) {
	localctx = NewValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TempoParserRULE_valueType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserASYNC {
		{
			p.SetState(52)
			p.Match(TempoParserASYNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(55)
		p.Ident()
	}
	{
		p.SetState(56)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.RoleType()
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

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TempoParserLSQUARE:
		localctx = NewRoleTypeSharedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(TempoParserLSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Ident()
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(61)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(62)
				p.Ident()
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(68)
			p.Match(TempoParserRSQUARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TempoParserLPAREN, TempoParserID:
		localctx = NewRoleTypeNormalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TempoParserID:
			{
				p.SetState(70)
				p.Ident()
			}

		case TempoParserLPAREN:
			{
				p.SetState(71)
				p.Match(TempoParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(72)
				p.Ident()
			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == TempoParserCOMMA {
				{
					p.SetState(73)
					p.Match(TempoParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.Ident()
				}

				p.SetState(79)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(80)
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
	p.EnterRule(localctx, 8, TempoParserRULE_struct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(TempoParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.RoleType()
	}
	{
		p.SetState(89)
		p.Ident()
	}
	{
		p.SetState(90)
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
	p.EnterRule(localctx, 10, TempoParserRULE_structFieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(93)
			p.StructField()
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(94)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(95)
				p.StructField()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(103)
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
	p.EnterRule(localctx, 12, TempoParserRULE_structField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Ident()
	}
	{
		p.SetState(106)
		p.Match(TempoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
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
	p.EnterRule(localctx, 14, TempoParserRULE_interface)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(TempoParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.RoleType()
	}
	{
		p.SetState(112)
		p.Ident()
	}
	{
		p.SetState(113)
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
	p.EnterRule(localctx, 16, TempoParserRULE_interfaceMethodsList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TempoParserFUNC {
		{
			p.SetState(116)
			p.InterfaceMethod()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
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
	p.EnterRule(localctx, 18, TempoParserRULE_interfaceMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.FuncSig()
	}
	{
		p.SetState(125)
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
	p.EnterRule(localctx, 20, TempoParserRULE_func)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.FuncSig()
	}
	{
		p.SetState(128)
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
	p.EnterRule(localctx, 22, TempoParserRULE_funcSig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TempoParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(TempoParserROLE_AT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.RoleType()
	}
	{
		p.SetState(133)

		var _x = p.Ident()

		localctx.(*FuncSigContext).name = _x
	}
	{
		p.SetState(134)

		var _x = p.FuncParamList()

		localctx.(*FuncSigContext).params = _x
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserASYNC || _la == TempoParserID {
		{
			p.SetState(135)

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
	p.EnterRule(localctx, 24, TempoParserRULE_funcParamList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(139)
			p.FuncParam()
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(140)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(141)
				p.FuncParam()
			}

			p.SetState(146)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(149)
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
	p.EnterRule(localctx, 26, TempoParserRULE_funcParam)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Ident()
	}
	{
		p.SetState(152)
		p.Match(TempoParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
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
	p.EnterRule(localctx, 28, TempoParserRULE_funcArgList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(TempoParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412316875904) != 0 {
		{
			p.SetState(156)
			p.expr(0)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(157)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(158)
				p.expr(0)
			}

			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(166)
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
	p.EnterRule(localctx, 30, TempoParserRULE_scope)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412316876208) != 0 {
		{
			p.SetState(169)
			p.Stmt()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(175)
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
	p.EnterRule(localctx, 32, TempoParserRULE_stmt)
	var _la int

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmtVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(TempoParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Ident()
		}
		{
			p.SetState(179)
			p.Match(TempoParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.ValueType()
		}
		{
			p.SetState(181)
			p.Match(TempoParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.expr(0)
		}
		{
			p.SetState(183)
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
			p.SetState(185)
			p.Match(TempoParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.expr(0)
		}
		{
			p.SetState(187)

			var _x = p.Scope()

			localctx.(*StmtIfContext).thenScope = _x
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TempoParserELSE {
			{
				p.SetState(188)
				p.Match(TempoParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(189)

				var _x = p.Scope()

				localctx.(*StmtIfContext).elseScope = _x
			}

		}

	case 3:
		localctx = NewStmtReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.Match(TempoParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412316875904) != 0 {
			{
				p.SetState(193)
				p.expr(0)
			}

		}
		{
			p.SetState(196)
			p.Match(TempoParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStmtAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(197)
			p.Ident()
		}
		{
			p.SetState(198)
			p.Match(TempoParserIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.expr(0)
		}
		{
			p.SetState(200)
			p.Match(TempoParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStmtExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
			p.expr(0)
		}
		{
			p.SetState(203)
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

type ExprBoolContext struct {
	ExprContext
}

func NewExprBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprBoolContext {
	var p = new(ExprBoolContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBoolContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TempoParserTRUE, 0)
}

func (s *ExprBoolContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TempoParserFALSE, 0)
}

func (s *ExprBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprBool(s)
	}
}

func (s *ExprBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprBool(s)
	}
}

func (s *ExprBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprBool(s)

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

type ExprNumContext struct {
	ExprContext
}

func NewExprNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNumContext {
	var p = new(ExprNumContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TempoParserNUMBER, 0)
}

func (s *ExprNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExprNum(s)
	}
}

func (s *ExprNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExprNum(s)
	}
}

func (s *ExprNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TempoVisitor:
		return t.VisitExprNum(s)

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

func (p *TempoParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TempoParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, TempoParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(208)
			p.Match(TempoParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TempoParserTRUE || _la == TempoParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewExprAwaitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.Match(TempoParserAWAIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.expr(7)
		}

	case 4:
		localctx = NewExprStructContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(212)
			p.Ident()
		}
		{
			p.SetState(213)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.RoleType()
		}
		{
			p.SetState(215)
			p.ExprStructField()
		}

	case 5:
		localctx = NewExprIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(217)
			p.IdentAccess()
		}

	case 6:
		localctx = NewExprComContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(218)

			var _x = p.RoleType()

			localctx.(*ExprComContext).sender = _x
		}
		{
			p.SetState(219)
			p.Match(TempoParserCOM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)

			var _x = p.RoleType()

			localctx.(*ExprComContext).receiver = _x
		}
		{
			p.SetState(221)
			p.expr(2)
		}

	case 7:
		localctx = NewExprGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Match(TempoParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.expr(0)
		}
		{
			p.SetState(225)
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
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprBinOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExprBinOpContext).lhs = _prevctx

				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(230)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147221504) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(231)

					var _x = p.expr(11)

					localctx.(*ExprBinOpContext).rhs = _x
				}

			case 2:
				localctx = NewExprCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(233)
					p.FuncArgList()
				}

			case 3:
				localctx = NewExprFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TempoParserRULE_expr)
				p.SetState(234)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(235)
					p.Match(TempoParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(236)
					p.Ident()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, TempoParserRULE_exprStructField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(TempoParserLCURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TempoParserID {
		{
			p.SetState(243)
			p.Ident()
		}
		{
			p.SetState(244)
			p.Match(TempoParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.expr(0)
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TempoParserCOMMA {
			{
				p.SetState(246)
				p.Match(TempoParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(247)
				p.Ident()
			}
			{
				p.SetState(248)
				p.Match(TempoParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(249)
				p.expr(0)
			}

			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 38, TempoParserRULE_identAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Ident()
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(261)
			p.Match(TempoParserROLE_AT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
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

func (p *TempoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
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
