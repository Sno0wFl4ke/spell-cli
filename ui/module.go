package ui

import (
	"github.com/charmbracelet/bubbletea"
)

type ModuleModel struct {
	ModuleName string
	Commands   []string
}

func NewModuleModel(moduleName string, commands []string) ModuleModel {
	return ModuleModel{
		ModuleName: moduleName,
		Commands:   commands,
	}
}

func (m ModuleModel) Init() tea.Cmd {
	return nil
}

func (m ModuleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ModuleModel) View() string {

	title := TitleStyle.Render("✭ SPELL CLI")
	module := modelTitle.Render(m.ModuleName)
	commands := ""
	for _, command := range m.Commands {
		commands += "\n" + ResponseTitle.Render("• "+command)
	}

	return "" +
		"\n" + title + " " +
		"\n\n" +
		authorStyle.Render("Your simple tool for day to day dev ops") + "\n" +
		authorStyle.Render("Use") + commandStyle.Render(" spell <module> <command> <args>") +
		"\n\n" +
		module +
		commands +
		"\n"
}
