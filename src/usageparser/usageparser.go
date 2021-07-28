package usageparser

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{})

type UsageSection struct {
	Exists bool
	Steps  []UsageStep
}

type UsageStep struct {
	Name string
}

func ExtractStep(fileContent string) UsageStep {
	logger.Info("test")
	return UsageStep{}
}

func ExtractSection(fileContent string) UsageSection {
	logger.Info("test")
	return UsageSection{
		Exists: hasUsage(fileContent),
	}
}

func hasUsage(fileContent string) bool {
	return stringUtils.Contains(fileContent, "Usage*")
}
