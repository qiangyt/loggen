package config

import (
	"fmt"
)

const (
	DefaultLoggerWeight   uint32 = 1
	DefaultMessagFilePath        = "res:/message.default.txt"
)

type LoggerT struct {
	Name    string
	Weight  uint32
	Message string
}

type Logger = *LoggerT

func NewLogger() Logger {
	return &LoggerT{}
}

func (me Logger) Normalize(hint string) {
	me.NormalizeName(hint)
	me.NormalizeWeight()
	me.NormalizeMessage()
}

func (me Logger) NormalizeName(hint string) {
	if len(me.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (me Logger) NormalizeWeight() {
	if me.Weight == 0 {
		me.Weight = DefaultLoggerWeight
	}
}

func (me Logger) NormalizeMessage() {
	if len(me.Message) == 0 {
		me.Message = DefaultMessagFilePath
	}
}
