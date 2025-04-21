grammar Chorego;

/*
 * Parser Rules
 */

func: FUNC roleTypeNormal ident funcParamList scope;

funcParamList: LPAREN (funcParam (COMMA funcParam)*)? RPAREN;

funcParam: ident COLON valueType;

scope: LCURLY RCURLY;

valueType: ident roleType;
roleType: roleTypeNormal | roleTypeShared;

roleTypeShared: ROLE_AT (LSQUARE ident (COMMA ident)* RSQUARE);

roleTypeNormal:
	ROLE_AT (ident | (LPAREN ident (COMMA ident)* RPAREN));

ident: ID;

/*
 * Lexer Rules
 */

// Keywords

FUNC: 'func';

// Parenthesis

LPAREN: '(';
LSQUARE: '[';
LCURLY: '{';

RPAREN: ')';
RSQUARE: ']';
RCURLY: '}';

// Symbols

ROLE_AT: '@';
COMMA: ',';
COLON: ':';

ID: [a-zA-Z_][a-zA-Z_0-9]*;
NUMBER: [0-9]+;
WHITESPACE: [ \t\r\n]+ -> skip;