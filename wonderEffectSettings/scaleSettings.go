package wonderEffectSettings

const scaleOnDefault = 0.2

type ScaleSettings struct {
	scaleOn float64
	*BaseSettings
}

func NewScaleSettings() *ScaleSettings {
	return &ScaleSettings{
		scaleOn:      scaleOnDefault,
		BaseSettings: NewBaseSettings(DirectionNormal),
	}
}

func (s *ScaleSettings) Direction() int {
	return s.direction
}

func (s *ScaleSettings) SetDirection(direction int) {
	s.direction = direction
}

func (s *ScaleSettings) ScaleOn() float64 {
	return s.scaleOn
}

func (s *ScaleSettings) SetScaleOn(scaleOn float64) {
	s.scaleOn = scaleOn
}
