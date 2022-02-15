package main

import (
	"github.com/qiangyt/loggen/bunyan"
	"github.com/qiangyt/loggen/common"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string
)

func main() {
	ok, options := common.OptionsWithCommandLine(Version)
	if !ok || options == nil {
		return
	}

	if options.LogType() == common.LogType_bunyan {
		ok, bunyanOptions := bunyan.OptionsWithCommandLine(options)
		if !ok || bunyanOptions == nil {
			return
		}
		g := bunyan.NewGenerator(bunyanOptions)
		g.Generate()
	}
}
