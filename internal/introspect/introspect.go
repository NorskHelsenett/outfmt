package introspect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func isArray(object any) bool {
	if reflect.TypeOf(object).Kind() == reflect.Array || reflect.TypeOf(object).Kind() == reflect.Slice {
		return true
	}

	return false
}

func Strip(from reflect.Type) reflect.Type {
	if from.Kind() == reflect.Array || from.Kind() == reflect.Slice || from.Kind() == reflect.Pointer {
		return Strip(from.Elem())
	}

	return from
}

func convertToString(val reflect.Value) string {
	switch val.Type().Kind() {
	case reflect.Bool:
		if val.Bool() {
			return "true"
		}

		return "false"
	case reflect.Int:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Int8:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Int16:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Int32:
		return fmt.Sprintf("%d", val.Int())
	case reflect.Int64:
		return fmt.Sprintf("%d", val.Int())

	case reflect.Uint:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Uint8:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Uint16:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Uint32:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)

	case reflect.Float32:
		return fmt.Sprintf("%f", val.Float())
	case reflect.Float64:
		return fmt.Sprintf("%f", val.Float())
	case reflect.String:
		return val.String()
	case reflect.Pointer:
		return convertToString(val.Elem())

	case reflect.Struct:
	case reflect.Array:
	case reflect.Slice:
	case reflect.Uintptr:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
		panic(fmt.Sprintf("unsupported type: \"%s\"", val.Type().String()))
	case reflect.UnsafePointer:
	}

	return ""
}

func extractFieldValue(val reflect.Value, idx []int) string {
	field := val.Field(idx[0])

	if field.Type().Kind() == reflect.Struct {
		// TODO: Should more types get special treatment, if so, what types?
		if field.Type().Name() == "Time" {
			return field.Addr().Interface().(*time.Time).String()
		}

		return extractFieldValue(field, idx[1:])
	}

	return convertToString(field)
}

func extractFieldsValue(val reflect.Value, fields [][]int) []string {
	values := []string{}
	for _, field := range fields {
		values = append(values, extractFieldValue(val, field))
	}

	return values
}

func ExtractFields(object any, fields [][]int) []string {
	val := reflect.ValueOf(object)
	if val.Type().Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return extractFieldsValue(val, fields)
}

func extractFieldsOfArray(data any, fields [][]int) [][]string {
	if !isArray(data) {
		panic("data is not array")
	}

	val := reflect.ValueOf(data)

	rows := make([][]string, 0)

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		rows = append(rows, extractFieldsValue(elem, fields))
	}

	return rows
}

func IntrospectFieldsWithPath(data any, paths [][]int) [][]string {
	if !isArray(data) {
		return [][]string{ExtractFields(data, paths)}
	}

	return extractFieldsOfArray(data, paths)
}

func getFieldTypePath(of reflect.Type, fields []string) []int {
	path := []int{}

	child, found := of.FieldByName(fields[0])
	if !found {
		// TODO: Possibily report this better
		panic(fmt.Sprintf("field: %s was not found", fields[0]))
	}

	path = append(path, child.Index...)
	if len(fields) > 1 {
		path = append(path, getFieldTypePath(child.Type, fields[1:])...)
	}

	return path
}

func IntrospectFieldPath(data any, field string) []int {
	fields := strings.Split(field, ".")
	of := Strip(reflect.TypeOf(data))
	return getFieldTypePath(of, fields)
}
