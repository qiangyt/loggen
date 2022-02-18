package config

import (
	"time"
)

// ------------
type LoggerStateT struct {
	Config  Logger
	Message string
}

type LoggerState = *LoggerStateT

// ------------
type HostStateT struct {
	Config Host
}

type HostState = *HostStateT

// ------------
const (
	LevelTrace uint32 = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// ------------
type AppStateT struct {
	Config App

	Level  uint32
	Pid    uint32
	Logger LoggerState
	Host   HostState
}

type AppState = *AppStateT

type StateT struct {
	Config Config

	Timestamp time.Time
	App       AppState
}

type State = *StateT

func NewState(cfg Config) State {
	return &StateT{Config: cfg}
}
