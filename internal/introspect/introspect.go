package introspect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Header struct {
	Name string
	Path []int
	Wide bool
}

func introspectHeadersType(of reflect.Type, base []int) []Header {
	if of.Kind() != reflect.Struct {
		panic("input must be a struct")
	}

	headers := make([]Header, 0)

	for i := 0; i < of.NumField(); i++ {
		r := of.Field(i)

		path := base

		if r.Type.Kind() == reflect.Struct {
			path = append(path, r.Index...)
			headers = append(headers, introspectHeadersType(r.Type, path)...)
			continue
		}

		if r.Tag.Get("outfmt") == "" {
			continue
		}

		tags := strings.Split(r.Tag.Get("outfmt"), ",")

		path = append(path, r.Index...)

		wide := false
		if len(tags) == 2 && tags[1] == "wide" {
			wide = true
		}

		headers = append(headers, Header{Name: tags[0], Wide: wide, Path: path})
	}

	return headers
}

func isArray(object any) bool {
	if reflect.TypeOf(object).Kind() == reflect.Array || reflect.TypeOf(object).Kind() == reflect.Slice {
		return true
	}

	return false
}

func IntrospectHeaders(object any) []Header {
	if isArray(object) {
		return introspectHeadersType(reflect.TypeOf(object).Elem(), []int{})
	}

	return introspectHeadersType(reflect.TypeOf(object), []int{})
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

	case reflect.Array:
	case reflect.Slice:
	case reflect.Struct:
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
	if val.Field(idx[0]).Type().Kind() == reflect.Struct {
		return extractFieldValue(val.Field(idx[0]), idx[1:])
	}

	return convertToString(val.Field(idx[0]))
}

func extractFieldsValue(val reflect.Value, fields [][]int) []string {
	values := []string{}
	for _, field := range fields {
		values = append(values, extractFieldValue(val, field))
	}

	return values
}

func ExtractFields(object any, fields [][]int) []string {
	return extractFieldsValue(reflect.ValueOf(object), fields)
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

// TODO: actually use wide
func IntrospectFields(data any, headers []Header, wide bool) [][]string {
	fields := [][]int{}
	for _, header := range headers {
		fields = append(fields, header.Path)
	}

	return IntrospectFieldsWithPath(data, fields)
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

	of := reflect.TypeOf(data)
	if isArray(data) {
		of = of.Elem()
	}

	return getFieldTypePath(of, fields)
}
