package config

import (
	"fmt"
)

const (
	DefaultHostWeight uint32 = 1
)

type HostT struct {
	Name   string
	Weight uint32
}

type Host = *HostT

func NewHost() Host {
	return &HostT{}
}

func (me Host) Normalize(hint string) {
	me.NormalizeName(hint)
	me.NormalizeWeight()
}

func (me Host) NormalizeName(hint string) {
	if len(me.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (me Host) NormalizeWeight() {
	if me.Weight == 0 {
		me.Weight = DefaultHostWeight
	}
}
