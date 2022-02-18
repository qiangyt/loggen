package gen

import (
	wr "github.com/mroth/weightedrand"
	"github.com/qiangyt/loggen/pkg/config"
)

type LoggerGeneratorT struct {
	config         config.Logger
	messageChooser *wr.Chooser
}

type LoggerGenerator = *LoggerGeneratorT

func NewLoggerGenerator(cfg config.Logger) LoggerGenerator {
	return &LoggerGeneratorT{
		config: cfg,
	}
}

func (me LoggerGenerator) Next() config.LoggerState {
	return &config.LoggerStateT{
		Config: me.config,
	}
}
