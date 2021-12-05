package wonderEffectOptions

const distanceDefault = 100

type MoveOptions struct {
	distance     float64
	pathFunction string
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

func (m *MoveOptions) PathFunction() string {
	return m.pathFunction
}

func (m *MoveOptions) SetPathFunction(pathFunction string) {
	m.pathFunction = pathFunction
}
