package ui

import (
	"github.com/charmbracelet/bubbletea"
	"spell/ui"
)

type ConfigJsonModel struct {
	Status    string
	File      string
	Formatted string
}

func NewConfigJsonModel(status, file, formatted string) ConfigJsonModel {
	return ConfigJsonModel{
		Status:    status,
		File:      file,
		Formatted: formatted,
	}
}

func (m ConfigJsonModel) Init() tea.Cmd {
	return nil
}

func (m ConfigJsonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ConfigJsonModel) View() string {

	status := ""
	title := ui.TitleStyle.Render("✭ SPELL CLI")
	formatted := ui.FormattedReadJsonStyle.Render(m.Formatted)

	if m.Status == "success" {
		status += ui.StatusReadJsonSuccessStyle.Render(m.Status)
	}
	if m.Status == "error" {
		status += ui.StatusReadJsonErrorStyle.Render(m.Status)
	}

	return "" +
		"\n" + title +
		"\n\n" +
		status + " File " + m.File +
		"\n\n" + ui.ResponseTitle.Render("JSON") + "\n" + formatted + "\n"
}
