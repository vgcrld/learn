package messages

import (
	log "github.com/sirupsen/logrus"
)

var (
	NoConfigurationFound = log.Fields{
		"id":      1,
		"message": "this is a message",
		"prog":    "bob.dll",
	}
)
