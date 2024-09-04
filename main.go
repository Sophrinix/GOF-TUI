package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Constants for the game
var width, height int

type cell struct {
	alive bool
}

type model struct {
	grid    [][]cell
	running bool
}

func initialModel() model {
	m := model{
		grid: make([][]cell, height),
	}
	for i := 0; i < height; i++ {
		m.grid[i] = make([]cell, width)
		for j := 0; j < width; j++ {
			m.grid[i][j].alive = rand.Intn(2) == 1
		}
	}
	return m
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.running = !m.running
		}
	case tickMsg:
		if m.running {
			m.grid = nextGeneration(m.grid)
		}
		return m, tickCmd()
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if m.grid[i][j].alive {
				b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#212829")).Background(lipgloss.Color("#3b7ed2")).Render("  "))
			} else {
				b.WriteString("  ")
			}
		}
		if i < height-1 {
			b.WriteRune('\n')
		}
	}
	return b.String() + "\n\nPress 'r' to toggle simulation, 'q' or 'ctrl+c' to quit."
}

func nextGeneration(grid [][]cell) [][]cell {
	height := len(grid)
	width := len(grid[0])
	newGrid := make([][]cell, height)
	for i := 0; i < height; i++ {
		newGrid[i] = make([]cell, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			aliveNeighbors := countAliveNeighbors(grid, i, j)
			if grid[i][j].alive && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				newGrid[i][j] = cell{alive: true}
			} else if !grid[i][j].alive && aliveNeighbors == 3 {
				newGrid[i][j] = cell{alive: true}
			} else {
				newGrid[i][j] = cell{alive: false}
			}
		}
	}
	return newGrid
}

func countAliveNeighbors(grid [][]cell, x, y int) int {
	height := len(grid)
	width := len(grid[0])
	aliveCount := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Skip the current cell itself
			if i == 0 && j == 0 {
				continue
			}

			// Neighbor's coordinates
			neighborX, neighborY := x+i, y+j

			// Only count neighbors within bounds of the grid
			if neighborX >= 0 && neighborX < height && neighborY >= 0 && neighborY < width {
				if grid[neighborX][neighborY].alive {
					aliveCount++
				}
			}
		}
	}

	return aliveCount
}

func isValidCoordinate(x, y int) bool {
	return x >= 0 && x < height && y >= 0 && y < width
}

type tickMsg struct{}

func tickCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(250 * time.Millisecond)
		return tickMsg{}
	}
}

func main() {
	// Command-line flags for width, height, and seed input
	flag.IntVar(&width, "width", 40, "width of the grid")
	flag.IntVar(&height, "height", 25, "height of the grid")
	var seed int64
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "seed for random generation")

	// Parse the command-line flags
	flag.Parse()

	// Set the random seed
	rand.Seed(seed)

	// Create and start the Bubble Tea program
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
