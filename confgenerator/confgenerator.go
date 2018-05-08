package confgenerator

import (
	"io/ioutil"
)

func ConfGenerator(conf string) error {
	err := ioutil.WriteFile("./confgenerator/run.click", []byte(conf), 0777)
	return err
}
