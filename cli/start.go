package cli

import(
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)
var (
	options = []string{
	"* asdf *",
	"* docker *",
	"* docker-compose *",
	"* git *",
	"* neovim *",
	"* postgres *",
	"* teamviewer *",
	"* tmux *",
	"* zsh *(disabled)*",
    "* htop *",
    "* nvtop *",
    "* neofetch *",
    "* tldr *",
	"* wget *",
	"* curl *",
	"* bat *",
	"* exa *",
	"* fd *",
	"* thefuck *",
    "* fzf *",
	}
    optionsInit = []string{
		"1. instalação padrão (manjaro install apps)",
		"2. instalação cli (manjaro install cli)",
		"3. instalação web/dev (manjaro install web)",
		"4. instalação servidor (no working)",

})

func StartViews() {
    initialModel := model{ Choice: 0, 
		SelectedChoices: make(map[int]bool),
		Ticks: 10, 
		Frames: 0, 
		Progress: 0, 
		Loaded: false, 
		Quitting: false,
		options:optionsInit,
		}
    p := tea.NewProgram(initialModel)
    if _, err := p.Run(); err != nil {
        fmt.Println("could not start program:", err)
    }
	fmt.Print("printando valor depos da execuçaõ :::",initialModel)
	

	// initialModel2 := model{ Choice: 0, 
	// 	SelectedChoices: make(map[int]bool),
	// 	Ticks: 10, 
	// 	Frames: 0, 
	// 	Progress: 0, 
	// 	Loaded: false, 
	// 	Quitting: false,
	// 	options:options,
	// 	}
    // q := tea.NewProgram(initialModel2)
    // if _, err := q.Run(); err != nil {
    //     fmt.Println("could not start program:", err)
    // }
}
