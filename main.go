package main

import (
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

}
