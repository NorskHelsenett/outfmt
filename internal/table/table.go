package table

import (
	"fmt"
	"strings"
	"text/tabwriter"

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

func (t *Table) addHeaders(headers []introspect.Header) {
	collection := []string{}
	for _, header := range headers {
		collection = append(collection, header.Name)
	}

	t.AddRow(collection...)
}

func (t *Table) addValues(values [][]string) {
	for _, row := range values {
		t.AddRow(row...)
	}
}
func (t *Table) For(data any, wide bool) {
	headers := introspect.IntrospectHeaders(data)
	values := introspect.IntrospectFields(data, headers, wide)

	t.addHeaders(headers)
	t.addValues(values)
}

func (t *Table) Field(data any, field string) {
	path := introspect.IntrospectFieldPath(data, field)
	values := introspect.IntrospectFieldsWithPath(data, [][]int{path})

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
