package modules

import (
	"fmt"
	"github.com/spf13/cobra"
	"spell/modules/api"
	"spell/modules/db"
	"spell/modules/env"
	"spell/modules/json"
	"spell/ui"
)
import _ "spell/modules/api"

var rootCmd = &cobra.Command{
	Use:   "spell",
	Short: "Spell is a simple tool for dev operations",
	Run: func(cmd *cobra.Command, args []string) {
		model := ui.NewMainModel("1.0.0", "Philip Langenbrink")
		fmt.Print(model.View())

	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: "Help about any command",
		Long:  "Help about any command",
		Run: func(cmd *cobra.Command, args []string) {
			model := ui.NewHelpModel("1.0.0", "Philip Langenbrink")
			fmt.Print(model.View())
			err := cmd.Help()
			if err != nil {
				return
			}

		},
	})

	rootCmd.AddCommand(api.ModuleCmd)
	rootCmd.AddCommand(json.ModuleCmd)
	rootCmd.AddCommand(env.ModuleCmd)
	rootCmd.AddCommand(db.ModuleCmd)
}
