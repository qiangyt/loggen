package gen

import (
	"fmt"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/qiangyt/loggen/pkg/config"
)

type Generator interface {
	NextPid() uint32
	NextTimestamp(timestamp *time.Time) string
	NextLevel() uint32
	App() config.App
	NextLogger() string
}

type GeneratorBuilder func(config config.Config, app config.App) Generator

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

func BuildGenerator(config config.Config, app config.App) Generator {
	builder, found := generatorBuilders[app.Type]
	if !found {
		panic(fmt.Errorf("invalid app(name=%s) type: %s", app.Name, app.Type))
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

func CreateAppChooser(cfg config.Config) *wr.Chooser {
	choices := []wr.Choice{}
	for _, app := range cfg.Apps {
		choices = append(choices, wr.Choice{
			Item:   BuildGenerator(cfg, app),
			Weight: uint(app.Weight),
		})
	}

	r, _ := wr.NewChooser(choices...)
	return r
}
