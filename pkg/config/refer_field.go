package config

import (
	"fmt"
	"strings"

	_ "github.com/qiangyt/loggen/res/statik"
)

type ReferFieldT struct {
	Name     string
	Values   interface{}
	Chooser  Chooser
	Children map[string]Field
	Target   Field
}

type ReferField = *ReferFieldT

func NewReferFieldT(target Field, name string, data map[string]interface{}) ReferField {
	panic("TODO")
}

func TryToNormalizeStringReferFieldData(path FieldPath, presetFields map[string]Field, data string) (bool, map[string]interface{}) {
	if strings.IndexAny(data, "$") == 0 {
		refer := data[1:] // remove the leading '$'
		dataM := map[string]interface{}{}
		return true, NormalizeRefFieldData(path, presetFields, refer, dataM)
	}
	return false, map[string]interface{}{}
}

func TryToNormalizeMapRefFieldData(path FieldPath, presetFields map[string]Field, data map[string]interface{}) (bool, map[string]interface{}) {
	fType := data["type"].(FieldType)
	if len(fType) > 0 {
		if fType != FieldType_Refer {
			return false, data
		}
	}

	refer := data["refer"].(string) //TODO: must be string
	if len(refer) == 0 {
		refer = path.Name()
	}

	return true, NormalizeRefFieldData(path, presetFields, refer, data)
}

func NormalizeRefFieldData(path FieldPath, presetFields map[string]Field, refer string, data map[string]interface{}) map[string]interface{} {
	var presetField Field
	if presetField = presetFields[refer]; presetField == nil {
		panic(fmt.Errorf("%s: preset field %s not found", path, refer))
	}

	data["type"] = FieldType_Refer
	data["refer"] = presetField

	return data
}

func (me ReferField) Normalize(hint string) {
	panic("TODO")
}

func (me ReferField) GetName() string {
	panic("TODO")
}

func (me ReferField) GetValue() interface{} {
	panic("TODO")
}

func (me ReferField) GetChooser() Chooser {
	panic("TODO")
}

func (me ReferField) GetChildren() map[string]Field {
	panic("TODO")
}
