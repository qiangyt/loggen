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

func TryBuildStringReferFieldData(path FieldPath, presetFields map[string]Field, data string) (bool, FieldData) {
	if refer, yes := str.TrimPrefix(data, "$"); yes {
		return true, _buildRefFieldData(path, presetFields, refer, map[string]interface{}{})
	}
	return false, nil
}

func TryBuildMapRefFieldData(path FieldPath, presetFields map[string]Field, data map[string]interface{}) (bool, FieldData) {
	if fType, found := _map.OptionalString(data, "type", path.Path()); found {
		if fType != FieldType_Refer {
			return false, nil
		}
	}

	refer := _map.DefaultString(data, "refer", path.Name(), path.Path())
	return true, _buildRefFieldData(path, presetFields, refer, data)
}

func _buildRefFieldData(path FieldPath, presetFields map[string]Field, refer string, data map[string]interface{}) FieldData {
	var presetField Field
	if presetField = presetFields[refer]; presetField == nil {
		panic(fmt.Errorf("%s: preset field %s not found", path, refer))
	}

	data["type"] = FieldType_Refer

	return BuildFieldDataWithMap(path, presetFields, data)
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
