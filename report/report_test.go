package report

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type reportTestCase struct {
	expectErr   bool
	expects     any
	name        string
	setupErr    error
	setupReport func(testCase *reportTestCase) *report
}

type reportSuite struct {
	suite.Suite
}

func (rt *reportTestCase) setup() *report {
	return rt.setupReport(rt)
}

func TestReport(t *testing.T) {
	suite.Run(t, new(reportSuite))
}

func (rs *reportSuite) TestAddRowData() {
	testCases := []*reportTestCase{
		{
			name:    "should add placeholders for missing data",
			expects: []any{"apple", "-", "-"},
			setupReport: func(tc *reportTestCase) *report {
				rpt := NewReport()

				rpt.AddHeaders("h0", "h1", "h2")
				rpt.AddRowData("apple")

				return rpt
			},
		},
		{
			name:    "should truncate inputted datapoints if greater than the number of columns",
			expects: []any{"apple", "orange", "pear"},
			setupReport: func(tc *reportTestCase) *report {
				rpt := NewReport()

				rpt.AddHeaders("h0", "h1", "h2")
				rpt.AddRowData("apple", "orange", "pear", "grape", "dragonfruit")

				return rpt
			},
		},
	}

	for _, scenario := range testCases {
		rs.Run(scenario.name, func() {
			r := scenario.setup()
			rs.Equal(scenario.expects, r.rows[0])
		})
	}
}

func (rs *reportSuite) TestAddRowDataMaxColumnWidths() {
	testCases := []*reportTestCase{
		{
			name:    "should equal data value widths",
			expects: []int{5, 6, 9},
			setupReport: func(tc *reportTestCase) *report {
				rpt := NewReport()

				rpt.AddHeaders("h0", "h1", "h2")
				rpt.AddRowData("apple", "orange", "pineapple")

				return rpt
			},
		},
		{
			name:    "should equal greatest data value widths",
			expects: []int{5, 13, 9},
			setupReport: func(tc *reportTestCase) *report {
				rpt := NewReport()

				rpt.AddHeaders("h0", "h1", "h2")
				rpt.AddRowData("apple", "orange", "pineapple")
				rpt.AddRowData("key", "transactionid", "message")

				return rpt
			},
		},
		{
			name:    "can be overriden",
			expects: []int{20, 13, 9},
			setupReport: func(tc *reportTestCase) *report {
				rpt := NewReport()

				rpt.AddHeaders("h0", "h1", "h2")
				rpt.AddRowData("apple", "orange", "pineapple")
				rpt.AddRowData("key", "transactionid", "message")
				tc.setupErr = rpt.SetColumnMaxWidth("h0", 20)

				return rpt
			},
		},
	}

	for _, scenario := range testCases {
		rs.Run(scenario.name, func() {
			r := scenario.setup()
			rs.Nil(scenario.setupErr)
			expectedMaxWidths := scenario.expects.([]int)
			for i, column := range r.columns {
				rs.Equal(expectedMaxWidths[i], column.maxWidth)
			}
		})
	}
}

