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
)

var apiGetCmd = &cobra.Command{
	Use:   "GET <url>",
	Short: "Make a GET request to the specified URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		resp, err := http.Get(url)
		if err != nil {
			cmd.PrintErr("Error making GET request:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var prettyJSON interface{}
		err = json.Unmarshal(body, &prettyJSON)
		var formatted string
		if err != nil {
			formatted = string(body)
		} else {
			bytes, _ := json.MarshalIndent(prettyJSON, "", "  ")
			formatted = string(bytes)
		}

		status := fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		model := ui.NewAPIModel(url, status, formatted)

		if err := tea.NewProgram(model).Start(); err != nil {
			cmd.PrintErrln("UI Error:", err)
		}
	},
}

func init() {
	ModuleCmd.AddCommand(apiGetCmd)
}
