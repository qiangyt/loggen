package config

import (
	"fmt"
	"strconv"

	"github.com/qiangyt/loggen/pkg/util/str"
	_ "github.com/qiangyt/loggen/res/statik"
)

type FieldPathT struct {
	path   string
	parent string
	name   string
	weight int
}

type FieldPath = *FieldPathT

func (me FieldPath) String() string {
	return me.Path()
}

func (me FieldPath) Parent() string {
	return me.parent
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

func (me FieldPath) Child(name string) FieldPath {
	childPath := fmt.Sprintf("%s.%s", me.Path(), name)
	return NewFieldPath(childPath)
}

func HasWeight(weight int) bool {
	return weight >= 0
}

func NewFieldPath(path string) FieldPath {
	name, parent, _ := str.CutLast(path, ".", false, false, path)

	if newName, weight := ParseFieldString(path, name); HasWeight(weight) {
		return &FieldPathT{
			path:   path,
			parent: parent,
			name:   newName,
			weight: weight,
		}
	}

	return &FieldPathT{
		path:   path,
		parent: parent,
		name:   name,
		weight: -1,
	}
}

func ParseFieldString(hint string, s string) (string, int) {
	if name, weightS, found := str.CutLast(s, "*", false, false, hint); found {
		//TODO: escape * using **
		if weight, err := strconv.Atoi(weightS); err == nil {
			if weight < 0 {
				panic(fmt.Errorf("%s - invalid weight %s (should be >= 0)", hint, weightS))
			}
			return name, weight
		}
	}

	return s, -1
}
