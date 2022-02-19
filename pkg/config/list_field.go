package config

import (
	"fmt"
	"reflect"

	_ "github.com/qiangyt/loggen/res/statik"
)

type ListFieldT struct {
	Name     string
	Values   interface{}
	Chooser  Chooser
	Children map[string]Field
}

type ListField = *ListFieldT

func NewListField(path string, name string, data map[string]interface{}) ListField {
	values, hasValues := data["values"]
	if !hasValues {
		panic(fmt.Errorf("%s.%s: missing values", path, name))
	}

	reflect.TypeOf(values)
	return &ListFieldT{
		Name:   name,
		Values: []interface{}{values},
	}
}

func NormalizeStringListFieldData(path FieldPath, data []string) map[string]interface{} {
	values := []map[string]interface{}{}
	chooser := RandomChooser

	for _, s := range data {
		s, valueWeight := ParseFieldString(path.Path(), s)
		if HasWeight(valueWeight) {
			chooser = WeightedChooser
			values = append(values, map[string]interface{}{
				"weight": valueWeight,
				"name":   s,
			})
			continue
		}

		values = append(values, map[string]interface{}{
			"name": s,
		})
	}

	if chooser == WeightedChooser {
		for _, value := range values {
			if _, has := value["weight"]; !has {
				value["weight"] = DefaultWeight
			}
		}
	}

	r := map[string]interface{}{
		"type":    FieldType_List,
		"chooser": chooser,
		"values":  values,
	}
	return r
}

func NormalizeMapListFieldData(path FieldPath, presetFields map[string]Field, data map[string]interface{}) map[string]interface{} {
	values, found := data["values"]
	if !found {
		panic(fmt.Errorf("%s: missing values", path))
	}

	//chooser := RandomChooser

	alreadyNormalized := false

	valuesType := reflect.TypeOf(values)
	if valuesType.Kind() == reflect.Map {
		// convert it to array
		newValues := []map[string]interface{}{}
		for valueName, valueData := range values.(map[string]interface{}) {
			valuePath := NewFieldPath(fmt.Sprintf("%s.values.%s", path.Path(), valueName))
			valueDataM := NormalizeFieldData(valuePath, presetFields, valueData)
			newValues = append(newValues, valueDataM)
		}
		values = newValues

		alreadyNormalized = true
	}

	valuesType = reflect.TypeOf(values)
	switch valuesType.Kind() {
	case reflect.Array, reflect.Slice:
		if !alreadyNormalized {
			valuesA := values.([]map[string]interface{})
			for i, valueData := range valuesA {
				valuesA[i] = NormalizeFieldData(path, presetFields, valueData)
			}
		}
	default:
		panic(fmt.Errorf("%s: values must be either array/slice or map/object", path))
	}

	data["type"] = FieldType_List
	data["values"] = values

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
	return me.Value
}

func (me ListField) GetChooser() Chooser {
	return me.Chooser
}

func (me ListField) GetChildren() map[string]Field {
	return me.Children
}
