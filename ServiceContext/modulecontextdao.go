package ServiceContext

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetServiceContextFromModuleList(moduleList []string) ([]string, []ServiceContext, error) {

	moduleList = append(moduleList, "ControlSocket")

	content, err := ioutil.ReadFile("./ServiceContext/elementmap.xml")
	if err != nil {
		log.Fatal("modulecontextdao 15 read elementmap.xml error : ", err)
	}

	var res ElementMapXML
	err = xml.Unmarshal(content, &res)
	if err != nil {
		log.Fatal("modulecontextdao 21 parse xml error : ", err)
	}

	fullContext := GetFullContextFromModule(moduleList, res)
	compileList, serviceContext := GetServiceContextFromFullContext(fullContext, res)

	fmt.Println("CompileList as follow")
	for _, v := range compileList {
		fmt.Println(v)
	}

	fmt.Println("ServiceContext as follow")
	for _, v := range serviceContext {
		fmt.Println(v.ToString())
	}

	return compileList, serviceContext, err
}

func GetFullContextFromModule(moduleList []string, res ElementMapXML) map[string]bool {
	var result = make(map[string]bool)

	for i := 0; i < len(moduleList); i++ {
		for _, line := range res.EntryString {
			if strings.EqualFold(moduleList[i], line.ClassName) {
				result[line.ClassName] = true
				if line.Dependency != "" {
					dep := strings.Split(line.Dependency, " ")
					for _, req := range dep {
						if result[req] != true {
							result[req] = true
						}
					}
				}
			}
		}
	}

	return result
}

func GetServiceContextFromFullContext(fullContext map[string]bool, res ElementMapXML) ([]string, []ServiceContext) {
	var serviceContextResult []ServiceContext
	var compileListResult []string

	for key, _ := range fullContext {
		for _, line := range res.EntryString {
			if strings.EqualFold(key, line.ClassName) || strings.EqualFold(key, line.Provides) {
				if line.Provides != "" {
					var sc ServiceContext
					if line.HeaderFilePath[0] == byte('<') {
						sc = ServiceContext{line.SourceFilePath, line.HeaderFilePath, ""}
					} else {
						sc = ServiceContext{line.SourceFilePath, `"` + line.HeaderFilePath + `"`, ""}
					}
					serviceContextResult = append(serviceContextResult, sc)
					compileListResult = append(compileListResult, line.CompileName)
				} else {
					var sc ServiceContext
					if line.HeaderFilePath[0] == byte('<') {
						sc = ServiceContext{line.SourceFilePath, line.HeaderFilePath, line.ClassName + "-" + line.ClassName}
					} else {
						sc = ServiceContext{line.SourceFilePath, `"` + line.HeaderFilePath + `"`, line.ClassName + "-" + line.ClassName}
					}

					serviceContextResult = append(serviceContextResult, sc)
					compileListResult = append(compileListResult, line.CompileName)
				}
			}
		}
	}

	return compileListResult, serviceContextResult
}

func InsertModuleContext(sc []ServiceContext) {
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	for _, v := range sc {
		mc := &ModuleContext{
			ClassName:      v.ClassName,
			CompileName:    strings.ToLower(v.ClassName),
			HeaderFilePath: v.HeaderFilePath,
			SourceFilePath: v.SourceFilePath}
		enc.Encode(mc)
		buf.WriteString("\n")
	}
	buf.WriteString("</elementmap>")
	elementMap, errMap := ioutil.ReadFile("./ServiceContext/elementmap.xml")
	if errMap != nil {
		log.Fatal("read elementmap.xml error : ", errMap)
	}
	str := strings.Replace(string(elementMap), "</elementmap>", buf.String(), -1)
	err := ioutil.WriteFile("./ServiceContext/elementmap.xml", []byte(str), 0777)
	if err != nil {
		log.Fatal("write elementmap.xml error : ", err)
	}
}
