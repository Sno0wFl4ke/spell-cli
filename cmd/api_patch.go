package cmd

import (
	"encoding/json"
	"fmt"
	_ "github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"spell/ui"
	"strings"
)

var apiPatchCmd = &cobra.Command{
	Use:   "PATCH <url> <body>",
	Short: "Make a PATCH request to the specified URL with a JSON body",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body := args[1]

		req, err := http.NewRequest("PATCH", url, strings.NewReader(body))
		if err != nil {
			cmd.PrintErr("Error creating PATCH request:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			cmd.PrintErr("Error making PATCH request:", err)
			return
		}
		defer resp.Body.Close()

		bodyResponse, _ := ioutil.ReadAll(resp.Body)

		var prettyJSON interface{}
		err = json.Unmarshal(bodyResponse, &prettyJSON)
		var formatted string
		if err != nil {
			formatted = string(bodyResponse)
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
	rootCmd.AddCommand(apiPatchCmd)
}
