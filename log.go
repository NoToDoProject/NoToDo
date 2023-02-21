package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// set text output
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // time format
		ForceColors:     true,                  // force show color
		//FullTimestamp:   true,                  // show full time format
	})

	// set output to stdout, default is stderr
	// can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}
