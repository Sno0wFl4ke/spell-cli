package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "spell",
	Short: "Spell is a simple is dev tool for api operations",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
