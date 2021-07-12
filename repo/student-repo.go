package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/Cumbercubie/model"
	"go.mongodb.org/mongo-driver/bson"
)

// func GetStudentById(id int32) *model.Student {
// 	return &model.Student{Id: id, Name: "CaoThienHuan", Age: 12, BOD: "12/02/21", Grade: 10}
// 	// return model.TestDB
// }

func GetUserByFirstname(firstName string) ([]model.Student, error) {
	list := make([]model.Student, 0)
	result, qErr := model.UserDB.Col.Find(context.TODO(), bson.M{})
	if qErr != nil {
		log.Println("Error finding", qErr.Error())
		return list, qErr
	}
	fmt.Println(result == nil)
	// var user model.Student
	// err := result.Decode(&user)
	err := result.All(context.TODO(), &list)

	if err != nil {
		log.Println("Error decoding", err.Error())
		return list, err
	}
	// list = append(list, user)
	return list, nil
}
