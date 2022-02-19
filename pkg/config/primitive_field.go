package config

import (
	_ "github.com/qiangyt/loggen/res/statik"
)

type PrimitiveFieldT struct {
	Name  string
	Value interface{}
}

type PrimitiveField = *PrimitiveFieldT

func NewPrimitiveField(name string, data map[string]interface{}) PrimitiveField {
	return &PrimitiveFieldT{
		Name:  name,
		Value: data["value"],
	}
}

func NormalizePrimitiveFieldData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"type":  FieldType_Primitive,
		"value": data,
	}
}

func (me PrimitiveField) GetType() FieldType {
	return FieldType_Primitive
}

func (me PrimitiveField) Normalize(hint string) {
	// nothing to do
}

func (me PrimitiveField) GetName() string {
	return me.Name
}

func (me PrimitiveField) GetValue() interface{} {
	return me.Value
}

func (me PrimitiveField) GetChooser() Chooser {
	return HardcodedChooser
}

func (me PrimitiveField) GetChildren() map[string]Field {
	return nil
}
