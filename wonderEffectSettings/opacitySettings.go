package wonderEffectSettings

type OpacitySettings struct {
	*BaseSettings
}

func NewOpacitySettings() *OpacitySettings {
	return &OpacitySettings{
		BaseSettings: NewBaseSettings(DirectionNormal),
	}
}

func (o *OpacitySettings) Direction() int {
	return o.direction
}

func (o *OpacitySettings) SetDirection(direction int) {
	o.direction = direction
}
