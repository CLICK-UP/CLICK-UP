package ClickDirver

/*
*  Author : @ychuang
*  Create date : 2018-4-23
*  Input :
*		linkList : elements in Click configure file needed to link together
*		serviceContext : in order to generate and compile elements.cc
*		user_defined_element : in order to generate and compile udf.cc
*  Output :
*		err : the error information of three methods above
 */

import (
	"ServiceContext"
	"strings"
)

type User_defined_element struct {
	Ele_name string
	Click_hh string
	Click_cc string
}

func ExecutableClickGenerator(linkList []string, serviceContext []ServiceContext, user_defined_element []User_defined_element) error {
	err1 := UDFCompiler(user_defined_element)
	if err1 != nil {
		return err1
	}
	for _, ele := range user_defined_element {
		linkList = append(linkList, strings.ToLower(ele.Ele_name))
		eleName := strings.ToLower(ele.Ele_name)
		headerFilePath := UDFPATH + " " + eleName + ".hh"
		sourceFilePath := UPFPATH + " " + eleName + ".cc"
		sc = serviceContext{sourceFilePath, headerFilePath, eleName + "-" + eleName}
		serviceContext = append(serviceContext, sc)
	}
	err2 := SCCompiler(serviceContext)
	if err2 != nil {
		return err2
	}
	err3 := Linker(linkList)
	if err3 != nil {
		return err3
	}

}
