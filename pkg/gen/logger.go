package gen

import (
	"math/rand"
	"strings"

	"github.com/qiangyt/loggen/pkg/config"
	_io "github.com/qiangyt/loggen/pkg/io"
	"github.com/qiangyt/loggen/pkg/res"
)

type LoggerGeneratorT struct {
	config   config.Logger
	messages []string
}

type LoggerGenerator = *LoggerGeneratorT

func NewLoggerGenerator(cfg config.Logger) LoggerGenerator {
	return &LoggerGeneratorT{
		config:   cfg,
		messages: BuildMessageArray(cfg.Message),
	}
}

func (me LoggerGenerator) Next() config.LoggerState {
	i := rand.Intn(len(me.messages))
	msg := me.messages[i]

	return &config.LoggerStateT{
		Config:  me.config,
		Message: msg,
	}
}

func BuildMessageArray(messageFilePath string) []string {
	var msg string
	if res.IsResourceUrl(messageFilePath) {
		msg = res.NewResourceWithUrl(messageFilePath).ReadString()
	} else {
		msg = _io.ReadTextFile(messageFilePath)
	}

	r := []string{}

	for _, msg := range strings.Split(msg, "\n") {
		msg = strings.Trim(msg, " \n\t\r")
		if len(msg) > 0 {
			r = append(r, msg)
		}
	}

	return r
}
