package report

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type processOptionsTestCase struct {
	expect     any
	name       string
	option     reportOption
	optionType string
}

type ReportSettingsSuite struct {
	suite.Suite
}

func TestReportSettings(t *testing.T) {
	suite.Run(t, new(ReportSettingsSuite))
}

func (rs *ReportSettingsSuite) TestNewReportSettingsHasRequestedOptionsStateSetToNotRequested() {
	s := newReportSettings()

	rs.Greater(len(s.requestedOptionsState), 0)
	for _, requested := range s.requestedOptionsState {
		rs.False(requested)
	}
}

func (rs *ReportSettingsSuite) TestProcessOptions() {
	testCases := []processOptionsTestCase{
		{
			name:       "sets column padding option",
			expect:     99,
			option:     ColumnPadding(99),
			optionType: OptTypeColumnPadding,
		},
		{
			name:       "sets placeholder option",
			expect:     "∆",
			option:     Placeholder("∆"),
			optionType: OptTypeDataPlaceholder,
		},
		{
			name:       "sets writer option",
			expect:     os.Stderr,
			option:     Writer(os.Stderr),
			optionType: OptTypeWriter,
		},
	}

	for _, scenario := range testCases {
		rs.Run(scenario.name, func() {
			s := newReportSettings()

			s.processOptions(scenario.option)
			rs.True(s.requestedOptionsState[scenario.optionType])
			switch scenario.optionType {
			case OptTypeColumnPadding:
				rs.Equal(scenario.expect, s.columnPadding)
			case OptTypeDataPlaceholder:
				rs.Equal(scenario.expect, s.dataPlaceholder)
			case OptTypeWriter:
				rs.Equal(scenario.expect, s.writer)
			}
		})
	}
}
