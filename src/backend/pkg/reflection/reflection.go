package reflection

import (
	"fmt"
	"reflect"
	"strconv"
)

func GetType(mod any) string {
	return reflect.TypeOf(mod).Name()
}

func GetArrayOfFields(mod any) []string {
	var fields []string
	numFields := reflect.TypeOf(mod).NumField()
	for i := 0; i < numFields; i++ {
		fields = append(fields, reflect.TypeOf(mod).Field(i).Name)
	}
	return fields
}

func GetValueOfStruct(model any, field string) string {
	result := reflect.ValueOf(model).FieldByName(field)

	switch result.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(result.Int(), 10)
	case reflect.Bool:
		return fmt.Sprintf("%t", result.Bool())
	case reflect.Float64, reflect.Float32:
		return fmt.Sprintf("%f", result.Float()) //result.Float()
	case reflect.String:
		return fmt.Sprintf("'%s'", result.String())
	case reflect.Interface:
		return fmt.Sprintf("'%s'", result.Interface())
	}
	return fmt.Sprintf("'%s'", result.Interface()) //result.Interface()
}
