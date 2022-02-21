package str

import (
	"fmt"
	"reflect"
	"strings"
)

func TrimPrefix(s string, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):], true
	}
	return s, false
}

func Must(v interface{}, hint ...string) string {
	if r, is := v.(string); is {
		return r
	}

	if len(hint) == 0 {
		panic(fmt.Errorf("expect a string, but got %v (%v)", reflect.TypeOf(v), v))
	}
	panic(fmt.Errorf("%s: expect a string, but got %v (%v)", hint[0], reflect.TypeOf(v), v))
}

func CutLast(s string, sep string, allowAsPrefix bool, allowAsSuffix bool, hint ...string) (string, string, bool) {
	indexOfDot := strings.LastIndexAny(s, sep)
	if indexOfDot < 0 {
		return s, "", false
	}

	if indexOfDot == 0 {
		if !allowAsPrefix {
			if len(hint) == 0 {
				panic(fmt.Errorf("%s is not allowed as prefix: %s", sep, s))
			}
			panic(fmt.Errorf("%s: %s is not allowed as prefix: %s", hint, sep, s))
		}
	} else if indexOfDot == len(s)-1 {
		if allowAsSuffix {
			return s, "", true
		}
		if len(hint) == 0 {
			panic(fmt.Errorf("%s is not allowed as suffix: %s", sep, s))
		}
		panic(fmt.Errorf("%s: %s is not allowed as suffix: %s", hint, sep, s))
	}

	return s[:indexOfDot], s[indexOfDot+1:], true
}
