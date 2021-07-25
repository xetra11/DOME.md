package usageparser

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var logger = logrus.WithFields(logrus.Fields{})

type UsageStep struct {
	name string
}

func ExtractUsageStep(fileContent string) UsageSteps {


}
