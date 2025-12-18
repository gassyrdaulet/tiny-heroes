package physics

type Body struct {
	X, Y          float64
	VX, VY        float64
	Width, Height float64
	OnGround      bool
	Weight        float64
}

func (b *Body) Position() (x, y float64) {
	return b.X, b.Y
}

func (b *Body) SetPosition(x, y float64) {
	b.X = x
	b.Y = y
}

func (b *Body) Velocity() (vx, vy float64) {
	return b.VX, b.VY
}

func (b *Body) SetVelocity(vx, vy float64) {
	b.VX = vx
	b.VY = vy
}

func (b *Body) VXValue() float64 {
	return b.VX
}

func (b *Body) VYValue() float64 {
	return b.VY
}

func (b *Body) SetVX(vx float64) {
	b.VX = vx
}

func (b *Body) SetVY(vy float64) {
	b.VY = vy
}

func (b *Body) Size() (w, h float64) {
	return b.Width, b.Height
}

func (b *Body) SetSize(w, h float64) {
	b.Width = w
	b.Height = h
}

func (b *Body) IsOnGround() bool {
	return b.OnGround
}

func (b *Body) SetOnGround(onGround bool) {
	b.OnGround = onGround
}

func (b *Body) GetWeight() float64 {
	return b.Weight
}

func (b *Body) SetWeight(weight float64) {
	b.Weight = weight
}
