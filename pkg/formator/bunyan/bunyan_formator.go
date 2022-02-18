package bunyan

import (
	"encoding/json"
	"fmt"

	"github.com/qiangyt/loggen/pkg/config"
	"github.com/qiangyt/loggen/pkg/formator"
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
	formator.RegisterFormator("bunyan", &FormatorT{})
}

func (me Formator) Format(state config.State) string {
	appState := state.App
	loggerStage := appState.Logger

	obj := map[string]interface{}{
		"time":     state.Timestamp.Format(TIMESTAMP_LAYOUT),
		"level":    FormatLevel(appState.Level),
		"pid":      appState.Pid,
		"v":        0,
		"id":       loggerStage.Config.Name,
		"name":     appState.Config.Name,
		"hostname": appState.Host.Config.Name,
		"path":     loggerStage.Config.Path,
		"msg":      loggerStage.Message,
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
