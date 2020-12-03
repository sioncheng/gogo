package main

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	appName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|
		syslog.LOG_SYSLOG|
		syslog.LOG_MAIL|
		syslog.LOG_LOCAL7,
		appName)
	if err != nil {
		log.Fatal(err)
	}

	sysLog.Info("Loggin in Go\n")
}
