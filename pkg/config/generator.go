package config

import (
	"fmt"
	"time"
)

type Generator interface {
	NextPid() uint32
	NextTimestamp(timestamp *time.Time) string
	NextLevel() uint32
	App() App
	NextLogger() string
}

type GeneratorBuilder func(config Config, app App) Generator

var (
	generatorBuilders map[string]GeneratorBuilder
)

func init() {
	generatorBuilders = make(map[string]GeneratorBuilder)
}

func RegisterGenerator(name string, builder GeneratorBuilder) {
	if HasGenerator(name) {
		panic(fmt.Errorf("duplicated generator: %s", name))
	}
	generatorBuilders[name] = builder
}

func HasGenerator(name string) bool {
	_, found := generatorBuilders[name]
	return found
}

func BuildGenerator(config Config, app App) Generator {
	builder, found := generatorBuilders[app.Generator]
	if !found {
		panic(fmt.Errorf("invalid app(name=%s) generator: %s", app.Name, app.Generator))
	}
	return builder(config, app)
}

func EnumerateGeneratorNames() []string {
	r := []string{}
	for name, _ := range generatorBuilders {
		r = append(r, name)
	}
	return r
}

func IsValidGeneratorName(name string) bool {
	names := EnumerateGeneratorNames()
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}
