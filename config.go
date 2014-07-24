package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/imdario/mergo"
	"github.com/mitchellh/mapstructure"
)

// Config options for Rigger server.
type Config struct {

	// Path to main configuration file, every option below should be tweakable
	// via this file. CLI flags may overwrite these options
	ConfigFile string

	// Whether the server should run in daemon mode or not
	Daemon bool `mapstructure:"daemon"`

	// Log file
	LogFile string `mapstructure:"log_file"`

	// Server name, defaults to Hostname
	NodeName string `mapstructure:"node_name"`

	// Pidfile location for daemonized server
	PidFile string `mapstructure:"pid_file"`

	// Unix socket file. Any third-party application (including rigger's own API)
	// should use this socket for communication with the master server.
	SocketFile string `mapstructure:"socket_file"`
}

// Return a configuration object with default values.
func DefaultConfig() *Config {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error determining hostname: %s", err)
	}
	// Default config options
	return &Config{
		NodeName: hostname,
	}
}

// Merges two configuration objects. Values in C2 will overwrite values in C1.
func MergeConfig(c1, c2 *Config) *Config {
	var result Config = *c1

	if err := mergo.Merge(&result, c2); err != nil {
		log.Fatalf("Error merging Configuration objects: %s", err)
	}

	return &result
}

// Reads the configuration file found at the given path. Only files are
// accepted, directories will be rejected and an error will be thrown. If the
// file exists, it will be passed to the DecodeConfig method for processing.
func ReadConfigPath(path string) (*Config, error) {
	result := new(Config)

	// Try opening config file at path
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}

	// Try reading config file
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}

	// We are expecting a config file, reject directories
	if fi.IsDir() {
		return nil, fmt.Errorf("'%s' is a directory, expected a file", path)
	}

	config, err := DecodeConfig(f)
	f.Close()
	if err != nil {
		return nil, fmt.Errorf("Error parsing config file at '%s': %s", path, err)
	}

	result = MergeConfig(result, config)
	return result, nil
}

// Parses the config file. Content of the file must be valid JSON, an error will
// be thrown otherwise.
func DecodeConfig(r io.Reader) (*Config, error) {
	var raw interface{}
	var result Config

	// Create new JSON decoder
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}

	// Parse config mapstructure metadata
	var md mapstructure.Metadata

	// Match config file fields to Config attributes
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	})
	if err != nil {
		return nil, err
	}

	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}

	return &result, nil
}
