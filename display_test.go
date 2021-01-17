package userio
import (
  "io/ioutil"
  "os"
  "strings"
  "testing"

  "github.com/stretchr/testify/suite"
)

type UserIODisplaySuite struct {
  genericContent string
  suite.Suite
}

// need some setup to call output; normally we use os.Stdout but can't do that here
func (uiods *UserIODisplaySuite) callOutput(content string, padding int, hasNewline bool) []byte {
  // generate some "output" to a temporary file
  tempFile, _ := ioutil.TempFile("", "")
  output(ColorStandard, content, padding, hasNewline, tempFile)
  tempFile.Close()

  // open generated file
  f, openFileErr := os.Open(tempFile.Name())
  defer f.Close()
  if openFileErr != nil {
    uiods.Fail("could not open temp file " + tempFile.Name())
  }

  // read generated file; we're gonna work with bytes
  fileContents, readErr := ioutil.ReadAll(f)
  if readErr != nil {
    uiods.Fail("could not read temp file " + tempFile.Name())
  }

  return []byte(fileContents)
}

func (suite *UserIODisplaySuite) SetupTest() {
  suite.genericContent = "hello"
}

func TestUserIODisplay(t *testing.T) {
  suite.Run(t, new(UserIODisplaySuite))
}

func (suite *UserIODisplaySuite) TestPadWithSpace() {
  expectedSpaces := 3
  result := padWithSpace(expectedSpaces)

  suite.Equal(expectedSpaces, len(result))
  suite.Equal(expectedSpaces, strings.Count(result, " "))
}

func (suite *UserIODisplaySuite) TestOutputStructure() {
  result := suite.callOutput(suite.genericContent, standardPadding, writeNewline)
  contentStart := strings.Index(string(result), suite.genericContent)
  contentEnd := contentStart+len(suite.genericContent)

  // analyze output's effect
  // -----
  // printed standard padding?
  suite.Equal(padWithSpace(standardPadding), string(result[:2]))
  // has ansii escape code for color?
  suite.Equal(ColorStandard, string(result[2:contentStart]))
  // has specified content (string message)?
  suite.Equal(suite.genericContent, string(result[contentStart:contentEnd]))
  // content is followed by escape code for color reset?
  suite.Equal(ColorReset, string(result[contentEnd:len(result)-1]))
  // we specified output to write a newline, make sure it's there
  suite.Equal(byte('\n'), result[len(result)-1])
}

func (suite *UserIODisplaySuite) TestOutputWithoutNewline() {
  result := suite.callOutput(suite.genericContent, standardPadding, !writeNewline)

  suite.NotEqual(byte('\n'), result[len(result)-1])
}

func (suite *UserIODisplaySuite) TestOutputWithVariousPaddingAmounts() {
  bigPadding := 5
  result1 := suite.callOutput(suite.genericContent, bigPadding, writeNewline)
  result2 := string(suite.callOutput(suite.genericContent, 0, writeNewline))

  suite.Equal(padWithSpace(bigPadding), string(result1[:bigPadding]))
  suite.Equal(false, strings.Contains(result2, " "))
}
