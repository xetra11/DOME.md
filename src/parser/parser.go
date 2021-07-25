package parser

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"src/usageparser"
)

var logger = logrus.WithFields(logrus.Fields{})

type ParseResult struct {
	InstallationSection InstallationSection
	UsageSection        UsageSection
}

type InstallationSection struct {
	exists bool
}

type UsageSection struct {
	exists bool
	steps []UsageStep
}

func Parse(path string) ParseResult {
	logger.Info("Parsing README.md...")

	fileData, err := ioutil.ReadFile(path)
	check(err)

	fileContent := string(fileData)

	return ParseResult{
		InstallationSection: InstallationSection{
			exists: hasInstallation(fileContent),
		},
		UsageSection: UsageSection{
			exists: hasUsage(fileContent),
		},
	}
}


func hasInstallation(fileContent string) bool {
	return stringUtils.Contains(fileContent, "Installation*")
}

func hasUsage(fileContent string) bool {
	return stringUtils.Contains(fileContent, "Usage*")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
