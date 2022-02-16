package config

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
