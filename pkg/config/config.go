package config

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	_io "github.com/qiangyt/loggen/pkg/io"
	"github.com/qiangyt/loggen/pkg/res"
	_ "github.com/qiangyt/loggen/res/statik"
	"gopkg.in/yaml.v2"
)

const (
	DefaultNumber uint32 = 10
)

type ConfigT struct {
	Number uint32
	Fields map[string]interface{}
	Apps   map[string]App

	timestampField       Timestamp
	levelField           Level
	randomMessageField   RandomMessage
	weightedMessageField WeightedMessage
	pidField             Pid
}

type Config = *ConfigT

func NewConfigWithOptions(options Options) Config {
	var yamlText string

	configFilePath := options.ConfigFilePath
	if len(configFilePath) == 0 {
		exeDir := _io.ExeDirectory()
		configFilePath = filepath.Join(exeDir, "loggen.yaml")
		if !_io.FileExists(configFilePath) {
			configFilePath = filepath.Join(exeDir, "loggen.yml")
		}

		if _io.FileExists(configFilePath) {
			yamlText = _io.ReadTextFile(configFilePath)
		} else {
			yamlText = res.ReadDefaultConfigYaml()
		}
	}

	r := NewConfigWithYaml(yamlText)

	if options.Number > 0 {
		r.Number = options.Number
	}

	if !options.TimeBegin.IsZero() {
		timestampF := r.Fields["timestamp"].(Timestamp)
		timestampF.Begin = options.TimeBegin
	}

	if len(options.AppName) > 0 {
		var app App
		for _, a := range r.Apps {
			if a.Name == options.AppName {
				app = a
				break
			}
		}
		if app == nil {
			panic(fmt.Errorf("app %s not found", options.AppName))
		}

		r.Apps = []App{app}
	}

	r.Normalize()

	return r
}

func NewConfig(input map[string]interface{}) Config {
	r := &ConfigT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "failed decode config: %v", input))
	}
	return r
}

func NewConfigWithYaml(yamlText string) Config {
	m := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(yamlText), &m)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse yaml"))
	}

	r := NewConfig(m)
	r.Normalize()
	return r
}

func (me Config) Normalize() {
	me.NormalizeNumber()
	me.NormalizeFields()
	me.NormalizeGenerators()
	me.NormalizeApps()
}

func (me Config) NormalizeFields() {
	for n, f := range me.Fields {
		if n == "timestamp" {
			//TODO: check type cast
			me.timestampField = NewTimestamp("timestamp", f.(map[string]interface{}))
		} else if n == "level" {
			//TODO: check type cast
			me.levelField = NewLevel("level", f.(map[string]interface{}))
		} else if n == "timestamp" {
			//TODO: check type cast
			me.timestampField = NewTimestamp("timestamp", f.(map[string]interface{}))
		} else if n == "timestamp" {
			//TODO: check type cast
			me.timestampField = NewTimestamp("timestamp", f.(map[string]interface{}))
		} else if n == "timestamp" {
			//TODO: check type cast
			me.timestampField = NewTimestamp("timestamp", f.(map[string]interface{}))
		}
	}

	me.NormalizeTimestamp()

	if me.levelField == nil {
		me.levelField = NewLevel("level", map[string]interface{}{})
	}
	me.Fields["level"] = me.levelField
	me.levelField.Normalize("level")

	me.NormalizeRandomMessage()
	me.NormalizeWeightedMessage()
	me.NormalizePid()
	me.NormalizeOtherFields()
}

func (me Config) NormalizeTimestamp() {
	var f Timestamp
	m, found := me.Fields["timestamp"]
	if !found {
		f = NewTimestamp("timestamp", map[string]interface{}{})
	} else {
		//TODO: check type cast
		f = NewTimestamp("timestamp", m.(map[string]interface{}))
	}

	f.Normalize("timestamp")
	me.timestampField = f
	me.Fields["timestamp"] = f
}

func (me Config) NormalizeLevel() {
	f, found := me.Fields["level"]
	if !found {
		me.levelField = NewLevel("level", map[string]interface{}{})
	} else {
		//TODO: check type cast
		me.levelField = f.(Level)
	}

	me.levelField.Normalize("level")
}

func (me Config) NormalizeRandomMessage() {
	f, found := me.Fields["random-message"]
	if !found {
		me.randomMessageField = NewRandomMessage("random-message", map[string]interface{}{})
	} else {
		//TODO: check type cast
		me.randomMessageField = f.(RandomMessage)
	}

	me.randomMessageField.Normalize("random-message")
}

func (me Config) NormalizeWeightedMessage() {
	f, found := me.Fields["weighted-message"]
	if !found {
		me.weightedMessageField = NewWeightedMessage("weighted-message", map[string]interface{}{})
	} else {
		//TODO: check type cast
		me.weightedMessageField = f.(WeightedMessage)
	}

	me.weightedMessageField.Normalize("weighted-message")
}

func (me Config) NormalizePid() {
	f, found := me.Fields["pid"]
	if !found {
		me.pidField = NewPid("pid", map[string]interface{}{})
	} else {
		//TODO: check type cast
		me.pidField = f.(Pid)
	}

	me.pidField.Normalize("pid")
}

func (me Config) NormalizeNumber() {
	if me.Number == 0 {
		me.Number = DefaultNumber
	}
}

func (me Config) NormalizeApps() {
	if len(me.Apps) == 0 {
		panic(errors.New("at least 1 app is required"))
	}

	byNames := map[string]App{}
	for idx, app := range me.Apps {
		hint := fmt.Sprintf("apps[%d]", idx)

		name := app.Name
		if _, found := byNames[name]; found {
			panic(fmt.Errorf("%s.name(%s) is duplicated", hint, name))
		}
		byNames[name] = app

		app.Normalize(me, hint)
	}
}

func (me Config) NormalizeGenerators() {
	//TODO
}
