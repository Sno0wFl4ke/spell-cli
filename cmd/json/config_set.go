package json

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"spell/ui"
	"strings"
)

var configSetCmd = &cobra.Command{
	Use:     "write <file> <key> <value>",
	Short:   "Write a value in a JSON file",
	Aliases: []string{"set", "w"},
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		jsonFilePath := args[0]
		key := args[1]
		value := args[2]

		jsonFile, err := os.ReadFile(jsonFilePath)
		if err != nil {
			model := ui.NewConfigJsonModel("error", jsonFilePath, "Failed to read file: "+err.Error())
			fmt.Println(model.View())
			return
		}

		var jsonData map[string]interface{}
		if err := json.Unmarshal(jsonFile, &jsonData); err != nil {
			model := ui.NewConfigJsonModel("error", jsonFilePath, "Error unmarshaling JSON: "+err.Error())
			fmt.Println(model.View())
			return
		}

		keyParts := strings.Split(key, ".")
		current := jsonData

		for i, part := range keyParts {
			if i == len(keyParts)-1 {
				current[part] = value
			} else {
				if next, ok := current[part].(map[string]interface{}); ok {
					current = next
				} else {
					newMap := make(map[string]interface{})
					current[part] = newMap
					current = newMap
				}
			}
		}

		updatedJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			model := ui.NewConfigJsonModel("error", jsonFilePath, "Error marshaling updated JSON: "+err.Error())
			fmt.Println(model.View())
			return
		}

		if err := os.WriteFile(jsonFilePath, updatedJSON, 0644); err != nil {
			model := ui.NewConfigJsonModel("error", jsonFilePath, "Failed to write updated JSON to file: "+err.Error())
			fmt.Println(model.View())
			return
		}

		model := ui.NewConfigJsonModel("success", jsonFilePath, fmt.Sprintf("Successfully set key '%s' to value '%s'", key, value))
		fmt.Println(model.View())
	},
}

func init() {
	ModuleCmd.AddCommand(configSetCmd)
}
