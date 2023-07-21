package tictactoe

import (
	"bufio"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestTictactoe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tictactoe Suite")
}

var _ = Describe("TicTacToe", func() {

	Describe("NewGame", func() {
		It("should initialize a new game", func() {
			game := NewGame()
			Expect(game).ToNot(BeNil())
			Expect(game.currentMark).To(Equal(Player1Mark))
		})
	})

	Describe("NewPlayer", func() {
		It("should initialize a new player", func() {
			player := NewPlayer("TestPlayer", Player1Mark)
			Expect(player).ToNot(BeNil())
			Expect(player.Name).To(Equal("TestPlayer"))
			Expect(player.Mark).To(Equal(Player1Mark))
		})
	})

	Describe("MakeMove", func() {
		var (
			game    *Game
			player1 *Player
		)

		BeforeEach(func() {
			game = NewGame()
			player1 = NewPlayer("Player 1", Player1Mark)
		})

		It("should not return an error when a valid move is made", func() {
			err := game.MakeMove(player1, Move{0, 0})
			Expect(err).To(BeNil())
		})

		It("should return an error when an invalid move is made", func() {
			game.MakeMove(player1, Move{0, 0})
			err := game.MakeMove(player1, Move{0, 0})
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("CheckWin", func() {
		var (
			game    *Game
			player1 *Player
			player2 *Player
		)

		BeforeEach(func() {
			game = NewGame()
			player1 = NewPlayer("Player 1", Player1Mark)
			player2 = NewPlayer("Player 2", Player2Mark)
			game.MakeMove(player1, Move{0, 0})
			game.MakeMove(player2, Move{1, 0})
			game.MakeMove(player1, Move{0, 1})
			game.MakeMove(player2, Move{1, 1})
			game.MakeMove(player1, Move{0, 2})
			game.MakeMove(player2, Move{2, 0})
		})

		It("should return true when a player has won", func() {
			Expect(game.CheckWin(player1)).To(BeTrue())
		})
	})

	Describe("CheckDraw", func() {
		var (
			game    *Game
			player1 *Player
			player2 *Player
		)

		Describe("when the game is a draw", func() {
			BeforeEach(func() {
				game = NewGame()
				player1 = NewPlayer("Player 1", Player1Mark)
				player2 = NewPlayer("Player 2", Player2Mark)
				game.MakeMove(player1, Move{Row: 0, Col: 0})
				game.MakeMove(player2, Move{Row: 1, Col: 1})
				game.MakeMove(player1, Move{Row: 0, Col: 1})
				game.MakeMove(player2, Move{Row: 0, Col: 2})
				game.MakeMove(player1, Move{Row: 1, Col: 2})
				game.MakeMove(player2, Move{Row: 1, Col: 0})
				game.MakeMove(player1, Move{Row: 2, Col: 0})
				game.MakeMove(player2, Move{Row: 2, Col: 1})
				game.MakeMove(player1, Move{Row: 2, Col: 2})

			})

			It("should return true", func() {
				Expect(game.CheckDraw()).To(BeTrue())
			})
		})

		Describe("when the game is not a draw", func() {
			BeforeEach(func() {
				game = NewGame()
				player1 = NewPlayer("Player 1", Player1Mark)
				player2 = NewPlayer("Player 2", Player2Mark)
				game.MakeMove(player1, Move{Row: 0, Col: 0})
				game.MakeMove(player2, Move{Row: 1, Col: 1})
				game.MakeMove(player1, Move{Row: 0, Col: 1})
				game.MakeMove(player2, Move{Row: 0, Col: 2})
				game.MakeMove(player1, Move{Row: 1, Col: 2})
				game.MakeMove(player2, Move{Row: 1, Col: 0})
				game.MakeMove(player1, Move{Row: 2, Col: 0})
			})

			It("should return false", func() {
				Expect(game.CheckDraw()).To(BeFalse())
			})
		})
	})

	Describe("parseMove", func() {
		// Use a test table to test multiple cases
		tableEntries := []TableEntry{
			Entry("should return error when the input is empty", "", true),
			Entry("should return error when the input is non-numeric", "a b", true),
			Entry("should return error when the input is outside the valid range", "3 3", true),
			Entry("should return a valid Move when the input is within the valid range", "1 2", false),
		}

		DescribeTable("parseMove",
			func(input string, expectErr bool) {
				// Mock the scanner
				r := strings.NewReader(input)
				scanner := bufio.NewScanner(r)
				move, err := parseMove(scanner)

				if expectErr {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).To(BeNil())
					Expect(move).To(Equal(&Move{1, 2}))
				}
			},
			tableEntries...,
		)
	})

})
