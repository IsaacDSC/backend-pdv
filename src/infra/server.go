package infra

import (
	"backend-pdv/src/infra/routes"
	"fmt"
	"net/http"
)

func StartServer() {
	routes.StartRoutes()
	fmt.Println("[ * ] started application listen in port 3333")
	http.ListenAndServe(":3333", nil)
}
