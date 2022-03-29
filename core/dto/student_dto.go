package dto

type InputStudentDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       uint8  `json:"age"`
}

type OutputStudentDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
