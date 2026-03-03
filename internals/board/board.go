package board

import "fmt"

type Board struct {
	size    int
	snakes  map[int]int
	ladders map[int]int
}

func NewBoard(size int) *Board {
	return &Board{
		size:    size,
		snakes:  make(map[int]int),
		ladders: make(map[int]int),
	}
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) AddLadder(start, end int) error {
	if start <= 0 || start >= b.size {
		return fmt.Errorf("invalid ladder start")
	}
	if end <= start || end > b.size {
		return fmt.Errorf("invalid ladder end")
	}
	if _, exists := b.ladders[start]; exists {
		return fmt.Errorf("ladder already exists at this start")
	}
	if _, exists := b.snakes[start]; exists {
		return fmt.Errorf("snake already exists at this cell")
	}

	b.ladders[start] = end
	return nil
}

func (b *Board) AddSnake(head, tail int) error {
	if head <= 1 || head >= b.size {
		return fmt.Errorf("invalid snake head")
	}
	if tail <= 0 || tail >= head {
		return fmt.Errorf("invalid snake tail")
	}
	if _, exists := b.snakes[head]; exists {
		return fmt.Errorf("snake already exists at this head")
	}
	if _, exists := b.ladders[head]; exists {
		return fmt.Errorf("ladder already exists at this cell")
	}

	b.snakes[head] = tail
	return nil
}

// chain
// Ladder 5 -> 20
// Snake 20 -> 8
func (b *Board) GetFinalPosition(pos int) int {
	for {
		if newPos, ok := b.ladders[pos]; ok {
			fmt.Println("Ladder found 😎")
			pos = newPos
			continue
		}
		if newPos, ok := b.snakes[pos]; ok {
			fmt.Println("Sanke Bite 😭")
			pos = newPos
			continue
		}
		break
	}
	return pos
}
