package cache

import (
	"reflect"
)

type Entry struct {
	Key   string
	Field string
}

type Fields map[string][]Entry

var cache map[reflect.Type]Fields

func init() {
	cache = make(map[reflect.Type]Fields)
}

func Set(key reflect.Type, fields Fields) {
	cache[key] = fields
}

func Has(key reflect.Type) bool {
	if _, ok := cache[key]; ok {
		return true
	}

	return false
}

func Get(key reflect.Type) Fields {
	return cache[key]
}
