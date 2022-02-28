package config

import (
	_map "github.com/qiangyt/loggen/pkg/util/map"
	_ "github.com/qiangyt/loggen/res/statik"
)

var DefaultWeight = 1

type ListFieldT struct {
	Name       string
	Candidates interface{}
	Chooser    Chooser
	Children   map[string]Field
}

type ListField = *ListFieldT

func NewListField(path string, name string, data map[string]interface{}) ListField {
	candidates := _map.Must(data, "candidates", path)

	//TODO? reflect.TypeOf(candidates)
	return &ListFieldT{
		Name:       name,
		Candidates: candidates,
	}
}

func BuildFieldDataWithStringSlice(path FieldPath, presetFields map[string]Field, data []string) FieldData {
	candidates := []map[string]interface{}{}
	for _, candidate := range data {
		candidates = append(candidates, map[string]interface{}{"candidate": candidate})
	}

	r := &FieldDataT{
		Path: path,
		Type: FieldType_List,
	}

	r.NormalizeCandidates(presetFields, candidates)
	return r
}

func BuildFieldDataWithAnySlice(path FieldPath, presetFields map[string]Field, data []interface{}) FieldData {
	candidates := []map[string]interface{}{}
	for _, candidate := range data {
		candidates = append(candidates, map[string]interface{}{"candidate": candidate})
	}

	r := &FieldDataT{
		Path: path,
		Type: FieldType_List,
	}

	r.NormalizeCandidates(presetFields, candidates)
	return r
}

func (me ListField) GetType() FieldType {
	return FieldType_List
}

func (me ListField) Normalize(hint string) {
	//TODO
}

func (me ListField) GetName() string {
	return me.Name
}

func (me ListField) GetCandidate() interface{} {
	return me.Candidates
}

func (me ListField) GetChooser() Chooser {
	return me.Chooser
}

func (me ListField) GetChildren() map[string]Field {
	return me.Children
}
