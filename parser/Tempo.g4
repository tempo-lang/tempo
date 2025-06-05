grammar Tempo;

/*
 * Parser Rules
 */

// source file
sourceFile: (func | struct | interface)* EOF;

// identifier
ident: ID;

// type declaration
valueType: ASYNC? (ident ROLE_AT roleType | closureType);

roleType:
	(LSQUARE ident (COMMA ident)* RSQUARE)				# roleTypeShared
	| (ident | (LPAREN ident (COMMA ident)* RPAREN))	# roleTypeNormal;

// closure
closureType:
	FUNC ROLE_AT roleType params = closureParamList returnType = valueType?;

closureParamList: LPAREN (valueType (COMMA valueType)*)? RPAREN;

closureSig:
	FUNC ROLE_AT roleType params = funcParamList returnType = valueType?;

// struct
struct: STRUCT ROLE_AT roleType ident structFieldList;

structFieldList:
	LCURLY (structField (COMMA structField)*)? RCURLY;

structField: ident COLON valueType;

// interface
interface:
	INTERFACE ROLE_AT roleType ident interfaceMethodsList;

interfaceMethodsList: LCURLY interfaceMethod* RCURLY;

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
	LET ident (COLON valueType)? IS expr END				# stmtVarDecl
	| IF expr thenScope = scope (ELSE elseScope = scope)?	# stmtIf
	| RETURN expr? END										# stmtReturn
	| ident IS expr END										# stmtAssign
	| expr END												# stmtExpr;

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
	| AWAIT expr										# exprAwait
	| closureSig scope									# exprClosure
	| ident ROLE_AT roleType exprStructField			# exprStruct
	| expr funcArgList									# exprCall
	| expr DOT ident									# exprFieldAccess
	| identAccess										# exprIdent
	| sender = roleType COM receiver = roleType expr	# exprCom
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

STRING: '"' (ESC | SAFECODEPOINT)* '"';
fragment ESC: '\\' (["\\nrt]);
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];

ID: ('_' [a-zA-Z_0-9]+ | [a-zA-Z][a-zA-Z_0-9]*);
NUMBER: [0-9]+;
END: ';';
WHITESPACE: [ \t\r\n]+ -> skip;

LINE_COMMENT: '//' ~( '\r' | '\n')* -> skip;
BLOCK_COMMENT: '/*' .*? '*/' -> skip;