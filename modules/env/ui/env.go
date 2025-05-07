package ui

import (
	"github.com/charmbracelet/bubbletea"
	"spell/ui"
	"strings"
)

type EnvModel struct {
	File      string
	Key       string
	Value     string
	Variables []string // New field for a list of variables
}

func NewEnvModel(file, key, value string) EnvModel {
	return EnvModel{
		File:  file,
		Key:   key,
		Value: value,
	}
}

func NewEnvListModel(file string, variables []string) EnvModel {
	return EnvModel{
		File:      file,
		Variables: variables,
	}
}

func (m EnvModel) Init() tea.Cmd {
	return nil
}

func (m EnvModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m EnvModel) View() string {
	title := ui.TitleStyle.Render("✭ SPELL CLI")
	if len(m.Variables) > 0 {
		var styledVars []string
		for _, envVar := range m.Variables {
			parts := strings.SplitN(envVar, "=", 2)
			if len(parts) == 2 {
				key := ui.KeyStyle.Render(parts[0])
				value := ui.ValueStyle.Render(parts[1])
				styledVars = append(styledVars, key+ui.EqualsStyle.Render("=")+value)
			}
		}
		varList := strings.Join(styledVars, "\n")

		return "" +
			"\n" + title +
			"\n" +
			"File: " + m.File +
			"\n\n" +
			ui.EnvTitle.Render("Environment Variables") + "\n" + varList + "\n"
	}

	// Render a single key-value pair
	key := ui.KeyStyle.Render(m.Key)
	value := ui.ValueStyle.Render(m.Value)
	return "" +
		"\n" + title +
		"\n" +
		"File: " + m.File + "\n \n" +
		ui.EnvTitle.Render("Environment Variable") + "\n" +
		key + ui.EqualsStyle.Render("=") + value + "\n"
}
