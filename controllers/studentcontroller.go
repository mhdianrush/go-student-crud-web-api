package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mhdianrush/go-student-crud-web-api/entities"
	"github.com/mhdianrush/go-student-crud-web-api/libraries"
	"github.com/mhdianrush/go-student-crud-web-api/models"
)

var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {
	student, err := models.NewStudentModel().FindAll()
	if err != nil {
		panic(err)
	}
	data := map[string]any{
		"student": student,
	}

	temp, err := template.ParseFiles("views/student/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/student/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
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

		var data = make(map[string]any)
		// validation check
		vErrors := validation.Struct(student)
		if vErrors != nil {
			// so that the value will keep saving while click the save button without fullfilled all input
			data["student"] = student
			
			data["validation"] = vErrors
		} else {
			data["message"] = "data has been saved"
			models.NewStudentModel().Create(student)
		}

		temp, err := template.ParseFiles("views/student/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Running")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete Running")
}
