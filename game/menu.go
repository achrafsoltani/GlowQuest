package game

type MenuOption struct {
	Label    string
	Disabled bool
}

type MenuState struct {
	Options       []MenuOption
	SelectedIndex int
}

func NewMenuState(hasSave bool) MenuState {
	return MenuState{
		Options: []MenuOption{
			{Label: "NEW GAME", Disabled: false},
			{Label: "CONTINUE", Disabled: !hasSave},
		},
		SelectedIndex: 0,
	}
}

func (m *MenuState) MoveUp() {
	for i := 0; i < len(m.Options); i++ {
		m.SelectedIndex--
		if m.SelectedIndex < 0 {
			m.SelectedIndex = len(m.Options) - 1
		}
		if !m.Options[m.SelectedIndex].Disabled {
			return
		}
	}
}

func (m *MenuState) MoveDown() {
	for i := 0; i < len(m.Options); i++ {
		m.SelectedIndex++
		if m.SelectedIndex >= len(m.Options) {
			m.SelectedIndex = 0
		}
		if !m.Options[m.SelectedIndex].Disabled {
			return
		}
	}
}
