package table

import (
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/eskpil/outfmt/internal/cache"
	"github.com/eskpil/outfmt/internal/introspect"
)

type Table struct {
	out    *strings.Builder
	writer *tabwriter.Writer
}

func New() *Table {
	t := new(Table)

	t.out = new(strings.Builder)
	t.writer = tabwriter.NewWriter(t.out, 0, 0, 2, ' ', 0)

	return t
}

func (t *Table) addHeaders(headers []string) {
	t.AddRow(headers...)
}

func (t *Table) addValues(values [][]string) {
	for _, row := range values {
		t.AddRow(row...)
	}
}

func (t *Table) For(data any, condition string) {
	if !cache.Has(introspect.Strip(reflect.TypeOf(data))) {
		panic(fmt.Sprintf("type: %s has not been registered", introspect.Strip(reflect.TypeOf(data))))
	}

	fields := cache.Get(introspect.Strip(reflect.TypeOf(data)))

	headers := make([]string, 0)
	paths := make([][]int, 0)

	for _, e := range fields[condition] {
		headers = append(headers, e.Key)
		paths = append(paths, introspect.IntrospectFieldPath(data, e.Field))
	}

	values := introspect.IntrospectFieldsWithPath(data, paths)

	t.addHeaders(headers)
	t.addValues(values)
}

func (t *Table) Field(data any, field string) {
	paths := make([][]int, 0)
	for _, field := range strings.Split(field, ",") {
		paths = append(paths, introspect.IntrospectFieldPath(data, field))
	}

	values := introspect.IntrospectFieldsWithPath(data, paths)

	t.addValues(values)
}

func (t *Table) AddRow(values ...string) {
	row := new(strings.Builder)
	for _, val := range values {
		fmt.Fprintf(row, "%s\t", val)
	}
	fmt.Fprintln(t.writer, row.String())
}

func (t *Table) Flush() ([]byte, error) {
	err := t.writer.Flush()
	return []byte(t.out.String()), err
}
