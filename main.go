package main

import (
	"math/rand"
	"time"

	"github.com/qiangyt/loggen/pkg/config"
	_ "github.com/qiangyt/loggen/pkg/formator/bunyan"
	"github.com/qiangyt/loggen/pkg/gen"
	_ "github.com/qiangyt/loggen/pkg/res"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string
)

func main() {
	rand.Seed(time.Now().Unix())

	ok, options := config.NewOptionsWithCommandLine(Version)
	if !ok || options == nil {
		return
	}

	cfg := config.NewConfigWithOptions(options)
	cfg.Normalize()

	generator := gen.NewGenerator(cfg)
	generator.Generate()
}
