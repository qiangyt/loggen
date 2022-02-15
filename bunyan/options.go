package bunyan

import (
	"fmt"

	"github.com/qiangyt/loggen/common"
)

const (
	LogLevel_TRACE = 10
	LogLevel_DEBUG = 20
	LogLevel_INFO  = 30
	LogLevel_WARN  = 40
	LogLevel_ERROR = 50
	LogLevel_FATAL = 60
)

type OptionsT struct {
	parent common.Options
}

type Options = *OptionsT

func OptionsWithCommandLine(parent common.Options) (bool, Options) {

	r := &OptionsT{
		parent: parent,
	}

	args := parent.SubArgs()
	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "-h" || arg == "--help" {
			r.PrintHelp()
			return false, nil
		}
	}

	return true, r
}

func (i Options) PrintHelp() {
	fmt.Println("bunyan help: TODO")
	fmt.Println()

	fmt.Println("bunyan usage:")
	fmt.Println("  loggen bunyan [option...]")
	fmt.Println()

	fmt.Println("bunyan options:")
	fmt.Println()
}
