package bunyan

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/qiangyt/loggen/pkg/config"
)

const (
	LevelTrace uint32 = 10
	LevelDebug uint32 = 20
	LevelInfo  uint32 = 30
	LevelWarn  uint32 = 40
	LevelError uint32 = 50
	LevelFatal uint32 = 60
)

const (
	TIMESTAMP_LAYOUT = "2006-01-02T15:04:05.000Z"
)

type FormatorT struct {
}

type Formator = *FormatorT

func init() {
	config.RegisterFormator("bunyan", &FormatorT{})
}

func (i Formator) Format(cfg config.Config, timestamp time.Time, level uint32, app config.App) string {
	logger := app.NextLogger()

	obj := map[string]interface{}{
		"time":     timestamp.Format(TIMESTAMP_LAYOUT),
		"level":    FormatLevel(level),
		"pid":      app.NextPid(),
		"v":        0,
		"id":       logger.Name,
		"name":     app.Name,
		"hostname": "db9c2f8e0b7c",
		"path":     "/usr/src/app/config/config.json",
		"msg":      "no json configuration file",
	}
	r, _ := json.Marshal(obj)
	return string(r)
}

func FormatLevel(level uint32) uint32 {
	switch level {
	case config.LevelTrace:
		return LevelTrace
	case config.LevelDebug:
		return LevelDebug
	case config.LevelInfo:
		return LevelInfo
	case config.LevelWarn:
		return LevelWarn
	case config.LevelError:
		return LevelError
	case config.LevelFatal:
		return LevelFatal
	default:
		panic(fmt.Errorf("unexpected log level value: %d", level))
	}
}
