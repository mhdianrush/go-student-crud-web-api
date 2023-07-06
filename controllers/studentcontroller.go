package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Running")
}

func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Add Running")
}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Running")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Running")
}
