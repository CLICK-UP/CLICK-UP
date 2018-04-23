package ClickDirver

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
	"strings"
	"os/exec"
	"io/ioutil"
	"log"
)

const (
	NAME := "/bin/echo"
	UDFPATH := "../../udf/"
	CLICKBUILDTOOL := "click-buildtool"
	CLICK2EXPORT := "elem2export < elements.conf > "
	CXX := "g++ -DHAVE_CONFIG_H -I../include -I../include -I. -I..  -DCLICK_USERLEVEL -g -O2 -W -Wall -MD -MP -c "
	CXXFLAG := " -o " 
	CXXELEMENTS := "g++ -DHAVE_CONFIG_H -I../include -I../include -I. -I..  -DCLICK_USERLEVEL -g -O2 -W -Wall -MD -MP -c elements.cc -o elements.o"
)

func UDFCompiler(user_defined_element []User_defined_element) error {
	var err error
	//write source code and header code to strings.ToLowr(ele.eleName).cc&&.hh, and then compile 
	for _, ele := range user_defined_element{
		eleName := strings.ToLower(ele.Ele_name)
		headerFilePath := UDFPATH + " " + eleName + ".hh"
		sourceFilePath := UPFPATH + " " + eleName + ".cc"
		headerFileByte := []byte(ele.Click_hh)
		sourceFileByte := []byte(ele.Click_cc)
		err1 := ioutil.WriteFile(headerFilePath, headerFileByte, 0777)
		if err1 != nil {
			log.Fatal(err1)
		}
		err2 := ioutil.WriteFile(sourceFilePath, headerFileByte, 0777)
		if err2 != nil {
			log.Fatal(err2)
		}
		compileCmd := CXX + UDFPATH + eleName + ".cc" + CXXFLAG + eleName + ".o" 
		cmd := exec.Command(NAME, compileCmd)
		stdout, stdoutErr1 := cmd.StdoutPipe()
		if stdoutErr1 != nil {
			log.Fatal(stdoutErr1)
		}
		defer stdout.Close()
		if err = cmd.Start(); err != nil {
			log.Fatal(err)
		}
		opBytes, stdoutErr2 := ioutil.ReadAll(stdout)
		if stdoutErr2 != nil {
			log.Fatal(stdoutErr2)
		}
		log.Println(string(opBytes))
	}
	return err
}

func SCCompiler(serviceContext []ServiceContext) err {
	var SCByte []byte
	var err error

	for _, sc := range serviceContext {
		tempscByte := []byte(sc.ToString() + "\n")
		SCByte = append(SCByte, tempscByte...)
	}
	errIO := ioutil.WriteFile("elements.conf", SCByte, 0777)
	if errIO != nil {
		log.Fatal(errIO)
	}

	click2exportCmd := exec.Command(CLICKBULIDTOOL, CLICK2EXPORT)
	click2exportStdout, click2exportErr := click2exportCmd.StdoutPipe()
	if click2exportErr != nil {
		log.Fatal(click2exportErr)
	}
	defer click2exportStdout.Close()
	if exportErr := click2exportCmd.Start(); exportErr != nil {
		log.Fatal(exportErr)
	}
	exportOpBytes, click2exportErr2 := ioutil.ReadAll(click2exportStdout)
	if click2exportErr2 != nil {
		log.Fatal(click2exportErr2)
	}
	log.Println(string(exportOpBytes))

	cmd := exec.Command(NAME, CXXELEMENTS)
	stdout, stdoutErr1 := cmd.StdoutPipe()
	if stdoutErr1 != nil {
		log.Fatal(stdoutErr1)
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, stdoutErr2 := ioutil.ReadAll(stdout)
	if stdoutErr2 != nil {
		log.Fatal(stdoutErr2)
	}
	log.Println(string(opBytes))

	return err
}
