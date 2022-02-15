package bunyan

import (
	"encoding/json"
	"fmt"
)

type PatternTimeFunc func() string

type PatternT struct {
}

type Pattern = *PatternT

func Generate(options Options) {
	lineObj := map[string]interface{}{
		"time":     "2020-06-16T18:54:52.242Z",
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
