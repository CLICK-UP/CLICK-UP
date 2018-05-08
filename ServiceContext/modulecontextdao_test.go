package ServiceContext

import (
	"testing"
)

/*func TestGetServiceContextFromModuleList(t *testing.T) {
	moduleList := []string{"FromDevice", "Classifier", "Strip", "CheckIPHeader", "IPFilter", "IPPrint", "Discard"}
	GetServiceContextFromModuleList(moduleList)
}*/

func TestInsertModuleContext(t *testing.T) {
	sc1 := ServiceContext{
		ClassName:      "Firewall",
		HeaderFilePath: "../../udf/firewall.hh",
		SourceFilePath: "../../udf/firewall.cc"}
	sc2 := ServiceContext{
		ClassName:      "Log",
		HeaderFilePath: "../../udf/log.hh",
		SourceFilePath: "../../udf/log.cc"}

	var sc []ServiceContext
	sc = append(sc, sc1)
	sc = append(sc, sc2)
	InsertModuleContext(sc)
}
