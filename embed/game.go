package embed

type Position struct {
	poxX float64
	poxY float64
}

func (p *Position) Move(x, y float64) {
	p.poxX += x
	p.poxY += y
}

func (p *Position) Teleport(x, y float64) {
	p.poxX = x
	p.poxY = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}
