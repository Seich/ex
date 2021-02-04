package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
)

type ScoreEvent struct {
	StudentId string  `json:"studentId"`
	Exam      int     `json:"exam"`
	Score     float64 `json:"score"`
}

type Event struct {
	name string
	data ScoreEvent
}

func NewEventStream(url string, events chan Event, client *http.Client) {
	defer close(events)

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var ev = Event{}
	for scanner.Scan() {
		scanline := strings.SplitN(scanner.Text(), ":", 2)

		if len(scanline) < 2 {
			if ev != (Event{}) {
				events <- ev
			}

			continue
		}

		field := scanline[0]
		value := scanline[1]

		if field == "event" {
			ev.name = strings.TrimSpace(value)
		}

		if field == "data" {
			ev.data = ScoreEvent{}
			if err := json.Unmarshal([]byte(value), &ev.data); err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
