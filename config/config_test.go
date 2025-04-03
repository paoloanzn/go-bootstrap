package config

import (
	"testing"
)

// TestCfgInitialization tests the initialization of the exported variable Cfg.
// This test checks that Cfg is not nil and that its default fields are set as expected.
func TestCfgInitialization(t *testing.T) {
	// Happy path: Expect Cfg to be already initialized
	if Cfg == nil {
		t.Error("Expected Cfg to be initialized, got nil")
	}

	// Edge case: Check the default value of ProjectName, which should be an empty string
	expectedProjectName := ""
	if Cfg.ProjectName != expectedProjectName {
		t.Errorf("Expected Cfg.ProjectName to be '%s', got '%s'", expectedProjectName, Cfg.ProjectName)
	}
}

// TestVersionConstant tests the exported constant VERSION.
// It verifies that the version is as expected.
func TestVersionConstant(t *testing.T) {
	// Happy path: The VERSION constant should be equal to "0.1"
	expectedVersion := "0.1"
	if VERSION != expectedVersion {
		t.Errorf("Expected VERSION to be '%s', got '%s'", expectedVersion, VERSION)
	}
}

// TestConfigStruct tests the Config struct behavior by manually creating an instance.
// Although there is no constructor, this test demonstrates using the struct with custom values.
func TestConfigStruct(t *testing.T) {
	// Happy path: Creating a Config with a valid ProjectName
	cfg := &Config{ProjectName: "SampleProject"}
	if cfg.ProjectName != "SampleProject" {
		t.Errorf("Expected ProjectName to be 'SampleProject', got '%s'", cfg.ProjectName)
	}

	// Edge case: Setting an empty string as ProjectName
	cfgEmpty := &Config{ProjectName: ""}
	if cfgEmpty.ProjectName != "" {
		t.Errorf("Expected ProjectName to be empty, got '%s'", cfgEmpty.ProjectName)
	}
}

// Additional tests can be added here as needed to further validate behavior or error conditions
// in functions/methods when they are added to the package.
