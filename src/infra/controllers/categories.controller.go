package controllers

import (
	"backend-pdv/lib"
	"backend-pdv/src/app/helpers"
	shared_instances "backend-pdv/src/shared/instances"
	shared_types "backend-pdv/src/shared/types"
	"encoding/json"
	"io"
	"net/http"
)

func CategoryController(res http.ResponseWriter, req *http.Request) {
	lib.HandlerMethodsHttp(lib.Methods{
		Res:  res,
		Req:  req,
		POST: post_category,
	})
}

func post_category(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	var body shared_types.CreateCategory_DTO
	err = json.Unmarshal(data, &body)
	if err != nil {
		res.WriteHeader(http.StatusNotAcceptable)
		return
	}
	service := shared_instances.GetInstanceCategoryService(body)
	client, list_errors := service.CreateCategory()
	if len(list_errors) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		output := helpers.BuildResponseListError(list_errors)
		output_json, _ := json.Marshal(output)
		res.Write(output_json)
		return
	}
	output, _ := json.Marshal(client)
	res.WriteHeader(http.StatusCreated)
	res.Write(output)
}
