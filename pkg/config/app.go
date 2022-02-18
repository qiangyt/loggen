package config

import (
	"fmt"
)

const (
	DefaultAppWeight uint32 = 1
)

type AppT struct {
	Name    string
	Format  string
	Level   Level
	Pid     Pid
	Weight  uint32
	Loggers []Logger
	Hosts   []Host

	//Formator formator.Formator `yaml:"-"`
}

type App = *AppT

func NewApp() App {
	return &AppT{}
}

func (me App) Normalize(cfg Config, hint string) {
	me.NormalizeName(hint)
	hint = fmt.Sprintf("%s(name=%s)", hint, me.Name)

	me.NormalizeFormat(hint)
	me.NormalzieLevel()
	me.NormalizePid(hint)
	me.NormalizeWeight()
	me.NormalizeLoggers(hint)
	me.NormalizeHosts(hint)
}

func (me App) NormalizeName(hint string) {
	if len(me.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (me App) NormalizeFormat(hint string) {
	format := me.Format

	if len(format) == 0 {
		panic(fmt.Errorf("missing %s.generator", hint))
	}
}

func (me App) NormalizeWeight() {
	if me.Weight == 0 {
		me.Weight = DefaultAppWeight
	}
}

func (me App) NormalizePid(hint string) {
	if me.Pid == nil {
		me.Pid = NewPid()
	}
	me.Pid.Normalize(hint)
}

func (me App) NormalzieLevel() {
	if me.Level == nil {
		me.Level = NewLevel()
	}
	me.Level.Normalize()
}

func (me App) NormalizeLoggers(hint string) {
	if len(me.Loggers) == 0 {
		panic(fmt.Errorf("%s: at least 1 logger is required", hint))
	}

	byNames := map[string]Logger{}
	for idx, logger := range me.Loggers {
		loggerHint := fmt.Sprintf("%s.loggers[%d]", hint, idx)

		name := logger.Name
		if _, found := byNames[name]; found {
			panic(fmt.Errorf("%s.name(%s) is duplicated", loggerHint, name))
		}
		byNames[name] = logger

		logger.Normalize(loggerHint)
	}
}

func (me App) NormalizeHosts(hint string) {
	if len(me.Hosts) == 0 {
		panic(fmt.Errorf("%s: at least 1 host is required", hint))
	}

	byNames := map[string]Host{}
	for idx, host := range me.Hosts {
		hostHint := fmt.Sprintf("%s.hosts[%d]", hint, idx)

		name := host.Name
		if _, found := byNames[name]; found {
			panic(fmt.Errorf("%s.name(%s) is duplicated", hostHint, name))
		}
		byNames[name] = host

		host.Normalize(hostHint)
	}
}
