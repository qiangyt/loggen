package common

import (
	"fmt"
	"os"
)

const (
	LogType_bunyan = "bunyan"
)

type SubOptions interface {
}

// ----------------------------------
type OptionsT struct {
	debug   bool
	logType string
	subArgs []string
	version string
}

type Options = *OptionsT

func (i Options) Debug() bool {
	return i.debug
}

func (i Options) LogType() string {
	return i.logType
}

func (i Options) SubArgs() []string {
	return i.subArgs
}

func (i Options) Version() string {
	return i.version
}

func OptionsWithCommandLine(version string) (bool, Options) {

	r := &OptionsT{
		debug:   false,
		subArgs: []string{},
		version: version,
	}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		var isOption bool
		if arg[0:1] == "-" {
			isOption = true
		} else {
			isOption = false
		}

		if isOption {
			if arg == "-h" || arg == "--help" {
				r.PrintHelp()
				return false, nil
			} else if arg == "-V" || arg == "--version" {
				r.PrintVersion()
				return false, nil
			} else if arg == "-d" || arg == "--debug" {
				r.debug = true
			} else if i == 1 {
				fmt.Println("please input log type argument")
				r.PrintVersion()
				return false, nil
			} else {
				r.subArgs = append(r.subArgs, arg)
			}
		} else {
			switch arg {
			case LogType_bunyan:
				r.logType = LogType_bunyan
			default:
				fmt.Printf("log type '%s' is not supported", arg)
				r.PrintVersion()
				return false, nil
			}
		}
	}

	return true, r
}

// PrintVersion ...
func (i Options) PrintVersion() {
	fmt.Println(i.Version())
}

// PrintHelp ...
func (i Options) PrintHelp() {
	fmt.Println("\nloggen: Log generator for jog(https://github.com/qiangyt/jog) developement only.")
	i.PrintVersion()
	fmt.Println()

	fmt.Println("Global options:")
	fmt.Printf("  -d,  --debug                    Print more error detail\n")
	fmt.Printf("  -h,  --help                     Display this information\n")
	fmt.Printf("  <log type> -h,  server --help   Display log type specific help information\n")
	fmt.Printf("  -V,  --version                  Display app version information\n")
	fmt.Println()

	fmt.Println("Supported log types:")
	fmt.Println("  " + LogType_bunyan)
}
