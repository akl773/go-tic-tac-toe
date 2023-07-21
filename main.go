package main

import (
	"bufio"
	"fmt"
	ttt "github.com/akl773/go-tic-tac-toe/tictactoe"
	"os"
	"strconv"
	"strings"
)

func main() {
	game := ttt.NewGame()
	game.PrintInstruction()

	player1 := ttt.NewPlayer("Player 1", 1)
	player2 := ttt.NewPlayer("Player 2", -1)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		game.Print()

		for _, player := range []*ttt.Player{player1, player2} {
			fmt.Printf("%s's move: ", player.Name)

			inputs := make([]int, 2)
			for i := 0; i < 2; i++ {
				scanner.Scan()
				input, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
				inputs[i] = input
			}

			if err := game.MakeMove(player, inputs[0], inputs[1]); err != nil {
				fmt.Println(err)
				continue
			}

			if game.CheckWin(player) {
				fmt.Printf("%s wins!\n", player.Name)
				game.Print()
				os.Exit(0)
			}
		}
	}
}
