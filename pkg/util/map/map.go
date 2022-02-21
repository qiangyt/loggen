package _map

import (
	"fmt"

	"github.com/qiangyt/loggen/pkg/util/list"
	"github.com/qiangyt/loggen/pkg/util/str"
)

//TOOD: generic

func Must(m map[string]interface{}, key string, hint ...string) interface{} {
	if r, found := m[key]; found {
		return r
	}

	if len(hint) == 0 {
		panic(fmt.Errorf("missing %s", key))
	}
	panic(fmt.Errorf("%s: missing %s", hint[0], key))
}

func Default(m map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if r, found := m[key]; found {
		return r
	}
	return defaultValue
}

func Optional(m map[string]interface{}, key string) (interface{}, bool) {
	if r, found := m[key]; found {
		return r, true
	}
	return nil, false
}

func MustString(m map[string]interface{}, key string, hint ...string) string {
	r := Must(m, key, hint...)
	return str.Must(r, hint...)
}

func MustAnySlice(m map[string]interface{}, key string, hint ...string) []interface{} {
	v := Must(m, key, hint...)
	return list.MustAnySlice(v, hint...)
}

func DefaultString(m map[string]interface{}, key string, defaultValue string, hint ...string) string {
	r := Default(m, key, defaultValue)
	return str.Must(r, hint...)
}

func OptionalString(m map[string]interface{}, key string, hint string) (string, bool) {
	if r, found := Optional(m, key); found {
		return str.Must(r, hint), true
	}
	return "", false
}
