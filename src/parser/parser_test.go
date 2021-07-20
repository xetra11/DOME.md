package parser

import (
	"testing"
)

func TestShouldDiscoverInstallationHeader(TEST *testing.T) {
	result := Parse("../../test/test.md")

	if !result.InstallationSection.exists {
		TEST.Fail()
	}

}
