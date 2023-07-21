package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Path string
)

// NewStashPopCmd creates a new token command.
func NewStashPopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "pop",
		Short:   "",
		Long:    "",
		Example: "stash pop",
		Aliases: []string{"unstash", "get"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Args:         cobra.NoArgs,
		SilenceUsage: true,
	}

	cmd.Flags().StringVarP(&Name, "name", "n", "", "Name of a stash. Should be a simple identifier akin to a job name.")

	err := cobra.MarkFlagRequired(cmd.Flags(), "name")
	if err != nil {
		panic(err)
	}

	return cmd
}
