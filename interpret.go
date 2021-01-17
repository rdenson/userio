package userio
import (
  "fmt"
  "os"
  "strings"
)


const (
  ExpectedLongFormat = "[verb] [subject words] [preposition] [noun]..."
  ExpectedShortFormat = "[verb] [subject words]..."
  InterpretationErr = "not sure I understand"
)

//  find id-cloud by email in demo with scott.

//  list k8s versions registered to demo
//  list k8s versions by tenant

/*
  what are we talking about? enter: structured question
  main descriptor starter (MDS) [verb] - describes partially what the user wants
  subject [subject words] - the noun or nouns that follow the MDS

    we know we're done with the subject when we come across a preposition or another verb
*/

func InitiateRequest(withQuestionToUser string) *UserRequest {
  const MainVerbLoc = 0

  userRequest := getUserRequest(withQuestionToUser)
  prepositionFound, prepositionLoc := FindPreposistion(userRequest.Words)
  if prepositionFound {
    userRequest.Subject = strings.Join(userRequest.Words[MainVerbLoc + 1:prepositionLoc], " ")
    userRequest.ActionPhrase = fmt.Sprintf("%s %s", userRequest.Words[MainVerbLoc], strings.Join(userRequest.Words[prepositionLoc:prepositionLoc+2], " "))
    userRequest.RawArguments = userRequest.Words[prepositionLoc+2:len(userRequest.Words)]
  } else {
    userRequest.Subject = strings.Join(userRequest.Words[MainVerbLoc + 1:len(userRequest.Words)], " ")
    userRequest.ActionPhrase = userRequest.Words[MainVerbLoc]
  }

  userRequest.Words = []string{}

  return userRequest
}

func SubsequentRequest(withQuestionToUser string) *UserRequest {
  return getUserRequest(withQuestionToUser)
}

func ShowExpectations() {
  WriteInstruction(fmt.Sprintf(
    "hey there %s, just to let you know... I expect your requests to be in the following formats:",
    EmojiWave,
  ))
  WriteInstruction(fmt.Sprintf("  - short form %s", HighlightInstruction(ExpectedShortFormat)))
  WriteInstruction(fmt.Sprintf("  - long form %s", HighlightInstruction(ExpectedLongFormat)))
  fmt.Println()
}

func FindKnownVerbs(words []string) (found bool, index int) {
  const VerbList = "registered"

  currentPosition := 0
  for !found && currentPosition < len(words) {
    if strings.Contains(VerbList, words[currentPosition]) {
      found = true
      index = currentPosition
    } else {
      currentPosition++
    }
  }

  return
}

func FindPreposistion(words []string) (found bool, index int) {
  const PrepositionList = "by,for,from,in,to,with"

  currentPosition := 0
  for !found && currentPosition < len(words) {
    if strings.Contains(PrepositionList, words[currentPosition]) {
      found = true
      index = currentPosition
    } else {
      currentPosition++
    }
  }

  return
}

func getUserRequest(promptToUser string) *UserRequest {
  var promptErr error
  var userReq *UserRequest = new(UserRequest)
  secondaryPrompt := "try it again:"

  // get user response to prompt
  userReq.Raw, promptErr = promptLine(promptToUser, os.Stdin, os.Stdout)
  if promptErr != nil {
    WriteError(fmt.Sprintf("something went wrong... %s", promptErr))
    // attempt retry on failure
    getUserRequest(secondaryPrompt)
  }

  // massage
  userReq.Raw = strings.TrimSpace(strings.Replace(userReq.Raw, "\n", "", 1))
  userReq.Words = strings.Split(userReq.Raw, " ")
  if !userReq.requestIsAcceptable() {
    Write(InterpretationErr)
    getUserRequest(secondaryPrompt)
  }

  return userReq
}
