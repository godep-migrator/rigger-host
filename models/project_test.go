package models

import (
	"testing"
)

func TestNewProject_return_valid_project(t *testing.T) {
	project, err := NewProject("rigger-host", "rigger-dot-io", "github.com", "git", "git://github.com/rigger-dot-io/rigger-host")
	if err != nil {
		t.Fatalf("Error creating new project: %s", err)
	}
	if project.Name != "rigger-host" {
		t.Fatalf("Project names don't match: %s", project.Name)
	}
}

func TestNewProject_creates_a_valid_rsa_key_for_the_project(t *testing.T) {
	project, err := NewProject("rigger-host", "rigger-dot-io", "github.com", "git", "git://github.com/rigger-dot-io/rigger-host")
	if err != nil {
		t.Fatalf("Error creating new project: %s", err)
	}
	if project.DeployKeyPrivate == "" {
		t.Fatalf("Empty private key")
	}
	if project.DeployKeyPublic == "" {
		t.Fatalf("Empty private key")
	}
}
