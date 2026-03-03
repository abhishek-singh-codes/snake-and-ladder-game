package person

type Person struct {
	name            string
	currentPosition int
}

func NewPerson(n string) *Person {
	return &Person{
		name:            n,
		currentPosition: 0,
	}
}

func (p *Person) GetPlayerPosition() int {
	return p.currentPosition
}

func (p *Person) SetPlayerPosition(pos int) {
	p.currentPosition = pos
}

func (p *Person) GetPlayerName() string {
	return p.name
}
