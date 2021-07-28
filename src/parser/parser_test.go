package parser

import (
	"testing"
)

func TestShouldDiscoverInstallationSection(TEST *testing.T) {
	result := Parse("../../test/test.md")

	if !result.InstallationSection.Exists {
		TEST.Fail()
	}

	if !result.UsageSection.Exists {
		TEST.Fail()
	}
}
