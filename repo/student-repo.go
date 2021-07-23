package repo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Cumbercubie/model"
	"go.mongodb.org/mongo-driver/bson"
)

// func GetStudentById(id ingot32) *model.Student {
// 	return &model.Student{Id: id, Name: "CaoThienHuan", Age: 12, BOD: "12/02/21", Grade: 10}
// 	// return model.TestDB
// }
func GetUserByFirstname(firstName string) ([]model.Student, error) {
	list := make([]model.Student, 0)
	fmt.Println(firstName)
	result := model.UserDB.Col.FindOne(context.TODO(), bson.M{"firstName": firstName})
	if result == nil {
		return list, errors.New("Không tìm thấy user này")
	}
	var foundStudent model.Student
	err := result.Decode(&foundStudent)
	if err != nil {
		return list, err
	}
	list = append(list, foundStudent)
	return list, nil
}

func GetAllUser(firstName string) ([]model.Student, error) {
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
