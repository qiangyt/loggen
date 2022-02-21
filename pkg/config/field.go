package config

import (
	"fmt"

	_ "github.com/qiangyt/loggen/res/statik"
)

type Chooser string

const (
	HardcodedChooser Chooser = "hardcoded"
	WeightedChooser  Chooser = "weighted"
	RandomChooser    Chooser = "random"
)

type FieldType string

const (
	FieldType_Primitive = "primitive"
	FieldType_Refer     = "refer"
	FieldType_List      = "url"
	FieldType_File      = "file"
)

type Field interface {
	GetType() FieldType
	Normalize(hint string)
	GetName() string
	GetValue() interface{}
	GetChooser() Chooser
	GetChildren() map[string]Field
}

//
func BuildField(path FieldPath, presetFields map[string]Field, data interface{}) Field {
	dataM := NormalizeFieldData(presetFields, name, data)
	fType := dataM["type"].(FieldType)

	switch fType {
	case FieldType_Primitive:
		return NewPrimitiveField(name, dataM)
	case FieldType_Refer:
		return BuildRefField(presetFields, name, dataM)
	case FieldType_List:
		return NewListField(path, name, dataM)
	case FieldType_File:
		return BuildFileField(name, dataM)
	default:
		panic(fmt.Errorf("unexpected field: %s", fType))
	}
}

func NormalizeFieldData(path FieldPath, presetFields map[string]Field, data interface{}) map[string]interface{} {
	switch data.(type) {
	case []string: //TODO: other type of lices
		dataM := StringSliceToListFieldValues(data.([]string))
		return NormalizeFieldData(path, presetFields, dataM)
	case []interface{}:
		dataM := AnySliceToListFieldValues(data.([]interface{}))
		return NormalizeFieldData(path, presetFields, dataM)
	case string:
		{
			dataS := data.(string)
			if yes, dataM := TryToNormalizeStringReferFieldData(path, presetFields, dataS); yes {
				return dataM
			}
			if yes, dataM := TryToNormalizeStringFileFieldData(dataS); yes {
				return dataM
			}
			return NormalizePrimitiveFieldData(dataS)
		}
	case map[string]interface{}:
		{
			dataM := data.(map[string]interface{})
			if yes, dataM := TryToNormalizeMapRefFieldData(path, presetFields, dataM); yes {
				return dataM
			}
			return NormalizeMapListFieldData(path, presetFields, dataM)
		}
	default:
		return NormalizePrimitiveFieldData(data)
	}
}

func BuildRefField(presetFields map[string]Field, name string, data map[string]interface{}) Field {
	panic("TODO")
}

func BuildFileField(name string, data map[string]interface{}) Field {
	panic("TODO")
}
