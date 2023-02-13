package userio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Optimistic boolean parser. A positive response from stdin yields true o/w false.
func PromptForBool(content string) bool {
	var (
		formattedUserInput bool
		validTrueResponses map[string]int = map[string]int{
			"true": 1,
			"y":    1,
			"yes":  1,
			"1":    1,
		}
	)

	userInput := strings.ToLower(promptToken(content, os.Stdin, os.Stdout))
	if _, found := validTrueResponses[userInput]; len(userInput) > 0 && found {
		formattedUserInput = true
	}

	return formattedUserInput
}

// Choice parser. Choices here correspond to a numerical list (positive integers).
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

// Simple input parser.
//
// Simple because we only expect a single string from stdin.
func PromptForString(content string) string {
	return promptToken(content, os.Stdin, os.Stdout)
}

func promptLine(content string, in, out *os.File) (string, error) {
	const Newline = '\n'
	var inputReader *bufio.Reader = bufio.NewReader(in)

	output(ColorStandard, content+" ", standardPadding, !writeNewline, out)
	inputRead, inputReadErr := inputReader.ReadString(Newline)
	if inputReadErr != nil {
		return "", inputReadErr
	}

	return inputRead, nil
}

// Scan a file for a token (string).
//
// Looks for a single token, whitespace indicates another token. If the file contains
// more than one string, this function will return the first string. If the file
// has no tokens (only EOF), an empty string is returned.
func promptToken(content string, in, out *os.File) string {
	var userInput string
	output(ColorStandard, content+" ", standardPadding, !writeNewline, out)
	tokensScanned, scanErr := fmt.Fscanln(in, &userInput)

	// if we've received an error here, it's most likely because "in" did not process
	// the exact amount of values we anticipated
	if scanErr != nil {
		// we may have more values than arguments, flush "in"
		if tokensScanned != 0 {
			bufio.NewReader(in).ReadString('\n')
		}
	}

	return userInput
}
