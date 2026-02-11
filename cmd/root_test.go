package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRootCommandWritesCommitMessageToOutputFile(t *testing.T) {
	// Create a temp git repo with a known commit
	repoDir := t.TempDir()
	repo, err := git.PlainInit(repoDir, false)
	if err != nil {
		t.Fatalf("failed to init repo: %v", err)
	}

	// Create a file and commit it
	testFile := filepath.Join(repoDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	wt, err := repo.Worktree()
	if err != nil {
		t.Fatalf("failed to get worktree: %v", err)
	}
	if _, err := wt.Add("test.txt"); err != nil {
		t.Fatalf("failed to add file: %v", err)
	}
	commitMsg := "Initial commit with test content"
	commitHash, err := wt.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@test.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		t.Fatalf("failed to commit: %v", err)
	}

	// Set up output file
	outputDir := t.TempDir()
	outputPath := filepath.Join(outputDir, "output.md")

	// Run the command with injected repo opener
	root := &RootCmd{
		RepoOpener: func() (*git.Repository, error) {
			return repo, nil
		},
	}
	cmd := root.Command()
	cmd.SetArgs([]string{commitHash.String(), "-o", outputPath})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the output file contains the commit message
	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	if string(content) != commitMsg {
		t.Errorf("expected %q, got %q", commitMsg, string(content))
	}
}
