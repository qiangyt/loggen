package config

import (
	"fmt"
	"strconv"
	"strings"

	_ "github.com/qiangyt/loggen/res/statik"
)

type FieldPathT struct {
	path   string
	name   string
	weight int
}

type FieldPath = *FieldPathT

func (me FieldPath) String() string {
	return me.Path()
}

func (me FieldPath) Path() string {
	return me.path
}

func (me FieldPath) Name() string {
	return me.name
}

func (me FieldPath) Weight() int {
	return me.weight
}

func (me FieldPath) HasWeight() bool {
	return HasWeight(me.Weight())
}

func HasWeight(weight int) bool {
	return weight >= 0
}

func NewFieldPath(path string) FieldPath {
	var name string

	indexOfDot := strings.LastIndexAny(path, ".")
	if indexOfDot < 0 {
		name = path
	} else {
		if indexOfDot == len(path)-1 {
			panic(fmt.Errorf("invalid path: %s - '.' cannot be last character", path))
		}
		name = path[indexOfDot+1:]
	}

	name, weight := ParseFieldString(path, name)
	if HasWeight(weight) {
		return &FieldPathT{
			path:   path,
			name:   name,
			weight: weight,
		}
	}

	return &FieldPathT{
		path:   path,
		name:   name,
		weight: -1,
	}
}

func ParseFieldString(hint string, s string) (string, int) {
	indexOfStar := strings.LastIndexAny(s, "*") //TODO: escape * using **
	if 0 < indexOfStar && indexOfStar < len(s)-1 {
		if weight, err := strconv.Atoi(s[indexOfStar+1:]); err == nil {
			if weight < 0 {
				panic(fmt.Errorf("%s - invalid weight %d (should be >= 0)", hint, weight))
			}
			s = s[:indexOfStar]
			return s, weight
		}
	}

	return s, -1
}
