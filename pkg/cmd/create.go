package cmd

import (
	"github.com/spf13/cobra"
)

var (
	AllowEmpty         bool
	Excludes           string
	Includes           string
	UseDefaultExcludes bool
)

// NewStashCmd creates a new token command.
func NewStashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Creates a stash containing the referenced files",
		Long:    "",
		Example: "stash create ...",
		Aliases: []string{"c", "new"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Args:         cobra.NoArgs,
		SilenceUsage: true,
	}

	cmd.Flags().StringVarP(&Name, "name", "n", "", "Name of a stash. Should be a simple identifier akin to a job name.")
	cmd.Flags().BoolVarP(&AllowEmpty, "allowEmpty", "", false, "Create stash even if no files are included. If false (default), an error is raised when the stash does not contain files.")
	cmd.Flags().StringVarP(&Excludes, "excludes", "", "", "Optional set of Ant-style exclude patterns. Use a comma separated list to add more than one expression. If blank, no files will be excluded.")
	cmd.Flags().StringVarP(&Includes, "includes", "", "", "Optional set of Ant-style include patterns. Use a comma separated list to add more than one expression. If blank, treated like **: all files. The current working directory is the base directory for the saved files, which will later be restored in the same relative locations, so if you want to use a subdirectory wrap this in dir.")
	cmd.Flags().BoolVarP(&UseDefaultExcludes, "useDefaultExcludes", "", false, "If selected, use the default excludes from Ant - see here for the list. Defaults to true.")

	err := cobra.MarkFlagRequired(cmd.Flags(), "name")
	if err != nil {
		panic(err)
	}

	return cmd
}
