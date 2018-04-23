package ServiceContext

import (
	"testing"
)

func TestGetServiceContextFromModuleList(t *testing.T) {
	moduleList := []string{"IPRewriter"}
	GetServiceContextFromModuleList(moduleList)
}
