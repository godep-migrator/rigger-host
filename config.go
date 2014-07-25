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

// Load default values into Configuration object.
func (config *Config) LoadDefaultConfig() error {
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("Error determining hostname: %s", err)
	}
	// Default config options
	config.NodeName = hostname
	return nil
}

// Merges configuration objects.
func (config *Config) MergeWith(c *Config) error {
	if err := mergo.Merge(config, c); err != nil {
		return fmt.Errorf("Error merging Configuration objects: %s", err)
	}

	return nil
}

// Reads the configuration file found at the given path. Only files are
// accepted, directories will be rejected and an error will be thrown. If the
// file exists, it will be passed to the DecodeConfig method for processing.
func (config *Config) LoadConfigFromPath(path string) error {
	// Try opening config file at path
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error reading '%s': %s", path, err)
	}

	// Try reading config file
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return fmt.Errorf("Error reading '%s': %s", path, err)
	}

	// We are expecting a config file, reject directories
	if fi.IsDir() {
		return fmt.Errorf("'%s' is a directory, expected a file", path)
	}

	configFromFile, err := decodeConfig(f)
	f.Close()

	if err != nil {
		return fmt.Errorf("Could not decode configuration file: %s", err)
	}

	config.MergeWith(configFromFile)

	log.Printf("Loaded configuration from: %s", path)

	return nil
}

// Parses the config file. Content of the file must be valid JSON, an error will
// be thrown otherwise.
func decodeConfig(r io.Reader) (*Config, error) {
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
