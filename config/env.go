package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	// LogLevelEnvVar is the name of the environment variable that controls
	// the log level of the application logger
	LogLevelEnvVar = "LOG_LEVEL"

	// PortEnvVar is the name of the environment variable that controls the
	// value of the port the service should listen on
	PortEnvVar = "PORT"

	defaultLogLevel = logrus.DebugLevel // used if LOG_LEVEL not set
	defaultPort     = 3000              // used if PORT not set
)

// LogLevel returns the log level set in the environment, or debug if not defined
func LogLevel() logrus.Level {
	var (
		level logrus.Level
		err   error
	)

	if level, err = logrus.ParseLevel(os.Getenv(LogLevelEnvVar)); err != nil {
		return defaultLogLevel
	}

	return level
}

// Port returns the port the service should listen on, or 3000 if not defined or
// is not a valid port
func Port() int {
	var (
		rawPort string
		found   bool
		port    int
		err     error
	)

	if rawPort, found = os.LookupEnv(PortEnvVar); !found {
		return defaultPort
	}

	if port, err = strconv.Atoi(rawPort); err != nil {
		return defaultPort
	}

	return port
}
