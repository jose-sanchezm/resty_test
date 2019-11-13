package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Initializing")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Sending request without logrus")
	resty.New().SetDebug(true).R().Get("www.google.com")
	logrus.Debug("Request sent")

	logrus.Debug("Sending request with logrus")
	resty.New().SetLogger(logrus.New()).SetDebug(true).R().Get("www.google.com")
	logrus.Debug("Request sent")
}
