package config

import (
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/qiangyt/loggen/common"
	_io "github.com/qiangyt/loggen/pkg/io"
	_ "github.com/qiangyt/loggen/res/statik"
	"gopkg.in/yaml.v2"
)

type ConfigT struct {
	Timestamp Timestamp
	Number    uint32
	Apps      []App
}

type Config = *ConfigT

func NewConfig() Config {
	return &ConfigT{
		Timestamp: NewTimestamp(),
		Apps:      []App{NewApp()},
	}
}

func NewConfigWithOptions(options common.Options) Config {
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
			yamlText = GetDefaultConfigYaml()
		}
	}

	return NewConfigWithYaml(yamlText)
}

func NewConfigWithYaml(yamlText string) Config {
	r := NewConfig()
	err := yaml.Unmarshal([]byte(yamlText), &r)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse yaml"))
	}
	return r
}
