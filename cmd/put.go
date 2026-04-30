package cmd

import "github.com/spf13/cobra"

var putCmd = &cobra.Command{
	Use:   "put <url>",
	Short: "Send a PUT request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeRequest("PUT", args[0], flagBody)
	},
}

func init() {
	putCmd.Flags().StringVarP(&flagBody, "body", "b", "", "JSON request body")
	rootCmd.AddCommand(putCmd)
}
