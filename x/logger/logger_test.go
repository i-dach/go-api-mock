package logger

import (
	"testing"
)

func TestPrint(t *testing.T) {
	data := []struct {
		name string
		op   *Option
		ex   int
	}{
		{
			name: "debug - no format",
			op:   &Option{"id1", "DEBUG", "", "test string"},
			ex:   1,
		},
		{
			name: "info - no format",
			op:   &Option{"id1", "INFO", "", "test string"},
			ex:   0,
		},
		{
			name: "warn - no format",
			op:   &Option{"id1", "WARN", "", "test string"},
			ex:   1,
		},
		{
			name: "emerg - no format",
			op:   &Option{"id1", "EMERG", "", "test string"},
			ex:   1,
		},
		{
			name: "err - no format",
			op:   &Option{"id1", "ERROR", "", "test string"},
			ex:   0,
		},
		{
			name: "debug",
			op:   &Option{"id1", "DEBUG", "format %s", "test string"},
			ex:   1,
		},
		{
			name: "info",
			op:   &Option{"id1", "INFO", "format %s", "test string"},
			ex:   0,
		},
		{
			name: "warn",
			op:   &Option{"id1", "WARN", "format %s", "test string"},
			ex:   1,
		},
		{
			name: "emerg",
			op:   &Option{"id1", "EMERG", "format %s", "test string"},
			ex:   1,
		},
		{
			name: "err",
			op:   &Option{"id1", "ERROR", "format %s", "test string"},
			ex:   0,
		}}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			if d.ex != 1 {
				t.Log(Printer(d.op))
			}
		})
	}
}
