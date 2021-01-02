/*
Package hooks houses the different configuration hooks we have for logrus.

Currently it only has an elasticsearch hook to allow saving the logs in containerized environments.
*/
package hooks

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"hextechdocs-be/configuration"
	"hextechdocs-be/logger"
	"os"
)

// AddElasticHook adds an elastic logger hook to the system so we can do centralized logging
func AddElasticHook(config *configuration.Configuration) {
	var client *elastic.Client
	var err error
	if config.ShouldUseElasticAuth {
		client, err = elastic.NewClient(elastic.SetURL(config.ElasticHosts...), elastic.SetBasicAuth(config.ElasticUsername, config.ElasticPassword))
	} else {
		client, err = elastic.NewClient(elastic.SetURL(config.ElasticHosts...))
	}

	if err != nil {
		logger.HexLogger.Panic(err)
	}

	hook, err := elogrus.NewAsyncElasticHook(client, os.Getenv("ELASTIC_HOSTNAME"), logrus.DebugLevel, config.ElasticIndex)
	if err != nil {
		logger.HexLogger.Panic(err)
	}
	logger.HexLogger.Hooks.Add(hook)
}
