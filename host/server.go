package host

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	utils "github.com/rigger-dot-io/rigger-host/utils"
)

// Server
type Server struct {
	// Configuration object
	config *Config
}

// Starts a master server using a given Configuration object
func (s *Server) Start(config *Config) {
	// Set server config
	s.config = config

	// If daemon mode is turned on, we need to create a pid file
	if config.Daemon {
		log.Println("Starting daemon...")

		if err := utils.CreatePidFile(config.PidFile); err != nil {
			fmt.Errorf("Could not create pid file: %s", err)
		}
	}

	s.initInterruptListener()

	RPCListen(config.Port)
}

// Stops the server
func (s *Server) Stop() {
	// Remove pid file if daemon mode is turned on
	if s.config.Daemon {
		log.Println("Removing pid file...")
		utils.RemovePidFile(s.config.PidFile)
	}

	os.Exit(0)
}

// Init hook to catch terminate signals
func (s *Server) initInterruptListener() {
	// Capture Signals
	signalChannel := make(chan os.Signal, 1)
	// Listen to terminate signals
	signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	// Stop server when received a terminate signal
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Captured %v, stopping server and exiting...", sig)
		s.Stop()
	}(signalChannel)
}
