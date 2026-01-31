package main

import (
	"testing"
)

var text = []string{
	"Hello World!",
	"    This line has spaces in front.",
	"A slightly shorter line",
	"This is the last line of text.",
}

func TestCursorMoveUp(t *testing.T) {
	t.Run("Row Available", func(t *testing.T) {
		c := Cursor{row: 1, column: 0}
		c.MoveUp(text)

		want := Cursor{row: 0, column: 0}

		if c.row != want.row {
			t.Errorf("Cursor row is not decremented correctly! got %v, want %v\n", c.row, want.row)
		}
	})

	t.Run("Row Unavailable", func(t *testing.T) {
		c := Cursor{row: len(text) - 1, column: 0}

		for i := c.row; i > 0; i-- {
			c.MoveUp(text)
		}

		want := Cursor{row: 0, column: 0}

		if c.row != want.row {
			t.Errorf("Cursor row is not satured to 0! got %v, want %v\n", c.row, want.row)
		}
	})

	t.Run("Column Available", func(t *testing.T) {
		c := Cursor{row: 1, column: 5}
		c.MoveUp(text)

		want := Cursor{row: 0, column: 5}

		if c.column != want.column {
			t.Errorf("Cursor column is not preserved correctly! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Unavailable", func(t *testing.T) {
		c := Cursor{row: 1, column: len(text[1]) - 1}
		c.MoveUp(text)

		want := Cursor{row: 0, column: len(text[0]) - 1}

		if c.column != want.column {
			t.Errorf("Cursor column is not adjusted to new max line length! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Empty", func(t *testing.T) {
		text := []string{
			"",
			"Hello World!",
		}

		c := Cursor{row: 1, column: 0}
		c.MoveUp(text)

		want := Cursor{row: 0, column: len(text[0])}

		if c.column != want.column {
			t.Errorf("Cursor column is not saturated 0! got %v, want %v\n", c.column, want.column)
		}
	})
}

func TestCursorMoveDown(t *testing.T) {
	t.Run("Row Available", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}
		c.MoveDown(text, len(text))

		want := Cursor{row: 1, column: 0}

		if c.row != want.row {
			t.Errorf("Cursor row is not incremented correctly! got %v, want %v\n", c.row, want.row)
		}
	})

	t.Run("Row Unavailable", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}

		for i := c.row; i < len(text); i++ {
			c.MoveDown(text, len(text))
		}

		want := Cursor{row: len(text) - 1, column: 0}

		if c.row != want.row {
			t.Errorf("Cursor row is not satured to max index! got %v, want %v\n", c.row, want.row)
		}
	})

	t.Run("Column Available", func(t *testing.T) {
		c := Cursor{row: 0, column: 5}
		c.MoveDown(text, len(text))

		want := Cursor{row: 1, column: 5}

		if c.column != want.column {
			t.Errorf("Cursor column is not preserved correctly! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Unavailable", func(t *testing.T) {
		c := Cursor{row: 1, column: len(text[1]) - 1}
		c.MoveDown(text, len(text))

		want := Cursor{row: 2, column: len(text[2]) - 1}

		if c.column != want.column {
			t.Errorf("Cursor column is not adjusted to new max line length! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Colunm Empty", func(t *testing.T) {
		text := []string{
			"Hello World!",
			"",
		}

		c := Cursor{row: 0, column: 0}
		c.MoveDown(text, len(text))

		want := Cursor{row: 1, column: len(text[1])}

		if c.column != want.column {
			t.Errorf("Cursor column is not saturated 0! got %v, want %v\n", c.column, want.column)
		}
	})
}

func TestCursorMoveLeft(t *testing.T) {
	t.Run("Column Available", func(t *testing.T) {
		c := Cursor{row: 0, column: 5}
		c.MoveLeft(text)

		want := Cursor{row: 0, column: 4}

		if c.column != want.column {
			t.Errorf("Cursor column is not decremented correctly! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Unavailable", func(t *testing.T) {
		c := Cursor{row: 0, column: len(text[0]) - 1}

		for i := c.column; i > 0; i-- {
			c.MoveLeft(text)
		}

		want := Cursor{row: 0, column: 0}

		if c.column != want.column {
			t.Errorf("Cursor column is not satured to 0! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Empty", func(t *testing.T) {
		text := []string{""}
		c := Cursor{row: 0, column: 0}
		c.MoveLeft(text)

		want := Cursor{row: 0, column: 0}

		if c.column != want.column {
			t.Errorf("Cursor column is not satured to 0! got %v, want %v\n", c.column, want.column)
		}
	})
}

func TestCursorMoveRight(t *testing.T) {
	t.Run("Column Available", func(t *testing.T) {
		c := Cursor{row: 0, column: 5}
		c.MoveRight(text)

		want := Cursor{row: 0, column: 6}

		if c.column != want.column {
			t.Errorf("Cursor column is not incremented correctly! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Unavailable", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}

		for i := c.column; i < len(text[0]); i++ {
			c.MoveRight(text)
		}

		want := Cursor{row: 0, column: len(text[0]) - 1}

		if c.column != want.column {
			t.Errorf("Cursor column is satured to max length! got %v, want %v\n", c.column, want.column)
		}
	})

	t.Run("Column Empty", func(t *testing.T) {
		text := []string{""}
		c := Cursor{row: 0, column: 0}
		c.MoveRight(text)

		want := Cursor{row: 0, column: 0}

		if c.column != want.column {
			t.Errorf("Cursor column is not satured to 0! got %v, want %v\n", c.column, want.column)
		}
	})
}

func TestCursorEmptyFile(t *testing.T) {
	text := []string{} // Simulate empty file

	t.Run("Move Up", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}
		c.MoveUp(text)

		want := Cursor{row: 0, column: 0}

		if c != want {
			t.Errorf("Cursor is unintentionally moving! got %v, want %v\n", c, want)
		}
	})

	t.Run("Move Down", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}
		c.MoveDown(text, len(text))

		want := Cursor{row: 0, column: 0}

		if c != want {
			t.Errorf("Cursor is unintentionally moving! got %v, want %v\n", c, want)
		}
	})

	t.Run("Move Left", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}
		c.MoveLeft(text)

		want := Cursor{row: 0, column: 0}

		if c != want {
			t.Errorf("Cursor is unintentionally moving! got %v, want %v\n", c, want)
		}
	})

	t.Run("Move Right", func(t *testing.T) {
		c := Cursor{row: 0, column: 0}
		c.MoveRight(text)

		want := Cursor{row: 0, column: 0}

		if c != want {
			t.Errorf("Cursor is unintentionally moving! got %v, want %v\n", c, want)
		}
	})
}
