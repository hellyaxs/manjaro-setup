package cli

import (
	"time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
	"slices"
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
		switch m.Choice {
		    case 0:
				m.options = options
		}
		return updateChosen(msg, m)
	}
}


// Sub-update functions

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > len(m.options)-1 {
				m.Choice = 0
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = len(m.options)-1
			}
		case " ":
			if slices.Contains(m.SelectedChoices, m.Choice) {
				m.SelectedChoices = removeElement(m.SelectedChoices,m.Choice)
			} else {
				m.SelectedChoices = append(m.SelectedChoices, m.Choice)
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