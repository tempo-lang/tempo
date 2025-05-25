package epp

import (
	"fmt"
	"tempo/parser"
	"tempo/projection"
	"tempo/type_check"
)

type epp struct {
	info               *type_check.Info
	tmpValue           int
	callParamRoleSubst map[string]string
	callArgRoleSubst   map[string]string
}

func newEpp(info *type_check.Info) *epp {
	return &epp{
		info:               info,
		tmpValue:           0,
		callParamRoleSubst: map[string]string{},
		callArgRoleSubst:   map[string]string{},
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
