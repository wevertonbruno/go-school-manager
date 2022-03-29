package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"students-manager/adapters/repositories"
	"students-manager/core/dto"
	"students-manager/core/services"
)

func main() {
	repository := &repositories.InMemoryStudentRepository{}
	service := &services.StudentService{
		StudentRepository: repository,
	}

	http.HandleFunc("/students", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "POST":
			decoder := json.NewDecoder(req.Body)
			var input dto.InputStudentDTO
			err := decoder.Decode(&input)
			if err != nil {
				panic(err)
			}
			studentDTO, err := service.Create(input)
			if err != nil {
				panic(err)
			}

			response, err := json.Marshal(studentDTO)
			if err != nil {
				panic(err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.Write(response)

		case "GET":
			response, err := json.Marshal(service.FindAll())
			if err != nil {
				panic(err)
			}

			res.Header().Set("Content-Type", "application/json")
			res.Write(response)
		default:
			fmt.Fprintf(res, "Sorry, only GET and POST methods are supported.")
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
