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

func EndpointProject(info *type_check.Info, sourceFile parser.ISourceFileContext) *projection.SourceFile {
	return newEpp(info).eppSourceFile(sourceFile)
}

func (epp *epp) nextTmpName() string {
	name := fmt.Sprintf("tmp%d", epp.tmpValue)
	epp.tmpValue += 1
	return name
}
