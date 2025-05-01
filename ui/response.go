package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle         = lipgloss.NewStyle().Background(lipgloss.Color("#FF4081")).Foreground(lipgloss.Color("#FFFFFF")).Bold(false).Padding(0, 1)
	statusOKStyle      = lipgloss.NewStyle().Background(lipgloss.Color("#00E676")).Padding(0, 1)
	statusErrorStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#FF3D00")).Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF"))
	statusPendingStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FFC400")).Padding(0, 1)
	responseTitle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4081")).Bold(true)
)

type APIModel struct {
	URL       string
	Status    string
	Formatted string
}

func NewAPIModel(url, status, formatted string) APIModel {
	return APIModel{
		URL:       url,
		Status:    status,
		Formatted: formatted,
	}
}

func (m APIModel) Init() tea.Cmd {
	return nil
}

func (m APIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// qhen preccing q or ctrl+c, exit the program
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m APIModel) View() string {

	status := ""
	title := titleStyle.Render("✭ SPELL CLI")
	response := responseTitle.Render("Response")

	if m.Status == "200 OK" {
		status += statusOKStyle.Render(m.Status)

	}
	if m.Status == "404 Not Found" {
		status += statusErrorStyle.Render(m.Status)
	}
	if m.Status == "500 Internal Server Error" {
		status += statusErrorStyle.Render(m.Status)
	}
	if m.Status == "Pending" {
		status += statusPendingStyle.Render(m.Status)
	}
	if m.Status == "Error" {
		status += statusErrorStyle.Render(m.Status)
	}
	if m.Status == "Success" {
		status += statusOKStyle.Render(m.Status)
	}

	return "" +
		"\n" + title +
		"\n\n" +
		status + " " + m.URL +
		"\n\n" + response + "\n" + m.Formatted + "\n"
}
