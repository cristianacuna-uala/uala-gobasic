package point

type Point struct {
	X, Y int
}

func (p *Point) Swap(){
	if (p.X >= p.Y) {
		aux := p.X
		p.X = p.Y
		p.Y = aux
	}
}