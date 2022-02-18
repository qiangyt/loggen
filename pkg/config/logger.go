package config

import (
	"fmt"

	wr "github.com/mroth/weightedrand"
)

const (
	DefaultLoggerWeight uint32 = 1
)

type LoggerT struct {
	Name   string
	Weight uint32

	messageChooser *wr.Chooser `yaml:"-"`
}

type Logger = *LoggerT

func NewLogger() Logger {
	return &LoggerT{}
}

func (i Logger) Initialize() {
	// TODO
}

func (i Logger) Normalize(hint string) {
	i.NormalizeName(hint)
	i.NormalizeWeight()
	i.NormalizeMessage()
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

func (i Logger) NormalizeMessage() {
	//loggerChooser: BuilderLoggerChooser(app.Loggers),
	//"message: res:/message.default.txt"
}

func (i Logger) BuildChooser() wr.Choice {
	return wr.Choice{
		Item:   i,
		Weight: uint(i.Weight),
	}
}

func BuildLoggerChooser(loggers []Logger) *wr.Chooser {
	loggerChoices := []wr.Choice{}
	for _, logger := range loggers {
		loggerChoices = append(loggerChoices, logger.BuildChooser())
	}
	r, _ := wr.NewChooser(loggerChoices...)
	return r
}
