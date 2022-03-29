package services

import (
	"github.com/google/uuid"
	"students-manager/core/dto"
	"students-manager/core/entities"
	"students-manager/core/repositories"
)

type StudentService struct {
	StudentRepository repositories.StudentRepository
}

func (c StudentService) Create(input dto.InputStudentDTO) (dto.OutputStudentDTO, error) {
	student := entities.Student{
		ID:        uuid.New().String(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Age:       input.Age,
	}

	student, err := c.StudentRepository.Save(student)
	if err != nil {
		return dto.OutputStudentDTO{}, err
	}

	output := dto.OutputStudentDTO{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}

	return output, err
}

func (c StudentService) FindAll() []dto.OutputStudentDTO {
	students := c.StudentRepository.FindAll()
	output := make([]dto.OutputStudentDTO, len(students))

	for key, value := range students {
		output[key] = dto.OutputStudentDTO{ID: value.ID, FirstName: value.FirstName, LastName: value.LastName}
	}

	return output
}

func (c StudentService) FindById(ID string) (dto.OutputStudentDTO, error) {
	student, err := c.StudentRepository.FindById(ID)
	if err != nil {
		return dto.OutputStudentDTO{}, err
	}

	studentDTO := dto.OutputStudentDTO{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}

	return studentDTO, nil
}
