package formator

import (
	"fmt"
)

var (
	formators map[string]Formator
)

func init() {
	formators = make(map[string]Formator)
}

func RegisterFormator(name string, formator Formator) {
	if HasFormator(name) {
		panic(fmt.Errorf("duplicated formator: %s", name))
	}
	formators[name] = formator
}

func HasFormator(name string) bool {
	_, found := formators[name]
	return found
}

func GetFormator(appName string, name string) Formator {
	r, found := formators[name]
	if !found {
		panic(fmt.Errorf("invalid app(name=%s) generator: %s", appName, name))
	}
	return r
}

func EnumerateFormatorNames() []string {
	r := []string{}
	for name, _ := range formators {
		r = append(r, name)
	}
	return r
}

func IsValidFormatorName(name string) bool {
	names := EnumerateFormatorNames()
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}
