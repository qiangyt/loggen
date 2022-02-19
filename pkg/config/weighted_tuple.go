package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

const DefaultWeight = 1 //TODO: move elsewhere

type WeightedTupleT struct {
	Key    string
	Weight uint32
	Fields map[string]interface{} `mapstructure:",remain"`
}

type WeightedTuple = *WeightedTupleT

func NewWeightedTuple(input map[string]interface{}) WeightedTuple {
	r := &WeightedTupleT{}
	if err := mapstructure.Decode(input, &r); err != nil {
		panic(errors.Wrapf(err, "failed decode weighted tuple: %v", input))
	}
	return r
}

func (me WeightedTuple) Normalize(hint string) {
	if len(me.Key) == 0 {
		panic(fmt.Errorf("missing %s.key", hint))
	}
	if me.Weight == 0 {
		me.Weight = DefaultWeight
	}
}

func (me WeightedTuple) GetKey() string {
	return me.Key
}

func (me WeightedTuple) GetWeight() uint32 {
	return me.Weight
}
