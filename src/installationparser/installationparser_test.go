package insallationparser

import (
	"io/ioutil"
	"testing"
)

func TestShouldDiscoverInstallationSection(TEST *testing.T) {
	fileData, err := ioutil.ReadFile("../../test/test.md")
	check(err)

	fileContent := string(fileData)

	actual := ExtractSection(fileContent)

	if !actual.Exists {
		TEST.Fail()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
