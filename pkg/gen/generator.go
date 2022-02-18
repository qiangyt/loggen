package gen

import (
	"fmt"

	wr "github.com/mroth/weightedrand"
	"github.com/qiangyt/loggen/pkg/config"
	"github.com/qiangyt/loggen/pkg/formator"
)

type GeneratorT struct {
	config config.Config

	timestamp TimestampGenerator
	apps      *wr.Chooser
}

type Generator = *GeneratorT

func NewGenerator(cfg config.Config) Generator {
	return &GeneratorT{
		config:    cfg,
		timestamp: NewTimestampGenerator(cfg.Timestamp),
		apps:      BuildAppChooser(cfg.Apps),
	}
}

func (me Generator) Next(prev config.State) config.State {
	return &config.StateT{
		Config:    me.config,
		Timestamp: me.timestamp.Next(prev.Timestamp),
		App:       me.apps.Pick().(AppGenerator).Next(),
	}
}

func (me Generator) Generate() {
	cfg := me.config
	state := config.NewState(cfg)

	var n uint32
	for n = 0; n < cfg.Number; n++ {
		state = me.Next(state)

		app := state.App.Config
		fmtor := formator.GetFormator(app.Name, app.Format)

		line := fmtor.Format(state)
		fmt.Println(string(line))
	}
}

func BuildAppChooser(apps []config.App) *wr.Chooser {
	choices := []wr.Choice{}
	for _, app := range apps {
		choices = append(choices, wr.Choice{
			Item:   NewAppGenerator(app),
			Weight: uint(app.Weight),
		})
	}

	r, _ := wr.NewChooser(choices...)
	return r
}
