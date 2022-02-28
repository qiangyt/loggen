package config

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	_ "github.com/qiangyt/loggen/res/statik"
)

type FieldDataT struct {
	Path    FieldPath
	Type    FieldType
	Values  interface{} // is slice, actually
	Name    string
	Chooser Chooser
	Others  map[string]interface{} `mapstructure:",remain"`
}

type FieldData = *FieldDataT

func (me FieldData) NormalizeValues(presetFields map[string]Field, values []map[string]interface{}) {
	//chooser := RandomChooser
	valuesType := reflect.TypeOf(values)
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

	me.Type = FieldType_List
	me.Values = valuesData
}

func NewFieldData(path FieldPath, presetFields map[string]Field, data interface{}) FieldData {
	switch data.(type) {
	//TODO: other type of lices
	case []string:
		return BuildFieldDataWithStringSlice(path, presetFields, data.([]string))
	case []interface{}:
		return BuildFieldDataWithAnySlice(path, presetFields, data.([]interface{}))
	case string:
		return BuildFieldDataWithString(path, presetFields, data.(string))
	case map[string]interface{}:
		return BuildFieldDataWithMap(path, presetFields, data.(map[string]interface{}))
	default:
		return BuildPrimitiveFieldData(path, presetFields, data)
	}
}

func BuildFieldDataWithMap(path FieldPath, presetFields map[string]Field, data map[string]interface{}) FieldData {
	if yes, r := TryBuildMapRefFieldData(path, presetFields, data); yes {
		return r
	}

	r := &FieldDataT{Path: path}
	if err := mapstructure.Decode(data, &r); err != nil {
		panic(errors.Wrapf(err, "%s: failed to decode map -> %v", path.Path(), data))
	}

	if r.Values != nil {
		valuesData := NewFieldData(path.Child("values"), presetFields, r.Values)
		values := valuesData.Values.([]map[string]interface{})
		r.NormalizeValues(presetFields, values)
	}

	return r
}

func BuildFieldDataWithString(path FieldPath, presetFields map[string]Field, data string) FieldData {
	if yes, r := TryBuildStringReferFieldData(path, presetFields, data); yes {
		return r
	}
	if yes, r := TryToBuildStringFileFieldData(path, presetFields, data); yes {
		return r
	}
	return BuildPrimitiveFieldData(path, presetFields, data)
}
