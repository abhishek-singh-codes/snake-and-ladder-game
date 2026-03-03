package main

import (
	"snakeAndLadder/internals/board"
	"snakeAndLadder/internals/dice"
	"snakeAndLadder/internals/game"
	"snakeAndLadder/internals/person"
)

func main() {
	//1. create the board
	board := board.NewBoard(100) // size is 100

	//2. create the dice
	dice := dice.NewDice(6) // 6 faces (1,2,3,4,5,6)

	//3. Add snakes
	board.AddSnake(99, 10)
	board.AddSnake(60, 40)
	board.AddSnake(30, 2)

	//4. Add ladders
	board.AddLadder(50, 80)
	board.AddLadder(20, 75)
	board.AddLadder(80, 95)

	//5. create players
	p1 := person.NewPerson("Abhishek")
	p2 := person.NewPerson("Rahul")

	players := []*person.Person{p1, p2}

	// create Game
	g := game.NewGame(dice, board, players)

	// start the game
	g.Start()
}
