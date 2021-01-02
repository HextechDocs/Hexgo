package main

import (
	"github.com/sirupsen/logrus"
	"hextechdocs-be/configuration"
	"hextechdocs-be/flags"
	"hextechdocs-be/importer"
	"hextechdocs-be/logger"
	"hextechdocs-be/logger/hooks"
	"hextechdocs-be/model"
	"hextechdocs-be/web"
)

func main() {
	logger.InitLogger()
	logger.HexLogger.Info("HextechDocs Backend")
	flags.InitFlags()

	config, err := configuration.InitializeConfiguration()
	if err != nil {
		logger.HexLogger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Unable to fetch configuration! The process will terminate!")
		return
	}

	if config.ShouldLogToElastic {
		hooks.AddElasticHook(config)
	}

	err = model.Init(config)
	if err != nil {
		logger.HexLogger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Unable to connect to the database! The process will terminate!")
		return
	}

	if flags.Importer {
		logger.HexLogger.Info("Launching in importer mode")
		importer.ParseFolders()
	} else {
		logger.HexLogger.Info("Launching in web mode")
		web.ServeWeb()
	}
}
