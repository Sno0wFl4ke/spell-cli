package env

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
	"spell/ui"
)

var envGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of an environment variable",
	Run: func(cmd *cobra.Command, args []string) {
		var file = args[0]
		var key = args[1]

		err := godotenv.Load(file)
		if err != nil {
			model := ui.NewEnvModel(file, key, "Error loading .env file")
			fmt.Println(model.View())
			return
		}

		value := os.Getenv(key)
		if value == "" {
			model := ui.NewEnvModel(file, key, "Environment variable not found")
			fmt.Println(model.View())
		} else {
			model := ui.NewEnvModel(file, key, value)
			fmt.Println(model.View())
		}
	},
}

func init() {
	ModuleCmd.AddCommand(envGetCmd)
}
