package config

import (
	"time"
)

type ConfigJwt struct {
	Exp time.Time
}

var Timeexp int

var DefaultConfigJwt = ConfigJwt{
	Exp: time.Now().Add(time.Hour * time.Duration(Timeexp)),
}
