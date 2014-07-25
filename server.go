package main

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

	// Unix Socket connection
	socketConn *SocketConnection
}

// Starts a master server using a given Configuration object
func (s *Server) Start(config *Config) {
	// Capture Signals
	signalChannel := make(chan os.Signal, 1)
	// Listen to terminate signals
	signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	// Stop server when received a terminate signal
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Captured %v, stopping server and exiting...", sig)
		server.Stop()
	}(signalChannel)

	// Set server config
	s.config = config

	// If daemon mode is turned on, we need to create a pid file
	if config.Daemon {
		log.Println("Starting daemon...")

		if err := utils.CreatePidFile(config.PidFile); err != nil {
			fmt.Errorf("Could not create pid file: %s", err)
		}
	}

	// Start listening to socket connections
	s.socketConn = new(SocketConnection)
	if err := s.socketConn.StartSocketConnection(config.SocketFile); err != nil {
		log.Println("Could not start socket listener: %s", err)
		s.Stop()
	}
}

// Stops the server
func (s *Server) Stop() {
	// Remove pid file if daemon mode is turned on
	if s.config.Daemon {
		log.Println("Removing pid file...")
		utils.RemovePidFile(s.config.PidFile)
	}

	// Close all socket connections before exiting so the socket file can be
	// removed
	log.Println("Closing connections...")
	if err := s.socketConn.StopSocketConnection(); err != nil {
		log.Fatalf("Could not close socket connection: %s", err)
	}

	os.Exit(0)
}
