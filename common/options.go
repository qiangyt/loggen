package common

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
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

	timeBegin       time.Time
	timeIntervalMin uint32
	timeIntervalMax uint32

	levelWeightTrace uint32
	levelWeightDebug uint32
	levelWeightInfo  uint32
	levelWeightWarn  uint32
	levelWeightError uint32
	levelWeightFatal uint32

	number uint32
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

func (i Options) TimeBegin() time.Time {
	return i.timeBegin
}

func (i Options) TimeIntervalMin() uint32 {
	return i.timeIntervalMin
}

func (i Options) TimeIntervalMax() uint32 {
	return i.timeIntervalMax
}

func (i Options) Number() uint32 {
	return i.number
}

func (i Options) LevelWeightTrace() uint32 {
	return i.levelWeightTrace
}

func (i Options) LevelWeightDebug() uint32 {
	return i.levelWeightDebug
}

func (i Options) LevelWeightInfo() uint32 {
	return i.levelWeightInfo
}

func (i Options) LevelWeightWarn() uint32 {
	return i.levelWeightWarn
}

func (i Options) LevelWeightError() uint32 {
	return i.levelWeightError
}

func (i Options) LevelWeightFatal() uint32 {
	return i.levelWeightFatal
}

func OptionsWithCommandLine(version string) (bool, Options) {

	r := &OptionsT{
		debug:            false,
		subArgs:          []string{},
		version:          version,
		timeBegin:        time.Now(),
		timeIntervalMin:  10,        // 10 ms
		timeIntervalMax:  10 * 1000, // 10 seconds
		levelWeightTrace: 5,
		levelWeightDebug: 5,
		levelWeightInfo:  70,
		levelWeightWarn:  10,
		levelWeightError: 5,
		levelWeightFatal: 5,
		number:           10,
	}

	args := os.Args
	for i := 1; i < len(args); i++ {
		arg := args[i]

		argValue := ""
		if i+1 < len(args) {
			argValue = args[i+1]
		}

		var isOption bool
		if arg[0:1] == "-" {
			isOption = true
		} else {
			isOption = false
		}

		if isOption {
			if (arg == "-h" || arg == "--help") && len(r.logType) == 0 {
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
			} else if arg == "--time-begin" {
				if i+1 >= len(args) {
					panic(errors.New("missing --time-begin argument value"))
				}
				r.timeBegin = ParseTimestamp(argValue)
				i++
			} else if arg == "--time-interval-min" {
				if i+1 >= len(args) {
					panic(errors.New("missing --time-interval-min argument value"))
				}
				r.timeIntervalMin = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--time-interval-max" {
				if i+1 >= len(args) {
					panic(errors.New("missing --time-interval-max argument value"))
				}
				r.timeIntervalMax = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--level-weight-trace" {
				if i+1 >= len(args) {
					panic(errors.New("missing --level-weight-trace argument value"))
				}
				r.levelWeightTrace = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--level-weight-debug" {
				if i+1 >= len(args) {
					panic(errors.New("missing --level-weight-debug argument value"))
				}
				r.levelWeightDebug = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--level-weight-info" {
				if i+1 >= len(args) {
					panic(errors.New("missing --level-weight-info argument value"))
				}
				r.levelWeightInfo = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--level-weight-warn" {
				if i+1 >= len(args) {
					panic(errors.New("missing --level-weight-warn argument value"))
				}
				r.levelWeightWarn = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--level-weight-error" {
				if i+1 >= len(args) {
					panic(errors.New("missing --level-weight-error argument value"))
				}
				r.levelWeightError = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "--level-weight-fatal" {
				if i+1 >= len(args) {
					panic(errors.New("missing --level-weight-fatal argument value"))
				}
				r.levelWeightFatal = uint32(ParseUint(argValue, 20))
				i++
			} else if arg == "-n" || arg == "--number" {
				if i+1 >= len(args) {
					panic(fmt.Errorf("missing %s argument value", arg))
				}
				r.number = uint32(ParseUint(argValue, 31))
				i++
			} else {
				r.subArgs = append(r.subArgs, arg)
			}
		} else if i == 1 {
			switch arg {
			case LogType_bunyan:
				r.logType = LogType_bunyan
			default:
				fmt.Printf("log type '%s' is not supported\n", arg)
				r.PrintVersion()
				return false, nil
			}
		} else {
			r.subArgs = append(r.subArgs, arg)
		}
	}

	if len(r.logType) == 0 {
		fmt.Printf("missing log type argument\n")
		r.PrintVersion()
		return false, nil
	}

	if r.timeIntervalMin > r.timeIntervalMax {
		panic(fmt.Errorf("--time-interval-min (%d) cannot be great than --time-interval-max (%d)",
			r.timeIntervalMin, r.timeIntervalMax))
	}

	return true, r
}

// PrintVersion ...
func (i Options) PrintVersion() {
	fmt.Printf("version: %s\n", i.Version())
}

// PrintHelp ...
func (i Options) PrintHelp() {
	fmt.Println("\nloggen: Log generator for jog(https://github.com/qiangyt/jog) development only.")
	i.PrintVersion()
	fmt.Println()

	fmt.Println("Global options:")
	fmt.Printf("  -d,  --debug                                                  Print more error detail\n")
	fmt.Printf("  -h,  --help                                                   Display this information\n")
	fmt.Printf("  <log type> -h,  server --help                                 Display log type specific help information\n")
	fmt.Printf("  -V,  --version                                                Display app version information\n")
	fmt.Printf("  --time-begin <begin time>                                     Timestamp of first log line, default is now \n")
	fmt.Printf("  --time-interval-min <minimal time interval by milliseconds>   Default is 10 \n")
	fmt.Printf("  --time-interval-max <maximal time interval by milliseconds>   Default is 10000 (10 sec) \n")

	fmt.Printf("  --level-weight-trace <random weight of TRACE level>           Default is 5 \n")
	fmt.Printf("  --level-weight-debug <random weight of DEBUG level>           Default is 5 \n")
	fmt.Printf("  --level-weight-info <random weight of INFO level>             Default is 70 \n")
	fmt.Printf("  --level-weight-warn <random weight of WARN level>             Default is 10 \n")
	fmt.Printf("  --level-weight-error <random weight of ERROR level>           Default is 5 \n")
	fmt.Printf("  --level-weight-fatal <random weight of FATAL level>           Default is 5 \n")

	fmt.Printf("  -n,  --number <number of log lines to generate                Default is 10 \n")
	fmt.Println()

	fmt.Println("Supported log types:")
	fmt.Println("  " + LogType_bunyan)
}
