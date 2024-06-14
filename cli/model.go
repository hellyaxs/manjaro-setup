package cli

import (
	"github.com/charmbracelet/lipgloss"
	"setup/cli/utils"
    tea "github.com/charmbracelet/bubbletea"
)

const (
    progressBarWidth  = 71
    progressFullChar  = "█"
    progressEmptyChar = "░"
    dotChar           = " • "
)

var (
    keywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
    subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
    ticksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
    checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
    progressEmpty = subtleStyle.Render(progressEmptyChar)
    dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
    mainStyle     = lipgloss.NewStyle().MarginLeft(2)
    locationScripts = [...]string{
        "../scripts/manjaro_install_web_dev.sh",
        "../scripts/manjaro_install_cli.sh",
        "../scripts/manjaro_install_apps.sh",
        "no file",
    }
    ramp = utils.MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

type (
    tickMsg  struct{}
    frameMsg struct{}
)

type model struct {
    Choice          int
    SelectedChoices map[int]bool
	Chosen          bool
    Ticks           int
    Frames          int
    Progress        float64
    Loaded          bool
    Quitting        bool
}

func (m model) Init() tea.Cmd {
    return tick()
}

