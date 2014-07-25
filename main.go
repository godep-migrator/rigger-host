// Rigger CI
//
// Opinionated, docker-based continuous integration platform designed with
// performance in mind.
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	flag "github.com/dotcloud/docker/pkg/mflag"
)

func main() {
	// Halt process if host OS is not *nix type
	if runtime.GOOS == "windows" {
		log.Fatalf("Rigger daemon is only supported on *nix platforms")
	}

	// Check security level
	if os.Geteuid() != 0 {
		log.Fatalf("Rigger daemon needs to be run as root")
	}

	// Boot server
	config := readConfig()
	server := new(Server)
	server.Start(config)
}

func readConfig() *Config {

	var cmdConfig Config

	var version = flag.Bool([]string{"v", "-version"},
		false, "Print version information and quit")
	flag.BoolVar(&cmdConfig.Daemon, []string{"d", "-daemon"},
		false, "Enable daemon mode")
	flag.StringVar(&cmdConfig.PidFile, []string{"p", "-pidfile"},
		"/var/run/rigger.pid", "Path to use for PID file")
	flag.StringVar(&cmdConfig.SocketFile, []string{"s", "-socket"},
		"/var/run/rigger.sock", "Use this file as the rigger socket")
	flag.StringVar(&cmdConfig.ConfigFile, []string{"c", "-config"},
		"/etc/rigger.conf", "Load configuration from file")
	flag.StringVar(&cmdConfig.LogFile, []string{"l", "-logfile"},
		"/var/log/rigger.log", "Path to rigger log file")

	flag.Parse()

	// Handle version flag separate from others
	if *version {
		fmt.Printf("Rigger Host v%s%s, build %v\n", Version, VersionPrerelease, GitCommit)
		os.Exit(0)
	}

	// New empty configuration object
	var config = new(Config)

	// Load defaults
	if err := config.LoadDefaultConfig(); err != nil {
		log.Fatalf("Could not load default config: %s", err)
	}

	// Overwrite defaults with options from command-line
	if err := config.MergeWith(&cmdConfig); err != nil {
		log.Fatalf("Could not load config: %s", err)
	}

	// Load config file
	if err := config.LoadConfigFromPath(config.ConfigFile); err != nil {
		log.Fatalf("Could not load config: %s", err)
	}

	log.Printf("Node name: %v\n", config.NodeName)

	return config
}
