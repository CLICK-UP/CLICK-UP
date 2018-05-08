package ClickDriver

/*
*  Author : @ychuang
*  Create date : 2018-4-23
*  Input :
*		serviceContext : in order to generate and compile elements.cc
*		user_defined_element : in order to generate and compile udf.cc
*  Output :
*		err : the error information of two methods above
 */

import (
	"ServiceContext"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

const (
	NAME           = "/bin/sh"
	SECONDTAG      = "-c"
	UDFFILEPATH    = "./udf/"
	UDFPATH        = "../../udf/"
	CLICKDIR       = "./click/userlevel/"
	CLICKBUILDTOOL = "click-buildtool elem2export < elements.conf > ./elements.cc"
	CXX            = "g++ -DHAVE_CONFIG_H -I../include -I../include -I. -I..  -DCLICK_USERLEVEL -g -O2 -W -Wall -MD -MP -c "
	CXXFLAG        = " -o "
	CXXELEMENTS    = "g++ -DHAVE_CONFIG_H -I../include -I../include -I. -I..  -DCLICK_USERLEVEL -g -O2 -W -Wall -MD -MP -c elements.cc -o elements.o"
)

func UDFCompiler(user_defined_element []User_defined_element) error {
	var err error
	//write source code and header code to strings.ToLowr(ele.eleName).cc&&.hh, and then compile
	for _, ele := range user_defined_element {
		eleName := strings.ToLower(ele.Ele_name)
		headerFilePath := UDFFILEPATH + eleName + ".hh"
		sourceFilePath := UDFFILEPATH + eleName + ".cc"
		headerFileByte := []byte(ele.Click_hh)
		sourceFileByte := []byte(ele.Click_cc)
		err1 := ioutil.WriteFile(headerFilePath, headerFileByte, 0777)
		if err1 != nil {
			log.Fatal("compiler 44 write header file error : ", err1)
		}
		err2 := ioutil.WriteFile(sourceFilePath, sourceFileByte, 0777)
		if err2 != nil {
			log.Fatal("compiler 48 write source file error : ", err2)
		}
		compileCmd := CXX + UDFPATH + eleName + ".cc" + CXXFLAG + eleName + ".o"
		cmd := exec.Command(NAME, SECONDTAG, compileCmd)
		cmd.Dir = CLICKDIR
		stdout, stdoutErr1 := cmd.StdoutPipe()
		if stdoutErr1 != nil {
			log.Fatal("compiler 56 open stdout error : ", stdoutErr1)
		}
		defer stdout.Close()
		if err = cmd.Start(); err != nil {
			log.Fatal("compiler 60 compile udf cmd execute error : ", err)
		}
		log.Println("compiler 62 udf compile complete")
	}
	return err
}

func SCCompiler(serviceContext []ServiceContext.ServiceContext) error {
	var SCByte []byte
	var err error

	for _, sc := range serviceContext {
		tempscByte := []byte(sc.ToString() + "\n")
		SCByte = append(SCByte, tempscByte...)
	}
	errIO := ioutil.WriteFile(CLICKDIR+"elements.conf", SCByte, 0777)
	if errIO != nil {
		log.Fatal("compiler 77 read elements.conf error : ", errIO)
	}

	click2exportCmd := exec.Command(NAME, SECONDTAG, CLICKBUILDTOOL)
	click2exportCmd.Dir = CLICKDIR
	click2exportStdout, click2exportErr := click2exportCmd.StdoutPipe()
	if click2exportErr != nil {
		log.Fatal("compiler 84 execute click2export error : ", click2exportErr)
	}
	defer click2exportStdout.Close()
	if exportErr := click2exportCmd.Start(); exportErr != nil {
		log.Fatal("compiler 88 execute click2export error : ", exportErr)
	}

	log.Println("compiler 91 click2export complete")

	cmd := exec.Command(NAME, SECONDTAG, CXXELEMENTS)
	cmd.Dir = CLICKDIR
	stdout, stdoutErr1 := cmd.StdoutPipe()
	if stdoutErr1 != nil {
		log.Fatal("compiler 97 execute compile elements.cc error : ", stdoutErr1)
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		log.Fatal("compiler 101 execute compile elements.cc error : ", err)
	}
	log.Println("compiler 103 elements.cc compile complete")

	return err
}
