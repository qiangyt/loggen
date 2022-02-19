package config

type Tuple interface {
	Normalize(hint string)
	GetKey() string
	GetWeight() uint32
}
