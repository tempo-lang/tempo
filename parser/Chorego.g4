grammar Chorego;

/*
 * Parser Rules
 */

// source file
sourceFile: func*;

// identifier
ident: ID;

// type declaration
valueType: ident roleType;
roleType: roleTypeNormal | roleTypeShared;

roleTypeShared: ROLE_AT (LSQUARE ident (COMMA ident)* RSQUARE);

roleTypeNormal:
	ROLE_AT (ident | (LPAREN ident (COMMA ident)* RPAREN));

// function
func: FUNC roleTypeNormal ident funcParamList scope;

funcParamList: LPAREN (funcParam (COMMA funcParam)*)? RPAREN;

funcParam: ident COLON valueType;

// scope
scope: LCURLY statement* RCURLY;

// statements
statement: stmtVarDecl;

stmtVarDecl: LET ident COLON valueType EQUAL expression END;

// expressions
expression: NUMBER | ident;

/*
 * Lexer Rules
 */

// Keywords

FUNC: 'func';
LET: 'let';

// Parenthesis

LPAREN: '(';
LSQUARE: '[';
LCURLY: '{';

RPAREN: ')';
RSQUARE: ']';
RCURLY: '}';

// Symbols

EQUAL: '=';
ROLE_AT: '@';
COMMA: ',';
COLON: ':';

ID: ('_' [a-zA-Z_0-9]+ | [a-zA-Z][a-zA-Z_0-9]*);
NUMBER: [0-9]+;
END: ';';
WHITESPACE: [ \t\r\n]+ -> skip;