func (rs *reportSuite) TestSetColumnMaxWidth() {
	testCases := []*reportTestCase{
		{
			name:    "should set for known header",
			expects: 17,
			setupReport: func(tc *reportTestCase) *report {
				headerName := "h0"
				rpt := NewReport()

				rpt.AddHeaders(headerName)
				tc.setupErr = rpt.SetColumnMaxWidth(
					headerName,
					tc.expects.(int),
				)

				return rpt
			},
		},
		{
			name:      "should return error for unknown header",
			expectErr: true,
			setupReport: func(tc *reportTestCase) *report {
				badHeaderName := "doesNotExist"
				headerName := "h0"
				rpt := NewReport()

				rpt.AddHeaders(headerName)
				tc.setupErr = rpt.SetColumnMaxWidth(badHeaderName, 9)
				tc.expects = fmt.Errorf(ErrUnknownHeader, badHeaderName)

				return rpt
			},
		},
		{
			name:      "should return error for negative column width",
			expectErr: true,
			expects:   ErrNegagiveColumnWidth,
			setupReport: func(tc *reportTestCase) *report {
				headerName := "h0"
				rpt := NewReport()

				rpt.AddHeaders(headerName)
				tc.setupErr = rpt.SetColumnMaxWidth(headerName, -1)

				return rpt
			},
		},
		{
			name:    "should not be overridden if called before AddRowData and column data is less than the max width",
			expects: 1000,
			setupReport: func(tc *reportTestCase) *report {
				headerName := "h0"
				rpt := NewReport()

				rpt.AddHeaders(headerName)
				tc.setupErr = rpt.SetColumnMaxWidth(headerName, tc.expects.(int))
				rpt.AddRowData("value")

				return rpt
			},
		},
		{
			name: "should be overridden if called before AddRowData and column data is greater than the max width",
			setupReport: func(tc *reportTestCase) *report {
				headerName := "h0"
				datapoint := "value"
				rpt := NewReport()

				tc.expects = len(datapoint)

				rpt.AddHeaders(headerName)
				tc.setupErr = rpt.SetColumnMaxWidth(headerName, 3)
				rpt.AddRowData(datapoint)

				return rpt
			},
		},
	}

	for _, scenario := range testCases {
		rs.Run(scenario.name, func() {
			rpt := scenario.setup()

			if !scenario.expectErr {
				rs.Nil(scenario.setupErr)
				rs.Equal(scenario.expects.(int), rpt.columns[0].maxWidth)
			} else {
				rs.Equal(scenario.expects.(error), scenario.setupErr)
			}
		})
	}
}

func (rs *reportSuite) TestWritesToSpecifiedWriter() {
	var reportLines [][]byte = make([][]byte, 1)
	buf := new(bytes.Buffer)
	rpt := NewReport(Writer(buf))
	expectedHeaders := []string{"h0", "h1", "h2"}
	expectedReportLineCount := 3

	// fill report
	rpt.AddHeaders(expectedHeaders...)
	rpt.AddRowData("hello", "test", "report")
	// write report
	rpt.PrintHeader()
	rpt.PrintBody()

	// inspection/testing below
	// rs.T().Logf("%+v", buf.Bytes())
	newlineCounter := 0
	reportLines[0] = make([]byte, 0)
	for bpos, b := range buf.Bytes() {
		if b == '\n' {
			newlineCounter++
			if bpos < buf.Len()-1 {
				reportLines = append(reportLines, make([]byte, 0))
			}
		} else {
			reportLines[newlineCounter] = append(reportLines[newlineCounter], b)
		}
	}

	rs.Equal(expectedReportLineCount, len(reportLines))
	// for lineNumber, line := range reportLines {
	// 	fmt.Printf("line %d (%2d bytes): %s\n", lineNumber, len(line), line)
	// }

	headerWords := regexp.MustCompile(`\w+`)
	headerMatches := headerWords.FindAllString(string(reportLines[0]), -1)
	for i, actualHeader := range headerMatches {
		rs.Equal(expectedHeaders[i], actualHeader)
	}

	headerSpaceGroups := regexp.MustCompile(`[ ]+`)
	spaceGroupMatches := headerSpaceGroups.FindAllString(string(reportLines[0]), -1)
	expctedSpaces := rpt.columns[0].maxWidth + rpt.settings.columnPadding - len(rpt.columns[0].header)
	rs.Equal(expctedSpaces, len(spaceGroupMatches[0]))

	headerBorders := regexp.MustCompile(headerBorder + `+`)
	headerMatches = headerBorders.FindAllString(string(reportLines[1]), -1)
	rs.Equal(len(expectedHeaders), len(headerMatches))
}

/*
line 0 (19 bytes): h0       h1      h2
line 1 (19 bytes): --       --      --
line 2 (23 bytes): hello    test    report
*/
