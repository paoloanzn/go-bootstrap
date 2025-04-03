package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/paoloanzn/go-bootstrap/config"
)

// helper function to run the main.go file as a subprocess using 'go run'.
// This allows us to capture exit codes and output without directly invoking os.Exit in the test process.
func runMain(args ...string) (output string, exitCode int, err error) {
	// Prepare command: assume main.go is in the same directory as this test file.
	cmdArgs := append([]string{"run", "main.go"}, args...)
	cmd := exec.Command("go", cmdArgs...)

	// Capture combined output
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &outBuf

	err = cmd.Run()
	output = outBuf.String()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
			return output, exitCode, err
		}
		// if error is not an ExitError, set exit code to -1
		exitCode = -1
		return output, exitCode, err
	}
	// If no error, exit code is 0
	exitCode = 0
	return output, exitCode, nil
}

// TestMainNoArgs tests the case where no command-line arguments are provided.
// Expected outcome: the program should exit with code 1.
func TestMainNoArgs(t *testing.T) {
	// Running without any additional arguments.
	output, exitCode, err := runMain()
	// We expect an error because os.Exit(1) is called when len(os.Args) < 2.
	if err == nil {
		t.Fatalf("Expected error for missing arguments, got nil")
	}
	if exitCode != 1 {
		t.Fatalf("Expected exit code 1, got %d, output: %s", exitCode, output)
	}
	// Inline comment: This test verifies that the application fails early if no arguments are provided.
}

// TestMainInitMissingArg tests the 'init' command without providing the required JSON template argument.
// Expected outcome: the program should exit with code 1.
func TestMainInitMissingArg(t *testing.T) {
	// Running with 'init' but missing the JSON template argument.
	output, exitCode, err := runMain("init")
	if err == nil {
		t.Fatalf("Expected error due to missing JSON template argument, got nil")
	}
	if exitCode != 1 {
		t.Fatalf("Expected exit code 1 for missing argument, got %d, output: %s", exitCode, output)
	}
	// Inline comment: This test ensures that the program does not proceed without required parameters for the 'init' command.
}

// TestMainInitFailParse simulates a failure in parsing the JSON template by providing a non-existent template file.
// Expected outcome: the program should log a fatal error and exit with a non-zero exit code.
func TestMainInitFailParse(t *testing.T) {
	// Using a filename that likely does not exist to force a parse error.
	fakeTemplate := "non_existing_template.json"
	output, exitCode, err := runMain("init", fakeTemplate)
	if err == nil {
		t.Fatalf("Expected error due to failing to parse template, got nil")
	}
	if exitCode == 0 {
		t.Fatalf("Expected non-zero exit code due to parse failure, got %d, output: %s", exitCode, output)
	}
	// Inline comment: This test validates error handling when parsing fails.
}

// TestMainDefault tests the default case when a command other than 'init' is provided.
// Expected outcome: the program should print the version string from the config package.
func TestMainDefault(t *testing.T) {
	// Using a dummy command that is not 'init' to trigger the default case.
	dummyCommand := "version"
	output, exitCode, err := runMain(dummyCommand)
	if err != nil {
		t.Fatalf("Did not expect an error for default command, got: %v, output: %s", err, output)
	}
	if exitCode != 0 {
		t.Fatalf("Expected exit code 0 for default command, got %d, output: %s", exitCode, output)
	}

	expected := fmt.Sprintf("version %s\n", config.VERSION)
	if strings.TrimSpace(output) != strings.TrimSpace(expected) {
		t.Fatalf("Expected output '%s', got '%s'", expected, output)
	}
	// Inline comment: This test ensures that the version information is correctly output in the default case.
}
