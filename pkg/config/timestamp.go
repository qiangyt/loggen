package config

import "time"

type TimestampT struct {
	Begin       time.Time
	IntervalMin uint32
	IntervalMax uint32
}

type Timestamp = *TimestampT

func NewTimestamp() Timestamp {
	return &TimestampT{}
}
