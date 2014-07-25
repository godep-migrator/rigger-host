package utils

import (
	"fmt"
	"log"
	"os"
)

func CreatePidFile(pidfile string) error {
	fi, err := os.Stat(pidfile)
	if (fi != nil) && (err == nil) {
		return fmt.Errorf("Pid file found, please make sure rigger is not running or delete %s", pidfile)
	}

	file, err := os.Create(pidfile)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	return err
}

func RemovePidFile(pidfile string) {
	if err := os.Remove(pidfile); err != nil {
		log.Printf("Error removing %s: %s\n", pidfile, err)
	}
}
