package main

import (
	"testing"
)

func setGridDimensions(w, h int) {
	width = w
	height = h
}

func TestInitialModel(t *testing.T) {
	// Set width and height for the test
	setGridDimensions(5, 5)

	m := initialModel()

	// Check if the grid has the correct dimensions
	if len(m.grid) != height {
		t.Errorf("Expected grid height to be %d, got %d", height, len(m.grid))
	}
	for i := 0; i < height; i++ {
		if len(m.grid[i]) != width {
			t.Errorf("Expected grid width to be %d, got %d", width, len(m.grid[i]))
		}
	}

	// Check if every cell is either alive (true) or dead (false)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if !(m.grid[i][j].alive == true || m.grid[i][j].alive == false) {
				t.Errorf("Expected grid[%d][%d] to be either alive or dead, got %v", i, j, m.grid[i][j])
			}
		}
	}
}

func TestNextGeneration(t *testing.T) {
	// Set width and height for the test grid to match the test input size
	setGridDimensions(3, 3)

	input := [][]cell{
		{{alive: false}, {alive: true}, {alive: false}},
		{{alive: true}, {alive: true}, {alive: false}},
		{{alive: false}, {alive: false}, {alive: true}},
	}

	expected := [][]cell{
		{{alive: true}, {alive: true}, {alive: false}},
		{{alive: true}, {alive: true}, {alive: true}},
		{{alive: false}, {alive: true}, {alive: false}},
	}

	result := nextGeneration(input)
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if result[i][j].alive != expected[i][j].alive {
				t.Errorf("Expected grid[%d][%d] to be %v, got %v", i, j, expected[i][j].alive, result[i][j].alive)
			}
		}
	}
}

func TestCountAliveNeighbors(t *testing.T) {
	setGridDimensions(3, 3)
	grid := [][]cell{
		{{alive: false}, {alive: true}, {alive: false}},
		{{alive: true}, {alive: true}, {alive: false}},
		{{alive: false}, {alive: false}, {alive: true}},
	}
	tests := []struct {
		x, y     int
		expected int
	}{
		{0, 0, 3}, // Top-left corner (should have 3 alive neighbors)
		{0, 1, 2}, // Top-middle (should have 2 alive neighbors)
		{1, 1, 3}, // Center (should have 3 alive neighbors)
		{2, 2, 1}, // Bottom-right corner (should have 1 alive neighbor)
	}
	for _, test := range tests {
		result := countAliveNeighbors(grid, test.x, test.y)
		if result != test.expected {
			t.Errorf("For cell (%d, %d), expected %d alive neighbors, got %d", test.x, test.y, test.expected, result)
		}
	}
}

func TestCustomGridSize(t *testing.T) {
	// Test with a custom grid size
	setGridDimensions(10, 10)

	m := initialModel()

	if len(m.grid) != 10 {
		t.Errorf("Expected grid height to be 10, got %d", len(m.grid))
	}
	for i := 0; i < 10; i++ {
		if len(m.grid[i]) != 10 {
			t.Errorf("Expected grid width to be 10, got %d", len(m.grid[i]))
		}
	}
}
