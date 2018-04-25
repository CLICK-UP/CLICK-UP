package udfgenerator

import (
	"ClickDirver"
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

type UserDefinedElement struct {
	Element_name string
	Atom_name    []string
}

type Atom struct {
	port           string
	Name           string
	Include        string
	PubilcFunc     string
	PublicFuncImpl string
	value          string
	Conf           string
	Const          string
	Input          string
	Action         string
}

func Udfgenerator(udf []UserDefinedElement) ([]ClickDriver.User_defined_element, error) {
	//使用udf的element_name替换模板中的$ELEMENTNAME, $CLASSNAME, $COMPILE_NAME
	//使用Atom_action去atomlib中生成Atom对象，使用atom对象对模板进行内容替换
	var result []ClickDriver.User_defined_element
	var err error
	for _, v := range udf {

		headerTemp, errHead := ioutil.ReadFile("headertemplate.txt")
		if errHead != nil {
			log.Fatal(err)
			return errHead
		}
		sourceTemp, errSrc := ioutil.ReadFile("sourcetemplate.txt")
		if errSrc != nil {
			log.Fatal(errSrc)
			return errSrc
		}
		//headerNameReplace := strings.Replace(string(headerTemp), "$ELEMENTNAME", strings.ToUpper(v.Element_name), -1)
		//headerNameReplace = strings.Replace(headerNameReplace, "$CLASSNAME", v.Element_name, -1)
		//sourceNameReplace := strings.Replace(string(sourceTemp), "$CLASSNAME", v.Element_name, -1)
		//sourceNameReplace = strings.Replace(sourceNameReplace, "$COMPILE_NAME", strings.ToLower(v.Element_name), -1)
		var includeStr bytes.Buffer
		var port bytes.Buffer
		var publicFunc bytes.Buffer
		var atomValue bytes.Buffer
		var constDef bytes.Buffer
		var atomValueInit bytes.Buffer
		var publicFuncImpl bytes.Buffer
		var elementInput bytes.Buffer
		var atomAction bytes.Buffer
		for _, atom := range v.Atom_name {
			atomTemp, errAtom := getAtomFromAtomName(atom)
			if errAtom != nil {
				log.Fatal(errAtom)
				return errAtom
			}
			includeStr.WriteString(atomTemp.Include)
			port.WriteString(atomTemp.port)
			publicFunc.WriteString(atomTemp.PubilcFunc)
			atomValue.WriteString(atomTemp.value)
			constDef.WriteString(atomTemp.Const)
			atomValueInit.WriteString(atomTemp.Conf)
			publicFuncImpl.WriteString(atomTemp.PublicFuncImpl)
			elementInput.WriteString(atomTemp.Input)
			atomAction.WriteString(atomTemp.Action)
		}
		click_hh := strings.Replace(string(headerTemp), "$ELEMENTNAME", strings.ToUpper(v.Element_name), -1)
		click_hh = strings.Replace(click_hh, "$CLASSNAME", v.Element_name, -1)
		click_hh = strings.Replace(click_hh, "$INCLUDE", string(includeStr), -1)
		click_hh = strings.Replace(click_hh, "$PORT", string(port), -1)
		click_hh = strings.Replace(click_hh, "$PUBLICFUNCTION", string(publicFunc), -1)
		click_hh = strings.Replace(click_hh, "$ATOMVALUE", string(atomValue), -1)
		click_cc := strings.Replace(string(sourceTemp), "$CONST_DEFINE", string(constDef), -1)
		click_cc = strings.Replace(click_cc, "$ATOM_VALUE_INIT", string(atomValueInit), -1)
		click_cc = strings.Replace(click_cc, "$PUBLIC_FUNCTION_IMPL", string(publicFuncImpl), -1)
		click_cc = strings.Replace(click_cc, "$ELEMENT_INPUT", string(elementInput), -1)
		click_cc = strings.Replace(click_cc, "$ATOM_ACTION", string(atomAction), -1)
		click_cc = strings.Replace(click_cc, "$CLASSNAME", v.Element_name, -1)
		click_cc = strings.Replace(click_cc, "$COMPILE_NAME", strings.ToLower(v.Element_name), -1)
		result = append(ClickDriver.user_defined_element{v.Element_name, click_hh, click_cc})
	}
	return result, err
}

func getAtomFromAtomName(atom string) (Atom, error) {

}
