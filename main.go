package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/qiangyt/loggen/pkg/config"
	_ "github.com/qiangyt/loggen/pkg/formator/bunyan"
	"github.com/qiangyt/loggen/pkg/options"
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
	cfg.Normalize()
	cfg.Initialize()

	timestamp := time.Time{}
	var n uint32

	for n = 0; n < cfg.Number; n++ {
		app := cfg.ChooseApp()
		g := app.Formator

		cfg.Timestamp.Next(&timestamp)
		level := app.NextLevel()

		line := g.Format(cfg, timestamp, level, app)
		fmt.Println(string(line))
	}

}
