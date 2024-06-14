package cli

import (
	"fmt"
	"strings"
)

func ChoicesViewSetupCli(m model) string {
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
func ChoicesViewSetupApps(m model) string {
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

func ChoicesViewSetupWeb(m model) string {
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


func ChoicesViewSetupCliDe(m model) string {
    
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
        choicesBuilder.WriteString(checkbox(option, c == i))
        if i < len(options)-1 {
            choicesBuilder.WriteString("\n")
        }
    }

    choices := choicesBuilder.String()

    return fmt.Sprintf(tpl, choices)
}