package model

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id    int32   `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	BOD   string  `json:"bod,omitempty"`
	Age   int32   `json:"age,omitempty"`
	Grade float32 `json:"grade,omitempty"`
}

func (s Student) String() string {
	sjson, _ := json.Marshal(s)
	return fmt.Sprint(sjson)

}
