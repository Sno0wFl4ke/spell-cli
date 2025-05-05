package json

import "github.com/spf13/cobra"

var ModuleCmd = &cobra.Command{
	Use:   "json",
	Short: "The json module of Spell",
	Long:  "The json module of Spell is used to manage the json configs",
}
