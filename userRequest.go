package userio
import (
  "strings"
)

type UserRequest struct {
  Cancelled bool
  Subject string
  ActionPhrase string
  Raw string
  RawArguments []string
  Resolved bool
  Words []string
}

var exitKeywords []string = []string{
  "cancel",
  "exit",
  "nevermind",
  "nvm",
}

func (ur *UserRequest) requestIsAcceptable() bool {
  if len(ur.Raw) == 0 || phraseContains(ur.Raw, exitKeywords) {
    // no response from user or exit keyword spotted
    ur.Cancelled = true
    return true
  } else if len(ur.Words) < 2 {
    // not enough words for us to figure it out
    return false
  }

  return true
}

func phraseContains(phrase string, keywords []string) (found bool) {
  idx := 0
  for !found && idx < len(keywords) {
    found = strings.Contains(phrase, keywords[idx])
    idx++
  }

  return
}
