package repo

import "github.com/Cumbercubie/model"

func GetStudentById(id int32) *model.Student {
	return &model.Student{Id: id, Name: "CaoThienHuan", Age: 12, BOD: "12/02/21", Grade: 10}
}
