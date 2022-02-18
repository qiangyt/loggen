package gen

import (
	wr "github.com/mroth/weightedrand"
	"github.com/qiangyt/loggen/pkg/config"
)

type AppGeneratorT struct {
	config config.App
	level  LevelGenerator
	pid    PidGenerator

	loggers *wr.Chooser
	hosts   *wr.Chooser
}

type AppGenerator = *AppGeneratorT

func NewAppGenerator(app config.App) AppGenerator {
	return &AppGeneratorT{
		config:  app,
		level:   NewLevelGenerator(app.Level),
		pid:     NewPidGenerator(app.Pid),
		loggers: BuildLoggerChooser(app.Loggers),
		hosts:   BuildHostChooser(app.Hosts),
	}
}

func (me AppGenerator) Next() config.AppState {
	return &config.AppStateT{
		Config: me.config,
		Level:  me.level.Next(),
		Pid:    me.pid.Next(),
		Logger: me.loggers.Pick().(LoggerGenerator).Next(),
		Host:   me.hosts.Pick().(config.HostState),
	}
}

func BuildLoggerChooser(loggers []config.Logger) *wr.Chooser {
	choices := []wr.Choice{}
	for _, logger := range loggers {
		choices = append(choices, wr.Choice{
			Item:   NewLoggerGenerator(logger),
			Weight: uint(logger.Weight),
		})
	}

	r, _ := wr.NewChooser(choices...)
	return r
}

func BuildHostChooser(hosts []config.Host) *wr.Chooser {
	choices := []wr.Choice{}
	for _, host := range hosts {
		choices = append(choices, wr.Choice{
			Item:   &config.HostStateT{Config: host},
			Weight: uint(host.Weight),
		})
	}

	r, _ := wr.NewChooser(choices...)
	return r
}
