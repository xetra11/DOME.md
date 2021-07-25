package usageparser

import (
	"testing"
)

func TestShouldDiscoverUsageSectionSteps(TEST *testing.T) {
	result := Parse("../../test/test.md")

	if len(result.UsageSection.steps) <= 0 {
		TEST.Fail()
	}
}

func TestShouldDiscoverUsageSection(TEST *testing.T) {
	result := Parse("../../test/test.md")

	if !result.UsageSection.exists {
		TEST.Fail()
	}
}
