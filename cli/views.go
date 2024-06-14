package cli

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	options = []string{
	"1. instalação padrão (manjaro install apps) tesnaddo",
	"2. instalação cli (manjaro install cli) 1",
	"3. instalação web/dev (manjaro install web) 1",
	"4. instalação servidor (no working)",
})

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
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("1. instalação padrão (manjaro install apps)", c == 0),
		checkbox("2. instalação cli (manjaro install cli)", c == 1),
		checkbox("3. instalação web/dev (manjaro install web)", c == 2),
		checkbox("4. instalação servidor (no working)", c == 3),
	)

	return fmt.Sprintf(tpl, choices)
}

// The second view, after a task has been chosen
func chosenView(m model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("preparando setup para [instalação de apps]?\n\nCool, we'll need %s and %s...", keywordStyle.Render("libgarden"), keywordStyle.Render("vegeutils"))
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
		ExecScripts(locationScripts[m.Choice])
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.Progress) + "%"
}

func checkbox(label string, checked bool) string {
	if checked {
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