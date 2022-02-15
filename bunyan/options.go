package bunyan

import (
	"fmt"
	"os"

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
	common.Options
}

type Options = *OptionsT

func OptionsWithCommandLine(parent common.Options) (bool, Options) {

	r := &OptionsT{
		parent,
	}

	for i := 0; i < len(parent.SubArgs()); i++ {
		arg := os.Args[i]

		if arg == "-h" || arg == "--help" {
			r.PrintHelp()
			return false, nil
		} else {

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
	fmt.Printf("  -c,  --config <server config file path>                     Specify server config YAML file path. The default is ./jog.server.yaml or $HOME/.jog/jog.server.yaml \n")
	fmt.Printf("  -t,  --template                                             Print a server config YAML file template\n")
	fmt.Println()
}
