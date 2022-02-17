package bunyan

import (
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/qiangyt/loggen/pkg/config"
	"github.com/qiangyt/loggen/pkg/gen"
)

const (
	LogLevel_TRACE uint32 = 10
	LogLevel_DEBUG uint32 = 20
	LogLevel_INFO  uint32 = 30
	LogLevel_WARN  uint32 = 40
	LogLevel_ERROR uint32 = 50
	LogLevel_FATAL uint32 = 60
)

const (
	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05.000Z"
)

type GeneratorT struct {
	config config.Config
	app    config.App

	levelChooser  *wr.Chooser
	pIdArray      []uint32
	loggerChooser *wr.Chooser
}

type Generator = *GeneratorT

func init() {
	gen.RegisterGenerator("bunyan", NewGenerator)
}

func BuildLevelChooser(level config.Level) *wr.Chooser {
	r, _ := wr.NewChooser(
		wr.Choice{Item: LogLevel_TRACE, Weight: uint(level.WeightTrace)},
		wr.Choice{Item: LogLevel_DEBUG, Weight: uint(level.WeightDebug)},
		wr.Choice{Item: LogLevel_INFO, Weight: uint(level.WeightInfo)},
		wr.Choice{Item: LogLevel_WARN, Weight: uint(level.WeightWarn)},
		wr.Choice{Item: LogLevel_ERROR, Weight: uint(level.WeightError)},
		wr.Choice{Item: LogLevel_FATAL, Weight: uint(level.WeightFatal)},
	)
	return r
}

func BuildPidArray(pid config.Pid) []uint32 {
	r := []uint32{}
	for i := 0; i < int(pid.Amount); i++ {
		pIdArange := int32(pid.End - pid.Begin)
		pId := pid.Begin + uint32(rand.Int31n(pIdArange))
		r = append(r, pId)
	}
	return r
}

func BuilderLoggerChooser(loggers []config.Logger) *wr.Chooser {
	loggerChoices := []wr.Choice{}
	for _, logger := range loggers {
		loggerChoices = append(loggerChoices, wr.Choice{
			Item:   logger,
			Weight: uint(logger.Weight),
		})
	}
	r, _ := wr.NewChooser(loggerChoices...)
	return r
}

func NewGenerator(config config.Config, app config.App) gen.Generator {
	return &GeneratorT{
		config:        config,
		app:           app,
		levelChooser:  BuildLevelChooser(app.Level),
		pIdArray:      BuildPidArray(app.Pid),
		loggerChooser: BuilderLoggerChooser(app.Loggers),
	}
}

func (i Generator) App() config.App {
	return i.app
}

func (i Generator) NextPid() uint32 {
	index := rand.Intn(len(i.pIdArray))
	return i.pIdArray[index]
}

func (i Generator) NextTimestamp(timestamp *time.Time) string {
	cfg := i.config.Timestamp

	if timestamp.IsZero() {
		*timestamp = cfg.Begin
	} else {
		intervalDeta := cfg.IntervalMax - cfg.IntervalMin
		interval := cfg.IntervalMin + uint32(rand.Int31n(int32(intervalDeta)))
		dura := time.Duration(interval * 1000 * 1000)

		*timestamp = timestamp.Add(dura)
	}

	return timestamp.Format(TIMESTAMP_LAYOUT)
}

func (i Generator) NextLevel() uint32 {
	return i.levelChooser.Pick().(uint32)
}

func (i Generator) NextLogger() string {
	logger := i.loggerChooser.Pick().(config.Logger)
	return logger.Name
}
