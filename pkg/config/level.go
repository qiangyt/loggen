package config

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
}

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
