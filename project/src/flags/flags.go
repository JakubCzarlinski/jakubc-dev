package flags

import (
	"github.com/JakubCzarlinski/go-logging"
)

const Name string = "jakubc"

const AssestsDir string = "./dist/assets/"
const DisableCache bool = false
const UseGinDefault bool = false
const UseGzip bool = true
const UseHttps bool = false
const UseLiveReload bool = false

func init() {
	logging.MinLogLevel = logging.DEBUG
	logging.UseLogger = true
	logging.UseTimestamp = true
	logging.UseLineLabels = true
	logging.UsePrefix = true
}
