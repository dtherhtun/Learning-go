package pomodoro

import (
	"errors"
	"time"
)

// Category constants
const (
	CategoryPomodoro   = "Pomodoro"
	CategoryShortBreak = "ShortBreak"
	CategoryLongBreak  = "LongBreak"
)

// State constants
const (
	StateNotStarted = iota
	StateRunning
	StatePaused
	StateDone
	StateCancelled
)

type Interval struct {
	ID              int64
	StartTime       time.Time
	PlannedDuration time.Duration
	ActualDuration  time.Duration
	Category        string
	State           int
}

type Repository interface {
	Create(i Interval) (int64, error)
	Update(i Interval) error
	ByID(id int64) (Interval, error)
	Last() (Interval, error)
	Breaks(n int) ([]Interval, error)
}

var (
	ErrNoIntervals        = errors.New("no intervals")
	ErrIntervalNotRunning = errors.New("interval not running")
	ErrIntervalCompleted  = errors.New("interval is completed or cancelled")
	ErrInvalidState       = errors.New("invalid State")
	ErrInvalidID          = errors.New("invalid id")
)

type IntervalConfig struct {
	repo               Repository
	PomodoroDuration   time.Time
	ShortBreakDuration time.Time
	LongBreakDuration  time.Duration
}