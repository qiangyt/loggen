package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
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

func CreateAppChooser(cfg config.Config) *wr.Chooser {
	appChoices := []wr.Choice{}
	for _, app := range cfg.Apps {
		appChoices = append(appChoices, wr.Choice{
			Item:   gen.BuildGenerator(cfg, app),
			Weight: uint(app.Weight),
		})
	}

	r, _ := wr.NewChooser(appChoices...)
	return r
}

func main() {
	rand.Seed(time.Now().Unix())

	ok, options := options.WithCommandLine(Version)
	if !ok || options == nil {
		return
	}

	cfg := config.NewConfigWithOptions(options)

	appChooser := CreateAppChooser(cfg)

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
			"id":       "Config",
			"name":     g.App().Name,
			"hostname": "db9c2f8e0b7c",
			"path":     "/usr/src/app/config/config.json",
			"msg":      "no json configuration file",
		}
		lineTxt, _ := json.Marshal(lineObj)

		fmt.Println(string(lineTxt))
	}

}
