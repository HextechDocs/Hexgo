package web

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"hextechdocs-be/flags"
	"hextechdocs-be/logger"
	"hextechdocs-be/web/handler"
	"log"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	logger.HexLogger.WithFields(logrus.Fields{
		"path": r.URL.Path,
		"method": r.Method,
	}).Trace("Incoming HTTP request")

	if r.URL.Path == "/graphql" {
		switch method := r.Method; method {
		case "POST":
			handler.HandleGraphQL(w, r)
		case "OPTIONS":
			handler.WritePreflight(w)
		default:
			handler.WriteMethodNotAllowed(w)
		}
	} else if r.URL.Path == "/health" && r.Method == "GET" {
		handler.HandleHealthcheck(w)
	} else {
		handler.WriteNotFound(w)
	}
}

func ServeWeb() {
	http.HandleFunc("/", handleRequest)

	addr := fmt.Sprintf(":%d", flags.ServerPort)
	if os.Getenv("NOMAD_HOST_PORT_http") != "" {
		addr = fmt.Sprintf(":%s", os.Getenv("NOMAD_HOST_PORT_http"))
	}

	log.Fatal(http.ListenAndServe(addr, nil))
}
