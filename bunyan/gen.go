package bunyan

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const (
	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05.000Z"
)

type GeneratorT struct {
	options   Options
	timestamp time.Time
}

type Generator = *GeneratorT

func NewGenerator(options Options) Generator {
	return &GeneratorT{options: options}
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

func (i Generator) Generate() {
	var n uint32
	for n = 0; n < i.options.parent.Number(); n++ {
		lineObj := map[string]interface{}{
			"time":     i.NextTimestamp(),
			"level":    LogLevel_TRACE,
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
