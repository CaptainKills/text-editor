package main

import tea "charm.land/bubbletea/v2"

type Cursor struct {
	row    int
	column int
	offset int
}

func (c *Cursor) GetCursor() *tea.Cursor {
	return &tea.Cursor{
		Position: tea.Position{X: c.column + 7, Y: c.row - c.offset},
		Blink:    true,
	}
}

func (c *Cursor) MoveUp(buffer []string) {
	if len(buffer) == 0 {
		c.row = 0
		c.column = 0
		return
	}

	// Update Row
	min_row_index := 0
	next_row_index := c.row - 1

	c.row = max(next_row_index, min_row_index)

	// Update Column
	max_col_index := max(0, len(buffer[c.row])-1)
	next_col_index := c.column

	c.column = min(next_col_index, max_col_index)

	// Update Offset
	if c.row < c.offset {
		c.offset = c.row
	}
}

func (c *Cursor) MoveDown(buffer []string, height int) {
	if len(buffer) == 0 {
		c.row = 0
		c.column = 0
		return
	}

	// Update Row
	max_row_index := len(buffer) - 1
	next_row_index := c.row + 1

	c.row = min(next_row_index, max_row_index)

	// Update Column
	max_col_index := max(0, len(buffer[c.row])-1)
	next_col_index := c.column

	c.column = min(next_col_index, max_col_index)

	// Update Offset
	if c.row >= c.offset+height {
		c.offset = c.row - height + 1
	}
}

func (c *Cursor) MoveLeft(buffer []string) {
	if len(buffer) == 0 {
		c.row = 0
		c.column = 0
		return
	}

	// Update Column
	min_index := 0
	next_index := c.column - 1

	c.column = max(next_index, min_index)
}

func (c *Cursor) MoveRight(buffer []string) {
	if len(buffer) == 0 {
		c.row = 0
		c.column = 0
		return
	}

	// Update Column
	max_index := max(0, len(buffer[c.row])-1)
	next_index := c.column + 1

	c.column = min(next_index, max_index)
}
