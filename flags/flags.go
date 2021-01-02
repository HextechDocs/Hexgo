// Package flags is a simple utility package for every command line argument the software supports
package flags

import (
	"flag"
	"github.com/sirupsen/logrus"
	"hextechdocs-be/logger"
)

// Importer controls whether or not the software should launch in web or importer mode
var Importer 		bool
// ServerPort overrides the default port of 8080 (can be overwritten by the NOMAD_HOST_ADDR_http environment variable)
var ServerPort 		int

// InitFlags parses the command line arguments and fills in the data we might need
func InitFlags() {
	flag.BoolVar(&Importer, "importer", false, "Should the software launch in importer mode")
	flag.IntVar(&ServerPort, "port", 8080, "The port the webserver should bind to")
	flag.Parse()

	logger.HexLogger.WithFields(logrus.Fields{
		"importer": Importer,
		"serverPort": ServerPort,
	}).Trace("Flags have been parsed")
}
