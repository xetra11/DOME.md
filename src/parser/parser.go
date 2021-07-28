package parser

import (
	"github.com/sirupsen/logrus"
	"github.com/xetra11/DOME.md/src/installationparser"
	"github.com/xetra11/DOME.md/src/usageparser"
	"io/ioutil"
)

var logger = logrus.WithFields(logrus.Fields{})

type ParseResult struct {
	InstallationSection insallationparser.InstallationSection
	UsageSection        usageparser.UsageSection
}

func Parse(path string) ParseResult {
	logger.Info("Parsing README.md...")

	fileData, err := ioutil.ReadFile(path)
	check(err)

	fileContent := string(fileData)

	installationSection := insallationparser.ExtractSection(fileContent)
	usageSection := usageparser.ExtractSection(fileContent)

	return ParseResult{
		InstallationSection: installationSection,
		UsageSection: usageSection,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
