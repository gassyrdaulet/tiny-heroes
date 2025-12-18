package physics

type PhysicalBody interface {
	Position() (x, y float64)
	SetPosition(x, y float64)
	Velocity() (vx, vy float64)
	SetVelocity(vx, vy float64)
	VXValue() float64
	VYValue() float64
	SetVX(vx float64)
	SetVY(vy float64)
	Size() (w, h float64)
	SetSize(w, h float64)
	IsOnGround() bool
	SetOnGround(onGround bool)
	GetWeight() float64
	SetWeight(weight float64)
}