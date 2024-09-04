# Game of Life TUI

![Game of Life TUI Demo](vhs/game-of-life.gif)

This project is an exploration of building a Text User Interface (TUI) using the Bubble Tea framework, applied to the well-understood problem of Conway's Game of Life.

## Overview

The Game of Life, devised by mathematician John Conway, is a cellular automaton simulation that demonstrates how complex patterns can emerge from simple rules. This implementation brings the classic simulation to life in your terminal using Go and the Bubble Tea framework.

## Features

- **Bubble Tea TUI**: Utilizes the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework to create an interactive and visually appealing terminal interface.
- **Game of Life Simulation**: Implements the classic rules of Conway's Game of Life.
- **Interactive Controls**: Toggle the simulation on/off and quit the application with simple key commands.
- **Colorful Display**: Uses [lipgloss](https://github.com/charmbracelet/lipgloss) for styling, making the game board visually distinct and easy to read.

## How It Works

The application combines the simplicity of the Game of Life rules with the power of Bubble Tea's event-driven architecture:

1. **Grid Initialization**: The game starts with a randomly populated grid of cells.
2. **Simulation Loop**: In each tick, the next generation of cells is calculated based on the current state.
3. **User Interface**: Bubble Tea handles the rendering and user input, creating a responsive TUI.
4. **Game Logic**: The core Game of Life rules are implemented in the `nextGeneration` function.

## Learning Outcomes

This project serves as an excellent example of:

- Building interactive TUIs with Go
- Implementing classic algorithms in a visual, engaging manner
- Structuring an application using the Bubble Tea framework
- Balancing simplicity and functionality in software design

## Getting Started

To run the Game of Life TUI:

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Run `go mod tidy` to install dependencies.
4. Execute `go run main.go` to start the application.

## Controls

- Press `r` to toggle the simulation on/off
- Press `q` or `Ctrl+C` to quit the application

## Testing

Unit tests are provided to ensure the correctness of the Game of Life rules implementation. Run `go test` to execute the tests.

## Creating a Screen Recording (or how I did)

You can create a GIF or video of the Game of Life TUI in action using [VHS](https://github.com/charmbracelet/vhs). Here's how:

1. Install VHS:
   - On macOS, you can use Homebrew:
     ```
     brew install vhs
     ```
   - For other operating systems, follow the instructions on the [VHS GitHub page](https://github.com/charmbracelet/vhs).

2. Run the following command in your terminal:

   ```
   vhs  -o vhs/gof-tui.gif  vhs/gof-tui.tape
   ```
## Conclusion

This project demonstrates how a well-understood problem like the Game of Life can be reimagined as an interactive terminal application using modern Go libraries. It serves as both a learning tool and a starting point for more complex TUI applications.

Feel free to explore, modify, and expand upon this code to deepen your understanding of Go, Bubble Tea, and cellular automata!
