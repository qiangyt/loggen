package bunyan

import (
	"fmt"

	"github.com/qiangyt/loggen/common"
)

const (
	LogLevel_TRACE uint32 = 10
	LogLevel_DEBUG uint32 = 20
	LogLevel_INFO  uint32 = 30
	LogLevel_WARN  uint32 = 40
	LogLevel_ERROR uint32 = 50
	LogLevel_FATAL uint32 = 60
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
