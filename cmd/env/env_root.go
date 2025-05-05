package env

import "github.com/spf13/cobra"

var ModuleCmd = &cobra.Command{
	Use:   "env",
	Short: "The env module of Spell",
	Long:  "The env module of Spell is used to manage the environment variables.",
}
