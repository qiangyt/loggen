package bunyan

import (
	"encoding/json"
	"fmt"

	"github.com/qiangyt/loggen/pkg/config"
	"github.com/qiangyt/loggen/pkg/formator"
)

// 211.144.80.119 - - [18/Feb/2022:12:16:18 +0800] "POST /admin/agent/PullAgentAPI HTTP/1.1" 200 135 "-" "restify/1.5.2 (x64-linux; v8/6.2.414.78; OpenSSL/1.0.2s) node/8.17.0"
// 47.96.168.136 - - [18/Feb/2022:12:15:55 +0800] "POST /admin/agent/PullAgentAPI HTTP/1.1" 200 135 "-" "restify/1.5.2 (x64-linux; v8/6.2.414.78; OpenSSL/1.0.2s) node/8.17.0"
// 114.93.200.90 - - [18/Feb/2022:12:15:52 +0800] "GET /ip HTTP/1.1" 200 13 "-" "Go-http-client/1.1"
// 114.93.200.90 - - [18/Feb/2022:19:57:32 +0800] "GET /favicon.ico HTTP/1.1" 200 67646 "https://ingeniseq.xmkbio.com/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36"
// 114.93.200.90 - - [18/Feb/2022:20:05:55 +0800] "POST /admin/run/SearchRun HTTP/1.1" 200 13687 "https://ingeniseq.xmkbio.com/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36"
// 114.93.200.90 - - [18/Feb/2022:19:57:31 +0800] "GET /static/js/vendor.aa2f1db8f6fc9e3cd0aa.js HTTP/1.1" 200 756647 "https://ingeniseq.xmkbio.com/" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36"

// 2022/02/18 08:32:04 [crit] 8179#8179: *3480776 SSL_do_handshake() failed (SSL: error:1417D18C:SSL routines:tls_process_client_hello:version too low) while SSL handshaking, client: 192.241.206.75, server: 0.0.0.0:443
const (
	LevelTrace uint32 = 10
	LevelDebug uint32 = 20
	LevelInfo  uint32 = 30
	LevelWarn  uint32 = 40
	LevelError uint32 = 50
	LevelFatal uint32 = 60
)

const (
	TIMESTAMP_LAYOUT = "02/Jan/2006:15:04:05 +0800"
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
