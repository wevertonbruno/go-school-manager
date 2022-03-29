package repositories

import (
	"fmt"
	"students-manager/core/entities"
)

type InMemoryStudentRepository struct {
	Data []entities.Student
}

func (r *InMemoryStudentRepository) Save(s entities.Student) (entities.Student, error) {
	r.Data = append(r.Data, s)
	return s, nil
}

func (r InMemoryStudentRepository) FindAll() []entities.Student {
	return r.Data
}

func (r InMemoryStudentRepository) FindById(ID string) (entities.Student, error) {
	for _, value := range r.Data {
		if value.ID == ID {
			return value, nil
		}
	}

	return entities.Student{}, fmt.Errorf("Student with ID(%v) not found!", ID)
}
