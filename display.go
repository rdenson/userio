package userio
import (
  "fmt"
  "os"
  "strings"
)

const (
  EmojiWave = "\U0001F44B"
  TextReset = "\x1b[0m"
  standardPadding = 2
  writeNewline = true
)

func Highlight(content string) string {
  return fmt.Sprintf("%s %s %s", HighlightYellow, content, TextReset)
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

  for i:=0; i<numberOfSpaces; i++ {
    stringToFill.WriteString(" ")
  }

  return stringToFill.String()
}

func output(color, content string, numberOfLeadingSpaces int, terminateLine bool, out *os.File) {
  lineEnder := ""
  if strings.Contains(content, TextReset) {
    content = strings.ReplaceAll(content, TextReset, TextReset + color)
  }

  if terminateLine {
    lineEnder = "\n"
  }

  fmt.Fprintf(
    out,
    "%s%s%s%s%s",
    padWithSpace(numberOfLeadingSpaces),
    color,
    content,
    TextReset,
    lineEnder,
  )
}
