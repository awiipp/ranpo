package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/internal/store"
	"github.com/awiipp/ranpo/pkg/models"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage environments",
}

var envUseCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Switch the active environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		cfg.ActiveEnv = args[0]
		if err := config.Save(cfg); err != nil {
			return err
		}

		fmt.Printf("active environment → %s\n", args[0])
		return nil
	},
}

var envSetCmd = &cobra.Command{
	Use:   "set <env> <key> <value>",
	Short: "Set a variable in an environment",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		envName, key, value := args[0], args[1], args[2]

		env, _ := store.LoadEnv(envName)
		if env == nil {
			env = &models.Environment{
				Name:      envName,
				Variables: map[string]string{},
			}
		}

		env.Variables[key] = value

		if err := store.SaveEnv(env); err != nil {
			return err
		}

		fmt.Printf("[%s] %s = %s\n", envName, key, value)
		return nil
	},
}

var envShowCmd = &cobra.Command{
	Use:   "show [name]",
	Short: "Show environment variables",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _ := config.Load()
		name := cfg.ActiveEnv

		env, err := store.LoadEnv(name)
		if err != nil {
			return fmt.Errorf("environment %q not found", name)
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(env)
	},
}

var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environments",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, _ := config.Load()
		names, _ := store.ListEnvs()

		if len(names) == 0 {
			fmt.Println("no environment found")
			return nil
		}

		for _, name := range names {
			active := ""

			if name == cfg.ActiveEnv {
				active = "  ← active"
			}

			fmt.Printf("  %s%s\n", name, active)
		}

		return nil
	},
}

var envDeleteCmd = &cobra.Command{
	Use:   "delete <env> <key>",
	Short: "Delete a variable from an environment",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		env, err := store.LoadEnv(args[0])
		if err != nil {
			return fmt.Errorf("environment %q not found", args[0])
		}

		delete(env.Variables, args[1])

		if err := store.SaveEnv(env); err != nil {
			return err
		}

		fmt.Printf("deleted %s from [%s]\n", args[1], args[0])
		return nil
	},
}

func init() {
	envCmd.AddCommand(envUseCmd, envShowCmd, envSetCmd, envListCmd, envDeleteCmd)
	rootCmd.AddCommand(envCmd)
}
