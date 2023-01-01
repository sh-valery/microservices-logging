package util

import "time"

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

func NewRealClock() realClock { return realClock{} }
