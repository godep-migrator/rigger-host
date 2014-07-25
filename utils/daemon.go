package utils

import (
	"fmt"
	"log"
	"os"
)

// Creates a pid file at given path. If the file already exists the
// process will throw an error and exit. This method is only run when the
// rigger host is in daemon mode
func CreatePidFile(pidFile string) error {
	fi, err := os.Stat(pidFile)
	if (fi != nil) && (err == nil) {
		return fmt.Errorf("Pid file found, please make sure rigger is not running or delete %s", pidFile)
	}

	file, err := os.Create(pidFile)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	return err
}

// RemovePidFile will attempt to remove a pid file.
func RemovePidFile(pidFile string) {
	if err := os.Remove(pidFile); err != nil {
		log.Printf("Error removing %s: %s\n", pidFile, err)
	}
}
