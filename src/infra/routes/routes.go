package routes

import (
	"backend-pdv/src/infra/controllers"
	"net/http"
)

func StartRoutes() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("{\"status\":\" true\"}"))
	})
	http.HandleFunc("/enterprise", controllers.EnterpriseController)
	http.HandleFunc("/category", controllers.CategoryController)
	http.HandleFunc("/product", controllers.ProductController)
	http.HandleFunc("/client", controllers.ClientController)
}
