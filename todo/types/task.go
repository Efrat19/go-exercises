package types

import (
	"time"
)

type Task struct {
	Id        int
	Name      string
	TimeStamp time.Time
}
