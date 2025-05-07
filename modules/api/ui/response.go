package ui

import (
	"github.com/charmbracelet/bubbletea"
	"spell/ui"
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
	title := ui.TitleStyle.Render("✭ SPELL CLI")
	response := ui.ResponseTitle.Render("Response")

	if m.Status == "200 OK" {
		status += ui.StatusOKStyle.Render(m.Status)

	}
	if m.Status == "404 Not Found" {
		status += ui.StatusErrorStyle.Render(m.Status)
	}
	if m.Status == "500 Internal Server Error" {
		status += ui.StatusErrorStyle.Render(m.Status)
	}
	if m.Status == "Pending" {
		status += ui.StatusPendingStyle.Render(m.Status)
	}
	if m.Status == "Error" {
		status += ui.StatusErrorStyle.Render(m.Status)
	}
	if m.Status == "Success" {
		status += ui.StatusOKStyle.Render(m.Status)
	}

	return "" +
		"\n" + title +
		"\n\n" +
		status + " " + m.URL +
		"\n\n" + response + "\n" + m.Formatted + "\n"
}
