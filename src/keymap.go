package src

import "github.com/charmbracelet/bubbles/key"

// pre-defined keymaps for the TUI
type keyMap struct {
	Esc       key.Binding
	Tab       key.Binding
	List1     key.Binding
	List2     key.Binding
	Table     key.Binding
	InfoBox   key.Binding
	Enter     key.Binding
	Input     key.Binding
	CtrlTab   key.Binding
	ToggleBox key.Binding
	Help      key.Binding
}

/* newKeyMap
 * ---------
 * returns a keyMap consisting of keybinds
 * used a shared function but for the TUI
 */
func newKeyMap() keyMap {
	return keyMap{
		Input: key.NewBinding(
			key.WithKeys("!"),
			key.WithHelp("!", "focus input"),
		),
		List1: key.NewBinding(
			key.WithKeys("#"),
			key.WithHelp("#", "focus list one"),
		),
		List2: key.NewBinding(
			key.WithKeys("$"),
			key.WithHelp("$", "focus list two"),
		),
		Table: key.NewBinding(
			key.WithKeys("@"),
			key.WithHelp("@", "focus table"),
		),
		InfoBox: key.NewBinding(
			key.WithKeys("%"),
			key.WithHelp("%", "focus info box"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Action"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "quit"),
		),
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch tabs forward"),
		),
		CtrlTab: key.NewBinding(
			key.WithKeys("ctrl+tab"),
			key.WithHelp("ctrl+tab", "switch tabs backward"),
		),
		ToggleBox: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "toggle download/info box"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "show/hide help menu"),
		),
	}
}
