package model

import (
	"context"
	"log"

	"github.com/Cumbercubie/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test struct {
	TestId      int32  `json:"id,omitempty"`
	Name        string `json:"name"`
	QuestionNum int32  `json:"questionNum"`
	// TimeAllowed time.Time
}

var TestDB = &db.Instance{
	ColName: "test",
}

func CreateTest(test Test) error {
	// return &Test{TestId: 1, Name: "New test", QuestionNum: 1}
	_, err := TestDB.Col.InsertOne(context.TODO(), test)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func InitTestModel(m *mongo.Database) {
	TestDB.ApplyDatabase(m)
}
