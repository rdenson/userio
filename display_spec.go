package userio

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

// Specification for printing interpolated variables
type DisplaySpec struct {
	// specification destination (dest) needs to implement io.Writer
	Dest io.Writer
	// leading spaces on the line to be written
	PadSize int
	// add a newline character when DisplaySpec.Write() is called
	TerminateLine bool
	// color of template or content to be written
	TextColor Color
	// color of template values
	ValuesColor          Color
	ValuesAreHighlighted bool
}

const standardPadSize int = 2

var formatIdRegex *regexp.Regexp = regexp.MustCompile(`%[-0-9.*+# ]*[a-z]{1}`)
var standardDisplaySpec *DisplaySpec = &DisplaySpec{
	Dest:        os.Stdout,
	TextColor:   ColorStandard,
	ValuesColor: ColorInfo,
	PadSize:     standardPadSize,
}

// Write executes the specification against a template and values (if any exist).
//
// Args mirror fmt.Printf(), execpt the template will not be the
// final version to be printed.
func (spec *DisplaySpec) Write(template string, v ...any) {
	var format string
	var refinedTemplate string
	var vars []any

	if spec.ValuesColor != "" {
		format, vars = spec.colorize(template, v...)
	} else {
		format = template
		vars = v
	}

	switch {
	case spec.PadSize > 0 && spec.TerminateLine:
		refinedTemplate = fmt.Sprintf(
			"%*s%s%s%s\n",
			spec.PadSize, " ",
			spec.TextColor,
			format,
			TextReset,
		)
	case spec.PadSize > 0 && !spec.TerminateLine:
		refinedTemplate = fmt.Sprintf(
			"%*s%s%s%s", spec.PadSize, " ",
			spec.TextColor,
			format,
			TextReset,
		)
	case spec.TerminateLine:
		refinedTemplate = fmt.Sprintf(
			"%s%s%s\n",
			spec.TextColor,
			format,
			TextReset,
		)
	default:
		refinedTemplate = fmt.Sprintf(
			"%s%s%s",
			spec.TextColor,
			format,
			TextReset,
		)
	}

	if spec.Dest == nil {
		spec.Dest = os.Stdout
	}

	fmt.Fprintf(
		spec.Dest,
		refinedTemplate,
		vars...,
	)
}

// Colorize template operands
//
// template placeholders are located with a regular expression
func (spec *DisplaySpec) colorize(template string, v ...any) (expandedTemplate string, colorizedData []any) {
	highlightedVarsTemplate := "%s ${0} %s%s"
	standardVarsTemplate := "%s${0}%s"

	if spec.ValuesAreHighlighted {
		expandedTemplate = formatIdRegex.ReplaceAllString(template, highlightedVarsTemplate)
		for _, value := range v {
			colorizedData = append(colorizedData, spec.ValuesColor, value, TextReset, spec.TextColor)
		}
	} else {
		expandedTemplate = formatIdRegex.ReplaceAllString(template, standardVarsTemplate)
		for _, value := range v {
			colorizedData = append(colorizedData, spec.ValuesColor, value, spec.TextColor)
		}
	}

	return expandedTemplate, colorizedData
}

func Highlight(template string, v ...any) {
	spec := &DisplaySpec{
		Dest:                 standardDisplaySpec.Dest,
		TextColor:            standardDisplaySpec.TextColor,
		ValuesColor:          HighlightYellow,
		PadSize:              standardDisplaySpec.PadSize,
		TerminateLine:        true,
		ValuesAreHighlighted: true,
	}

	spec.Write(template, v...)
}

func Write(template string, v ...any) {
	standardDisplaySpec.Write(template, v...)
}

func Writeln(template string, v ...any) {
	spec := &DisplaySpec{
		Dest:          standardDisplaySpec.Dest,
		TextColor:     standardDisplaySpec.TextColor,
		ValuesColor:   standardDisplaySpec.ValuesColor,
		PadSize:       standardDisplaySpec.PadSize,
		TerminateLine: true,
	}

	spec.Write(template, v...)
}

func WriteError(e error) {
	spec := &DisplaySpec{
		Dest:          standardDisplaySpec.Dest,
		TextColor:     ColorRed,
		PadSize:       standardDisplaySpec.PadSize,
		TerminateLine: true,
	}

	spec.Write("%s", e.Error())
}
