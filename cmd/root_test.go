package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRootCommandWritesHelloWorldToOutputFile(t *testing.T) {
	// Create a temp directory for our output file
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "output.md")

	// Set up command args
	rootCmd.SetArgs([]string{outputPath})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Read the output file and verify contents
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	expected := "Hello, world!"
	if string(content) != expected {
		t.Errorf("expected %q, got %q", expected, string(content))
	}
}
