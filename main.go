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
	host "github.com/rigger-dot-io/rigger-host/host"
)

var (
	server *host.Server
	config *host.Config
)

func main() {
	// Halt process if host OS is not *nix type
	if runtime.GOOS == "windows" {
		log.Fatalf("Rigger is only supported on *nix platforms")
	}

	// Check security level
	if os.Geteuid() != 0 {
		log.Fatalf("Rigger needs to be run as root")
	}

	// Check GOMAXPROCS
	if runtime.GOMAXPROCS(0) == 1 {
		log.Printf("WARNING: It is highly recommended to set GOMAXPROCS higher than 1")
	}

	// Boot server
	config = readConfig()
	server = new(host.Server)
	server.Start(config)
}

// readConfig reads in any options passed through the command-line and merges
// these options with the base Configuration object. The order of importance
// is: CLI > Config file > Defaults
func readConfig() *host.Config {

	var cmdConfig host.Config

	var version = flag.Bool([]string{"v", "-version"},
		false, "Print version information and quit")
	flag.BoolVar(&cmdConfig.Daemon, []string{"d", "-daemon"},
		false, "Enable daemon mode")
	flag.StringVar(&cmdConfig.PidFile, []string{"-pidfile"},
		"/var/run/rigger.pid", "Path to use for PID file")
	flag.StringVar(&cmdConfig.ConfigFile, []string{"c", "-config"},
		"/etc/rigger.conf", "Load configuration from file")
	flag.StringVar(&cmdConfig.LogFile, []string{"l", "-logfile"},
		"/var/log/rigger.log", "Path to rigger log file")
	flag.IntVar(&cmdConfig.Port, []string{"p", "-port"}, 9876, "RPC Port")

	flag.Parse()

	// Handle version flag separate from others
	if *version {
		fmt.Printf("Rigger Host v%s%s, build %v\n", Version, VersionPrerelease, GitCommit)
		os.Exit(0)
	}

	// New empty configuration object
	var config = new(host.Config)

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

	return config
}
