package userio
import (
  "io"
  "io/ioutil"
  "os"
  "testing"

  "github.com/stretchr/testify/suite"
)

type UserIOPromptSuite struct {
  suite.Suite
}

func (uiops *UserIOPromptSuite) mockUserWrite(content string) (*os.File, error) {
  tempFile, tempFileErr := ioutil.TempFile("", "")
  if tempFileErr != nil {
    return nil, tempFileErr
  }

  _, writeErr := io.WriteString(tempFile, content)
  if writeErr != nil {
    return nil, writeErr
  }

  _, seekErr := tempFile.Seek(0, io.SeekStart)
  if seekErr != nil {
    return nil, seekErr
  }

  return tempFile, nil
}

func (uiods *UserIOPromptSuite) outToNull() *os.File {
  return os.NewFile(0, os.DevNull)
}

func TestUserIOPrompt(t *testing.T) {
  suite.Run(t, new(UserIOPromptSuite))
}

func (suite *UserIOPromptSuite) TestPromptLine() {
  expectedLine := "does it work?\n"
  nowhere := suite.outToNull()
  in, _ := suite.mockUserWrite(expectedLine)
  defer in.Close()
  promptResult, promptErr := promptLine("gimme some words!", in, nowhere)
  nowhere.Close()

  suite.NoError(nil, promptErr)
  suite.Equal(expectedLine, promptResult)
}

func (suite *UserIOPromptSuite) TestPromptLineFailure() {
  expectedLine := ""
  nowhere := suite.outToNull()
  defer nowhere.Close()
  in, _ := suite.mockUserWrite("do we get a failure?")
  defer in.Close()
  promptResult, promptErr := promptLine("gimme some words!", in, nowhere)

  suite.Equal(io.EOF, promptErr)
  suite.Equal(expectedLine, promptResult)
}

func (suite *UserIOPromptSuite) TestPromptToken() {
  expectedToken := "ok"
  tempFile, _ := ioutil.TempFile("", "")
  defer tempFile.Close()
  in, _ := suite.mockUserWrite(expectedToken + "\n")
  defer in.Close()
  promptResult := promptToken("gimme a string!", in, tempFile)


  suite.Equal(expectedToken, promptResult)
}
