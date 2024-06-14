package cli

import(
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

func StartViews() {
    initialModel := model{Choice: 0, SelectedChoices: make(map[int]bool), Ticks: 10, Frames: 0, Progress: 0, Loaded: false, Quitting: false}
    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Println("could not start program:", err)
    }
}
