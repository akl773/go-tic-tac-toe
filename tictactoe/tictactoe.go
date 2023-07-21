package tictactoe

import (
	"errors"
	"fmt"
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

func NewGame() *Game {
	return &Game{
		currentMark: 1, // Start with player 1's move
	}
}

func NewPlayer(name string, mark int) *Player {
	return &Player{name, mark}
}

func (g *Game) MakeMove(player *Player, row int, col int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 || g.board[row][col] != 0 {
		return errors.New("Invalid move")
	}

	if player.Mark != g.currentMark {
		return errors.New("Not your turn")
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
