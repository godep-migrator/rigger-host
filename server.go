package main

import ()

// Server
type Server struct {
	config *Config
}

// Starts a master server using a given Configuration
func (s *Server) Start(config *Config) error {
	// Set server config
	s.config = config
	return nil
}
