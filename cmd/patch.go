package cmd

import "github.com/spf13/cobra"

var patchCmd = &cobra.Command{
	Use:   "patch <url>",
	Short: "Send a PATCH request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeRequest("PATCH", args[0], flagBody)
	},
}

func init() {
	patchCmd.Flags().StringVarP(&flagBody, "body", "b", "", "JSON request body")
	rootCmd.AddCommand(patchCmd)
}
