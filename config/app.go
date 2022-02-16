package config

type AppT struct {
	Name  string
	Type  string
	Level Level
	Pid   Pid
}

type App = *AppT

func NewApp() App {
	return &AppT{
		Level: NewLevel(),
		Pid:   NewPid(),
	}
}
