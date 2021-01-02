package configuration

import (
	"github.com/sirupsen/logrus"
	"hextechdocs-be/logger"
	"os"
)

// InitializeConfiguration calls the appropriate internal function based on the presence of an environment variable
func InitializeConfiguration() (*Configuration, error) {
	shouldUseAnuEco := os.Getenv("ANU_DISABLE") == ""
	logger.HexLogger.WithFields(logrus.Fields{"shouldUseAnuEco": shouldUseAnuEco}).Trace("shouldUseAnuEco: ")

	if !shouldUseAnuEco {
		return initFileConfiguration()
	} else {
		return initConfiguration()
	}
}
