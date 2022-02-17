package config

import "fmt"

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

func NewPid() Pid {
	return &PidT{}
}

func (i Pid) Normalize(hint string) {
	i.NormalizeBeginEnd(hint)
	i.NormalizeAmount()
}

func (i Pid) NormalizeBeginEnd(hint string) {
	if i.Begin == 0 {
		i.Begin = DefaultPidBegin
		if i.End == 0 {
			i.End = DefaultPidEnd
		}
	} else {
		if i.End == 0 {
			i.End = i.Begin + (DefaultPidEnd - DefaultPidBegin)
		}
	}

	if i.End == 0 {
		i.End = DefaultPidEnd
	}

	if i.Begin > i.End {
		panic(fmt.Errorf("%s.begin must be <= %s.end", hint, hint))
	}
}

func (i Pid) NormalizeAmount() {
	if i.Amount == 0 {
		i.Amount = DefaultPidAmount
	}
}
