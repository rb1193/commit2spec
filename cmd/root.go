package cmd

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

type RepoOpener func() (*git.Repository, error)

type RootCmd struct {
	RepoOpener RepoOpener
}

func (r *RootCmd) Command() *cobra.Command {
	var outputPath string

	cmd := &cobra.Command{
		Use:   "commit2spec [commit-sha]",
		Short: "Generate markdown spec files from git commits",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			commitSHA := args[0]

			repo, err := r.RepoOpener()
			if err != nil {
				return err
			}

			hash, err := repo.ResolveRevision(plumbing.Revision(commitSHA))
			if err != nil {
				return err
			}

			commit, err := repo.CommitObject(*hash)
			if err != nil {
				return err
			}

			return os.WriteFile(outputPath, []byte(commit.Message), 0644)
		},
	}

	cmd.Flags().StringVarP(&outputPath, "output", "o", "", "output file path")
	cmd.MarkFlagRequired("output")

	return cmd
}

func Execute() {
	root := &RootCmd{
		RepoOpener: func() (*git.Repository, error) {
			return git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
		},
	}
	if err := root.Command().Execute(); err != nil {
		os.Exit(1)
	}
}
