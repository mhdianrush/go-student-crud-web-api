package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mhdianrush/go-student-crud-web-api/entities"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/student/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/student/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		var student entities.Student

		student.FullName = r.FormValue("full_name")
		student.StudentUniqueId = r.FormValue("student_unique_id")
		student.Gender = r.FormValue("gender")
		student.BirthPlace = r.FormValue("birth_place")
		student.BirthDay = r.FormValue("birth_day")
		student.Address = r.FormValue("address")
		student.PhoneNumber = r.FormValue("phone_number")

		fmt.Println(student)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Running")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Running")
}
