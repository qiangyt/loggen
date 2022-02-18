package gen

import (
	"math/rand"
	"time"

	"github.com/qiangyt/loggen/pkg/config"
)

const (
	DefaultTimestampIntervalMin uint32 = 10
	DefaultTimestampIntervalMax uint32 = 10000
)

type TimestampGeneratorT struct {
	config config.Timestamp
}

type TimestampGenerator = *TimestampGeneratorT

func NewTimestampGenerator(config config.Timestamp) TimestampGenerator {
	return &TimestampGeneratorT{
		config: config,
	}
}

func (me TimestampGenerator) Next(current time.Time) time.Time {
	cfg := me.config

	if current.IsZero() {
		return cfg.Begin
	}

	intervalDeta := cfg.IntervalMax - cfg.IntervalMin
	interval := cfg.IntervalMin + uint32(rand.Int31n(int32(intervalDeta)))
	dura := time.Duration(interval * 1000 * 1000)

	return current.Add(dura)
}
