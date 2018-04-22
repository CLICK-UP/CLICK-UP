package ServiceContext

import (
	"container/list"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func GetServiceContextFromModuleList(moduleList []string) []ModuleContext {
	result := make(map[ServiceContext]bool)
	content, err := ioutil.ReadFile("elementmap.xml")
	if err != nil {
		log.Fatal(err)
	}

	var res ElementMapXML
	err = xml.Unmarshal(content, &res)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(moduleList); i++ {
		for _, line := range res.EntryString {
			if strings.EqualFold(line.ClassName, moduleList[i]) {
				sc := &ServiceContext{line.SourceFilePath, line.HeaderFilePath, line.ClassName + "-" + line.ClassName}
				result[sc] = true
				//深度优先搜索找出所有的依赖
				if line.Provides != nil{
					dep := strings.Split(line.Provides, " ")
					for _, pro range dep {
						
						stack := NewStack()
						stack.Push(pro)
						for stack.Empty() != true {

						}
					}
				}
				
				result = append(result, line)
				fmt.Println(line.ToString())
			}
		}
	}
	return result
}
