package config

import (
	"fmt"

	wr "github.com/mroth/weightedrand"
)

const (
	DefaultAppWeight uint32 = 1
)

type AppT struct {
	Name      string
	Generator string
	Level     Level
	Pid       Pid
	Weight    uint32

	Loggers      []Logger
	GeneratorObj Generator `yaml:"-"`
}

type App = *AppT

func NewApp() App {
	return &AppT{}
}

func (i App) Normalize(cfg Config, hint string) {
	i.NormalizeName(hint)
	hint = fmt.Sprintf("%s(name=%s)", hint, i.Name)

	i.NormalizeGenerator(cfg, hint)
	i.NormalzieLevel()
	i.NormalizePid(hint)
	i.NormalizeWeight()
	i.NormalizeLoggers(hint)
}

func (i App) NormalizeName(hint string) {
	if len(i.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (i App) NormalizeGenerator(cfg Config, hint string) {
	name := i.Generator

	if len(name) == 0 {
		panic(fmt.Errorf("missing %s.generator", hint))
	}

	if !IsValidGeneratorName(name) {
		panic(fmt.Errorf("%s.generator: %s is not supported; availables: [%v]",
			hint, name, EnumerateGeneratorNames()))
	}

	i.GeneratorObj = BuildGenerator(cfg, i)
}

func (i App) BuildChoice() wr.Choice {
	return wr.Choice{
		Item:   i,
		Weight: uint(i.Weight),
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

func (i App) NormalizeLoggers(hint string) {
	if len(i.Loggers) == 0 {
		panic(fmt.Errorf("%s: at least 1 logger is required", hint))
	}

	byNames := map[string]Logger{}
	for idx, logger := range i.Loggers {
		loggerHint := fmt.Sprintf("%s.logger[%d]", hint, idx)

		name := logger.Name
		if _, found := byNames[name]; found {
			panic(fmt.Errorf("%s.name(%s) is duplicated", loggerHint, name))
		}
		byNames[name] = logger

		logger.Normalize(loggerHint)
	}
}
