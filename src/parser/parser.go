package parser

import (
	"fmt"
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

	data, err := ioutil.ReadFile(path)
	check(err)

	fmt.Print(data)

	return ParseResult{
		InstallationSection: InstallationSection{
			exists: false,
		},
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
