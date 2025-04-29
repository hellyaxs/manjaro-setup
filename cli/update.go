package cli

import (
	"slices"
	"time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
)

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}


// Main update function.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	if !m.Chosen {
		return updateChoices(msg, m)
	}else{
		nextView := viewState{
			options: options,
			SelectedChoices: []int{},
		}
		switch m.Choice {
		    case 0: // instalação padrão (manjaro install apps)
				m.viewState = append(m.viewState, nextView)
			case 1: // instalação cli (manjaro install cli)
				m.viewState = append(m.viewState, nextView)
			case 2: // instalação web/dev (manjaro install web)
				m.viewState = append(m.viewState, nextView)	
		}
		m.indexState++
		return updateChosen(msg, m)
	}
}


// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > len(m.viewState[m.indexState].options)-1 {
				m.Choice = 0
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = len(m.viewState[m.indexState].options)-1
			}
		case " ":
			if slices.Contains(m.viewState[m.indexState].SelectedChoices, m.Choice) {
				m.viewState[m.indexState].SelectedChoices = removeElement(m.viewState[m.indexState].SelectedChoices,m.Choice)
			} else {
				m.viewState[m.indexState].SelectedChoices = append(m.viewState[m.indexState].SelectedChoices, m.Choice)
			}
		case "tab":
			if(m.indexState < len(m.viewState)){
				m.indexState++
				return m, frame()
			}
		case "shift+tab":
			if(m.indexState > 0){
				m.indexState--
				return m, frame()
			}	
		case "enter":
			m.Chosen = true
			return m, frame()
		}
	}

	return m, nil
}

// Update loop for the second view after a choice has been made
func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		if !m.Loaded {
			m.Frames++
			m.Progress = ease.OutBounce(float64(m.Frames) / float64(100))
			if m.Progress >= 1 {
				m.Progress = 1
				m.Loaded = true
				m.Ticks = 3
				return m, tick()
			}
			return m, frame()
		}

	case tickMsg:
		if m.Loaded {
			if m.Ticks == 0 {
				// m.Quitting = true
				m.Chosen = false
				m.Loaded = false
				return m, frame()
			}
			m.Ticks--
			return m, tick()
		}
	}

	return m, nil
}



  func removeElement(slice []int, element int) []int {
    index := -1
    for i, v := range slice {
        if v == element {
            index = i
            break
        }
    }
    if index == -1 {
        // Elemento não encontrado, retornar a lista original
        return slice
    }
    return append(slice[:index], slice[index+1:]...)
}