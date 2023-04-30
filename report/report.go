package report

import (
	"errors"
	"fmt"
	"strings"
)

// repesentation of column metadata for  report
type column struct {
	// header name
	header string
	// max width measured in spaces
	maxWidth int
}

// generic report intended for command-line display
//
// holds reporting metadata and stores options set when
// initializing a report
type report struct {
	// column specification; id → index
	columns []*column
	// column hash of header names
	columnIds map[string]int
	rows      [][]any
	settings  *reportSettings
}

// report formatting helpers
const (
	columnSpace  string = " "
	headerBorder string = "-"
)

// report errors
var (
	ErrNegagiveColumnWidth error  = errors.New("bad column width, cannot use a negative value")
	ErrUnknownHeader       string = "header [%s] unrecognized"
)

// adds header names order by argument position
//
// Will append header names to a list of report columns.
func (r *report) AddHeaders(headers ...string) {
	currColumnId := len(r.columns)
	for _, header := range headers {
		r.columns = append(r.columns, &column{header: header})
		r.columnIds[header] = currColumnId
		currColumnId++
	}
}

// adds a set of datapoints as a new row in the report
//
// Number of datapoints should align with the number of columns. This
// function will append "data placeholders" if there are less datapoints
// than the expected number of columns. If there more datapoints than
// the expected number of columns, the set of datapoints will be truncated.
func (r *report) AddRowData(datapoints ...any) {
	r.rows = append(r.rows, make([]any, 0))
	currRow := len(r.rows)
	expectedDatapointsCount := len(r.columns)
	datapointsCount := len(datapoints)
	if currRow != 0 {
		currRow--
	}

	// expected number of datapoints was not entered
	if datapointsCount != expectedDatapointsCount {
		datapointsCountDiff := expectedDatapointsCount - datapointsCount
		if datapointsCountDiff < 0 {
			// truncate
			datapoints = datapoints[:expectedDatapointsCount]
		} else {
			// or add placeholders
			for i := 0; i < datapointsCountDiff; i++ {
				datapoints = append(datapoints, r.settings.dataPlaceholder)
			}
		}
	}

	// inspect datapoints array for max width
	for columnIndex, dataValue := range datapoints {
		currentMaxWidth := r.columns[columnIndex].maxWidth
		currentDataValueLength := len(fmt.Sprintf("%v", dataValue))
		if currentDataValueLength > currentMaxWidth {
			r.columns[columnIndex].maxWidth = currentDataValueLength
		}

		// fmt.Printf("max width for column %d: %d\n", columnIndex, r.columns[columnIndex].maxWidth)
	}

	r.rows[currRow] = append(r.rows[currRow], datapoints...)
}

// gets the format string for any row in the report
//
// With the exception of the last column, every column will have:
// data (type any) and right padding (number of spaces and a space).
//
//	%v%*s "inside row": data, flag for number of spaces, space
//	%v    "outside row": data only
func (r *report) GetRowFormatString() string {
	headerFormat := new(strings.Builder)
	for i := 0; i < len(r.columns); i++ {
		if i != len(r.columns)-1 {
			headerFormat.WriteString("%v%*s")
		} else {
			headerFormat.WriteString("%v")
		}
	}

	headerFormat.WriteString("\n")

	return headerFormat.String()
}

// outputs a formatted report header
//
// Header includes: column names and a border underlining each
// column name.
func (r *report) PrintHeader() {
	headerFmt := r.GetRowFormatString()
	// headerElements := ((len(r.columns) - 1) * 2) + len(r.columns)
	headerNames := make([]any, 0)
	headerBorders := make([]any, 0)
	for idx, column := range r.columns {
		if idx < len(r.columns)-1 {
			headerNames = append(headerNames, column.header)
			headerBorders = append(headerBorders, strings.Repeat(headerBorder, len(column.header)))
			headerNames = append(headerNames, column.maxWidth+r.settings.columnPadding-len(column.header))
			headerBorders = append(headerBorders, column.maxWidth+r.settings.columnPadding-len(column.header))
			headerNames = append(headerNames, columnSpace)
			headerBorders = append(headerBorders, columnSpace)
		} else {
			// no need to handle column spacing for last column
			headerNames = append(headerNames, column.header)
			headerBorders = append(headerBorders, strings.Repeat(headerBorder, len(column.header)))
		}
	}

	fmt.Println(">>> headerNames content")
	fmt.Printf(">>> format: %q\n", headerFmt)
	for i, v := range headerNames {
		fmt.Printf(">>>   %d | {%v} (%T)\n", i, v, v)
	}
	fmt.Fprintf(r.settings.writer, headerFmt, headerNames...)
	fmt.Fprintf(r.settings.writer, headerFmt, headerBorders...)
}

// output a formatted report body
func (r *report) PrintBody() {
	rowFmt := r.GetRowFormatString()
	for _, row := range r.rows {
		data := make([]any, 0)
		for columnIdx, columnData := range row {
			// need to convert columnData, type: any to a string... using fmt.Sprintf()
			dataAsString := fmt.Sprintf("%v", columnData)
			if columnIdx < len(r.columns)-1 {
				data = append(data, columnData)
				data = append(data, r.columns[columnIdx].maxWidth+r.settings.columnPadding-len(dataAsString))
				data = append(data, columnSpace)
			} else {
				// no need to handle column spacing for last column
				data = append(data, columnData)
			}
		}

		fmt.Fprintf(r.settings.writer, rowFmt, data...)
	}
}

// sets the max width (in spaces) for a column
//
// Columns are referenced by the header name. If headerName is not found,
// an error will be returned. Max column widths are computed as data is input
// into the report, see report.AddRowData(). Calling this function after data
// has been input can override the max width of a column. However, calling
// this function before adding data can cause the max width value to be overriden
// if a value encountered in report.AddRowData() exceeds the max width set here.
func (r *report) SetColumnMaxWidth(headerName string, maxWidth int) error {
	_, known := r.columnIds[headerName]
	if !known {
		return fmt.Errorf(ErrUnknownHeader, headerName)
	}

	if maxWidth < 0 {
		return ErrNegagiveColumnWidth
	}

	r.columns[r.columnIds[headerName]].maxWidth = maxWidth

	return nil
}

// initializes a new report
//
// Defaults will be used if options are not passed.
func NewReport(options ...reportOption) *report {
	r := &report{
		columns:   make([]*column, 0),
		columnIds: make(map[string]int),
		rows:      make([][]any, 0),
		settings:  newReportSettings(),
	}

	r.settings.processOptions(options...)

	return r
}
