package ServiceContext

import (
	"testing"
)

func TestGetModuleContextFromModuleList(t *testing.T) {
	moduleList := []string{"IPRewriterBase"}
	GetModuleContextFromModuleList(moduleList)
}
