package outfmt

import (
	"encoding/json"
	"reflect"

	"github.com/NorskHelsenett/outfmt/internal/cache"
	"github.com/NorskHelsenett/outfmt/internal/table"

	"gopkg.in/yaml.v2"
)

type SpecField struct {
	Key   string
	Field string
}

type Spec map[string][]SpecField

type OutputFormat int64

const (
	OutputFormatJSON OutputFormat = iota
	OutputFormatYAML
	OutputFormatTable
	OutputFormatField
	OutputFormatCondition
)

type Config struct {
	Format OutputFormat

	/// only used for OutputFormatField
	AdditionalField string
}

func Format(data any, config *Config) ([]byte, error) {
	switch config.Format {
	case OutputFormatJSON:
		return json.Marshal(data)
	case OutputFormatYAML:
		return yaml.Marshal(data)
	case OutputFormatTable:
		t := table.New()
		t.For(data, "default")
		return t.Flush()
	case OutputFormatCondition:
		t := table.New()
		t.For(data, config.AdditionalField)
		return t.Flush()
	case OutputFormatField:
		t := table.New()
		t.Field(data, config.AdditionalField)
		return t.Flush()
	}
	return nil, nil
}

func Register(of any, spec *Spec) {
	fields := make(cache.Fields)
	for c, condition := range *spec {
		fields[c] = []cache.Entry{}
		for _, e := range condition {
			fields[c] = append(fields[c], cache.Entry{
				Key:   e.Key,
				Field: e.Field,
			})
		}
	}

	cache.Set(reflect.TypeOf(of), fields)
}
