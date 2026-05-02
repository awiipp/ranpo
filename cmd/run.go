package cmd

import (
	"fmt"
	"strings"

	"github.com/awiipp/ranpo/internal/client"
	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/internal/renderer"
	"github.com/awiipp/ranpo/internal/store"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <request> or run <collection>/<request>",
	Short: "Run a saved request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		colName := "default"
		reqName := args[0]

		if parts := strings.SplitN(args[0], "/", 2); len(parts) == 2 {
			colName, reqName = parts[0], parts[1]
		}

		col, err := store.LoadCollection(colName)
		if err != nil {
			return fmt.Errorf("collection %q not found", colName)
		}

		for _, req := range col.Requests {
			if req.Name == reqName {
				cfg, _ := config.Load()
				env, _ := store.LoadEnv(cfg.ActiveEnv)
				resp, err := client.Execute(&req, env)

				if err != nil {
					return err
				}

				fmt.Print(renderer.RenderResponse(resp.StatusCode, resp.Status, resp.Body, resp.Duration))
				return nil
			}
		}

		return fmt.Errorf("request %q not found in collection %q", reqName, colName)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
