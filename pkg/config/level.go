package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

const (
	DefaultLevelWeightTrace uint32 = 5
	DefaultLevelWeightDebug uint32 = 5
	DefaultLevelWeightInfo  uint32 = 70
	DefaultLevelWeightWarn  uint32 = 10
	DefaultLevelWeightError uint32 = 5
	DefaultLevelWeightFatal uint32 = 5
)

type LevelT struct {
	Chooser     string
	WeightTrace uint32 `mapstructure:"weight-TRACE"`
	WeightDebug uint32 `mapstructure:"weight-DEBUG"`
	WeightInfo  uint32 `mapstructure:"weight-INFO"`
	WeightWarn  uint32 `mapstructure:"weight-WARN"`
	WeightError uint32 `mapstructure:"weight-ERROR"`
	WeightFatal uint32 `mapstructure:"weight-FATAL"`
}

type Level = *LevelT

func NewLevel(hint string, input map[string]interface{}) Level {
	r := &LevelT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "%s: failed decode level: %v", hint, input))
	}
	return r
}

func (me Level) Normalize(hint string) {
	if (me.WeightTrace + me.WeightDebug + me.WeightInfo + me.WeightWarn + me.WeightError + me.WeightFatal) == 0 {
		me.WeightTrace = DefaultLevelWeightTrace
		me.WeightDebug = DefaultLevelWeightDebug
		me.WeightInfo = DefaultLevelWeightInfo
		me.WeightWarn = DefaultLevelWeightWarn
		me.WeightError = DefaultLevelWeightError
		me.WeightFatal = DefaultLevelWeightFatal
	}
}
