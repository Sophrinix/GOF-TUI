package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var width, height int
var inputFile string

type cell struct {
	alive bool
}

type model struct {
	grid    [][]cell
	running bool
}

func readInputFile(filename string) ([][]cell, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]cell
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			width = len(line)
		} else if len(line) != width {
			return nil, fmt.Errorf("line %d has length %d, expected %d", lineNum+1, len(line), width)
		}
		row := make([]cell, width)
		for i, char := range line {
			row[i] = cell{alive: char == '*' || char == '1'}
		}
		grid = append(grid, row)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	height = len(grid)
	return grid, nil
}

func initialModel() model {
	var grid [][]cell
	var err error

	if inputFile != "" {
		grid, err = readInputFile(inputFile)
		if err != nil {
			fmt.Printf("Error reading input file: %v\n", err)
			os.Exit(1)
		}
	} else {
		grid = make([][]cell, height)
		for i := 0; i < height; i++ {
			grid[i] = make([]cell, width)
			for j := 0; j < width; j++ {
				grid[i][j].alive = rand.Intn(2) == 1
			}
		}
	}

	return model{
		grid: grid,
	}
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
	aliveCount := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			neighborX, neighborY := x+i, y+j
			if neighborX >= 0 && neighborX < height && neighborY >= 0 && neighborY < width {
				if grid[neighborX][neighborY].alive {
					aliveCount++
				}
			}
		}
	}
	return aliveCount
}

type tickMsg struct{}

func tickCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(250 * time.Millisecond)
		return tickMsg{}
	}
}

func main() {
	flag.IntVar(&width, "width", 40, "width of the grid (default 40, ignored if input file is provided)")
	flag.IntVar(&height, "height", 25, "height of the grid (default 25, ignored if input file is provided)")
	flag.StringVar(&inputFile, "input", "", "path to input file for initial state")
	var seed int64
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "seed for random generation")
	flag.Parse()

	rand.Seed(seed)

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
