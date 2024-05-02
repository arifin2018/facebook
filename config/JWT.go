package config

import (
	"time"
)

type ConfigJwt struct {
	Exp time.Time
}

var DefaultConfigJwt = ConfigJwt{
	Exp: time.Now().Add(time.Hour * 72),
}
