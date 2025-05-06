package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	versionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	authorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#F5F5F5"))
)

type HelpModel struct {
	Version string
	Author  string
}

func NewHelpModel(version string, author string) HelpModel {
	return HelpModel{
		Version: version,
		Author:  author,
	}
}

func (m HelpModel) Init() tea.Cmd {
	return nil
}

func (m HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m HelpModel) View() string {

	version := versionStyle.Render(m.Version)
	author := authorStyle.Render("by " + m.Author)

	title := titleStyle.Render("✭ SPELL CLI")

	return "" +
		"\n" + title + " " + version +
		"\n" + author + "\n\n"
}
