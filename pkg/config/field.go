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
