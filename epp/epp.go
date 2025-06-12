// This package implements endpoint projection.
// The [EndpointProject] function is responsible for converting a choreographic AST into projections (defined in the [projection] package).
package epp

import (
	"fmt"

	"github.com/tempo-lang/tempo/parser"
	"github.com/tempo-lang/tempo/projection"
	"github.com/tempo-lang/tempo/type_check"
)

type epp struct {
	info     *type_check.Info
	tmpValue int
}

func newEpp(info *type_check.Info) *epp {
	return &epp{
		info:     info,
		tmpValue: 0,
	}
}

// EndpointProject takes an [type_check.Info] object obtained from type checking, as well as the AST and returns a projection.
//
// It is undefined to run this function on an AST that produced any type errors.
// Doing so will likely cause a `panic`.
func EndpointProject(info *type_check.Info, sourceFile parser.ISourceFileContext) *projection.SourceFile {
	return newEpp(info).eppSourceFile(sourceFile)
}

func (epp *epp) nextTmpName() string {
	name := fmt.Sprintf("tmp%d", epp.tmpValue)
	epp.tmpValue += 1
	return name
}
