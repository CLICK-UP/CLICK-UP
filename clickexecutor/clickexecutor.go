package clickexecutor

import (
	"log"
	"os/exec"
	"strconv"
)

func ClickExecutor() (int, error) {
	var err error
	commend := "./click/userlevel/click ./confgenerator/run.click & "
	cmd := exec.Command("/bin/sh", "-c", commend)
	stdout, stdoutErr1 := cmd.StdoutPipe()
	if stdoutErr1 != nil {
		log.Fatal("clickexecutor 12 open stdout error : ", stdoutErr1)
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		log.Fatal("clickexecutor 16 execute click cmd error : ", err)
	}
	return cmd.Process.Pid, err
}

func ClickKill(processId int) error {
	var err error
	commend := "kill -9 " + strconv.Itoa(processId)
	cmd := exec.Command("/bin/sh", "-c", commend)
	stdout, stdoutErr1 := cmd.StdoutPipe()
	if stdoutErr1 != nil {
		log.Fatal("clickKill 28 open stdout error : ", stdoutErr1)
	}
	defer stdout.Close()
	if err = cmd.Run(); err != nil {
		log.Fatal("clickKill 32 kill click process error : ", err)
	}
	return err
}
