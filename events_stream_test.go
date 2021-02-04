package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEventStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("event: score\n"))
		rw.Write([]byte(`data: {"exam": 3, "studentId": "foo", "score": 0.991}`))
		rw.Write([]byte("\n\n"))
	}))

	defer server.Close()

	events := make(chan Event)
	go NewEventStream(server.URL, events, server.Client())

	ev := <-events

	if ev.name != "score" {
		t.Error("score != ", ev.name)
	}

	if ev.data != (ScoreEvent{
		StudentId: "foo",
		Exam:      3,
		Score:     0.991,
	}) {
		t.Error(ev.data)
	}
}
