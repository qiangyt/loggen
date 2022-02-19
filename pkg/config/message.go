package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

const (
	DefaultMessageRandomFileUrl   = "res:/message.random.txt"
	DefaultMessageWeightedFileUrl = "res:/message.weighted.csv"
	DefaultMessageChooser         = RandomChooser
)

type MessageT struct {
	Url string
}

// -----------------------
type RandomMessageT struct {
	MessageT
}

type RandomMessage = *RandomMessageT

func NewRandomMessage(hint string, input map[string]interface{}) RandomMessage {
	r := &RandomMessageT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "%s: failed decode random-message: %v", hint, input))
	}
	return r
}

func (me RandomMessage) Normalize(hint string) {
	if len(me.Url) == 0 {
		me.Url = DefaultMessageRandomFileUrl
	}
}

// -----------------------
type WeightedMessageT struct {
	MessageT
}

type WeightedMessage = *WeightedMessageT

func NewWeightedMessage(hint string, input map[string]interface{}) WeightedMessage {
	r := &WeightedMessageT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "%s: failed decode weighted-message: %v", hint, input))
	}
	return r
}

func (me WeightedMessage) Normalize(hint string) {
	if len(me.Url) == 0 {
		me.Url = DefaultMessageWeightedFileUrl
	}
}
