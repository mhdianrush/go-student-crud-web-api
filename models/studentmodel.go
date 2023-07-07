package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mhdianrush/go-student-crud-web-api/config"
	"github.com/mhdianrush/go-student-crud-web-api/entities"
)

type StudentModel struct {
	Connection *sql.DB
}

func NewStudentModel() *StudentModel {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	return &StudentModel{
		Connection: db,
	}
}

func (s *StudentModel) FindAll() ([]entities.Student, error) {
	rows, err := s.Connection.Query(
		`select * from student`,
	)
	if err != nil {
		return []entities.Student{}, err
	}
	defer rows.Close()

	var dataStudent []entities.Student

	for rows.Next() {
		var student entities.Student
		rows.Scan(
			&student.Id,
			&student.FullName,
			&student.StudentUniqueId,
			&student.Gender,
			&student.BirthPlace,
			&student.BirthDay,
			&student.Address,
			&student.PhoneNumber,
		)
		if student.Gender == "1" {
			student.Gender = "Male"
		}

		if student.Gender == "2" {
			student.Gender = "Female"
		}
		// // // 2006-01-02 ==> YYYY-MM-DD
		dateTime, _ := time.Parse("2006-01-02T00:00:00Z", student.BirthDay)
		// // // // // 02-01-2006 ==> DD-MM-YYYY (indonesian time)
		student.BirthDay = dateTime.Format("02-01-2006")

		dataStudent = append(dataStudent, student)
	}
	return dataStudent, nil
}

func (s *StudentModel) Create(student entities.Student) bool {
	result, err := s.Connection.Exec(
		`insert into student(full_name, student_unique_id, gender, birth_place, birth_day, address, phone_number)
		values(?, ?, ?, ?, ?, ?, ?)`,
		student.FullName, student.StudentUniqueId, student.Gender, student.BirthPlace, student.BirthDay,
		student.Address, student.PhoneNumber,
	)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
}
