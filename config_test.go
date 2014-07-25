package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	config := new(Config)
	config.LoadDefaultConfig()
	if config.NodeName == "" {
		t.Fatalf("should never be empty")
	}
}

func TestMergeWith(t *testing.T) {
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

	c1.MergeWith(c2)

	if !reflect.DeepEqual(c1, c2) {
		t.Fatalf("should be equal %v %v", c1, c2)
	}
}

func TestLoadConfigFromPath_badPath(t *testing.T) {
	c := new(Config)
	err := c.LoadConfigFromPath("/does/not/exist.json")
	if err == nil {
		t.Fatalf("should have thrown an error")
	}
}

func TestLoadConfigFromPath_folderPath(t *testing.T) {
	td, err := ioutil.TempDir("", "rigger")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	c := new(Config)
	err = c.LoadConfigFromPath(td)
	if err == nil {
		t.Fatalf("should have thrown an error")
	}

	defer os.RemoveAll(td)
}

func TestLoadConfigFromPath_correctPath(t *testing.T) {
	tf, err := ioutil.TempFile("", "rigger.json")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	tf.Write([]byte(`{"node_name":"foo", "daemon":true}`))
	tf.Close()

	c := new(Config)
	err = c.LoadConfigFromPath(tf.Name())
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if c.NodeName != "foo" {
		t.Fatalf("should be equal %s foo", c.NodeName)
	}

	if c.Daemon != true {
		t.Fatalf("should be equal %s true", c.Daemon)
	}

	defer os.Remove(tf.Name())
}
