package core_postgres_conn

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Database string        `envconfig:"DB" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func MustNewConfig() Config {

	var config Config

	if err := envconfig.Process("POSTGRES", &config); err != nil {
		panic(fmt.Sprintf("unable to load PostgreSQL config: %s", err))
	}

	return config
}
