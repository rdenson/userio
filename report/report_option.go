package report

import (
	"io"
)

// defines a report option
//
// option will be applied to apply to report.settings
type reportOption interface {
	ApplyOption(*reportSettings)
}

// report options
type (
	optionColumnPadding   int
	optionDataPlaceholder string
	optionWriter          struct {
		io.Writer
	}
)

// option name (type)
const (
	OptTypeColumnPadding   string = "report.optionColumnPadding"
	OptTypeDataPlaceholder string = "report.optionDataPlaceholder"
	OptTypeWriter          string = "report.optionWriter"
)

func (cp optionColumnPadding) ApplyOption(rs *reportSettings) {
	rs.columnPadding = int(cp)
}

func (dp optionDataPlaceholder) ApplyOption(rs *reportSettings) {
	rs.dataPlaceholder = string(dp)
}

func (w optionWriter) ApplyOption(rs *reportSettings) {
	rs.writer = w.Writer
}

// report option: column padding
//
// define the space inbetween the data from one column and the next
func ColumnPadding(n int) reportOption {
	return optionColumnPadding(n)
}

// report option: placeholder for absent data
//
// define the placeholder to display when data is absent
func Placeholder(dp string) reportOption {
	return optionDataPlaceholder(dp)
}

// report option: how to output the report
//
// can define the "writer" for the report
//
// intended for report output destination, eg: stdout, a file, etc.
func Writer(w io.Writer) reportOption {
	return optionWriter{w}
}
