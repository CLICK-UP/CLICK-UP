package ServiceContext

import (
	"encoding/xml"
)

type ElementMapXML struct {
	XMLNAME     xml.Name        `xml:"elementmap"`
	EntryString []ModuleContext `xml:"entry"`
}

type ModuleContext struct {
	XMLName        xml.Name `xml:"entry"`
	ClassName      string   `xml:"name,attr"`
	CompileName    string   `xml:"cxxclass,attr"`
	HeaderFilePath string   `xml:"headerfile,attr"`
	SourceFilePath string   `xml:"sourcefile,attr"`
	Dependency     string   `xml:"requires,attr"`
	Provides       string   `xml:"provides,attr"`
}

func (mc ModuleContext) ToString() string {
	return mc.ClassName + "\t" + mc.CompileName + "\t" + mc.HeaderFilePath + "\t" + mc.SourceFilePath + "\t" + mc.Dependency + "\t" + mc.Provides
}
