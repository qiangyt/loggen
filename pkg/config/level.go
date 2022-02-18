package config

import (
	wr "github.com/mroth/weightedrand"
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
	WeightTrace uint32 `yaml:"weightTrace"`
	WeightDebug uint32 `yaml:"weightDebug"`
	WeightInfo  uint32 `yaml:"weightInfo"`
	WeightWarn  uint32 `yaml:"weightWarn"`
	WeightError uint32 `yaml:"weightError"`
	WeightFatal uint32 `yaml:"weightFatal"`

	levelChooser *wr.Chooser `yaml:"-"`
}

const (
	LevelTrace uint32 = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Level = *LevelT

func NewLevel() Level {
	return &LevelT{}
}

func (i Level) Normalize() {
	if (i.WeightTrace + i.WeightDebug + i.WeightInfo + i.WeightWarn + i.WeightError + i.WeightFatal) == 0 {
		i.WeightTrace = DefaultLevelWeightTrace
		i.WeightDebug = DefaultLevelWeightDebug
		i.WeightInfo = DefaultLevelWeightInfo
		i.WeightWarn = DefaultLevelWeightWarn
		i.WeightError = DefaultLevelWeightError
		i.WeightFatal = DefaultLevelWeightFatal
	}
}

func (i Level) Initialize() {
	i.levelChooser, _ = wr.NewChooser(
		wr.Choice{Item: LevelTrace, Weight: uint(i.WeightTrace)},
		wr.Choice{Item: LevelDebug, Weight: uint(i.WeightDebug)},
		wr.Choice{Item: LevelInfo, Weight: uint(i.WeightInfo)},
		wr.Choice{Item: LevelWarn, Weight: uint(i.WeightWarn)},
		wr.Choice{Item: LevelError, Weight: uint(i.WeightError)},
		wr.Choice{Item: LevelFatal, Weight: uint(i.WeightFatal)},
	)
}

func (i Level) Next() uint32 {
	return i.levelChooser.Pick().(uint32)
}
