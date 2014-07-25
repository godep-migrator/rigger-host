package main

import (
	"fmt"
	"log"
	"net"
)

// SocketConnection is the main end-point of communication with the host
// server. It sets a up a unix socket listener and handles connections.
type SocketConnection struct {
	listener net.Listener
}

// StartSocketConnection will create a socket file and bind a listener to it.
func (s *SocketConnection) StartSocketConnection(socketFile string) error {
	listener, err := net.Listen("unix", socketFile)
	if err != nil {
		return fmt.Errorf("Failed to open socket file %s: %s", socketFile, err)
	}

	s.listener = listener
	log.Printf("Node listening for incoming connections on: %s\n", socketFile)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return fmt.Errorf("Error accepting message: %s", err)
		}
		go handleClient(conn)
	}
}

// StopSocketConnection will try to close all connections and clean up the
// listener so the socket can be freed.
func (s *SocketConnection) StopSocketConnection() error {
	if err := s.listener.Close(); err != nil {
		return fmt.Errorf("Could not stop listener: %s", err)
	}
	log.Printf("All connections closed")
	return nil
}

// Process incoming messages
func handleClient(c net.Conn) error {
	// Close connections after it's been used
	defer c.Close()
	// Main loop
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return err
		}

		data := buf[0:nr]
		// @TODO: relay messages to services
		log.Printf("Received: %s", data)
	}
}
