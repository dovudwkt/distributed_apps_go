package grades

import (
	"errors"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Average() float32 {
	var sum float32
	for _, grade := range s.Grades {
		sum += grade.Score
	}

	return sum / float32(len(s.Grades))
}

type Students []Student

var ErrStudentNotFound = errors.New("ErrStudentNotFound")

func (s Students) GetByID(id int) (*Student, error) {
	for i, student := range s {
		if student.ID == id {
			return &s[i], nil
		}
	}
	return nil, ErrStudentNotFound
}

var (
	students      Students
	studentsMutex sync.Mutex
)

type GradeType string

const (
	GradeTest     GradeType = "Test"
	GradeHomework GradeType = "Homework"
	GradeQuiz     GradeType = "Quiz"
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
