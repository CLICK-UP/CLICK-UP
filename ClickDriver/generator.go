package ClickDriver

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
	"log"
	"strings"
)

type User_defined_element struct {
	Ele_name string
	Click_hh string
	Click_cc string
}

func ExecutableClickGenerator(linkList []string, serviceContext []ServiceContext.ServiceContext, user_defined_element []User_defined_element) error {
	var err error
	err = UDFCompiler(user_defined_element)
	if err != nil {
		log.Fatal("generator 30 udf compile error : ", err)
	}
	var addSC []ServiceContext.ServiceContext
	for _, ele := range user_defined_element {
		linkList = append(linkList, strings.ToLower(ele.Ele_name))
		eleName := strings.ToLower(ele.Ele_name)
		headerFilePath := UDFPATH + eleName + ".hh"
		sourceFilePath := UDFPATH + eleName + ".cc"
		sc := ServiceContext.ServiceContext{sourceFilePath, `"` + headerFilePath + `"`, ele.Ele_name + "-" + ele.Ele_name}
		addSC = append(addSC, sc)
		serviceContext = append(serviceContext, sc)
	}
	err = SCCompiler(serviceContext)
	if err != nil {
		log.Fatal("generator 42 service context compile error : ", err)
	}
	err = Linker(linkList)
	if err != nil {
		log.Fatal("generator 46 link error : ", err)
	}

	ServiceContext.InsertModuleContext(addSC)
	return err
}
