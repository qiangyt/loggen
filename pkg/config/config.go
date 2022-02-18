package config

import (
	"fmt"
	"path/filepath"

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
	Timestamp Timestamp
	Number    uint32
	Apps      []App
}

type Config = *ConfigT

func NewConfig() Config {
	return &ConfigT{}
}

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
			yamlText = res.GetDefaultConfigYaml()
		}
	}

	r := NewConfigWithYaml(yamlText)

	if options.Number > 0 {
		r.Number = options.Number
	}

	if !options.TimeBegin.IsZero() {
		r.Timestamp.Begin = options.TimeBegin
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

	return r
}

func NewConfigWithYaml(yamlText string) Config {
	r := NewConfig()
	err := yaml.Unmarshal([]byte(yamlText), &r)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse yaml"))
	}

	return r
}

func (me Config) Normalize() {
	me.NormalizeTimestamp()
	me.NormalizeNumber()
	me.NormalizeApps()
}

func (me Config) NormalizeTimestamp() {
	if me.Timestamp == nil {
		me.Timestamp = NewTimestamp()
	}
	me.Timestamp.Normalize("timestamp")
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
