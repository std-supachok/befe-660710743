package main
import (
	"fmt"
	"errors"
)

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool {
	return s.GPA >= 3.5
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New(" name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("GPA must be between 0-4")
	}
	return nil
}

func main() {
	// var st Student = Student{ID:"1", Name:"Supachok", Email:"Hangnak_s@silpakorn.edu", Year:4, GPA:3.75}
	students := []Student{
		{ID: "1", Name: "Supachok", Email: "Hangnak_s@silpakorn.edu", Year: 4, GPA: 3.75},
		{ID: "2", Name: "alice", Email: "alice@silpakorn.edu", Year: 4, GPA: 2.75},

	}
	newStudent := Student{ID: "3", Name: "Trudy", Email:"Trudy@hotmail.com", Year:4, GPA:3.50}
	students = append(students, newStudent)

	for i, student:= range students {
		fmt.Printf("%d Honor = %v\n",i, student.IsHonor())
		fmt.Printf("%d Validation = %v\n",i, student.Validate())
	}
}