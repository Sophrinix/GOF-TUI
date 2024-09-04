package main

import (
	"testing"
)

func TestInitialModel(t *testing.T) {
	m := initialModel()
	// Check if every cell is either true or false
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if !(m.grid[i][j].alive == true || m.grid[i][j].alive == false) {
				t.Errorf("Expected grid[%d][%d] to be either alive or dead, got %v", i, j, m.grid[i][j])
			}
		}
	}
}

func TestNextGeneration(t *testing.T) {
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
				t.Errorf("Expected grid[%d][%d] to be %v, got %v", i, j, expected[i][j], result[i][j])
			}
		}
	}
}
