package outfmt

import (
	"encoding/json"

	"github.com/eskpil/outfmt/internal/table"

	"gopkg.in/yaml.v2"
)

type OutputFormat int64

const (
	OutputFormatJSON = iota
	OutputFormatYAML
	OutputFormatTable
	OutputFormatField
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
		t.For(data, false)
		return t.Flush()
	case OutputFormatField:
		t := table.New()
		t.Field(data, config.AdditionalField)
		return t.Flush()
	}

	return nil, nil
}
