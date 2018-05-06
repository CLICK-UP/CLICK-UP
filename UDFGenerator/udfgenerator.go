package udfgenerator

import (
	"ClickDriver"
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"strings"
)

type AtomMapXml struct {
	XMLNAME    xml.Name `xml:"atommap"`
	AtomString []Atom   `xml:"atom"`
}

type UserDefinedElement struct {
	Element_name string
	Atom_name    []string
}

type Atom struct {
	XMLName        xml.Name `xml:"atom"`
	Port           string   `xml:"port,attr"`
	Name           string   `xml:"name,attr"`
	Include        string   `xml:"include,attr"`
	PublicFunc     string   `xml:"publicfunc,attr"`
	PublicFuncImpl string   `xml:"publicfuncimpl,attr"`
	Value          string   `xml:"value,attr"`
	Conf           string   `xml:"conf,attr"`
	Const          string   `xml:"const,attr"`
	Input          string   `xml:"input,attr"`
	Action         string   `xml:"action,attr"`
}

func Udfgenerator(udf []UserDefinedElement) ([]ClickDriver.User_defined_element, error) {
	//使用udf的element_name替换模板中的$ELEMENTNAME, $CLASSNAME, $COMPILE_NAME
	//使用Atom_action去atomlib中生成Atom对象，使用atom对象对模板进行内容替换
	var result []ClickDriver.User_defined_element
	var err error
	for _, v := range udf {

		headerTemp, errHead := ioutil.ReadFile("./UDFGenerator/headertemplate.tmpl")
		if errHead != nil {
			log.Fatal("udfgenerator 45 read heardertemplate.tmpl error : ", err)
			return result, errHead
		}
		sourceTemp, errSrc := ioutil.ReadFile("./UDFGenerator/sourcetemplate.tmpl")
		if errSrc != nil {
			log.Fatal("udfgenerator 50 read sourcetemplate.tmpl error : ", errSrc)
			return result, errSrc
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
				log.Fatal("udfgenerator 69 get Atom From Atom_name error : ", errAtom)
				return result, errAtom
			}
			includeStr.WriteString(atomTemp.Include)
			port.WriteString(atomTemp.Port)
			publicFunc.WriteString(atomTemp.PublicFunc)
			atomValue.WriteString(atomTemp.Value)
			constDef.WriteString(atomTemp.Const)
			atomValueInit.WriteString(atomTemp.Conf)
			publicFuncImpl.WriteString(atomTemp.PublicFuncImpl)
			elementInput.WriteString(atomTemp.Input)
			atomAction.WriteString(atomTemp.Action)
		}
		click_hh := strings.Replace(string(headerTemp), "$ELEMENTNAME", strings.ToUpper(v.Element_name), -1)
		click_hh = strings.Replace(click_hh, "$CLASSNAME", v.Element_name, -1)
		click_hh = strings.Replace(click_hh, "$INCLUDE", includeStr.String(), -1)
		click_hh = strings.Replace(click_hh, "$PORT", port.String(), -1)
		click_hh = strings.Replace(click_hh, "$PUBLICFUNCTION", publicFunc.String(), -1)
		click_hh = strings.Replace(click_hh, "$ATOMVALUE", atomValue.String(), -1)
		click_cc := strings.Replace(string(sourceTemp), "$CONST_DEFINE", constDef.String(), -1)
		click_cc = strings.Replace(click_cc, "$ATOM_VALUE_INIT", atomValueInit.String(), -1)
		click_cc = strings.Replace(click_cc, "$PUBLIC_FUNCTION_IMPL", publicFuncImpl.String(), -1)
		click_cc = strings.Replace(click_cc, "$ELEMENT_INPUT", elementInput.String(), -1)
		click_cc = strings.Replace(click_cc, "$ATOM_ACTION", atomAction.String(), -1)
		click_cc = strings.Replace(click_cc, "$CLASSNAME", v.Element_name, -1)
		click_cc = strings.Replace(click_cc, "$COMPILE_NAME", strings.ToLower(v.Element_name), -1)
		userFunction := ClickDriver.User_defined_element{v.Element_name, click_hh, click_cc}
		result = append(result, userFunction)
	}
	return result, err
}

func getAtomFromAtomName(atom string) (Atom, error) {
	var atomStruct Atom
	var atomReader AtomMapXml
	content, err := ioutil.ReadFile("./UDFGenerator/atommap.xml")
	if err != nil {
		log.Fatal("udfgenerator 106 read atommap.xml error : ", err)
	}
	err = xml.Unmarshal(content, &atomReader)
	if err != nil {
		log.Fatal("udfgenerator 110 parse xml error : ", err)
	}
	for _, v := range atomReader.AtomString {
		if strings.EqualFold(v.Name, atom) {
			atomStruct = v
		}
	}

	atomStruct.Include = strings.Replace(atomStruct.Include, "@", "\n", -1)
	atomStruct.PublicFunc = strings.Replace(atomStruct.PublicFunc, "@", "\n", -1)
	atomStruct.PublicFuncImpl = strings.Replace(atomStruct.PublicFuncImpl, "@", "\n", -1)
	atomStruct.Value = strings.Replace(atomStruct.Value, "@", "\n", -1)
	atomStruct.Conf = strings.Replace(atomStruct.Conf, "@", "\n", -1)
	atomStruct.Const = strings.Replace(atomStruct.Const, "@", "\n", -1)
	atomStruct.Input = strings.Replace(atomStruct.Input, "@", "\n", -1)
	atomStruct.Action = strings.Replace(atomStruct.Action, "@", "\n", -1)
	return atomStruct, err
}
