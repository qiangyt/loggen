package config

type LevelT struct {
	WeightTrace uint32
	WeightDebug uint32
	WeightInfo  uint32
	WeightWarn  uint32
	WeightError uint32
	WeightFatal uint32
}

type Level = *LevelT

func NewLevel() Level {
	return &LevelT{}
}
