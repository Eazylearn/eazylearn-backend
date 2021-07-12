package model

import (
	"encoding/json"
	"fmt"

	"github.com/Cumbercubie/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	BirthDate string `json:"birthDate" bson:"birthDate,omitempty"`
}

// type Student struct {
// 	Id        int32   `json:"id,omitempty" bson:`
// 	Name      string  `json:"name,omitempty"`
// 	BOD       string  `json:"bod,omitempty"`
// 	Age       int32   `json:"age,omitempty"`
// 	Grade     float32 `json:"grade,omitempty"`
// 	IsTeacher bool    `json:"isTeacher"`
// }

func (s Student) String() string {
	sjson, _ := json.Marshal(s)
	return fmt.Sprint(sjson)

}

var UserDB = &db.Instance{
	ColName: "drivers",
}

func InitUserModel(m *mongo.Database) {
	UserDB.ApplyDatabase(m)
}
