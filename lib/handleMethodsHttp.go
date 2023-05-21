package lib

import (
	"net/http"
	"strings"
)

type Methods struct {
	Res    http.ResponseWriter
	Req    *http.Request
	GET    http.HandlerFunc
	POST   http.HandlerFunc
	PUT    http.HandlerFunc
	DELETE http.HandlerFunc
}

func HandlerMethodsHttp(routes Methods) {
	switch strings.ToUpper(strings.ToUpper(routes.Req.Method)) {
	case "POST":
		if routes.POST != nil {
			routes.POST(routes.Res, routes.Req)
			return
		}
		routes.Res.WriteHeader(http.StatusNotFound)
		return

	case "GET":
		if routes.GET != nil {
			routes.GET(routes.Res, routes.Req)
			return
		}
		routes.Res.WriteHeader(http.StatusNotFound)
		return
	case "PUT":
		if routes.PUT != nil {
			routes.PUT(routes.Res, routes.Req)
			return
		}
		routes.Res.WriteHeader(http.StatusNotFound)
		return
	case "DELETE":
		if routes.DELETE != nil {
			routes.DELETE(routes.Res, routes.Req)
			return
		}
		routes.Res.WriteHeader(http.StatusNotFound)
		return
	default:
		return
	}
}
