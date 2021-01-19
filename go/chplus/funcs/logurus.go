package funcs

import (
	"github.com/vgcrld/all-things-go/cmd/messages"

	log "github.com/sirupsen/logrus"
)

func RunLogurus() {
	log.SetLevel(log.DebugLevel)
	// log.SetLevel(log.InfoLevel)
	f := log.Fields{
		"name": "richard",
		"age":  50,
	}
	log.WithFields(f).Infof("Message: %v", "Getting the stuff")

	log.WithFields(messages.NoConfigurationFound).Debugf("ERROR: %v", "you failed big")
}
