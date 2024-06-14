package cli

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"slices"
	// tea "github.com/charmbracelet/bubbletea"
)

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n  Instalação concluida até a proxima!\n\n"
	}
	if !m.Chosen {
		s = choicesView(m)
	} else {
		s = chosenView(m)
	}
	
	return mainStyle.Render("\n" + s + "\n\n")
}

// The first view, where you're choosing a task
func choicesView(m model) string {
	
	c := m.Choice
    tpl := "Qual setup você deseja ?\n\n"
    tpl += "%s\n\n"
    tpl += ""
    tpl += subtleStyle.Render("UpArrow/DownArrow, up/down: select") + dotStyle +
        subtleStyle.Render("space: select/deselect") + dotStyle +
        subtleStyle.Render("enter: confirm") + dotStyle +
        subtleStyle.Render("q, esc: quit")

    var choicesBuilder strings.Builder

    for i, option := range m.options {
        choicesBuilder.WriteString(checkbox(option,  slices.Contains(m.SelectedChoices, i) ,slices.Contains(m.SelectedChoices, i) || c == i))
        if i < len(options)-1 {
            choicesBuilder.WriteString("\n")
        }
    }

    choices := choicesBuilder.String()

    return fmt.Sprintf(tpl, choices)
}

// The second view, after a task has been chosen
func chosenView(m model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("preparando setup para [instalação de apps]?\n\nOkay, carregando arquivo %s e %s...", keywordStyle.Render("integrando lista"), keywordStyle.Render("pacotes"))
	case 1:
		msg = fmt.Sprintf("preparando setup para [instalação de cli]?\n\nOkay, then we should install %s and %s...", keywordStyle.Render("marketkit"), keywordStyle.Render("libshopping"))
	case 2:
		msg = fmt.Sprintf("preparando setup para [instalação de web/dev]?\n\nOkay, cool, then we’ll need a library. Yes, an %s.", keywordStyle.Render("actual library"))
	default:
		msg = fmt.Sprintf("It’s always good to see friends.\n\nFetching %s and %s...", keywordStyle.Render("social-skills"), keywordStyle.Render("conversationutils"))
	}

	label := "preparando scripts..."
	if m.Loaded {
		label = fmt.Sprintf("Downloaded. Exiting in %s seconds...", ticksStyle.Render(strconv.Itoa(m.Ticks)))
	
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.Progress) + "%"
}

func checkbox(label string,isSelected bool, checked bool) string {
	if checked {
		if isSelected {
			return checkboxStyleSelected.Render("[x] " + label)
		}
		return checkboxStyle.Render("[x] " + label)
	}
	
	return fmt.Sprintf("[ ] %s", label)
}

func progressbar(percent float64) string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += ramp[i].Render(progressFullChar)
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}