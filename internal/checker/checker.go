package checker

import (
	"mnb/checker/internal/writer"
	"time"
)

type Checker struct {
	WebSites map[string]uint
	Writer   *writer.Writer
	Duration time.Duration
}

func New(w *writer.Writer, duration time.Duration) *Checker {
	return &Checker{
		Writer:   w,
		Duration: duration,
	}
}
