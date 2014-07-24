package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()
	if config.NodeName == "" {
		t.Fatalf("should never be empty")
	}
}

func TestMergeConfig(t *testing.T) {
	c1 := &Config{
		NodeName:   "rigger.dev",
		Daemon:     false,
		ConfigFile: "/etc/rigger.conf",
		LogFile:    "/var/log/rigger.log",
		PidFile:    "/var/run/rigger.pid",
		SocketFile: "/var/run/rigger.sock",
	}

	c2 := &Config{
		NodeName:   "rigger.prod",
		Daemon:     true,
		ConfigFile: "/etc/test.conf",
		LogFile:    "/var/log/test.log",
		PidFile:    "/var/run/test.pid",
		SocketFile: "/var/run/test.sock",
	}

	c := MergeConfig(c1, c2)

	if !reflect.DeepEqual(c, c2) {
		t.Fatalf("should be equal %v %v", c, c2)
	}
}

func TestReadConfigPath_badPath(t *testing.T) {
	_, err := ReadConfigPath("/does/not/exist.json")
	if err == nil {
		t.Fatalf("should have thrown an error")
	}
}

func TestReadConfigPath_folderPath(t *testing.T) {
	td, err := ioutil.TempDir("", "rigger")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	_, err = ReadConfigPath(td)
	if err == nil {
		t.Fatalf("should have thrown an error")
	}

	defer os.RemoveAll(td)
}

func TestReadConfigPath_correctPath(t *testing.T) {
	tf, err := ioutil.TempFile("", "rigger.json")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	tf.Write([]byte(`{"node_name":"foo", "daemon":true}`))
	tf.Close()

	config, err := ReadConfigPath(tf.Name())
	if (err != nil) || (config == nil) {
		t.Fatalf("err: %s", err)
	}

	if config.NodeName != "foo" {
		t.Fatalf("should be equal %s foo", config.NodeName)
	}

	if config.Daemon != true {
		t.Fatalf("should be equal %s true", config.Daemon)
	}

	defer os.Remove(tf.Name())
}
