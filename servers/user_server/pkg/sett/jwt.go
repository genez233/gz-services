package sett

import "time"

type JWTSetting struct {
	Issuer string
	Secret string
	Expire time.Duration
}
