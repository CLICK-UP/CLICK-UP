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
	cmd.Dir = CLICKDIR
	stdout, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		log.Fatal("linker 34 execute link cmd error : ", stdoutErr)
	}
	defer stdout.Close()
	if err = cmd.Run(); err != nil {
		log.Fatal("linker 38 execute link cmd error : ", err)
	}
	log.Println("linker 41 link complete")
	return err
}
