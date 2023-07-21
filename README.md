## Go Tic Tac Toe
Go Tic Tac Toe is a simple yet robust console-based game of Tic Tac Toe written in Go.

## Introduction
The application is a Go package that implements the game logic for Tic Tac Toe. It allows two players to play the game on a 3x3 grid, marking their respective positions in turns. The first player to align 3 marks wins the game.

## Package Structure
The package is composed of several Go types and functions:

`Game`: The main struct representing a game. It keeps track of the board state, current player turn, and helper data for checking win conditions. 

`Player`: Represents a player with a name and mark.

`Move`: Represents a player's move with row and column.

`NewGame()`: Initializes a new game.

`NewPlayer(name string, mark int)`: Initializes a new player.

`StartGame`(): Starts a new game and handles user input and game logic.


## Usage

```
package main

import "tictactoe"

func main() {
	tictactoe.StartGame()
}

```

This will start a new game of Tic Tac Toe.

The game will guide players through the process. Players should enter the coordinates of the cell they want to mark, separated by a space, for example 1 2. The top-left cell has coordinates 0 0.
