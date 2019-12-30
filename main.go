package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configurationFile = flag.String("config", "config.yaml", "where the configuration file is located")
	logLevel = flag.String("loglevel", "debug", "sets the log levels which will be written to the log file")
)


func main() {
	flag.Parse()
	logrus.SetLevel(func(level string) logrus.Level {
		var lvl logrus.Level

		if level == "debug" {
			lvl = logrus.DebugLevel
		} else if level == "info" {
			lvl = logrus.InfoLevel
		} else if level == "trace" {
			lvl = logrus.TraceLevel
		} else if level == "warn" {
			lvl = logrus.WarnLevel
		} else if level == "error" {
			lvl = logrus.ErrorLevel
		} else {
			lvl = logrus.FatalLevel
		}

		return lvl
	}(*logLevel))

	cfg := readConfigurationFile(*configurationFile)

	stop := make(chan os.Signal)

	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGKILL, syscall.SIGABRT)
	sig := <-stop

	gracefulShutdownDuration := 5 * time.Second
	logrus.Warnf("%v signal received", sig)
	logrus.Info("blog database service application shutting down. Waiting 5 seconds to finish outstanding operations.")
	<-time.After(gracefulShutdownDuration)
}
