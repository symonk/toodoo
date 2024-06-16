package model

import "time"

type Task struct {
	name        string
	description string
	recurring   bool
	schedule    time.Time
}
