package am

import "time"

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

func IsAM(clock Clock) bool {
	return clock.Now().Hour() < 12
}
