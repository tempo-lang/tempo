package type_check

import (
	"github.com/tempo-lang/tempo/sym_table"
	"github.com/tempo-lang/tempo/types"
)

func (tc *typeChecker) populateGlobalSymbols() {
	floatSym := sym_table.NewTypeSymbol(types.Float(nil), tc.info.GlobalScope)
	tc.info.GlobalScope.InsertSymbol(floatSym)

	intSym := sym_table.NewTypeSymbol(types.Int(nil), tc.info.GlobalScope)
	tc.info.GlobalScope.InsertSymbol(intSym)

	stringSym := sym_table.NewTypeSymbol(types.String(nil), tc.info.GlobalScope)
	tc.info.GlobalScope.InsertSymbol(stringSym)
}
