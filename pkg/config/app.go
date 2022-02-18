package config

import (
	"fmt"

	wr "github.com/mroth/weightedrand"
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

	Formator      Formator    `yaml:"-"`
	loggerChooser *wr.Chooser `yaml:"-"`
}

type App = *AppT

func NewApp() App {
	return &AppT{}
}

func (i App) Initialize(cfg Config) {
	i.Level.Initialize()
	i.Formator = GetFormator(i.Name, i.Format)
	i.Pid.Initialize()
	i.loggerChooser = BuildLoggerChooser(i.Loggers)
}

func (i App) NextLevel() uint32 {
	return i.Level.Next()
}

func (i App) NextPid() uint32 {
	return i.Pid.Next()
}

func (i App) NextLogger() Logger {
	logger := i.loggerChooser.Pick().(Logger)
	return logger
}

func (i App) Normalize(cfg Config, hint string) {
	i.NormalizeName(hint)
	hint = fmt.Sprintf("%s(name=%s)", hint, i.Name)

	i.NormalizeFormat(hint)
	i.NormalzieLevel()
	i.NormalizePid(hint)
	i.NormalizeWeight()
}

func (i App) NormalizeName(hint string) {
	if len(i.Name) == 0 {
		panic(fmt.Errorf("missing %s.name", hint))
	}
}

func (i App) NormalizeFormat(hint string) {
	format := i.Format

	if len(format) == 0 {
		panic(fmt.Errorf("missing %s.generator", hint))
	}

	if !IsValidFormatorName(format) {
		panic(fmt.Errorf("%s.format: %s is not supported; availables: [%v]",
			hint, format, EnumerateFormatorNames()))
	}
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

func BuildAppChooser(apps []App) *wr.Chooser {
	choices := []wr.Choice{}
	for _, app := range apps {
		choices = append(choices, app.BuildChoice())
	}

	r, _ := wr.NewChooser(choices...)
	return r
}
