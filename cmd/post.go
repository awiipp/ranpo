package cmd

import (
	"github.com/awiipp/ranpo/internal/tui"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post <url>",
	Short: "Send a POST request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagInteractive {
			return tui.LaunchWithRequest("POST", args[0])
		}
		return executeRequest("POST", args[0], flagBody)
	},
}

func init() {
	postCmd.Flags().StringVarP(&flagBody, "body", "b", "", "JSON request body")
	postCmd.Flags().BoolVarP(&flagInteractive, "interactive", "i", false, "Launch interactive TUI form")
	rootCmd.AddCommand(postCmd)
}
