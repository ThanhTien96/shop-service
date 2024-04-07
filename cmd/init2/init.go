package init2

import "flag"

const DEV_LOG_FILE = "./var/log/shop-dev.log"
const PRODUCTION_LOG_FILE = "./var/log/shop-production.log"

type (
	Args struct {
		ConfigFile string
		Debug      bool
	}
)

func LoadCommandArgs() *Args {
	argConfigFile := flag.String("config", "../.env-dev", "Configuration file")
	argIsDebug := flag.Bool("debug", false, "Enable debug output")

	flag.Parse()

	flags := &Args{
		*argConfigFile,
		*argIsDebug,
	}
	return flags
}