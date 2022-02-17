package config

import (
	"fmt"
)

const (
	DefaultLoggerWeight uint32 = 1
)

type LoggerT struct {
	Name   string
	Weight uint32
}

type Logger = *LoggerT

func NewLogger() Logger {
	return &LoggerT{}
}

func (i Logger) Normalize(hint string) {
	i.NormalizeName(hint)
	i.NormalizeWeight()

}

func (i Logger) NormalizeName(hint string) {
	if len(i.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (i Logger) NormalizeWeight() {
	if i.Weight == 0 {
		i.Weight = DefaultLoggerWeight
	}
}
