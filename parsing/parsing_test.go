package parsing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	// using log package since parsing.go uses log.Fatalf
)

// TestParseTemplateValid tests the happy path where a valid JSON file is provided.
func TestParseTemplateValid(t *testing.T) {
	// Create temporary file with valid JSON content
	tempFile, err := ioutil.TempFile("", "valid_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Prepare valid JSON data that matches JSONTemplate structure
	validJSON := `{
		"project": "example_project",
		"config": {"key": "value"}
	}`

	if _, err := tempFile.Write([]byte(validJSON)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Invoke the ParseTemplate function
	result, err := ParseTemplate(tempFile.Name())
	if err != nil {
		t.Errorf("Expected no error for valid JSON, got: %v", err)
	}

	// Verify the parsed content
	// Unmarshal validJSON into a map for easier comparison
	var expected map[string]interface{}
	if err := json.Unmarshal([]byte(validJSON), &expected); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}

	// Check the project field
	if result.Project != expected["project"] {
		t.Errorf("Expected project %v, got %v", expected["project"], result.Project)
	}

	// Check the config field
	expectedConfig, ok := expected["config"].(map[string]interface{})
	if !ok {
		t.Fatalf("Expected config to be a map[string]interface{}")
	}

	for key, val := range expectedConfig {
		if result.Config[key] != val {
			t.Errorf("For key %s expected value %v, got %v", key, val, result.Config[key])
		}
	}
}

// TestParseTemplateInvalidJSON tests the scenario where the JSON file content is invalid.
func TestParseTemplateInvalidJSON(t *testing.T) {
	// Create temporary file with invalid JSON content
	tempFile, err := ioutil.TempFile("", "invalid_*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	invalidJSON := "{ invalid json }"
	if _, err := tempFile.Write([]byte(invalidJSON)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Call ParseTemplate which should return an error due to invalid JSON
	result, err := ParseTemplate(tempFile.Name())
	if err == nil {
		t.Errorf("Expected error when parsing invalid JSON, got nil")
	}
	// Check if error message contains the expected substring
	if !strings.Contains(err.Error(), fmt.Sprintf("%s", tempFile.Name())) {
		t.Errorf("Error message does not mention file path. Got: %v", err.Error())
	}

	// Even with error, the result should be a valid pointer to JSONTemplate (with zeroed fields)
	if result == nil {
		t.Errorf("Expected non-nil result even when error occurs")
	}
}

// TestParseTemplateFileNotFound tests the scenario where the file does not exist.
// Since ParseTemplate uses log.Fatalf on file read error, we run it in a subprocess
// to capture the behavior without exiting the main test process.
func TestParseTemplateFileNotFound(t *testing.T) {
	// When the environment variable is set, call the function directly.
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		// This should trigger log.Fatalf inside ParseTemplate due to missing file.
		// The following call is expected to terminate the process.
		ParseTemplate("/nonexistent/path/to/file.json")
		// In case it returns, exit with code 0
		os.Exit(0)
	}

	// Prepare command to re-run the test binary with environment variable set
	cmd := exec.Command(os.Args[0], "-test.run=^TestParseTemplateFileNotFound$")
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("Expected subprocess to fail due to log.Fatal on file not found, but it succeeded. Output: %s", output)
	}
}

// Additional grouping can be added if more exported functions/methods are present.
// Each test case above includes descriptive comments regarding their purpose and expected outcome.

// Note: The tests for ParseTemplate cover the valid scenario, invalid JSON content scenario, and file-not-found (error) scenario.
// The file-not-found test uses a subprocess to safely capture the log.Fatalf behavior.
