package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/qiangyt/loggen/pkg/config"
	"github.com/qiangyt/loggen/pkg/gen"
	"github.com/qiangyt/loggen/pkg/options"

	_ "github.com/qiangyt/loggen/pkg/gen/bunyan"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string
)

func main() {
	rand.Seed(time.Now().Unix())

	ok, options := options.WithCommandLine(Version)
	if !ok || options == nil {
		return
	}

	cfg := config.NewConfigWithOptions(options)

	appChooser := gen.CreateAppChooser(cfg)

	timestamp := time.Time{}
	var n uint32

	for n = 0; n < cfg.Number; n++ {
		g := appChooser.Pick().(gen.Generator)
		timestampText := g.NextTimestamp(&timestamp)

		lineObj := map[string]interface{}{
			"time":     timestampText,
			"level":    g.NextLevel(),
			"pid":      g.NextPid(),
			"v":        0,
			"id":       g.NextLogger(),
			"name":     g.App().Name,
			"hostname": "db9c2f8e0b7c",
			"path":     "/usr/src/app/config/config.json",
			"msg":      "no json configuration file",
		}
		lineTxt, _ := json.Marshal(lineObj)

		fmt.Println(string(lineTxt))
	}

}
