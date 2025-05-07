package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle         = lipgloss.NewStyle().Background(lipgloss.Color("#F50057")).Foreground(lipgloss.Color("#FFFFFF")).Bold(false).Padding(0, 1)
	StatusOKStyle      = lipgloss.NewStyle().Background(lipgloss.Color("#00E676")).Padding(0, 1)
	StatusErrorStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#FF3D00")).Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF"))
	StatusPendingStyle = lipgloss.NewStyle().Background(lipgloss.Color("#FFC400")).Padding(0, 1)
	ResponseTitle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4081")).Bold(true)

	// General styles for the env module
	KeyStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4081")).Bold(true)
	ValueStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(false)
	EqualsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B0BEC5")).Bold(false)
	EnvTitle    = lipgloss.NewStyle().Background(lipgloss.Color("#448AFF")).Bold(true).Padding(0, 1)

	// General styles for the json module
	StatusReadJsonSuccessStyle = lipgloss.NewStyle().Background(lipgloss.Color("#00E676")).Padding(0, 1)
	StatusReadJsonErrorStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#FF3D00")).Padding(0, 1).Foreground(lipgloss.Color("#FFFFFF"))
	FormattedReadJsonStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F5F5F5")).Bold(true)
)
