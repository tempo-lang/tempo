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
scope: LCURLY stmt* RCURLY;

// statements
stmt:
	LET ident COLON valueType EQUAL expr END	# stmtVarDecl
	| ident EQUAL expr END						# stmtAssign;

// expressions
expr:
	expr PLUS expr			# exprAdd
	| NUMBER				# exprNum
	| (TRUE | FALSE)		# exprBool
	| ident					# exprIdent
	| LPAREN expr RPAREN	# exprGroup;

/*
 * Lexer Rules
 */

// Keywords

FUNC: 'func';
LET: 'let';

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

ID: ('_' [a-zA-Z_0-9]+ | [a-zA-Z][a-zA-Z_0-9]*);
NUMBER: [0-9]+;
END: ';';
WHITESPACE: [ \t\r\n]+ -> skip;