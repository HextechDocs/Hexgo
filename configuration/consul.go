package configuration

import (
	"encoding/json"
	"errors"
	consul "github.com/hashicorp/consul/api"
	"hextechdocs-be/logger"
	"os"
)

// initConfiguration loads in the configuration from Consul (or saves the example config if there isn't one)
func initConfiguration() (*Configuration, error) {
	consulClient, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

	pair, _, err := consulClient.KV().Get(os.Getenv("CONSUL_KV_PATH"), nil)
	if err != nil {
		return nil, err
	}

	if pair == nil {
		logger.HexLogger.Warn("Configuration doesn't exist on the consul instance!")
		logger.HexLogger.Warn("Saving a template and terminating...")

		data, _ := json.Marshal(exampleConfig())
		_, _ = consulClient.KV().Put(&consul.KVPair{
			Key:   os.Getenv("CONSUL_KV_PATH"),
			Value: data,
		}, nil)

		os.Exit(1)
		return nil, errors.New("unable to find configuration")
	}

	var config Configuration
	if err := json.Unmarshal(pair.Value, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
