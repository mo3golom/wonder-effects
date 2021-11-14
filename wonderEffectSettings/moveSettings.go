package wonderEffectSettings

const distanceDefault = 100

type MoveSettings struct {
	distance float64
	*BaseSettings
}

func NewMoveSettings() *MoveSettings {
	return &MoveSettings{
		distance:     distanceDefault,
		BaseSettings: NewBaseSettings(DirectionRight),
	}
}

func (m *MoveSettings) Distance() float64 {
	return m.distance
}

func (m *MoveSettings) SetDistance(distance float64) {
	m.distance = distance
}
