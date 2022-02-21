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
		Values: []interface{}{values},
	}
}

func StringSliceToListFieldValues(data []string) map[string]interface{} {
	values := []map[string]interface{}{}
	for _, value := range data {
		values = append(values, map[string]interface{}{"value": value})
	}

	return map[string]interface{}{"values": values}
}

func AnySliceToListFieldValues(data []interface{}) map[string]interface{} {
	values := []map[string]interface{}{}
	for _, value := range data {
		values = append(values, map[string]interface{}{"value": value})
	}

	return map[string]interface{}{"values": values}
}

func NormalizeMapListFieldData(path FieldPath, presetFields map[string]Field, data map[string]interface{}) map[string]interface{} {
	_map.MustAnySlice(data, "values", path.Path())
	data["type"] = FieldType_List
	return data
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
