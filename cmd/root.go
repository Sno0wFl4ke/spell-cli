package cmd

import (
	"github.com/spf13/cobra"
	"spell/cmd/api"
	"spell/cmd/json"
)
import _ "spell/cmd/api"

var rootCmd = &cobra.Command{
	Use:   "spell",
	Short: "Spell is a simple is dev tool for api_module operations",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(api.ModuleCmd)
	rootCmd.AddCommand(json.ModuleCmd)
}
