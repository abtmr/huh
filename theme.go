package huh

import "github.com/charmbracelet/lipgloss"

// Theme is the style for a form.
type Theme struct {
	Form      lipgloss.Style
	Group     lipgloss.Style
	Unfocused FieldStyles
	Focused   FieldStyles
}

// FieldStyles are the styles for input fields
type FieldStyles struct {
	Base        lipgloss.Style
	Title       lipgloss.Style
	Description lipgloss.Style
	Selector    lipgloss.Style // Selection indicator in selects and multi-selects
	Option      lipgloss.Style // Select options
	Cursor      lipgloss.Style // Cursor in textinputs and textareas
	Placeholder lipgloss.Style
}

// NewBaseTheme returns a new base theme with general styles to be inherited by
// other themes.
func NewBaseTheme() *Theme {
	var t Theme

	t.Unfocused = FieldStyles{
		Base: lipgloss.NewStyle().
			PaddingLeft(1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true),
	}

	t.Focused.Base = t.Unfocused.Base.Copy().
		BorderStyle(lipgloss.HiddenBorder())

	return &t
}