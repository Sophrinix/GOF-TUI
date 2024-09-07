# Set up the terminal
Set Shell "bash"
Set FontSize 16
Set Width 2400
Set Height 600
Set TypingSpeed 80ms

# Run the Game of Life
Type "go run main.go -input examples/gol/gol.gol"
Enter
Sleep 4s

# Toggle the simulation on and then off
Type "r"
Sleep 4s


# Exit the game
Type "q"

# Wait before finishing
Sleep 1s
