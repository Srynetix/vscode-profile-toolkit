package parser

import (
	"path/filepath"
	"testing"
)

func TestParsePath(t *testing.T) {
	parser := &ProfilePackParser{}
	pack := parser.ParsePath(filepath.Join("../../fixtures", "profiles", "Default.code-profile"))

	if pack.Name != "Default" {
		t.Fatalf("pack.Name should be \"Default\" and not \"%s\"", pack.Name)
	}
}

func TestParseFolder(t *testing.T) {
	parser := &ProfilePackParser{}
	pack := parser.ParseFolder(filepath.Join("../../fixtures", "extracted", "Default"))

	if pack.Name != "Default" {
		t.Fatalf("pack.Name should be \"Default\" and not \"%s\"", pack.Name)
	}
}
