package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"spell/modules/api/ui"
	"strings"
)

var apiPostCmd = &cobra.Command{
	Use:   "POST <url> <body...>",
	Short: "Make a POST request to the specified URL with a JSON body",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body := strings.Join(args[1:], " ")

		var jsonBody map[string]interface{}
		if err := json.Unmarshal([]byte(body), &jsonBody); err != nil {
			cmd.PrintErrln("Invalid JSON body:", err)
			return
		}

		resp, err := http.Post(url, "application/json", strings.NewReader(body))
		if err != nil {
			cmd.PrintErr("Error making POST request:", err)
			return
		}
		defer resp.Body.Close()

		responseBody, _ := ioutil.ReadAll(resp.Body)

		var prettyJSON interface{}
		err = json.Unmarshal(responseBody, &prettyJSON)
		var formatted string
		if err != nil {
			formatted = string(responseBody)
		} else {
			formattedBytes, _ := json.MarshalIndent(prettyJSON, "", "  ")
			formatted = string(formattedBytes)
		}

		status := fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		model := ui.NewAPIModel(url, status, formatted)

		if err := tea.NewProgram(model).Start(); err != nil {
			cmd.PrintErrln("UI Error:", err)
		}
	},
}

func init() {
	ModuleCmd.AddCommand(apiPostCmd)
}
