package ui

import (
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	keyStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4081")).Bold(true)
	valueStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(false)
	equalsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B0BEC5")).Bold(false)
	envTitle    = lipgloss.NewStyle().Background(lipgloss.Color("#448AFF")).Bold(true).Padding(0, 1)
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
	title := titleStyle.Render("✭ SPELL CLI")
	if len(m.Variables) > 0 {
		var styledVars []string
		for _, envVar := range m.Variables {
			parts := strings.SplitN(envVar, "=", 2)
			if len(parts) == 2 {
				key := keyStyle.Render(parts[0])
				value := valueStyle.Render(parts[1])
				styledVars = append(styledVars, key+equalsStyle.Render("=")+value)
			}
		}
		varList := strings.Join(styledVars, "\n")

		return "" +
			"\n" + title +
			"\n" +
			"File: " + m.File +
			"\n\n" +
			envTitle.Render("Environment Variables") + "\n" + varList + "\n"
	}

	// Render a single key-value pair
	key := keyStyle.Render(m.Key)
	value := valueStyle.Render(m.Value)
	return "" +
		"\n" + title +
		"\n" +
		"File: " + m.File + "\n \n" +
		envTitle.Render("Environment Variable") + "\n" +
		key + equalsStyle.Render("=") + value + "\n"
}
