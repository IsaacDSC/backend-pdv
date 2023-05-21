package controllers

import "net/http"

func OrderController(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	defer req.Body.Close()
}
