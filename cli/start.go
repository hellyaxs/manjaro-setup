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
		"2. instalação cli (ferrmentas cli uteis produtivas)",
		"3. instalação web/dev (ferramentas para desenvolmento web)",
		"4. Otimizações do sistema (no working)",
	}
)

func StartViews() {
	// Estado inicial com a primeira tela
	initialViewState := viewState{
		options: optionsInit,
		SelectedChoices: []int{},
	}

	// Modelo inicial com histórico de telas
	initialModel := model{ 
		Choice: 0,
		indexState: 0, 
		Ticks: 10, 
		Frames: 0, 
		Progress: 0, 
		Loaded: false, 
		Quitting: false,
		viewState: []viewState{initialViewState},
	}

	// Inicia o programa TUI
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
