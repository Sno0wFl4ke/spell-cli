package db

import "github.com/spf13/cobra"

var ModuleCmd = &cobra.Command{
	Use:   "db",
	Short: "The db module of Spell",
	Long:  "The db module of Spell is used to manage the database.",
}
