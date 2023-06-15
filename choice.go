package userio

import (
	"fmt"
	"strconv"
)

type (
	choice struct {
		Display string
		Id      string
		Payload interface{}
	}
	choiceSession struct {
		choices                  map[string]choice
		displayChoicesOnAdd      bool
		useOffsetsForNumericalId bool
	}
)

func (cs *choiceSession) AddChoice(identifier, content string, value interface{}) {
	cs.choices[identifier] = choice{
		Display: content,
		Id:      identifier,
		Payload: value,
	}

	if cs.displayChoicesOnAdd {
		// ListElementWithLabel(identifier, content)
		fmt.Println("placeholder")
	}
}

func (cs *choiceSession) AddChoiceFromArray(index int, content string, value interface{}) {
	var identifier string

	if cs.useOffsetsForNumericalId {
		identifier = strconv.Itoa(index + 1)
	} else {
		identifier = strconv.Itoa(index)
		// call to ListElementFromArray() below will manipulate the index
		// we'll account for that here
		index--
	}

	cs.choices[identifier] = choice{
		Display: content,
		Id:      identifier,
		Payload: value,
	}

	if cs.displayChoicesOnAdd {
		// ListElementFromArray(index, content)
		fmt.Println("placeholder")
	}
}

func (cs *choiceSession) SelectChoiceWithPrompt(prompt string) (choice, bool) {
	promptContent := "make a selection"
	if len(prompt) != 0 {
		promptContent = prompt
	}

	userChoice := PromptForString(promptContent)
	selection, isAvailable := cs.choices[userChoice]

	return selection, isAvailable
}

func (cs *choiceSession) SetDisplayChoicesOnAdd(b bool) {
	cs.displayChoicesOnAdd = b
}

func (cs *choiceSession) SetUseOffsetsForNumericalId(b bool) {
	cs.useOffsetsForNumericalId = b
}

func StartRecordingChoices() *choiceSession {
	session := new(choiceSession)

	session.choices = make(map[string]choice)
	session.displayChoicesOnAdd = true
	session.useOffsetsForNumericalId = true

	return session
}
