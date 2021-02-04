package main

import (
	"testing"

	"github.com/tidwall/buntdb"
)

func TestNewScoreStore(t *testing.T) {
	store := NewScoreStore()

	indexes, _ := store.db.Indexes()

	if l := len(indexes); l != 3 {
		t.Errorf("Expected: Indexes %d == 3", l)
	}
}

func TestAddScore(t *testing.T) {
	store := NewScoreStore()

	ev := Event{
		name: "score",
		data: ScoreEvent{
			StudentId: "student.1",
			Exam:      123,
			Score:     0.99,
		},
	}

	store.addScore(ev)

	store.db.View(func(tx *buntdb.Tx) error {
		if l, _ := tx.Len(); l != 3 {
			t.Errorf("Expected: length %d == 3", l)
		}

		return nil
	})
}

func TestGetStudents(t *testing.T) {
	store := buildAStore()
	students, _ := store.getStudents()

	if l := len(students); l != 2 {
		t.Errorf("Expected: students %d == 2", l)
	}
}

func TestGetStudent(t *testing.T) {
	store := buildAStore()
	student, _ := store.getStudent("student.2")

	if l := len(student.Exams); l != 2 {
		t.Errorf("Expected: exams %d == 2", l)
	}

	if student.Average != 0.25 {
		t.Errorf("Expected: average %f == 0.25", student.Average)
	}
}

func TestGetExams(t *testing.T) {
	store := buildAStore()
	exams, _ := store.getExams()

	if l := len(exams); l != 2 {
		t.Errorf("Expected: exams %d == 2", l)
	}
}

func TestGetExam(t *testing.T) {
	store := buildAStore()
	exam, _ := store.getExam("1")

	if l := len(exam.Results); l != 2 {
		t.Errorf("Expected: results %d == 2", l)
	}

	if exam.Average != 0.625 {
		t.Errorf("Expected: average %f == 0.625", exam.Average)
	}
}

