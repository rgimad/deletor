package clean

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"github.com/pashkov256/deletor/internal/tui/interfaces"
	"github.com/pashkov256/deletor/internal/tui/options"
	"github.com/pashkov256/deletor/internal/tui/styles"
)

// Define options in fixed order

type OptionsTab struct {
	model interfaces.CleanModel
}

// func (t *OptionsTab) View() string {
// 	var content strings.Builder

// 	for i, name := range options.DefaultCleanOption {
// 		optionStyle := styles.OptionStyle
// 		if t.model.GetFocusedElement() == fmt.Sprintf("clean_option_%d", i+1) {
// 			optionStyle = styles.OptionFocusedStyle
// 		} else {
// 			if t.model.GetOptionState()[name] {
// 				optionStyle = styles.SelectedOptionStyle
// 			}
// 		}

// 		emoji := options.GetEmojiByCleanOption(name)

// 		content.WriteString(zone.Mark(fmt.Sprintf("clean_option_%d", i+1), optionStyle.Render(fmt.Sprintf("[%s] %s %-20s", map[bool]string{true: "✓", false: "○"}[t.model.GetOptionState()[name]], emoji, name))))

// 		content.WriteString("\n")
// 	}

//		return content.String()
//	}

func (t *OptionsTab) View() string {
	var content strings.Builder

	for i, name := range options.DefaultCleanOption {
		content.WriteString(t.renderSimpleOption(i+1, name))
		content.WriteString("\n")
		if name == options.SecureDeleteFiles && t.model.GetOptionState()[options.SecureDeleteFiles] {
			content.WriteString(t.renderSecureDeletionSubmenu())
		}
	}

	return content.String()
}

func (t *OptionsTab) renderSimpleOption(index int, name string) string {
	style := styles.OptionStyle
	if t.model.GetFocusedElement() == fmt.Sprintf("clean_option_%d", index) {
		style = styles.OptionFocusedStyle
	} else if t.model.GetOptionState()[name] {
		style = styles.SelectedOptionStyle
	}

	emoji := options.GetEmojiByCleanOption(name)

	return zone.Mark(fmt.Sprintf("clean_option_%d", index),
		style.Render(fmt.Sprintf("[%s] %s %-20s",
			map[bool]string{true: "✓", false: "○"}[t.model.GetOptionState()[name]],
			emoji, name)))
}

func (t *OptionsTab) renderSecureDeletionSubmenu() string {
	var builder strings.Builder

	for i, algo := range options.SecureDeletionAlgos {

		subStyle := styles.SubOptionStyle
		if t.model.GetFocusedElement() == fmt.Sprintf("secure_deletion_algo_%d", i) {
			subStyle = styles.SubOptionFocusedStyle
		} else if algo == t.model.GetSecureDeletionAlgo() {
			subStyle = styles.SubOptionSelectedStyle
		}

		radio := "○"
		if algo == t.model.GetSecureDeletionAlgo() {
			radio = "●"
		}

		builder.WriteString(zone.Mark(fmt.Sprintf("secure_deletion_algo_%d", i),
			subStyle.Render(fmt.Sprintf("    [%s] %s", radio, algo))))
		builder.WriteString("\n")

	}

	return builder.String()
}

func (t *OptionsTab) Init() tea.Cmd {
	return nil
}

func (t *OptionsTab) Update(msg tea.Msg) tea.Cmd {
	return nil
}
