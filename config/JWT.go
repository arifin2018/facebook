package config

import (
	"time"
)

type ConfigJwt struct {
	Exp time.Time
}

var Timeexp int

var DefaultConfigJwt = &ConfigJwt{}
