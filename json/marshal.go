package json

import (
	"fmt"
	"reflect"
)

func Marshal(
	input string,
	typ reflect.Type,
) (string, error) {
	j := "{"
	a := ""

	for i := range typ.NumField() {
		field := typ.Field(i)
		fieldName := field.Name
		jsonFieldName := fieldName
		if tagValue, ok := field.Tag.Lookup("json"); ok {
			jsonFieldName = tagValue
		}
		fieldType := field.Type
		var verb string
		switch fieldType.Kind() {
		case reflect.Bool:
			verb = "%%b"
		case reflect.Int:
			verb = "%%d"
		case reflect.Int8:
			verb = "%%d"
		case reflect.Int16:
			verb = "%%d"
		case reflect.Int32:
			verb = "%%d"
		case reflect.Int64:
			verb = "%%d"
		case reflect.Uint:
			verb = "%%d"
		case reflect.Uint8:
			verb = "%%d"
		case reflect.Uint16:
			verb = "%%d"
		case reflect.Uint32:
			verb = "%%d"
		case reflect.Uint64:
			verb = "%%d"
		case reflect.Uintptr:
			verb = "%%d"
		case reflect.Float32:
			verb = "%%f"
		case reflect.Float64:
			verb = "%%f"
		// case reflect.Complex64: verb = ""
		// case reflect.Complex128: verb = ""
		// case reflect.Array: verb = ""
		// case reflect.Chan: verb = ""
		// case reflect.Func: verb = ""
		// case reflect.Interface: verb = ""
		// case reflect.Map: verb = ""
		// case reflect.Pointer: verb = ""
		// case reflect.Slice: verb = ""
		case reflect.String:
			verb = "%%q"
			// case reflect.Struct: verb = ""
			// case reflect.UnsafePointer: verb = ""
		}
		j += fmt.Sprintf("%q: "+verb, jsonFieldName)
		a += fmt.Sprintf(", this.%s", fieldName)
		if i+1 < typ.NumField() {
			j += ","
		}
	}

	j += "}"

	output := `package {{.Package}}

import "encoding/json"
import "fmt"

var _ json.Marshaler = (*{{.Type.Name}})(nil)
	
func (this {{.Type.Name}}) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(%q%s)), nil
}`

	output = fmt.Sprintf(output, j, a)

	return output, nil
}
