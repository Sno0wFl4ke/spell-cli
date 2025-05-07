package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"spell/ui"
)

var ModuleCmd = &cobra.Command{
	Use:   "api",
	Short: "The API module of Spell",
	Run: func(cmd *cobra.Command, args []string) {
		commands := []string{"GET <url>", "POST <url> <json>", "PATCH <url> <json>", "DELETE <url>"}
		module := ui.NewModuleModel("API", commands)

		// Print the module's information
		fmt.Println(module.View())
	},
}
