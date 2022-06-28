package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	handler := new(studentsHandler)

	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentsHandler struct{}

func (h studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")

	switch len(pathSegments) {
	case 2:
		h.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data, err := h.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("failed to serialize studentsL %q", err))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err == ErrStudentNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := h.toJSON(*student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("failed to serialize student: %q", err))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err == ErrStudentNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var grade Grade
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&grade)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("error decoding grade: %q", err))
		return
	}

	student.Grades = append(student.Grades, grade)

	w.WriteHeader(http.StatusCreated)
}

func (s studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var buff bytes.Buffer
	enc := json.NewEncoder(&buff)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
