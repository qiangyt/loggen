package list

import (
	"errors"
	"fmt"
	"reflect"
)

//TODO: is it possible to be array?
func MustAnySlice(v interface{}, hint ...string) []interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Interface {
		return v.([]interface{})
	}

	if len(hint) == 0 {
		panic(errors.New("must be []interface{}"))
	}
	panic(fmt.Errorf("%s: must be []interface{}", hint))
}
