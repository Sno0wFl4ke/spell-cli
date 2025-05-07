package json

import (
	"encoding/json"
	"fmt"
	"os"
	"spell/ui"

	"github.com/spf13/cobra"
)

var configReadCmd = &cobra.Command{
	Use:     "read <file>",
	Short:   "Read a json file in terminal",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jsonFile, err := os.ReadFile(args[0])
		if err != nil {
			model := ui.NewConfigJsonModel("error", args[0], "Failed to read file: "+err.Error())
			fmt.Println(model.View())
			return
		}

		var jsonData interface{}
		if err := json.Unmarshal(jsonFile, &jsonData); err != nil {
			model := ui.NewConfigJsonModel("error", args[0], "Error unmarshaling JSON: "+err.Error())
			fmt.Println(model.View())
			return
		}

		formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			model := ui.NewConfigJsonModel("error", args[0], "Error formatting JSON: "+err.Error())
			fmt.Println(model.View())
			return
		}

		model := ui.NewConfigJsonModel("success", args[0], string(formattedJSON))
		fmt.Println(model.View())
	},
}

func init() {
	ModuleCmd.AddCommand(configReadCmd)
}
