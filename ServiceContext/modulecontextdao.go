package ServiceContext

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetServiceContextFromModuleList(moduleList []string) ([]string, []ServiceContext) {

	content, err := ioutil.ReadFile("elementmap.xml")
	if err != nil {
		log.Fatal(err)
	}

	var res ElementMapXML
	err = xml.Unmarshal(content, &res)
	if err != nil {
		log.Fatal(err)
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

	return compileList, serviceContext
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
					sc := ServiceContext{line.SourceFilePath, line.HeaderFilePath, ""}
					serviceContextResult = append(serviceContextResult, sc)
					compileListResult = append(compileListResult, line.CompileName)
				} else {
					sc := ServiceContext{line.SourceFilePath, line.HeaderFilePath, line.ClassName + "-" + line.ClassName}
					serviceContextResult = append(serviceContextResult, sc)
					compileListResult = append(compileListResult, line.CompileName)
				}
			}
		}
	}

	return compileListResult, serviceContextResult
}
