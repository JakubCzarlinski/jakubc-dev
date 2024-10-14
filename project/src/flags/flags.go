package flags

import (
	"os"

	"github.com/JakubCzarlinski/go-logging"
)

const Name string = "jakubc"

const (
	AssestsDir    string = "./dist/assets/"
	UseGinDefault bool   = false
	UseHttps      bool   = false
)

var (
	DisableCache  bool = true
	UseGzip       bool = true
	UseLiveReload bool = true
)

func init() {
	logging.MinLogLevel = logging.DEBUG
	logging.UseLogger = true
	logging.UseTimestamp = true
	logging.UseLineLabels = true
	logging.UsePrefix = true

	// Check if PROD is set in the environment
	_, ok := os.LookupEnv("PROD")
	if !ok {
		return
	}
	if os.Getenv("PROD") == "true" {
		logging.MinLogLevel = logging.INFO
		DisableCache = false
		UseGzip = true
		UseLiveReload = false
	}
}
