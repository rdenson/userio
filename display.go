package userio

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	standardPadding = 2
	writeNewline    = true
)

// decorate value with color
//
// ver2
// Wraps some value with a Color and a "reset" escape code.
func Colorize(c Color, a any) string {
	return fmt.Sprintf("%s%v%s", c, a, TextReset)
}

// highlights a value
//
// ver2
// Nearly the same function as Colorize(), only it adds space around
// the value.
func Highlight(s string) string {
	return fmt.Sprintf("%s %s %s", HighlightYellow, s, TextReset)
}

// output a format specified string with the operands called out
//
// ver2
// Acts like fmt.Printf(), where the array of arguments after the format
// template are painted with the theme colors.
func WriteData(template string, data ...any) {
	formatIdRegex := regexp.MustCompile(`%[-0-9.*+# ]*[a-z]{1}`)
	colorLacedData := make([]any, 0)

	template = formatIdRegex.ReplaceAllString(template, "%s${0}%s")
	itr := 0
	for _, datapoint := range data {
		itr++
		colorLacedData = append(colorLacedData, ColorInfo, datapoint, ColorStandard)
	}

	fmt.Printf(template, colorLacedData...)
}

func ListElement(content string) {
	output(
		ColorCyan,
		content,
		5,
		writeNewline,
		os.Stdout,
	)
}

func ListElementFromArray(index int, content string) {
	offsetIndex := index + 1
	output(
		ColorCyan,
		fmt.Sprintf("%d) %s", offsetIndex, content),
		standardPadding,
		writeNewline,
		os.Stdout,
	)
}

func ListElementWithLabel(label, content string) {
	output(
		ColorCyan,
		fmt.Sprintf("%s) %s", label, content),
		standardPadding,
		writeNewline,
		os.Stdout,
	)
}

func Write(content string) {
	output(ColorStandard, content, standardPadding, writeNewline, os.Stdout)
}

func Writef(s string, a ...interface{}) {
	output(ColorStandard, fmt.Sprintf(s, a...), standardPadding, false, os.Stdout)
}

func WriteError(content string) {
	output(ColorRed, content, standardPadding, writeNewline, os.Stdout)
}

func WriteInfo(content string) {
	output(ColorInfo, content, standardPadding, writeNewline, os.Stdout)
}

func WriteInfof(s string, a ...interface{}) {
	output(ColorInfo, fmt.Sprintf(s, a...), standardPadding, writeNewline, os.Stdout)
}

func WriteInstruction(content string) {
	output(ColorInstruction, content, 4, writeNewline, os.Stdout)
}

func WriteResultListHeader(listCount int) {
	switch listCount {
	case 0:
		WriteInfo("no results could be found")
	case 1:
		WriteInfo("found 1 result:")
	default:
		WriteInfo(fmt.Sprintf("found %d results:", listCount))
	}
}

func padWithSpace(numberOfSpaces int) string {
	var stringToFill strings.Builder

	for i := 0; i < numberOfSpaces; i++ {
		stringToFill.WriteString(" ")
	}

	return stringToFill.String()
}

func output(c Color, content string, numberOfLeadingSpaces int, terminateLine bool, out *os.File) {
	lineEnder := ""
	if strings.Contains(content, string(TextReset)) {
		content = strings.ReplaceAll(content, string(TextReset), fmt.Sprintf("%s%s", TextReset, c))
	}

	if terminateLine {
		lineEnder = "\n"
	}

	fmt.Fprintf(
		out,
		"%s%s%s%s%s",
		padWithSpace(numberOfLeadingSpaces),
		c,
		content,
		TextReset,
		lineEnder,
	)
}
