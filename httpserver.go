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
	"ClickDriver"
	"ServiceContext"
	"clickexecutor"
	"confgenerator"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"udfgenerator"
)

var (
	ProcessId int
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
	User_defined_element []udfgenerator.UserDefinedElement
}

func main() {
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("resource"))))

	log.Println("starting httpserver... v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

	killErr := clickexecutor.ClickKill(ProcessId)
	if killErr != nil {
		fmt.Fprintf(w, "kill current click process error : %v", killErr)
	} else {
		fmt.Fprintf(w, "delete success")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./resource/index.tmpl")
	if err != nil {
		fmt.Fprintf(w, "httpserver 84 err: %v", err)
	}

	err = tmpl.Execute(w, r)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updateClick UpdateClick
	jsonErr := json.Unmarshal(body, &updateClick)
	if jsonErr != nil {
		log.Fatal("httpserver 99 parse json err:", jsonErr)
	}
	//fmt.Println(updateClick)

	//first step is to generate the user define element
	if strings.EqualFold(updateClick.Method, "update") {
		//kill process
		killErr := clickexecutor.ClickKill(ProcessId)
		if killErr != nil {
			fmt.Fprintf(w, "kill current click process error : %v", killErr)
		}
		//write confgure file
		err = confgenerator.ConfGenerator(updateClick.Params.Vnf_config)
		if err != nil {
			fmt.Fprintf(w, "write configure file error : ", err)
		}
		start := time.Now()
		fmt.Println("update start time : ", start)
		user_defined_element, udfgeneratorErr := udfgenerator.Udfgenerator(updateClick.Params.User_defined_element)
		if udfgeneratorErr != nil {
			fmt.Fprintf(w, "udfgeneratorErr: %v", udfgeneratorErr)
		}
		linkList, serviceContext, scErr := ServiceContext.GetServiceContextFromModuleList(updateClick.Params.Element_list)
		if scErr != nil {
			fmt.Fprintf(w, "scErr: %v", scErr)
		}
		genErr := ClickDriver.ExecutableClickGenerator(linkList, serviceContext, user_defined_element)

		//response to frontend (hasn'n define)
		fmt.Fprintf(w, "genErr: %v", genErr)
		end := time.Now()
		fmt.Println("update end time = ", end)
		sub := end.Sub(start)
		fmt.Println("end time - start time = ", sub)
		//start process
		var exeErr error
		ProcessId, exeErr = clickexecutor.ClickExecutor()
		if exeErr != nil {
			fmt.Fprintf(w, "exeErr: %v", exeErr)
		}
	} else if strings.EqualFold(updateClick.Method, "create") {
		err = confgenerator.ConfGenerator(updateClick.Params.Vnf_config)
		if err != nil {
			fmt.Fprintf(w, "write configure file error : ", err)
		}
		start := time.Now()
		fmt.Println("create start time : ", start)
		user_defined_element, udfgeneratorErr := udfgenerator.Udfgenerator(updateClick.Params.User_defined_element)
		if udfgeneratorErr != nil {
			fmt.Fprintf(w, "udfgeneratorErr: %v", udfgeneratorErr)
		}
		linkList, serviceContext, scErr := ServiceContext.GetServiceContextFromModuleList(updateClick.Params.Element_list)
		if scErr != nil {
			fmt.Fprintf(w, "scErr: %v", scErr)
		}
		genErr := ClickDriver.ExecutableClickGenerator(linkList, serviceContext, user_defined_element)

		//response to frontend (hasn'n define)
		fmt.Fprintf(w, "genErr: %v", genErr)
		end := time.Now()
		fmt.Println("create end time = ", end)
		sub := end.Sub(start)
		fmt.Println("end time - start time = ", sub)

		//start process
		var exeErr error
		ProcessId, exeErr = clickexecutor.ClickExecutor()
		if exeErr != nil {
			fmt.Fprintf(w, "exeErr: %v", exeErr)
		}

	} else {
		fmt.Fprintf(w, "error request")
	}

}
