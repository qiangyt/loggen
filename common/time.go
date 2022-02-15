package common

import (
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	"github.com/pkg/errors"
)

// ParseTimestamp ...
func ParseTimestamp(format string) time.Time {
	r, err := dateparse.ParseAny(format)
	if err != nil {
		panic(errors.Wrapf(err, "failed to parse time value: %s", format))
	}
	return r
}

func ParseUint(s string, bitSize int) uint64 {
	r, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		panic(errors.Wrapf(err, "failed to parse uint: %s (bit size=%d)", s, bitSize))
	}
	return r
}
