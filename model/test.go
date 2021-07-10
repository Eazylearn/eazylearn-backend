package model

type Test struct {
	TestId      int32  `json:"id,omitempty"`
	Name        string `json:"name"`
	QuestionNum int32  `json:"questionNum"`
	// TimeAllowed time.Time
}

func CreateTest() *Test {
	return &Test{TestId: 1, Name: "New test", QuestionNum: 1}
}
