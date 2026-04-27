package cmd

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use: "get <url>",
	Short: "Send a GET request",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeRequest("GET", args[0], "")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}