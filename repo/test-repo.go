package repo

import (
	"context"

	"github.com/Cumbercubie/model"
)

func CreateTest(test model.Test) error {
	// return &Test{TestId: 1, Name: "New test", QuestionNum: 1}
	_, err := model.TestDB.Col.InsertOne(context.TODO(), test)
	if err != nil {
		return err
	}
	return nil
}
