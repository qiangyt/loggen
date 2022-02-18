package config

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/qiangyt/loggen/pkg/res"
	_time "github.com/qiangyt/loggen/pkg/time"
)

// ----------------------------------
type OptionsT struct {
	Debug          bool
	AppName        string
	Version        string
	ConfigFilePath string
	TimeBegin      time.Time
	Number         uint32
}

type Options = *OptionsT

// PrintConfigTemplate ...
func (i Options) PrintConfigTemplate() {
	fmt.Println(res.ReadDefaultConfigYaml())
}

func NewOptionsWithCommandLine(version string) (bool, Options) {

	r := &OptionsT{
		Debug:   false,
		Version: version,
	}

	args := os.Args
	for i := 1; i < len(args); i++ {
		arg := args[i]

		argValue := ""
		if i+1 < len(args) {
			argValue = args[i+1]
		}

		if arg[0:1] != "-" {
			r.AppName = arg
		} else {
			if arg == "-c" || arg == "--config" {
				if i+1 >= len(args) {
					panic(errors.New("Missing config file path"))
				}

				r.ConfigFilePath = argValue
				i++
			} else if arg == "-h" || arg == "--help" {
				r.PrintHelp()
				return false, nil
			} else if arg == "-V" || arg == "--version" {
				r.PrintVersion()
				return false, nil
			} else if arg == "-d" || arg == "--debug" {
				r.Debug = true
			} else if arg == "--time-begin" {
				if i+1 >= len(args) {
					panic(errors.New("missing --time-begin argument value"))
				}
				r.TimeBegin = _time.ParseTimestamp(argValue)
				i++
			} else if arg == "-n" || arg == "--number" {
				if i+1 >= len(args) {
					panic(fmt.Errorf("missing %s argument value", arg))
				}
				r.Number = uint32(_time.ParseUint(argValue, 31))
				i++
			} else if arg == "-t" || arg == "--template" {
				r.PrintConfigTemplate()
				return false, nil
			} else {
				fmt.Printf("unknown option: %s", arg)
				return false, nil
			}
		}
	}

	return true, r
}

// PrintVersion ...
func (me Options) PrintVersion() {
	fmt.Printf("version: %s\n", me.Version)
}

// PrintHelp ...
func (me Options) PrintHelp() {
	fmt.Println("\nloggen: Log generator for jog(https://github.com/qiangyt/jog) development only.")
	me.PrintVersion()
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println(" loggen [options] [app]")
	fmt.Printf("   - [app]: only generate logs for specified app (listed in config yaml). All by default\n")

	fmt.Println("Options:")
	fmt.Printf("  -c,  --config <config file path>                       ./loggen.yaml or ./loggen.yml by default.\n")
	fmt.Printf("  -d,  --debug                                           Print more error detail\n")
	fmt.Printf("  -h,  --help                                            Display this information\n")
	fmt.Printf("  -t,  --template                                        Print a convertion config YAML file template\n")
	fmt.Printf("  -V,  --version                                         Display app version information\n")
	fmt.Printf("  -t,  --template                                        Print the default config YAML file\n")

	fmt.Printf("  --time-begin <begin time>                              Timestamp of first log line, default is now \n")
	fmt.Printf("  -n,  --number <number of log lines to generate>        Default is 10 \n")

	fmt.Println()
}
