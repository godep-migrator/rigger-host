package main

import (
	"testing"
)

func TestStartSocketConnection_invalidPath(t *testing.T) {
	sc := new(SocketConnection)
	if err := sc.StartSocketConnection("/invalid/path.sock"); err == nil {
		t.Fatalf("should have thrown an error")
	}
}
