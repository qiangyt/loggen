package config

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

const (
	DefaultTimestampIntervalMin uint32 = 10
	DefaultTimestampIntervalMax uint32 = 10000
)

type TimestampT struct {
	Begin       time.Time
	IntervalMin uint32 `mapstructure:"interval-min"`
	IntervalMax uint32 `mapstructure:"interval-max"`
}

type Timestamp = *TimestampT

func NewTimestamp(hint string, input map[string]interface{}) Timestamp {
	r := &TimestampT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "%s: failed decode timestamp: %v", hint, input))
	}
	return r
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
		panic(fmt.Errorf("%s.interval-min must be <= %s.interval-max", hint, hint))
	}
}
