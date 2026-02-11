package cmd

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRootCommandWritesCommitMessageToOutputFile(t *testing.T) {
	repo, commitHash, commitMsg := createTestRepo(t)

	tests := []struct {
		name string
		sha  string
	}{
		{"full SHA", commitHash.String()},
		{"short SHA", commitHash.String()[:7]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputPath := filepath.Join(t.TempDir(), "output.md")

			root := &RootCmd{
				RepoOpener: func() (*git.Repository, error) {
					return repo, nil
				},
			}
			cmd := root.Command()
			cmd.SetArgs([]string{tt.sha, "-o", outputPath})

			if err := cmd.Execute(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			content, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("failed to read output file: %v", err)
			}

			if string(content) != commitMsg {
				t.Errorf("expected %q, got %q", commitMsg, string(content))
			}
		})
	}
}

func createTestRepo(t *testing.T) (*git.Repository, plumbing.Hash, string) {
	t.Helper()

	repoDir := t.TempDir()
	repo, err := git.PlainInit(repoDir, false)
	if err != nil {
		t.Fatalf("failed to init repo: %v", err)
	}

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

	return repo, commitHash, commitMsg
}
