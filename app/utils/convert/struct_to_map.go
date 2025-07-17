package convertutil

import (
	"reflect"
	"strings"
)

func StructToMap(data any) map[string]any {
	result := make(map[string]any)
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// skip unexported fields
		if !fieldVal.CanInterface() {
			continue
		}

		// handle embedded structs recursively
		if field.Anonymous && fieldVal.Kind() == reflect.Struct {
			embeddedMap := StructToMap(fieldVal.Interface())
			for k, v := range embeddedMap {
				result[k] = v
			}
			continue
		}

		// skip zero value fields
		if isZeroValue(fieldVal) {
			continue
		}

		tag := field.Tag.Get("json")
		columnName := parseJSONTag(tag)
		if columnName == "" {
			columnName = field.Name
		}

		result[columnName] = fieldVal.Interface()
	}

	return result
}

func parseJSONTag(tag string) string {
	// json tag can be like "name,omitempty"
	parts := strings.Split(tag, ",")
	if len(parts) > 0 && parts[0] != "" && parts[0] != "-" {
		return parts[0]
	}
	return ""
}

func isZeroValue(v reflect.Value) bool {
	zero := reflect.Zero(v.Type()).Interface()
	current := v.Interface()
	return reflect.DeepEqual(current, zero)
}
