package controllers

import (
	"backend-pdv/lib"
	"backend-pdv/src/app/helpers"
	"backend-pdv/src/shared"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ClientController(res http.ResponseWriter, req *http.Request) {
	lib.HandlerMethodsHttp(lib.Methods{
		Res:  res,
		Req:  req,
		POST: post_client,
	})
}

func post_client(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	var body = shared.ClientCreateOnSet_DTO{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		res.WriteHeader(http.StatusNotAcceptable)
		return
	}
	service := shared.GetInstanceClientService(body)
	list_errors := service.CreateOnSet()
	if len(list_errors) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		output := helpers.BuildResponseListError(list_errors)
		output_json, _ := json.Marshal(output)
		res.Write(output_json)
		return
	}
	res.WriteHeader(http.StatusAccepted)
}
