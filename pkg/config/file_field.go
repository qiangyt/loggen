package config

import (
	"strings"

	_ "github.com/qiangyt/loggen/res/statik"
)

type FileFieldT struct {
}

type FileField = *FileFieldT

func NewFileField(name string, data map[string]interface{}) FileField {
	return &FileFieldT{
		Name:  name,
		Value: data["values"],
	}
}

func TryToNormalizeStringFileFieldData(data string) (bool, map[string]interface{}) {
	if strings.IndexAny(data, "^") == 0 {
		return true, map[string]interface{}{
			"type": FieldType_File,
			"url":  data[1:], // remove the leading '^'
		}
	}
	return false, map[string]interface{}{}
}

func (me FileField) GetType() FieldType {
	return FieldType_File
}

func (me FileField) Normalize(hint string) {
	// nothing to do
}

func (me FileField) GetName() string {
	panic("TOOD")
}

func (me FileField) GetValue() interface{} {
	panic("TOOD")
}

func (me FileField) GetChooser() Chooser {
	panic("TOOD")
}

func (me FileField) GetChildren() map[string]Field {
	panic("TOOD")
}
