package ClickDriver

/*
*  Author : @ychuang
*  Create date : 2018-4-23
*  Input :
*		linkList : elements in Click configure file needed to link together
*  Output :
*		err : the error information of link action
 */

import (
	"bytes"
	"io/ioutil"
	"log"
	"os/exec"
)

const (
	CXXLINK     = "g++ -g -O2 -W -Wall  -o click -rdynamic "
	CXXLINKFLAG = " elements.o click.o libclick.a `../bin/click-buildtool --otherlibs` "
)

func Linker(linkList []string) error {
	var linkBuffer bytes.Buffer
	var err error
	for _, v := range linkList {
		linkBuffer.WriteString(v + ".o ")
	}
	cmd := exec.Command(NAME, SECONDTAG, CXXLINK+linkBuffer.String()+CXXLINKFLAG)
	cmd.Path = CLICKDIR
	stdout, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		log.Fatal(stdoutErr)
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, readErr := ioutil.ReadAll(stdout)
	if readErr != nil {
		log.Fatal(readErr)
	}
	log.Println(string(opBytes))
	return err
}
