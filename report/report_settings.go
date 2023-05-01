package report

import (
	"fmt"
	"io"
	"os"
)

// settings container for a report
//
// Manages the settings for a report and keep track of settings
// explicitly requested.
type reportSettings struct {
	columnPadding         int
	dataPlaceholder       string
	debugEnabled          bool
	writer                io.Writer
	requestedOptionsState map[string]bool
}

// settings defaults
const (
	defaultColumnPadding   int    = 4
	defaultDataPlaceholder string = "-"
)

// initial state of options that can be requested
var defaulRequestedOptionsState map[string]bool = map[string]bool{
	OptTypeColumnPadding:   false,
	OptTypeDataPlaceholder: false,
	OptTypeDebugEnabled:    false,
	OptTypeWriter:          false,
}

// processes report options and persists to reportSettings
//
// Keeps track of options set or arguments passed. Just like report.NewReport
// the options to process are optional. Part of applying options is keeping
// track of which options were specified so we know which defaults to use.
func (rs *reportSettings) processOptions(o ...reportOption) {
	for _, option := range o {
		rs.requestedOptionsState[fmt.Sprintf("%T", option)] = true
		option.ApplyOption(rs)
	}

	for optionType, requested := range rs.requestedOptionsState {
		if !requested {
			// set the default for that option
			// bools defaulting to false will not be set here
			switch optionType {
			case OptTypeColumnPadding:
				ColumnPadding(defaultColumnPadding).ApplyOption(rs)
			case OptTypeDataPlaceholder:
				Placeholder(defaultDataPlaceholder).ApplyOption(rs)
			case OptTypeWriter:
				Writer(os.Stdout).ApplyOption(rs)
			}
		}
	}
}

// initialize report settings
//
// Will need to call reportSettings.processOptions() to set values and defaults.
func newReportSettings() *reportSettings {
	settings := &reportSettings{
		requestedOptionsState: make(map[string]bool),
	}

	// copy over default requested option state
	for knownOption, requestedState := range defaulRequestedOptionsState {
		settings.requestedOptionsState[knownOption] = requestedState
	}

	return settings
}
