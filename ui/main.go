package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	modelTitle   = lipgloss.NewStyle().Background(lipgloss.Color("#2979FF")).Padding(0, 1)
	modulesStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#29B6F6"))
	commandStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F50057"))
)

type MainModel struct {
	Version string
	Author  string
}

func NewMainModel(version string, author string) MainModel {
	return MainModel{
		Version: version,
		Author:  author,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m MainModel) View() string {

	version := versionStyle.Render(m.Version)
	author := authorStyle.Render("by " + m.Author)

	title := titleStyle.Render("✭ SPELL CLI")

	return "" +
		"\n" + title + " " + version +
		"\n" + author + "\n\n" +
		authorStyle.Render("Your simple tool for day to day dev ops") + "\n" +
		authorStyle.Render("Use") + commandStyle.Render(" spell <module> <command> <args>") +
		"\n\n" +
		modelTitle.Render("Modules") +
		"\n" + modulesStyle.Render("- api") +
		"\n" + modulesStyle.Render("- env") +
		"\n" + modulesStyle.Render("- json") +
		"\n" + modulesStyle.Render("- help") +
		"\n"
}
