package app

func (m *Model) moveCursor(delta int) {
	m.Cursor += delta
	if m.Cursor < 0 {
		m.Cursor = 0
	}
	max := len(Menu) - 1
	if m.Cursor > max {
		m.Cursor = max
	}
}

func (m *Model) selectPage(id PageID) {
	if !id.Valid() {
		return
	}
	m.ActivePage = id
	m.Cursor = int(id)
}

func (m *Model) openCursorPage() {
	m.ActivePage = PageID(m.Cursor)
}
