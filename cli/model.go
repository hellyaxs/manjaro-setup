package cli

import (
	"setup/cli/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "fmt"
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
    checkboxStyleSelected = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
    progressEmpty = subtleStyle.Render(progressEmptyChar)
    dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
    mainStyle     = lipgloss.NewStyle().MarginLeft(2)
    locationScripts = [...]string{
        "../../scripts/manjaro_install_web_dev.sh",
        "../../scripts/manjaro_install_cli.sh",
        "../../scripts/manjaro_install_apps.sh",
        "no file",
    }
    ramp = utils.MakeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

type (
    tickMsg  struct{}
    frameMsg struct{}
)

type model struct {
    indexState      int
    Choice          int
	Chosen          bool
    Ticks           int
    Frames          int
    Progress        float64
    Loaded          bool
    Quitting        bool
    viewState       []viewState
    currentView     int
}

func (m model) Init() tea.Cmd {
    return tick()
}

type viewState struct {
    SelectedChoices []int
    options []string
    title string
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//     switch msg := msg.(type) {
//     case tea.KeyMsg:
//         switch msg.String() {
//         case "ctrl+c", "q":
//             m.Quitting = true
//             return m, tea.Quit
//         case "up", "k":
//             if m.Choice > 0 {
//                 m.Choice--
//             }
//         case "down", "j":
//             if m.Choice < len(m.viewState[m.currentView].options)-1 {
//                 m.Choice++
//             }
//         case "enter", " ":
//             m.Chosen = true
//             // Adiciona nova tela baseada na escolha
//             if m.currentView == 0 {
//                 newView := viewState{
//                     options: options,
//                     SelectedChoices: []int{},
//                     title: "Selecione os pacotes para instalar",
//                 }
//                 m.viewState = append(m.viewState, newView)
//                 m.currentView++
//                 m.Choice = 0
//             }
//         case "esc":
//             // Volta para a tela anterior
//             if m.currentView > 0 {
//                 m.currentView--
//                 m.Choice = 0
//             }
//         }
//     case tickMsg:
//         m.Ticks++
//         return m, tick()
//     case frameMsg:
//         m.Frames++
//         return m, frame()
//     }
//     return m, nil
// }

// func (m model) View() string {
//     if m.Quitting {
//         return "Até logo!\n"
//     }

//     s := ""
//     if m.currentView == 0 {
//         s += "Selecione o tipo de instalação:\n\n"
//     } else {
//         s += m.viewState[m.currentView].title + "\n\n"
//     }

//     for i, choice := range m.viewState[m.currentView].options {
//         cursor := " "
//         if m.Choice == i {
//             cursor = ">"
//         }

//         checked := " "
//         if contains(m.viewState[m.currentView].SelectedChoices, i) {
//             checked = "x"
//         }

//         s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
//     }

//     s += "\nPressione q para sair, esc para voltar\n"
//     return s
// }

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

