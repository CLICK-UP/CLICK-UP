package ServiceContext

import (
	"testing"
)

func TestGetServiceContextFromModuleList(t *testing.T) {
	moduleList := []string{"FromDevice", "Classifier", "Strip", "CheckIPHeader", "IPFilter", "IPPrint", "Discard"}
	GetServiceContextFromModuleList(moduleList)
}
