package epp

import (
	"fmt"

	"github.com/tempo-lang/tempo/projection"
)

func (epp *epp) eppField(
	baseExpr projection.Expression,
	fieldName string,
	fieldType projection.Type,
) projection.Expression {

	switch baseExpr.Type().(type) {
	case *projection.AsyncType:
	case projection.BuiltinType:
	case *projection.ClosureType:
	case *projection.FunctionType:
	case *projection.InterfaceType:
		return projection.NewExprFieldAccess(baseExpr, fieldName, fieldType)
	case *projection.ListType:
		switch fieldName {
		case "length":
			return projection.NewExprListLength(baseExpr)
		}
	case *projection.StructType:
		return projection.NewExprFieldAccess(baseExpr, fieldName, fieldType)
	}

	panic(fmt.Sprintf("failed to epp unexpected field base type: %#v", baseExpr.Type()))
}
