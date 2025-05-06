package json

import (
	"fmt"
	"github.com/spf13/cobra"
	"spell/ui"
)

var ModuleCmd = &cobra.Command{
	Use:   "json",
	Short: "The json module of Spell",
	Long:  "The json module of Spell is used to manage the json configs",
	Run: func(cmd *cobra.Command, args []string) {
		commands := []string{"read <file>", "write <file> <key> <value>"}
		module := ui.NewModuleModel("JSON", commands)

		// Print the module's information
		fmt.Println(module.View())
	},
}
