grammar Tempo;

/*
 * Parser Rules
 */

// source file
sourceFile: (func | struct | interface)* EOF;

// identifier
ident: ID;

// type declaration
valueType: ASYNC? ident ROLE_AT roleType;

roleType:
	(LSQUARE ident (COMMA ident)* RSQUARE)				# roleTypeShared
	| (ident | (LPAREN ident (COMMA ident)* RPAREN))	# roleTypeNormal;

// struct
struct: STRUCT ROLE_AT roleType ident structFieldList;

structFieldList:
	LCURLY (structField (COMMA structField)*)? RCURLY;

structField: ident COLON valueType;

// interface
interface:
	INTERFACE ROLE_AT roleType ident interfaceMethodsList;

interfaceMethodsList:
	LCURLY (interfaceMethod (COMMA interfaceMethod)*)? RCURLY;

interfaceMethod: funcSig END;

// function
func: funcSig scope;

funcSig:
	FUNC ROLE_AT roleType name = ident params = funcParamList returnType = valueType?;

funcParamList: LPAREN (funcParam (COMMA funcParam)*)? RPAREN;

funcParam: ident COLON valueType;

funcArgList: LPAREN (expr (COMMA expr)*)? RPAREN;

// scope
scope: LCURLY stmt* RCURLY;

// statements
stmt:
	LET ident COLON valueType IS expr END	# stmtVarDecl
	| IF expr scope (ELSE scope)?			# stmtIf
	| RETURN expr END						# stmtReturn
	| ident IS expr END						# stmtAssign
	| expr END								# stmtExpr;

// expressions
expr:
	expr (
		PLUS
		| MINUS
		| MULTIPLY
		| DIVIDE
		| MODULO
		| EQUAL
		| NOT_EQUAL
		| LESS
		| LESS_EQ
		| GREATER
		| GREATER_EQ
		| AND
		| OR
	) expr										# exprBinOp
	| NUMBER									# exprNum
	| (TRUE | FALSE)							# exprBool
	| AWAIT expr								# exprAwait
	| roleType COM roleType expr				# exprCom
	| ident ROLE_AT roleType exprStructField	# exprStruct
	| expr funcArgList ROLE_AT roleType			# exprCall
	| expr DOT ident							# exprFieldAccess
	| ident										# exprIdent
	| LPAREN expr RPAREN						# exprGroup;

exprStructField:
	LCURLY (ident COLON expr (COMMA ident COLON expr)*)? RCURLY;

/*
 * Lexer Rules
 */

// Keywords

STRUCT: 'struct';
INTERFACE: 'interface';
FUNC: 'func';
RETURN: 'return';
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

PLUS: '+';
MINUS: '-';
MULTIPLY: '*';
DIVIDE: '/';
MODULO: '%';
EQUAL: '==';
NOT_EQUAL: '!=';
LESS: '<';
LESS_EQ: '<=';
GREATER: '>';
GREATER_EQ: '>=';
AND: '&&';
OR: '||';

IS: '=';
ROLE_AT: '@';
COMMA: ',';
DOT: '.';
COLON: ':';
COM: '->';

ID: ('_' [a-zA-Z_0-9]+ | [a-zA-Z][a-zA-Z_0-9]*);
NUMBER: [0-9]+;
END: ';';
WHITESPACE: [ \t\r\n]+ -> skip;