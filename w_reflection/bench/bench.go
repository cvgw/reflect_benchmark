package bench

import "reflect"

func MapFoo(f interface{}, fMap map[string]interface{}, fields []string) {
	val := reflect.ValueOf(f)

	e := val.Elem()

	for _, field := range fields {
		sField := e.FieldByName(field)
		fVal := sField.String()
		fMap[field] = fVal
	}
}
