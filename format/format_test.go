package format

import (
	"testing"

	"github.com/paoloanzn/go-bootstrap/config"
)

func TestFormatPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "./example", expected: "./example"},           // Already formatted relative path
		{input: "/absolute/path", expected: "/absolute/path"}, // Absolute path
		{input: "relative/path", expected: "./relative/path"}, // Relative path without './'
		{input: ".test", expected: "./.test"},                 // Single dot
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := FormatPath(test.input)
			if result != test.expected {
				t.Errorf("FormatPath(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestMatchWildCards(t *testing.T) {
	// Set up the configuration for testing
	config.Cfg.ProjectName = "TestProject"

	// Test cases
	tests := []struct {
		input    string
		expected string
	}{
		{"<main_package>", "TestProject"},                                              // Single wildcard
		{"This is <main_package>", "This is TestProject"},                              // Wildcard in a sentence
		{"No wildcards here", "No wildcards here"},                                     // No wildcards
		{"Multiple <main_package> <main_package>", "Multiple TestProject TestProject"}, // Multiple instances of the same wildcard
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := MatchWildCards(tt.input)
			if result != tt.expected {
				t.Errorf("MatchWildCards(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
