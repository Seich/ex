package main

func buildAStore() *ScoreStore {
	store := NewScoreStore()

	ev1 := Event{
		name: "score",
		data: ScoreEvent{
			StudentId: "student.1",
			Exam:      1,
			Score:     1.0,
		},
	}

	ev2 := Event{
		name: "score",
		data: ScoreEvent{
			StudentId: "student.2",
			Exam:      1,
			Score:     0.25,
		},
	}

	ev3 := Event{
		name: "score",
		data: ScoreEvent{
			StudentId: "student.2",
			Exam:      2,
			Score:     0.25,
		},
	}

	store.addScore(ev1)
	store.addScore(ev2)
	store.addScore(ev3)

	return store
}
