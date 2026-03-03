package dice

import (
	"math/rand"
	"time"
)

type Dice interface {
	Roll() int
}

type dice struct {
	face int
	rng  *rand.Rand
}

func NewDice(face int) Dice {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return &dice{
		face: face,
		rng:  r,
	}
}

func (d *dice) Roll() int {
	return d.rng.Intn(d.face) + 1
}
