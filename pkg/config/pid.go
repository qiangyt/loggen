package config

type PidT struct {
	Begin  uint32
	End    uint32
	Amount uint32
}

type Pid = *PidT

func NewPid() Pid {
	return &PidT{}
}
