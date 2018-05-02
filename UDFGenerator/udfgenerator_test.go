package udfgenerator

import (
	"fmt"
	"testing"
)

func TestUdfgenerator(t *testing.T) {
	var udf []UserDefinedElement
	var actions []string = []string{"state_machine"}
	udf = append(udf, UserDefinedElement{"Firewall", actions})
	UDF, err := Udfgenerator(udf)
	if err != nil {
		fmt.Println("error happened")
	}
	for _, v := range UDF {
		fmt.Println(v.Ele_name)
		fmt.Println(v.Click_hh)
		fmt.Println(v.Click_cc)
	}
}
