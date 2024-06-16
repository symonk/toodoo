package model

import "time"

type Task struct {
	name        string
	description string
	recurring   bool
	schedule    time.Time
}

func NewTask() *Task {
	return &Task{
		name:        "foo",
		description: "bar",
		recurring:   true,
		schedule:    time.Now().Add(time.Minute),
	}
}
