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

func (me Timestamp) Normalize(hint string) {
	me.NormalizeBegin()
	me.NormalizeInterval(hint)
}

func (me Timestamp) NormalizeBegin() {
	if me.Begin.IsZero() {
		me.Begin = time.Now()
	}
}

func (me Timestamp) NormalizeInterval(hint string) {
	if me.IntervalMin == 0 {
		me.IntervalMin = DefaultTimestampIntervalMin
		if me.IntervalMax == 0 {
			me.IntervalMax = DefaultTimestampIntervalMax
		}
	} else {
		if me.IntervalMax == 0 {
			me.IntervalMax = me.IntervalMin + (DefaultTimestampIntervalMax - DefaultTimestampIntervalMin)
		}
	}

	if me.IntervalMax == 0 {
		me.IntervalMax = DefaultTimestampIntervalMax
	}

	if me.IntervalMin > me.IntervalMax {
		panic(fmt.Errorf("%s.intervalMin must be <= %s.intervalMax", hint, hint))
	}
}
