grammar Tempo;

/*
 * Parser Rules
 */

// source file
sourceFile: func*;

// identifier
ident: ID;

// type declaration
valueType: ASYNC? ident ROLE_AT roleType;

roleType:
	(LSQUARE ident (COMMA ident)* RSQUARE)				# roleTypeShared
	| (ident | (LPAREN ident (COMMA ident)* RPAREN))	# roleTypeNormal;

// function
func:
	FUNC ROLE_AT roleType ident funcParamList valueType? scope;

funcParamList: LPAREN (funcParam (COMMA funcParam)*)? RPAREN;

funcParam: ident COLON valueType;

funcArgList: LPAREN (expr (COMMA expr)*)? RPAREN;

// scope
scope: LCURLY stmt* RCURLY;

// statements
stmt:
	LET ident COLON valueType EQUAL expr END	# stmtVarDecl
	| IF expr scope (ELSE scope)?				# stmtIf
	| ident EQUAL expr END						# stmtAssign
	| expr END									# stmtExpr;

// expressions
expr:
	expr PLUS expr							# exprAdd
	| NUMBER								# exprNum
	| (TRUE | FALSE)						# exprBool
	| AWAIT expr							# exprAwait
	| roleType COM roleType expr			# exprCom
	| ident funcArgList ROLE_AT roleType	# exprCall
	| ident									# exprIdent
	| LPAREN expr RPAREN					# exprGroup;

/*
 * Lexer Rules
 */

// Keywords

FUNC: 'func';
LET: 'let';
ASYNC: 'async';
AWAIT: 'await';
IF: 'if';
ELSE: 'else';

TRUE: 'true';
FALSE: 'false';

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
PLUS: '+';
COM: '->';

ID: ('_' [a-zA-Z_0-9]+ | [a-zA-Z][a-zA-Z_0-9]*);
NUMBER: [0-9]+;
END: ';';
WHITESPACE: [ \t\r\n]+ -> skip;