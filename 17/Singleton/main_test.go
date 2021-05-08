package Singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	var s Singleton
	s = GetInstance()
	if s == nil {
		t.Fatalf("First sigletone is nil")
	}

	s.SetTitle("First value")
	checkTitle := s.GetTitle()

	if checkTitle != "First value" {
		t.Errorf("First value is not setted")
	}

	var s2 Singleton
	s2 = GetInstance()
	if s2 != s {
		t.Error("New instance different")
	}

	s2.SetTitle("New title")

	newTitle := s.GetTitle()
	if newTitle != "New title" {
		t.Errorf("Title different after change")
	}
}
