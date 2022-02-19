package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

const (
	DefaultPidBegin  = 1000
	DefaultPidEnd    = 2000
	DefaultPidAmount = 1
)

type PidT struct {
	Begin  uint32
	End    uint32
	Amount uint32
}

type Pid = *PidT

func NewPid(hint string, input map[string]interface{}) Pid {
	r := &PidT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "%s: failed decode pid: %v", hint, input))
	}
	return r
}

func (me Pid) Normalize(hint string) {
	me.NormalizeBeginEnd(hint)
	me.NormalizeAmount()
}

func (me Pid) NormalizeBeginEnd(hint string) {
	if me.Begin == 0 {
		me.Begin = DefaultPidBegin
		if me.End == 0 {
			me.End = DefaultPidEnd
		}
	} else {
		if me.End == 0 {
			me.End = me.Begin + (DefaultPidEnd - DefaultPidBegin)
		}
	}

	if me.End == 0 {
		me.End = DefaultPidEnd
	}

	if me.Begin > me.End {
		panic(fmt.Errorf("%s.begin must be <= %s.end", hint, hint))
	}
}

func (me Pid) NormalizeAmount() {
	if me.Amount == 0 {
		me.Amount = DefaultPidAmount
	}
}
