package gen

import (
	wr "github.com/mroth/weightedrand"
	"github.com/qiangyt/loggen/pkg/config"
)

type LevelGeneratorT struct {
	chooser *wr.Chooser
}

type LevelGenerator = *LevelGeneratorT

func NewLevelGenerator(cfg config.Level) LevelGenerator {
	chooser, _ := wr.NewChooser(
		wr.Choice{Item: config.LevelTrace, Weight: uint(cfg.WeightTrace)},
		wr.Choice{Item: config.LevelDebug, Weight: uint(cfg.WeightDebug)},
		wr.Choice{Item: config.LevelInfo, Weight: uint(cfg.WeightInfo)},
		wr.Choice{Item: config.LevelWarn, Weight: uint(cfg.WeightWarn)},
		wr.Choice{Item: config.LevelError, Weight: uint(cfg.WeightError)},
		wr.Choice{Item: config.LevelFatal, Weight: uint(cfg.WeightFatal)},
	)
	return &LevelGeneratorT{chooser}
}

func (me LevelGenerator) Next() uint32 {
	return me.chooser.Pick().(uint32)
}
