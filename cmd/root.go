package cmd

import (
	"fmt"
	"os"

	"github.com/awiipp/ranpo/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ranpo",
	Short: "ranpo — TUI API testing tool",
	Long: `ranpo is a terminal API client.

Run without arguments to launch the alternate TUI.
Use subcommands for quick one-off requests or scripting.`,
	// No args -> lauch TUI
	RunE: func(cmd *cobra.Command, args []string) error {
		return tui.Launch()
	},
}

// The entry point called from main.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Shared flags
var (
	flagToken       string
	flagHeaders     []string
	flagSave        string
	flagCollection  string
	flagInteractive bool
	flagBody        string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&flagToken, "token", "t", "", "Bearer token (overrides env TOKEN")
	rootCmd.PersistentFlags().StringArrayVarP(&flagHeaders, "header", "H", nil, "Header in Key:Value format (repeatable)")
	rootCmd.PersistentFlags().StringVarP(&flagSave, "save", "s", "", "Save request with this name")
	rootCmd.PersistentFlags().StringVarP(&flagCollection, "collection", "c", "default", "Collection to save into")
}
