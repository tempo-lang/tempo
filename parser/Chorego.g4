grammar Chorego;

/*
 * Parser Rules
 */

func: FUNC role_type_normal ident func_param_list scope;

func_param_list:
	LPAREN (func_param (COMMA func_param)*)? RPAREN;

func_param: ident ( COLON value_type)?;

scope: LCURLY RCURLY;

value_type: ident? role_type | ident role_type?;
role_type: role_type_normal | role_type_shared;

role_type_normal:
	ROLE_AT (ident | (LPAREN ident (COMMA ident)* RPAREN));

role_type_shared:
	ROLE_AT (LSQUARE ident (COMMA ident)* RSQUARE);

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