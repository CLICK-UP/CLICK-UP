package main

/*
*  Author : @ychuang
*  Create date : 2018-4-24
*  Input :
*		{
*			“id” : 1,
*			“jsonrpc” : 2.0,
*			“method” : “update”,
*			“params” : {
*				“vnf_config” : “$click_config_file”,
*				“element_list” : [
*					“$element_name1”,
*					“$element_name2”,
*					“$element_namen”
*				],
*				“user_defined_element” : [
*					{
*						“element_name” : “$element_name1”,
*						“atom_action” : [
*							"$atom_action1",
*							"$atom_action2",
*							"$atom_actionn",
*						]
*					},
*					{
*						“element_name” : “$element_name2”,
*						“atom_action” : [
*							"$atom_action1",
*							"$atom_action2",
*							"$atom_actionn",
*						]
*					}
*				]
*			}
*		}
*  Output :
*		err : the error information of link action
 */

import (
	"ClickDirver"
	"ServiceContext"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"udfgenerator"
)

type UpdateClick struct {
	Id      int
	Jsonrpc float32
	Method  string
	Params  ParamsStruct
}

type ParamsStruct struct {
	Vnf_config           string
	Element_list         []string
	User_defined_element []UserDefinedElement
}

func main() {
	http.HandleFunc("/update", updateHander)

	log.Println("starting httpserver... v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func updateHander(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updateClick UpdateClick
	jsonErr := json.Unmarshal(body, &updateClick)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//first step is to generate the user define element
	user_defined_element, udfgeneratorErr := udfgenerator.Udfgenerator(updateClick.Params.User_defined_element)
	if udfgeneratorErr != nil {
		fmt.Fprintf(w, udfgeneratorErr)
	}
	linkList, serviceContext, scErr := ServiceContext.GetServiceContextFromModuleList(updateClick.Params.Element_list)
	if scErr != nil {
		fmt.Fprintf(w, scErr)
	}
	genErr := ClickDirver.ExecutableClickGenerator(linkList, serviceContext, user_defined_element)

	//response to frontend (hasn'n define)
	fmt.Fprintf(w, genErr)
}
