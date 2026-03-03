package game

import (
	"fmt"
	"snakeAndLadder/internals/board"
	"snakeAndLadder/internals/dice"
	"snakeAndLadder/internals/person"
)

type Game struct {
	dice    dice.Dice
	board   *board.Board
	players []*person.Person
	current int
}

func NewGame(d dice.Dice, b *board.Board, p []*person.Person) *Game {
	return &Game{
		dice:    d,
		board:   b,
		players: p,
		current: 0,
	}
}

func (g *Game) play() (*person.Person, bool) {
	if len(g.players) == 0 {
		return nil, false
	}
	// turn of this player
	player := g.players[g.current]
	playerName := player.GetPlayerName()
	playerCurrentPos := player.GetPlayerPosition()

	fmt.Printf("%s turn the dice\n", playerName)
	fmt.Printf("Current position of %s, is %d\n", playerName, playerCurrentPos)

	//player will roll the dice and get new position
	fmt.Println("Dice rolled...")
	roll := g.dice.Roll()
	fmt.Printf("%d number came in the dice\n", roll)
	newPos := playerCurrentPos + roll

	fmt.Printf("New position of %s, is %d\n", playerName, newPos)

	// check board limit
	if newPos <= g.board.Size() {
		newPos = g.board.GetFinalPosition(newPos)
		player.SetPlayerPosition(newPos)
	}

	// checkwin
	if newPos == g.board.Size() {
		return player, true
	}

	// rotate turn
	g.current = (g.current + 1) % len(g.players)

	return player, false
}

func (g *Game) Start() {
	for {
		player, won := g.play()
		if player == nil {
			fmt.Println("Player is not set to the game")
			return
		}

		fmt.Printf("%s moved to %d\n", player.GetPlayerName(), player.GetPlayerPosition())

		if won {
			fmt.Printf("Congratulation!!, %s Won the game", player.GetPlayerName())
			return
		}
	}
}
