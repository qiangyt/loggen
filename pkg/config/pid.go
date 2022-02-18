package config

import (
	"fmt"
	"math/rand"
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

	pIdArray []uint32 `yaml:"-"`
}

type Pid = *PidT

func NewPid() Pid {
	return &PidT{}
}

func (i Pid) Next() uint32 {
	index := rand.Intn(len(i.pIdArray))
	return i.pIdArray[index]
}

func (i Pid) Initialize() {
	r := []uint32{}
	for idx := 0; idx < int(i.Amount); idx++ {
		pIdArange := int32(i.End - i.Begin)
		pId := i.Begin + uint32(rand.Int31n(pIdArange))
		r = append(r, pId)
	}

	i.pIdArray = r
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
