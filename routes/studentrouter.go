package routes

import (
	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-student-crud-web-api/controllers"
)

func StudentRouter(r *mux.Router) {
	router := r.PathPrefix("/student").Subrouter()

	router.HandleFunc("", controllers.Index)
	router.HandleFunc("/add", controllers.Add)
	router.HandleFunc("/edit", controllers.Edit)
	router.HandleFunc("/delete", controllers.Delete)
}
