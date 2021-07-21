package parser

import (
	"testing"
)

func TestShouldDiscoverInstallationSection(TEST *testing.T) {
	result := Parse("../../test/test.md")

	if !result.InstallationSection.exists {
		TEST.Fail()
	}
}

func TestShouldDiscoverUsageSection(TEST *testing.T) {
	result := Parse("../../test/test.md")

	if !result.UsageSection.exists {
		TEST.Fail()
	}
}
