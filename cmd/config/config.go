package config

import (
	_ "embed"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App struct {
			Version int64
		}
		Listener string `default:":3000"`
		ServiceAPI struct {
			AuthService string
		}
		Logger   struct {
			Encoding string `default:"json"`
			Verbose  bool
		}
		DB struct {
			Host   string
			Port   string
			User   string
			Pass   string
			DBName string
		}
	}
)

func LoadConfigFromFile(fileName string) (*Config, error) {
	var cfg Config


	err := godotenv.Load(fileName)
	if err != nil {
		return nil, err
	}

	cfg.Listener = os.Getenv("LISTENER")
	cfg.Logger.Encoding = os.Getenv("ENDCODING")
	verbose, err := strconv.ParseBool(os.Getenv("VERSBOSE"))
	if err != nil {
		return nil, err
	}

	versionStr := os.Getenv("APP_VERSION")
	version, err := strconv.ParseInt(versionStr, 2, 64)
	if err != nil {
		return nil, err
	}

	cfg.App.Version = version

	cfg.Logger.Verbose = verbose
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Pass = os.Getenv("DB_PASS")
	cfg.DB.DBName = os.Getenv("DB_NAME")

	cfg.ServiceAPI.AuthService = os.Getenv("AUTH_SERVICE_URL")

	return &cfg, nil
}
