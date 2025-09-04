grammar Tempo;

/*
 * Parser Rules
 */

// source file
sourceFile: (func | struct | interface)* EOF;

// identifier
ident: ID;

roleIdent: ident (ROLE_AT roleType)?;

// type declaration
valueType:
	ASYNC inner = valueType														# asyncType
	| LSQUARE inner = valueType RSQUARE											# listType
	| FUNC ROLE_AT roleType params = closureParamList returnType = valueType?	# closureType
	| roleIdent																	# namedType;

roleType:
	(LSQUARE ident (COMMA ident)* RSQUARE)				# roleTypeShared
	| (ident | (LPAREN ident (COMMA ident)* RPAREN))	# roleTypeNormal;

// closure
closureParamList: LPAREN (valueType (COMMA valueType)*)? RPAREN;

closureSig:
	FUNC ROLE_AT roleType params = funcParamList returnType = valueType?;

// struct
struct:
	STRUCT (ROLE_AT roleType)? ident structImplements? body = structBody;

structImplements: IMPLEMENTS roleIdent (COMMA roleIdent)*;

structBody: LCURLY (structField | func)* RCURLY;

structField: ident COLON valueType END;

// interface
interface:
	INTERFACE (ROLE_AT roleType)? ident interfaceMethodsList;

interfaceMethodsList: LCURLY interfaceMethod* RCURLY;

interfaceMethod: funcSig END;

// function
func: funcSig scope;

funcSig:
	FUNC (ROLE_AT roleType)? name = ident params = funcParamList returnType = valueType?;

funcParamList: LPAREN (funcParam (COMMA funcParam)*)? RPAREN;

funcParam: ident COLON valueType;

funcArgList: LPAREN (expr (COMMA expr)*)? RPAREN;

// scope
scope: LCURLY stmt* RCURLY;

// statements
stmt:
	LET ident (COLON valueType)? IS expr END				# stmtVarDecl
	| IF expr thenScope = scope (ELSE elseScope = scope)?	# stmtIf
	| WHILE expr scope										# stmtWhile
	| RETURN expr? END										# stmtReturn
	| ident assignSpecifier* IS expr END					# stmtAssign
	| expr END												# stmtExpr;

assignSpecifier:
	DOT ident				# assignField
	| LSQUARE expr RSQUARE	# assignIndex;

// expressions
expr:
	lhs = expr (
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
	) rhs = expr										# exprBinOp
	| literal (ROLE_AT roleType)?						# exprPrimitive
	| closureSig scope									# exprClosure
	| roleIdent exprStructField							# exprStruct
	| expr funcArgList									# exprCall
	| expr DOT ident									# exprFieldAccess
	| baseExpr = expr LSQUARE indexExpr = expr RSQUARE	# exprIndex
	| LSQUARE (expr (COMMA expr)*)? RSQUARE				# exprList
	| identAccess										# exprIdent
	| sender = roleType COM receiver = roleType expr	# exprCom
	| AWAIT expr										# exprAwait
	| LPAREN expr RPAREN								# exprGroup;

exprStructField:
	LCURLY (ident COLON expr (COMMA ident COLON expr)*)? RCURLY;

identAccess: ident (ROLE_AT roleType)?;

literal:
	(NUMBER DOT NUMBER | NUMBER DOT | DOT NUMBER)	# float
	| NUMBER										# int
	| STRING										# string
	| (TRUE | FALSE)								# bool;

/*
 * Lexer Rules
 */

// Keywords
STRUCT: 'struct';
INTERFACE: 'interface';
IMPLEMENTS: 'implements';
FUNC: 'func';
RETURN: 'return';
LET: 'let';
ASYNC: 'async';
AWAIT: 'await';
IF: 'if';
ELSE: 'else';
WHILE: 'while';

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

STRING: '"' (ESC | SAFECODEPOINT)* '"';
fragment ESC: '\\' (["\\nrt]);
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];

ID: ('_' [a-zA-Z_0-9]+ | [a-zA-Z][a-zA-Z_0-9]*);
NUMBER: [0-9]+;
END: ';';
WHITESPACE: [ \t\r\n]+ -> skip;

LINE_COMMENT: '//' ~( '\r' | '\n')* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;