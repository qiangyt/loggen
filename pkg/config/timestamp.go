package config

import (
	"fmt"
	"time"
)

const (
	DefaultTimestampIntervalMin uint32 = 10
	DefaultTimestampIntervalMax uint32 = 10000
)

type TimestampT struct {
	Begin       time.Time
	IntervalMin uint32 `yaml:"intervalMin"`
	IntervalMax uint32 `yaml:"intervalMax"`
}

type Timestamp = *TimestampT

func NewTimestamp() Timestamp {
	return &TimestampT{}
}

func (i Timestamp) Normalize(hint string) {
	i.NormalizeBegin()
	i.NormalizeInterval(hint)
}

func (i Timestamp) NormalizeBegin() {
	if i.Begin.IsZero() {
		i.Begin = time.Now()
	}
}

func (i Timestamp) NormalizeInterval(hint string) {
	if i.IntervalMin == 0 {
		i.IntervalMin = DefaultTimestampIntervalMin
		if i.IntervalMax == 0 {
			i.IntervalMax = DefaultTimestampIntervalMax
		}
	} else {
		if i.IntervalMax == 0 {
			i.IntervalMax = i.IntervalMin + (DefaultTimestampIntervalMax - DefaultTimestampIntervalMin)
		}
	}

	if i.IntervalMax == 0 {
		i.IntervalMax = DefaultTimestampIntervalMax
	}

	if i.IntervalMin > i.IntervalMax {
		panic(fmt.Errorf("%s.intervalMin must be <= %s.intervalMax", hint, hint))
	}
}