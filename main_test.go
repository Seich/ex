package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetStudent(t *testing.T) {
	store := buildAStore()

	r, _ := http.NewRequest("GET", "/student/student.1", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(store.HandleGetStudent())
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: status %d == 200", status)
	}

	expected := `{"StudentId":"student.1","Exams":[{"ExamId":1,"Score":1}],"Average":1}`
	if rr.Body.String() != expected {
		t.Errorf("Expected: %v == %v", rr.Body.String(), expected)
	}
}

func TestHandleGetStudents(t *testing.T) {
	store := buildAStore()

	r, _ := http.NewRequest("GET", "/students", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(store.HandleGetStudents())
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: status %d == 200", status)
	}

	expected := `[{"StudentId":"student.1"},{"StudentId":"student.2"}]`
	if rr.Body.String() != expected {
		t.Errorf("Expected: %v == %v", rr.Body.String(), expected)
	}
}

func TestHandleGetExams(t *testing.T) {
	store := buildAStore()

	r, _ := http.NewRequest("GET", "/exams", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(store.HandleGetExams())
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: status %d == 200", status)
	}

	expected := `[{"ExamId":1},{"ExamId":2}]`
	if rr.Body.String() != expected {
		t.Errorf("Expected: %v == %v", rr.Body.String(), expected)
	}
}

func TestHandleGetExam(t *testing.T) {
	store := buildAStore()

	r, _ := http.NewRequest("GET", "/exam/1", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(store.HandleGetExam())
	handler.ServeHTTP(rr, r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected: status %d == 200", status)
	}

	expected := `{"ExamId":1,"Average":0.625,"Results":[{"Score":1,"StudentId":"student.1"},{"Score":0.25,"StudentId":"student.2"}]}`
	if rr.Body.String() != expected {
		t.Errorf("Expected: %v == %v", rr.Body.String(), expected)
	}
}
