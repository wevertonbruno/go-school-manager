package repositories

import (
	"students-manager/core/entities"
)

type StudentRepository interface {
	Save(s entities.Student) (entities.Student, error)
	FindAll() []entities.Student
	FindById(ID string) (entities.Student, error)
}
