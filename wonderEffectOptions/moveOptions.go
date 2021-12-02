package wonderEffectOptions

const distanceDefault = 100

type MoveOptions struct {
	distance float64
	*BaseOptions
}

func NewMoveOptions() *MoveOptions {
	return &MoveOptions{
		distance:    distanceDefault,
		BaseOptions: NewBaseOptions(DirectionRight),
	}
}

func (m *MoveOptions) Distance() float64 {
	return m.distance
}

func (m *MoveOptions) SetDistance(distance float64) {
	m.distance = distance
}
