package cli

import (
	"io"
	"os"
	"strconv"

	"github.com/SovereignEdgeEU-COGNIT/ai-orchestrator-env/pkg/build"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CheckError(err error) {
	if err != nil {
		log.WithFields(log.Fields{"BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error(err.Error())
		os.Exit(-1)
	}
}

func parseDBEnv() {
	DBHostEnv := os.Getenv("ENVSERVER_DB_HOST")
	if DBHostEnv != "" {
		DBHost = DBHostEnv
	}

	var err error
	DBPortEnvStr := os.Getenv("ENVSERVER__DB_PORT")
	if DBPortEnvStr != "" {
		DBPort, err = strconv.Atoi(DBPortEnvStr)
		CheckError(err)
	}

	if DBUser == "" {
		DBUser = os.Getenv("ENVSERVER_DB_USER")
	}

	if DBPassword == "" {
		DBPassword = os.Getenv("ENVSERVER_DB_PASSWORD")
	}

	initDBStr := os.Getenv("ENVSERVER_INITDB")
	if initDBStr == "true" {
		InitDB = true
	}
}

func parseEnv() {
	var err error

	ServerHostEnv := os.Getenv("ENVSERVER_HOST")
	if ServerHostEnv != "" {
		ServerHost = ServerHostEnv
	}

	ServerPortEnvStr := os.Getenv("ENVSERVER_PORT")
	if ServerPortEnvStr != "" {
		if ServerPort == -1 {
			ServerPort, err = strconv.Atoi(ServerPortEnvStr)
			if err != nil {
				log.Error("Failed to parse ENVSERVER_PORT")
			}
			CheckError(err)
		}
	}

	if !Verbose {
		VerboseEnv := os.Getenv("ENVSERVER_VERBOSE")
		if VerboseEnv == "true" {
			Verbose = true
		} else if VerboseEnv == "false" {
			Verbose = false
		}

		if Verbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
			gin.SetMode(gin.ReleaseMode)
			gin.DefaultWriter = io.Discard
		}
	}

	TLSEnv := os.Getenv("ENVSERVER_TLS")
	if TLSEnv == "true" {
		UseTLS = true
		Insecure = false
	} else if TLSEnv == "false" {
		UseTLS = false
		Insecure = true
	}
}
