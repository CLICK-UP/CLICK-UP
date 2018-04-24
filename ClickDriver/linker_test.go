package ClickDriver

import (
	"fmt"
	"testing"
)

func TestLinker(t *testing.T) {
	var linkList []string
	str := "iprewriter"
	linkList = append(linkList, str)

	err := Linker(linkList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Linker success!!!")
	}
}
