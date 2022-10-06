package userio
import (
  "fmt"
  "os"
  "strings"
  "time"
)

const (
  EmojiWave = "\U0001F44B"
  TextReset = "\x1b[0m"
  standardPadding = 2
  writeNewline = true
)

var EnableTimestamp bool = false

func Highlight(content string) string {
  return fmt.Sprintf("%s %s %s", HighlightYellow, content, TextReset)
}

func ListElement(content string) {
  outputNoTimestamp(
    ColorCyan,
    content,
    5,
    writeNewline,
    os.Stdout,
  )
}

func ListElementFromArray(index int, content string) {
  offsetIndex := index + 1
  outputNoTimestamp(
    ColorCyan,
    fmt.Sprintf("%d) %s", offsetIndex, content),
    standardPadding,
    writeNewline,
    os.Stdout,
  )
}

func ListElementWithLabel(label, content string) {
  outputNoTimestamp(
    ColorCyan,
    fmt.Sprintf("%s) %s", label, content),
    standardPadding,
    writeNewline,
    os.Stdout,
  )
}

func Write(content string) {
  canOutputTimestamp(ColorStandard, content, standardPadding, writeNewline, os.Stdout)
}

func Writef(s string, a ...interface{}) {
  canOutputTimestamp(ColorStandard, fmt.Sprintf(s, a...), standardPadding, false, os.Stdout)
}

func WriteError(content string) {
  canOutputTimestamp(ColorRed, content, standardPadding, writeNewline, os.Stdout)
}

func WriteInfo(content string) {
  canOutputTimestamp(ColorInfo, content, standardPadding, writeNewline, os.Stdout)
}

func WriteInfof(s string, a ...interface{}) {
  canOutputTimestamp(ColorInfo, fmt.Sprintf(s, a...), standardPadding, writeNewline, os.Stdout)
}

func WriteInstruction(content string) {
  outputNoTimestamp(ColorInstruction, content, 4, writeNewline, os.Stdout)
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

func canOutputTimestamp(color, content string, numberOfLeadingSpaces int, terminateLine bool, out *os.File) {
  output(
    color,
    content,
    numberOfLeadingSpaces,
    EnableTimestamp,
    terminateLine,
    out,
  )
}

func padWithSpace(numberOfSpaces int) string {
  var stringToFill strings.Builder

  for i:=0; i<numberOfSpaces; i++ {
    stringToFill.WriteString(" ")
  }

  return stringToFill.String()
}

func output(color, content string, numberOfLeadingSpaces int, includeTimestamp, terminateLine bool, out *os.File) {
  lineEnder := ""
  if strings.Contains(content, TextReset) {
    content = strings.ReplaceAll(content, TextReset, TextReset + color)
  }

  if terminateLine {
    lineEnder = "\n"
  }

  if includeTimestamp {
    fmt.Fprintf(
      out,
      "%s[%s]%s %s%s%s%s%s",
      ColorStandard,
      time.Now().Format("15:04"),
      TextReset,
      padWithSpace(numberOfLeadingSpaces),
      color,
      content,
      TextReset,
      lineEnder,
    )
  } else {
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
}

func outputNoTimestamp(color, content string, numberOfLeadingSpaces int, terminateLine bool, out *os.File) {
  output(
    color,
    content,
    numberOfLeadingSpaces,
    false,
    terminateLine,
    out,
  )
}
