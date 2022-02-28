package config

import (
	_map "github.com/qiangyt/loggen/pkg/util/map"
	_ "github.com/qiangyt/loggen/res/statik"
)

var DefaultWeight = 1

type ListFieldT struct {
	Name     string
	Values   interface{}
	Chooser  Chooser
	Children map[string]Field
}

type ListField = *ListFieldT

func NewListField(path string, name string, data map[string]interface{}) ListField {
	values := _map.Must(data, "values", path)

	//TODO? reflect.TypeOf(values)
	return &ListFieldT{
		Name:   name,
		Values: values,
	}
}

func BuildFieldDataWithStringSlice(path FieldPath, presetFields map[string]Field, data []string) FieldData {
	values := []map[string]interface{}{}
	for _, value := range data {
		values = append(values, map[string]interface{}{"value": value})
	}

	r := &FieldDataT{
		Path: path,
		Type: FieldType_List,
	}

	r.NormalizeValues(presetFields, values)
	return r
}

func BuildFieldDataWithAnySlice(path FieldPath, presetFields map[string]Field, data []interface{}) FieldData {
	values := []map[string]interface{}{}
	for _, value := range data {
		values = append(values, map[string]interface{}{"value": value})
	}

	r := &FieldDataT{
		Path: path,
		Type: FieldType_List,
	}

	r.NormalizeValues(presetFields, values)
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

func (me ListField) GetValue() interface{} {
	return me.Values
}

func (me ListField) GetChooser() Chooser {
	return me.Chooser
}

func (me ListField) GetChildren() map[string]Field {
	return me.Children
}
