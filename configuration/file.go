package configuration

import (
	"encoding/json"
	"errors"
	"hextechdocs-be/logger"
	"io/ioutil"
	"os"
)

// initFileConfiguration loads in the configuration from the filesystem (or saves the example config if there isn't one)
func initFileConfiguration() (*Configuration, error) {
	if _, err := os.Stat("config.json"); err == nil {
		contents, err := ioutil.ReadFile("config.json")
		if err != nil {
			return nil, err
		}

		var configuration Configuration
		if err := json.Unmarshal(contents, &configuration); err != nil {
			return nil, err
		}

		return &configuration, nil
	}

	logger.HexLogger.Warn("Configuration file doesn't exist but the use of consul is disabled!")
	logger.HexLogger.Warn("Saving a template file and terminating...")

	configContents, _ := json.MarshalIndent(exampleConfig(), "", " ")
	_ = ioutil.WriteFile("config.json", configContents, 0644)

	os.Exit(1)
	return nil, errors.New("unable to find configuration")
}
