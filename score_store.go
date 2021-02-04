package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/buntdb"
)

type ScoreStore struct {
	db *buntdb.DB
}

type ExamResults struct {
	ExamId  int64
	Average float64
	Results []Exam
}

type Exam struct {
	ExamId    int64   `json:",omitempty"`
	Score     float64 `json:",omitempty"`
	StudentId string  `json:",omitempty"`
}

type Student struct {
	StudentId string
	Exams     []Exam  `json:",omitempty"`
	Average   float64 `json:",omitempty"`
}

func NewScoreStore() *ScoreStore {
	s := ScoreStore{}
	db, err := buntdb.Open(":memory:")

	if err != nil {
		panic(err)
	}

	if err := db.CreateIndex("students", "student:*", buntdb.IndexString); err != nil {
		panic(err)
	}

	if err := db.CreateIndex("exams", "exam:*", buntdb.IndexString); err != nil {
		panic(err)
	}

	if err := db.CreateIndex("scores", "score:*", buntdb.IndexString); err != nil {
		panic(err)
	}

	s.db = db

	go s.keepScore()

	return &s
}

func (store *ScoreStore) keepScore() {
	events := make(chan Event)
	go NewEventStream("http://live-test-scores.herokuapp.com/scores", events, http.DefaultClient)

	for event := range events {
		store.addScore(event)
	}
}

func (store *ScoreStore) addScore(ev Event) {
	exam_key := fmt.Sprintf("exam:%d", ev.data.Exam)
	student_key := fmt.Sprintf("student:%s", ev.data.StudentId)
	score_key := fmt.Sprintf("score:%s:%d", ev.data.StudentId, ev.data.Exam)

	store.db.Update(func(tx *buntdb.Tx) error {
		tx.Set(exam_key, strconv.Itoa(ev.data.Exam), nil)
		tx.Set(student_key, ev.data.StudentId, nil)
		tx.Set(score_key, strconv.FormatFloat(ev.data.Score, 'g', -1, 64), nil)

		return nil
	})
}

func (store *ScoreStore) getStudents() ([]Student, error) {
	students := make([]Student, 0)
	err := store.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("students", func(key, value string) bool {
			students = append(students, Student{
				StudentId: value,
			})
			return true
		})

		return err
	})

	return students, err
}

func (store *ScoreStore) getStudent(studentId string) (*Student, error) {
	var s *Student
	err := store.db.View(func(tx *buntdb.Tx) error {
		total := 0.0
		count := 0.0
		exams := make([]Exam, 0)
		err := tx.AscendKeys(fmt.Sprintf("score:%s:*", studentId), func(key, value string) bool {
			score, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return false
			}

			examId, err := strconv.ParseInt(strings.Split(key, ":")[2], 0, 64)
			exams = append(exams, Exam{
				ExamId: examId,
				Score:  score,
			})

			count++
			total += score
			return true
		})

		if err != nil {
			return err
		}

		s = &Student{
			StudentId: studentId,
			Average:   total / count,
			Exams:     exams,
		}

		return nil
	})

	return s, err
}

func (store *ScoreStore) getExams() ([]Exam, error) {
	exams := make([]Exam, 0)
	err := store.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("exams", func(key, value string) bool {
			id, err := strconv.ParseInt(value, 0, 64)

			if err != nil {
				return false
			}

			exams = append(exams, Exam{
				ExamId: id,
			})

			return true
		})

		return err
	})

	return exams, err
}

func (store *ScoreStore) getExam(examId string) (*ExamResults, error) {
	var s *ExamResults
	err := store.db.View(func(tx *buntdb.Tx) error {
		total := 0.0
		count := 0.0
		exams := make([]Exam, 0)
		id, err := strconv.ParseInt(examId, 0, 64)

		if err != nil {
			return err
		}

		err = tx.AscendKeys(fmt.Sprintf("score:*:%s", examId), func(key, value string) bool {
			score, err := strconv.ParseFloat(value, 64)

			if err != nil {
				return false
			}

			exams = append(exams, Exam{
				Score: score,
				StudentId: strings.Split(key, ":")[1],
			})

			count++
			total += score
			return true
		})

		if err != nil {
			return err
		}

		s = &ExamResults{
			ExamId:  id,
			Average: total / count,
			Results: exams,
		}

		return nil
	})

	return s, err
}
