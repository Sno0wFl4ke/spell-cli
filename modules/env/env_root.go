package env

import (
	"fmt"
	"github.com/spf13/cobra"
	"spell/ui"
)

var ModuleCmd = &cobra.Command{
	Use:   "env",
	Short: "The env module of Spell",
	Long:  "The env module of Spell is used to manage the environment variables.",
	Run: func(cmd *cobra.Command, args []string) {
		commands := []string{"get <file> <key>", "list <file>", "set <file> <key> <value>"}
		module := ui.NewModuleModel("ENV", commands)

		// Print the module's information
		fmt.Println(module.View())
	},
}
