package main

import (
	"github.com/BenLubar/steamworks"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	le := logrus.NewEntry(log)

	le.Info("initializing steamworks client")
	if err := steamworks.InitClient(true); err != nil {
		le.WithError(err).Fatal("unable to init steamworks client")
		return
	}

	le.
		WithField("app-id", steamworks.GetAppID()).
		WithField("steam-id", steamworks.GetSteamID().String()).
		Info("steamworks client initialized!")
	steamworks.

	le.Info("shutting down steamworks")
	steamworks.Shutdown()
	le.Info("steamworks successfully shutdown")
}
