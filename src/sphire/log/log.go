package log

import (
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

var lpx = log.New()
var environ string

func Init(formatter string, environment string, log_file string) bool {
	switch formatter {
	case "text":
		lpx.Formatter = new(log.TextFormatter)
	case "json":
		lpx.Formatter = new(log.JSONFormatter)
	}

	// Output to stderr/file instead of stdout
	if log_file != "" {
		f_ptr, _ := os.OpenFile(log_file, os.O_RDWR|os.O_APPEND, os.ModePerm)
		lpx.Out = f_ptr
	} else {
		log.SetOutput(os.Stderr)
	}

	switch environment {
	case "DEV":
		lpx.Level = log.DebugLevel
	case "STG":
		lpx.Level = log.InfoLevel
	case "PRD":
		lpx.Level = log.WarnLevel
	}
	environ = environment

	return true
}

// Log logs a message to the defined logger
// var fields map[string]interface{} = make(map[string]interface{})
func Log(fields map[string]interface{}, message string, level string) bool {
	if environ == "DEV" {
		fmt.Println(level, fields, message)
	}

	switch level {
	case "debug":
		lpx.WithFields(fields).Debug(message)
		return true
	case "info":
		lpx.WithFields(fields).Info(message)
		return true
	case "warn":
		lpx.WithFields(fields).Warn(message)
		return true
	case "error":
		lpx.WithFields(fields).Error(message)
		return true
	case "fatal":
		lpx.WithFields(fields).Fatal(message)
		return true
	case "panic":
		lpx.WithFields(fields).Panic(message)
		return true
	}

	return false
}
