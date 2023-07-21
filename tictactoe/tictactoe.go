package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EmptyMark = iota
	Player1Mark
	Player2Mark
)

type Game struct {
	board       [3][3]int
	rowSum      [3]int
	colSum      [3]int
	diagSum     [2]int
	currentMark int
}

type Player struct {
	Name string
	Mark int
}

type Move struct {
	Row int
	Col int
}

// NewGame initializes a new game and returns its address.
func NewGame() *Game {
	return &Game{
		currentMark: Player1Mark, // Start with player 1's move
	}
}

// NewPlayer initializes a new player and returns its address.
func NewPlayer(name string, mark int) *Player {
	return &Player{name, mark}
}

// StartGame creates and starts a new game.
func StartGame() {
	game := NewGame()
	printInstruction()
	player1 := NewPlayer("Player 1", Player1Mark)
	player2 := NewPlayer("Player 2", Player2Mark)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		printBoard(game.board)

		for _, player := range []*Player{player1, player2} {
			fmt.Printf("%s's move (row col): ", player.Name)

			move, err := parseMove(scanner)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := game.MakeMove(player, *move); err != nil {
				fmt.Println(err)
				continue
			}

			if game.CheckWin(player) {
				fmt.Printf("%s wins!\n", player.Name)
				printBoard(game.board)
				os.Exit(0)
			}

			if game.CheckDraw() {
				fmt.Println("It's a draw!")
				printBoard(game.board)
				os.Exit(0)
			}
		}
	}
}

func parseMove(scanner *bufio.Scanner) (*Move, error) {
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	moveInput := strings.Split(input, " ")
	if len(moveInput) != 2 {
		return nil, errors.New("invalid input. Please enter row and column numbers separated by a space")
	}

	row, errRow := strconv.Atoi(moveInput[0])
	col, errCol := strconv.Atoi(moveInput[1])

	if errRow != nil || errCol != nil || row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, errors.New("invalid input. Please enter valid row and column numbers (0, 1, or 2)")
	}

	return &Move{Row: row, Col: col}, nil
}

// MakeMove makes a move for a player.
func (g *Game) MakeMove(player *Player, move Move) error {
	if move.Row < 0 || move.Row > 2 || move.Col < 0 || move.Col > 2 || g.board[move.Row][move.Col] != EmptyMark {
		return errors.New("invalid move")
	}

	if player.Mark != g.currentMark {
		return errors.New("not your turn")
	}

	g.board[move.Row][move.Col] = player.Mark
	g.rowSum[move.Row] += player.Mark
	g.colSum[move.Col] += player.Mark
	if move.Row == move.Col {
		g.diagSum[0] += player.Mark
	}
	if move.Row+move.Col == 2 {
		g.diagSum[1] += player.Mark
	}

	g.currentMark *= -1 // Switch player

	return nil
}

// CheckWin checks if a player has won the game.
func (g *Game) CheckWin(player *Player) bool {
	winSum := player.Mark * 3
	for i := 0; i < 3; i++ {
		if g.rowSum[i] == winSum || g.colSum[i] == winSum {
			return true
		}
	}
	return g.diagSum[0] == winSum || g.diagSum[1] == winSum
}

// CheckDraw checks if the game is a draw.
func (g *Game) CheckDraw() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == EmptyMark {
				return false
			}
		}
	}
	return true
}

func printInstruction() {
	fmt.Println("Welcome to Go Tic Tac Toe!")
	fmt.Println("The rules are simple, Player 1 is 'X' and Player 2 is 'O'.")
	fmt.Println("Take turns to enter a row and column number (0, 1, or 2) for your mark.")
	fmt.Println("The first player to get 3 of their marks in a row (up, down, across, or diagonally) is the winner.")
	fmt.Println("When entering your move, enter the row number and column number, without any space between them.")
	fmt.Println("For example, '0 1' will place your mark in the top-middle cell.")
	fmt.Println("\nHere are the coordinates of the board for your reference:")

	printCoordinates()

	fmt.Println("\n\nLet's get started!")
}

func printCoordinates() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" | %d,%d ", i, j)
		}
		fmt.Println("|")
		fmt.Println("----------------------")
	}
}

// printBoard prints the current state of the game board.
func printBoard(board [3][3]int) {
	for _, row := range board {
		for _, cell := range row {
			var mark string
			switch cell {
			case EmptyMark:
				mark = " "
			case Player1Mark:
				mark = "X"
			case Player2Mark:
				mark = "O"
			}
			fmt.Printf("| %s ", mark)
		}
		fmt.Println("|")
		fmt.Println("---------------")
	}
}
