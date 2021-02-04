package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
)

func main() {
	store := NewScoreStore()

	log.Println("Server starting: http://0.0.0.0:12933")

	http.HandleFunc("/students", store.HandleGetStudents())
	http.HandleFunc("/student/", store.HandleGetStudent())
	http.HandleFunc("/exams", store.HandleGetExams())
	http.HandleFunc("/exam/", store.HandleGetExam())
	log.Fatal(http.ListenAndServe(":12933", nil))
}

func (store *ScoreStore) HandleGetStudent() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		studentId := path.Base(r.URL.Path)
		student, err := store.getStudent(studentId)

		if student == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(student)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	})
}

func (store *ScoreStore) HandleGetStudents() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		students, err := store.getStudents()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(students)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	})
}

func (store *ScoreStore) HandleGetExams() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		students, err := store.getExams()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(students)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	})
}

func (store *ScoreStore) HandleGetExam() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		examId := path.Base(r.URL.Path)
		exam, err := store.getExam(examId)

		if exam == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(exam)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
	})
}
