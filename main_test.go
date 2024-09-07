package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func setGridDimensions(w, h int) {
	width = w
	height = h
}

func TestInitialModel(t *testing.T) {
	// Test random initialization
	setGridDimensions(5, 5)
	inputFile = ""
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

func TestInitialModelFromFileWithComments(t *testing.T) {
	// Create a temporary input file with comments and empty lines
	content := []byte(`
! This is a comment
*..*
! Another comment

.*.*
! Final comment
..*.
`)
	tmpfile, err := ioutil.TempFile("", "test_input")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Set the input file and initialize the model
	inputFile = tmpfile.Name()
	m := initialModel()

	// Check dimensions
	expectedHeight, expectedWidth := 3, 4
	if len(m.grid) != expectedHeight {
		t.Errorf("Expected grid height to be %d, got %d", expectedHeight, len(m.grid))
	}
	for i := 0; i < expectedHeight; i++ {
		if len(m.grid[i]) != expectedWidth {
			t.Errorf("Expected grid width to be %d, got %d", expectedWidth, len(m.grid[i]))
		}
	}

	// Check cell states
	expectedGrid := [][]bool{
		{true, false, false, true},
		{false, true, false, true},
		{false, false, true, false},
	}

	for i := 0; i < expectedHeight; i++ {
		for j := 0; j < expectedWidth; j++ {
			if m.grid[i][j].alive != expectedGrid[i][j] {
				t.Errorf("Expected grid[%d][%d] to be %v, got %v", i, j, expectedGrid[i][j], m.grid[i][j].alive)
			}
		}
	}
}

func TestInitialModelFromFile(t *testing.T) {
	// Create a temporary input file with a comment and an empty line
	content := []byte("! This is a comment\n*..\n\n.*.\n..*\n")
	tmpfile, err := ioutil.TempFile("", "test_input")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Set the input file and initialize the model
	inputFile = tmpfile.Name()
	m := initialModel()

	// Check dimensions
	expectedHeight, expectedWidth := 3, 3
	if len(m.grid) != expectedHeight {
		t.Errorf("Expected grid height to be %d, got %d", expectedHeight, len(m.grid))
	}
	for i := 0; i < expectedHeight; i++ {
		if len(m.grid[i]) != expectedWidth {
			t.Errorf("Expected grid width to be %d, got %d", expectedWidth, len(m.grid[i]))
		}
	}

	// Check cell states
	expectedGrid := [][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, true},
	}

	for i := 0; i < expectedHeight; i++ {
		for j := 0; j < expectedWidth; j++ {
			if m.grid[i][j].alive != expectedGrid[i][j] {
				t.Errorf("Expected grid[%d][%d] to be %v, got %v", i, j, expectedGrid[i][j], m.grid[i][j].alive)
			}
		}
	}
}

func TestNextGeneration(t *testing.T) {
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
	inputFile = "" // Ensure we're not using an input file
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
