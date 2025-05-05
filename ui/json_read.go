package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	statusReadJsonSuccessStyle = lipgloss.NewStyle().Background(lipgloss.Color("#00E676")).Padding(0, 1)
	statusReadJsonErrorStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#FF3D00")).Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF"))
	formattedReadJsonStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4081")).Bold(true)
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
	title := titleStyle.Render("✭ SPELL CLI")
	formatted := formattedReadJsonStyle.Render(m.Formatted)

	if m.Status == "success" {
		status += statusReadJsonSuccessStyle.Render(m.Status)
	}
	if m.Status == "error" {
		status += statusReadJsonErrorStyle.Render(m.Status)
	}

	return "" +
		"\n" + title +
		"\n\n" +
		status + " File " + m.File +
		"\n\n" + formatted + "\n"
}
