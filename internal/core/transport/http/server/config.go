package core_http_server

import (
	"fmt"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func MustNewConfig() Config {

	var config struct {
		Addr            string `envconfig:"ADDR" required:"true"`
		ShutdownTimeout string `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
	}

	if err := envconfig.Process("HTTP_SERVER", &config); err != nil {
		panic(fmt.Sprintf("unable to load HTTP server config: %s", err))
	}

	config.Addr = strings.Trim(config.Addr, " ")
	config.ShutdownTimeout = strings.Trim(config.ShutdownTimeout, " ")

	timeout, err := time.ParseDuration(config.ShutdownTimeout)
	if err != nil {
		panic(fmt.Sprintf("unable to load HTTP server config: unable to parse timeout: %q", config.ShutdownTimeout))
	}

	return Config{
		Addr:            config.Addr,
		ShutdownTimeout: timeout,
	}
}
