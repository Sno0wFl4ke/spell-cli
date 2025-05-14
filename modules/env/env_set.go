package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
	"spell/modules/env/ui"
)

var envSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set environment variables",
	Long:  "Set environment variables for the current shell session.",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		key := args[1]
		value := args[2]

		err := godotenv.Load(filePath)
		if err != nil {
			model := ui.NewEnvModel(filePath, key, "Error loading .env file")
			fmt.Println(model.View())
			return
		}

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			model := ui.NewEnvModel(filePath, key, "Error opening .env file")
			fmt.Println(model.View())
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "%s=%s\n", key, value)
		if err != nil {
			model := ui.NewEnvModel(filePath, key, "Error writing to .env file")
			fmt.Println(model.View())
			return
		}

		model := ui.NewEnvModel(filePath, key, value)
		fmt.Println(model.View())
	},
}

func init() {
	ModuleCmd.AddCommand(envSetCmd)
}
