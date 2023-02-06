package clock

import "time"

// Clocker
type Clocker interface {
	Now() time.Time
}

// RealClocker
type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

// FixedClocker
type FixedClocker struct{}

func (fc FixedClocker) Now() time.Time {
	return time.Date(2022, 5, 10, 12, 34, 56, 0, time.UTC)
}
