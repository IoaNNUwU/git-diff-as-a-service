package logger

import (
	"fmt"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level     string `envconfig:"LEVEL"  required:"true"`
	Format    string `envconfig:"Format" required:"true"`
	Stream    string `envconfig:"STREAM"`
	Folder    string `envconfig:"FOLDER"`
	AddSource bool   `envconfig:"ADD_SOURCE"`
}

func MustNewConfig() Config {
	var config Config
	if err := envconfig.Process("LOGGER", &config); err != nil {
		panic(fmt.Sprintf("unable to load logger config: %s", err))
	}

	config.Level = strings.Trim(config.Level, " ")
	config.Format = strings.Trim(config.Format, " ")
	config.Stream = strings.Trim(config.Stream, " ")
	config.Folder = strings.Trim(config.Folder, " ")

	if config.Folder == "" && config.Stream == "" {
		panic("logger config invalid: nor folder or an output stream selected. set LOGGER_STREAM=NONE to disable logging")
	}

	return config
}
