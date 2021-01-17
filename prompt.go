package userio
import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

// optimistic boolean parser; any positive response yields true o/w false
func PromptForBool(content string) bool {
  var (
    formattedUserInput bool
    validTrueResponses string = "true,y,yes,Y,Yes,1"
  )

  userInput := promptToken(content, os.Stdin, os.Stdout)
  if len(userInput) > 0 && strings.Contains(validTrueResponses, userInput) {
    formattedUserInput = true
  }

  return formattedUserInput
}

// choices here correspond to a numerical list (positive integers)
func PromptForChoice(content string) int {
  choice := -1
  continueToPrompt := true

  for continueToPrompt {
    parsedInput, conversionErr := strconv.Atoi(promptToken(content, os.Stdin, os.Stdout))
    if conversionErr != nil {
      WriteError("input not recognized, please enter a number")
    } else {
      choice = parsedInput - 1
      continueToPrompt = false
    }
  }

  return choice
}

// simple user prompt; looks for a string containing no spaces
func PromptForString(content string) string {
  return promptToken(content, os.Stdin, os.Stdout)
}

func promptLine(content string, in, out *os.File) (string, error) {
  const Newline = '\n'
  var inputReader *bufio.Reader = bufio.NewReader(in)

  output(ColorStandard, content + " ", standardPadding, !writeNewline, out)
  inputRead, inputReadErr := inputReader.ReadString(Newline)
  if inputReadErr != nil {
    return "", inputReadErr
  }

  return inputRead, nil
}

func promptToken(content string, in, out *os.File) string {
  var userInput string
  output(ColorStandard, content + " ", standardPadding, !writeNewline, out)
  fmt.Fscanln(in, &userInput)

  return userInput
}
