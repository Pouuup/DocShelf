package tui

import (
	"strings"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/lipgloss"
)

func renderCode(text string) []string {
	var output strings.Builder

	var codeStyle = lipgloss.NewStyle().
		PaddingLeft(4).
		Border(lipgloss.RoundedBorder())

	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, text)
	if err != nil {
		return nil
	}

	formatter := formatters.Get("terminal16m")

	style := styles.Get("monokai")

	err = formatter.Format(&output, style, iterator)
	if err != nil {
		return nil
	}

	formattedCode := codeStyle.Render(output.String())
	totalLines := strings.Split(formattedCode, "\n")

	return totalLines

}
