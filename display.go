package userio
import (
  "fmt"
  "os"
  "strings"
)

const (
  // red
  ColorError = "\u001b[31m"

  // yellow
  ColorInfo = "\u001b[33m"
  // yellow background with black text
  ColorHighlight = "\u001b[43m\u001b[30m"

  // magenta
  ColorInstruction = "\u001b[35m"
  // magenta background with white text
  ColorInstructionHighlight = "\u001b[45m\u001b[37m"

  // cyan
  ColorList = "\u001b[36m"

  // green
  ColorStandard = "\u001b[32m"

  ColorReset = "\u001b[0m"
  EmojiWave = "\U0001F44B"
  standardPadding = 2
  writeNewline = true
)

func Highlight(content string) string {
  return fmt.Sprintf("%s %s %s", ColorHighlight, content, ColorReset)
}

func HighlightInstruction(content string) string {
  return fmt.Sprintf("%s %s %s", ColorInstructionHighlight, content, ColorReset)
}

func ListElement(content string) {
  output(
    ColorList,
    content,
    5,
    writeNewline,
    os.Stdout,
  )
}

func NumberedListElement(index int, content string) {
  offsetIndex := index + 1
  output(
    ColorList,
    fmt.Sprintf("%d) %s", offsetIndex, content),
    standardPadding,
    writeNewline,
    os.Stdout,
  )
}

func Write(content string) {
  output(ColorStandard, content, standardPadding, writeNewline, os.Stdout)
}

func WriteError(content string) {
  output(ColorError, content, standardPadding, writeNewline, os.Stdout)
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
  if strings.Contains(content, ColorReset) {
    content = strings.ReplaceAll(content, ColorReset, ColorReset + color)
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
    ColorReset,
    lineEnder,
  )
}
