package controllers

import (
	"backend-pdv/lib"
	"backend-pdv/src/app/helpers"
	"backend-pdv/src/shared"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ProductController(res http.ResponseWriter, req *http.Request) {
	lib.HandlerMethodsHttp(lib.Methods{
		Res:  res,
		Req:  req,
		POST: post_product,
	})
}

func post_product(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	var body shared.CreateProduct_DTO
	err = json.Unmarshal(data, &body)
	if err != nil {
		res.WriteHeader(http.StatusNotAcceptable)
		return
	}
	service := shared.GetInstanceProductService(body)
	list_errors := service.CreateProduct()
	if len(list_errors) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		output := helpers.BuildResponseListError(list_errors)
		output_json, _ := json.Marshal(output)
		res.Write(output_json)
		return
	}
	res.WriteHeader(http.StatusAccepted)
}
