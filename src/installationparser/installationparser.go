package insallationparser

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{})

type InstallationSection struct {
	Exists bool
}

func ExtractSection(fileContent string) InstallationSection {
	logger.Info("Extracting installation section...")
	return InstallationSection{
			Exists: hasInstallation(fileContent),
		}
}

func hasInstallation(fileContent string) bool {
	return stringUtils.Contains(fileContent, "Installation*")
}
