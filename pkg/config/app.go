package config

import (
	"fmt"
)

const (
	DefaultAppWeight uint32 = 1
)

type AppT struct {
	Name   string
	Type   string
	Level  Level
	Pid    Pid
	Weight uint32
}

type App = *AppT

func NewApp() App {
	return &AppT{}
}

func (i App) Normalize(hint string) {
	i.NormalizeName(hint)
	hint = fmt.Sprintf("%s(name=%s)", hint, i.Name)

	i.NormalizeType(hint)
	i.NormalzieLevel()
	i.NormalizePid(hint)
	i.NormalizeWeight()

}

func (i App) NormalizeName(hint string) {
	if len(i.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (i App) NormalizeType(hint string) {
	if len(i.Type) == 0 {
		panic(fmt.Errorf("missing %s.type", hint))
	}
	if i.Type != "bunyan" {
		panic(fmt.Errorf("%s.type: %s is not supported", hint, i.Type))
	}
}

func (i App) NormalizeWeight() {
	if i.Weight == 0 {
		i.Weight = DefaultAppWeight
	}
}

func (i App) NormalizePid(hint string) {
	if i.Pid == nil {
		i.Pid = NewPid()
	}
	i.Pid.Normalize(hint)
}

func (i App) NormalzieLevel() {
	if i.Level == nil {
		i.Level = NewLevel()
	}
	i.Level.Normalize()
}
