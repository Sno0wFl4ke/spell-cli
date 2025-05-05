package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"spell/ui"
	"strings"

	"github.com/spf13/cobra"
)

var listEnvCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environment variables",
	Long:  "List all environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		err := godotenv.Load(file)
		if err != nil {
			model := ui.NewEnvModel(file, "", "Error loading .env file")
			fmt.Println(model.View())
			return
		}

		fileContent, err := os.ReadFile(file)
		if err != nil {
			model := ui.NewEnvModel(file, "", "Error reading .env file")
			fmt.Println(model.View())
			return
		}

		lines := strings.Split(string(fileContent), "\n")
		var envVars []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				if value, exists := os.LookupEnv(key); exists {
					envVars = append(envVars, fmt.Sprintf("%s=%s", key, value))
				}
			}
		}

		model := ui.NewEnvListModel(file, envVars)
		fmt.Println(model.View())
	},
}

func init() {
	ModuleCmd.AddCommand(listEnvCmd)
}
