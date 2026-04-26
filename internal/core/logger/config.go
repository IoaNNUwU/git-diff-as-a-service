package logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL"  required:"true"`
	Format string `envconfig:"Format" required:"true"`
	Stream string `envconfig:"STREAM"`
	Folder string `envconfig:"FOLDER"`
}

func MustNewConfig() Config {
	var config Config
	if err := envconfig.Process("LOGGER", &config); err != nil {
		panic(fmt.Sprintf("unable to load logger config: %w", err))
	}

	if config.Folder == "" && config.Stream == "" {
		panic("logger config invalid: nor folder or an output stream selected. set LOGGER_STREAM=NONE to disable logging")
	}

	return config
}
