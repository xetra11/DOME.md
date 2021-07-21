package parser

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var logger = logrus.WithFields(logrus.Fields{})

type ParseResult struct {
	InstallationSection InstallationSection
}

type InstallationSection struct {
	exists bool
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
	}
}

func hasInstallation(fileContent string) bool {
	return stringUtils.Contains(fileContent, "Installation*")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
