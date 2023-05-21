package controllers

import (
	"backend-pdv/lib"
	"backend-pdv/src/app/helpers"
	"backend-pdv/src/shared"
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
	var body shared.CreateCategory_DTO
	err = json.Unmarshal(data, &body)
	if err != nil {
		res.WriteHeader(http.StatusNotAcceptable)
		return
	}
	service := shared.GetInstanceCategoryService(body)
	list_errors := service.CreateCategory()
	if len(list_errors) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		output := helpers.BuildResponseListError(list_errors)
		output_json, _ := json.Marshal(output)
		res.Write(output_json)
		return
	}
	res.WriteHeader(http.StatusAccepted)
}
