package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:     "delete <url>",
	Aliases: []string{"del"},
	Short:   "Send a DELETE request",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeRequest("DELETE", args[0], "")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
