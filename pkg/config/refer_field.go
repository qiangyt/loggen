package config

import (
	"fmt"

	_map "github.com/qiangyt/loggen/pkg/util/map"
	"github.com/qiangyt/loggen/pkg/util/str"
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
	if refer, yes := str.TrimPrefix(data, "$"); yes {
		return true, _normalizeRefFieldData(path, presetFields, refer, map[string]interface{}{})
	}
	return false, map[string]interface{}{}
}

func TryToNormalizeMapRefFieldData(path FieldPath, presetFields map[string]Field, data map[string]interface{}) (bool, map[string]interface{}) {
	if fType, found := _map.OptionalString(data, "type", path.Path()); !found {
		if fType != FieldType_Refer {
			return false, data
		}
	}

	refer := _map.DefaultString(data, "refer", path.Name(), path.Path())
	return true, _normalizeRefFieldData(path, presetFields, refer, data)
}

func _normalizeRefFieldData(path FieldPath, presetFields map[string]Field, refer string, data map[string]interface{}) map[string]interface{} {
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
