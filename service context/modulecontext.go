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

/*func (mc ModuleContext) ClassName() string {
	return mc.className
}

func (mc *ModuleContext) SetClassName(className string) {
	mc.className = className
}

func (mc ModuleContext) CompileName() string {
	return mc.compileName
}

func (mc *ModuleContext) SetCompileName(compileName string) {
	mc.compileName = compileName
}

func (mc ModuleContext) HeaderFilePath() string {
	return mc.headerFilePath
}

func (mc *ModuleContext) SetHeaderFilePath(headerFilePath string) {
	mc.headerFilePath = headerFilePath
}

func (mc ModuleContext) SourceFilePath() string {
	return mc.sourceFilePath
}

func (mc *ModuleContext) SetSourceFilePath(sourceFilePath string) {
	mc.sourceFilePath = sourceFilePath
}

func (mc ModuleContext) Dependency() string {
	return mc.dependency
}

func (mc ModuleContext) SetDependency(dependency string) {
	mc.dependency = dependency
}*/
