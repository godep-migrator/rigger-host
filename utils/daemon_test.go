package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func TestCreatePidFile_invalidPath(t *testing.T) {
	err := CreatePidFile("/does/not/exist.pid")
	if err == nil {
		t.Fatalf("should have thrown an error")
	}
}

func TestCreatePidFile_existingPidFile(t *testing.T) {
	pf, err := ioutil.TempFile("/tmp", "rigger.pid")
	if err != nil {
		t.Fatalf("could not create temp file: %s", err)
	}

	err = CreatePidFile(pf.Name())
	if err == nil {
		t.Fatalf("should have thrown 'pidfile exists' error")
	}
}

func TestCreatePidFile_validPidPath(t *testing.T) {
	var pidFile = fmt.Sprintf("/tmp/rigger%d.pid", rand.Int())
	err := CreatePidFile(pidFile)
	if err != nil {
		t.Fatalf("should not have thrown error: %s", err)
	}
	pFile, err := os.Stat(pidFile)
	if (err != nil) || (pFile == nil) {
		t.Fatalf("pidfile does not exist at %s", pidFile)
	}
	defer os.Remove(pidFile)
}

func TestRemovePidFile(t *testing.T) {
	pf, err := ioutil.TempFile("/tmp", "rigger.pid")
	if err != nil {
		t.Fatalf("could not create temp file: %s", err)
	}

	RemovePidFile(pf.Name())

	pFile, err := os.Stat(pf.Name())
	if (pFile != nil) || (err == nil) {
		t.Fatalf("pidfile was not removed")
	}
}
