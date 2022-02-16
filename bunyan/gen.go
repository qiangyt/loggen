package bunyan

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	wr "github.com/qiangyt/loggen/weightedrand"
)

const (
	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05.000Z"
)

type GeneratorT struct {
	options   Options
	timestamp time.Time
	chooser   *wr.Chooser
}

type Generator = *GeneratorT

func NewGenerator(options Options) Generator {
	parentOptions := options.parent

	chooser, _ := wr.NewChooser(
		wr.Choice{Item: LogLevel_TRACE, Weight: uint(parentOptions.LevelWeightTrace())},
		wr.Choice{Item: LogLevel_DEBUG, Weight: uint(parentOptions.LevelWeightDebug())},
		wr.Choice{Item: LogLevel_INFO, Weight: uint(parentOptions.LevelWeightInfo())},
		wr.Choice{Item: LogLevel_WARN, Weight: uint(parentOptions.LevelWeightWarn())},
		wr.Choice{Item: LogLevel_ERROR, Weight: uint(parentOptions.LevelWeightError())},
		wr.Choice{Item: LogLevel_FATAL, Weight: uint(parentOptions.LevelWeightFatal())},
	)

	return &GeneratorT{
		options: options,
		chooser: chooser,
	}
}

func (i Generator) NextTimestamp() string {
	parentOptions := i.options.parent

	if i.timestamp.IsZero() {
		i.timestamp = parentOptions.TimeBegin()
	} else {
		intervalDeta := parentOptions.TimeIntervalMax() - parentOptions.TimeIntervalMin()
		interval := parentOptions.TimeIntervalMin() + uint32(rand.Int31n(int32(intervalDeta)))
		dura := time.Duration(interval * 1000 * 1000)

		i.timestamp = i.timestamp.Add(dura)
	}

	return i.timestamp.Format(TIMESTAMP_LAYOUT)
}

func (i Generator) NextLevel() uint32 {
	return i.chooser.Pick().(uint32)
}

func (i Generator) Generate() {
	var n uint32
	for n = 0; n < i.options.parent.Number(); n++ {
		lineObj := map[string]interface{}{
			"time":     i.NextTimestamp(),
			"level":    i.NextLevel(),
			"pid":      16,
			"v":        0,
			"id":       "Config",
			"name":     "tca_amplicon_admin",
			"hostname": "db9c2f8e0b7c",
			"path":     "/usr/src/app/config/config.json",
			"msg":      "no json configuration file",
		}
		lineTxt, _ := json.Marshal(lineObj)

		fmt.Println(string(lineTxt))
	}
}
