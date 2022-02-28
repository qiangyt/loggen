package config

import (
	"github.com/qiangyt/loggen/pkg/util/str"
	_ "github.com/qiangyt/loggen/res/statik"
)

type FileFieldT struct {
}

type FileField = *FileFieldT

func NewFileField(name string, data map[string]interface{}) FileField {
	return &FileFieldT{
		//Name:  name,
		//Value: data["values"],
	}
}

func TryToBuildStringFileFieldData(path FieldPath, presetFields map[string]Field, data string) (bool, FieldData) {
	if url, found := str.TrimPrefix(data, "^"); found {
		// TODO: validate the file
		r := &FieldDataT{
			Path:   path,
			Type:   FieldType_File,
			Others: map[string]interface{}{"url": url},
		}
		return true, r
	}
	return false, nil
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
