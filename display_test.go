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
// as we might clutter the test output
func (ds *UserIODisplaySuite) callOutput(content string, padding int, hasNewline bool) []byte {
	// generate some "output" to a temporary file
	tempFile, _ := ioutil.TempFile("", "")
	output(ColorStandard, content, padding, hasNewline, tempFile)
	tempFile.Close()

	// open generated file
	f, openFileErr := os.Open(tempFile.Name())
	defer f.Close()
	if openFileErr != nil {
		ds.Fail("could not open temp file " + tempFile.Name())
	}

	// read generated file; we're gonna work with bytes
	fileContents, readErr := ioutil.ReadAll(f)
	if readErr != nil {
		ds.Fail("could not read temp file " + tempFile.Name())
	}

	return []byte(fileContents)
}

func (suite *UserIODisplaySuite) SetupTest() {
	suite.genericContent = "hello"
}

func TestUserIODisplay(t *testing.T) {
	suite.Run(t, new(UserIODisplaySuite))
}

func (ds *UserIODisplaySuite) TestPadWithSpace() {
	expectedSpaces := 3
	result := padWithSpace(expectedSpaces)

	ds.Equal(expectedSpaces, len(result))
	ds.Equal(expectedSpaces, strings.Count(result, " "))
}

func (ds *UserIODisplaySuite) TestOutputStructure() {
	result := ds.callOutput(ds.genericContent, standardPadding, writeNewline)
	contentStart := strings.Index(string(result), ds.genericContent)
	contentEnd := contentStart + len(ds.genericContent)

	// analyze output's effect
	// -----
	// printed standard padding?
	ds.Equal(padWithSpace(standardPadding), string(result[:2]))
	// has ansii escape code for color?
	ds.Equal(string(ColorStandard), string(result[2:contentStart]))
	// has specified content (string message)?
	ds.Equal(ds.genericContent, string(result[contentStart:contentEnd]))
	// content is followed by escape code for color reset?
	ds.Equal(string(TextReset), string(result[contentEnd:len(result)-1]))
	// we specified output to write a newline, make sure it's there
	ds.Equal(byte('\n'), result[len(result)-1])
}

func (suite *UserIODisplaySuite) TestOutputWithoutNewline() {
	result := suite.callOutput(suite.genericContent, standardPadding, !writeNewline)

	suite.NotEqual(byte('\n'), result[len(result)-1])
}

func (ds *UserIODisplaySuite) TestOutputWithVariousPaddingAmounts() {
	bigPadding := 5
	result1 := ds.callOutput(ds.genericContent, bigPadding, writeNewline)
	result2 := string(ds.callOutput(ds.genericContent, 0, writeNewline))

	ds.Equal(padWithSpace(bigPadding), string(result1[:bigPadding]))
	ds.Equal(false, strings.Contains(result2, " "))
}
