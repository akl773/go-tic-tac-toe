package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func newGame() *Game {
	return &Game{
		currentMark: 1, // Start with player 1's move
	}
}

func NewPlayer(name string, mark int) *Player {
	return &Player{name, mark}
}

func StartGame() {
	game := newGame()
	game.PrintInstruction()
	player1 := NewPlayer("Player 1", 1)
	player2 := NewPlayer("Player 2", -1)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		game.Print()

		for _, player := range []*Player{player1, player2} {
			fmt.Printf("%s's move (row col): ", player.Name)

			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			move := strings.Split(input, "")
			if len(move) != 2 {
				fmt.Println("Invalid input. Please enter row and column numbers separated by a space.")
				continue
			}

			row, errRow := strconv.Atoi(move[0])
			col, errCol := strconv.Atoi(move[1])

			if errRow != nil || errCol != nil || row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid input. Please enter valid row and column numbers (0, 1, or 2).")
				continue
			}

			if err := game.MakeMove(player, row, col); err != nil {
				fmt.Println(err)
				continue
			}

			if game.CheckWin(player) {
				fmt.Printf("%s wins!\n", player.Name)
				game.Print()
				os.Exit(0)
			}

			if game.CheckDraw() {
				fmt.Println("It's a draw!")
				game.Print()
				os.Exit(0)
			}
		}
	}
}

func (g *Game) MakeMove(player *Player, row int, col int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 || g.board[row][col] != 0 {
		return errors.New("invalid move")
	}

	if player.Mark != g.currentMark {
		return errors.New("not your turn")
	}

	g.board[row][col] = player.Mark
	g.rowSum[row] += player.Mark
	g.colSum[col] += player.Mark
	if row == col {
		g.diagSum[0] += player.Mark
	}
	if row+col == 2 {
		g.diagSum[1] += player.Mark
	}

	g.currentMark *= -1 // Switch player

	return nil
}

func (g *Game) CheckWin(player *Player) bool {
	mark := player.Mark
	winSum := mark * 3
	for i := 0; i < 3; i++ {
		if g.rowSum[i] == winSum || g.colSum[i] == winSum {
			return true
		}
	}
	return g.diagSum[0] == winSum || g.diagSum[1] == winSum
}

func (g *Game) CheckDraw() bool {
	movesMade := 0
	for _, row := range g.board {
		for _, cell := range row {
			if cell != 0 {
				movesMade++
			}
		}
	}
	return movesMade == 9
}

func (g *Game) Print() {
	for _, row := range g.board {
		for _, cell := range row {
			var mark string
			switch cell {
			case 0:
				mark = " "
			case 1:
				mark = "X"
			case -1:
				mark = "O"
			}
			fmt.Printf("| %s ", mark)
		}
		fmt.Println("|")
		fmt.Println("-----------")
	}
}

func (g *Game) PrintInstruction() {
	fmt.Println("Welcome to Go Tic Tac Toe!")
	fmt.Println("The rules are simple, Player 1 is 'X' and Player 2 is 'O'.")
	fmt.Println("Take turns to enter a row and column number (0, 1, or 2) for your mark.")
	fmt.Println("The first player to get 3 of their marks in a row (up, down, across, or diagonally) is the winner.")
	fmt.Println("When entering your move, enter the row number and column number, without any space between them.")
	fmt.Println("For example, '0 1' will place your mark in the top-middle cell.")
	fmt.Println("\nHere are the coordinates of the board for your reference:")

	g.PrintCoordinates()

	fmt.Println("\n\nLet's get started!")
}

func (g *Game) PrintCoordinates() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("| %d,%d ", i, j)
		}
		fmt.Println("|")
		fmt.Println("--------------")
	}
}
