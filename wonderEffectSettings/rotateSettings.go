package wonderEffectSettings

type RotateSettings struct {
	angle float64
	*BaseSettings
}

func NewRotateSettings() *RotateSettings {
	return &RotateSettings{
		angle:        360,
		BaseSettings: NewBaseSettings(DirectionNormal),
	}
}

func (r *RotateSettings) Angle() float64 {
	return r.angle
}

func (r *RotateSettings) SetAngle(angle float64) {
	r.angle = angle
}

func (r *RotateSettings) Direction() int {
	return r.direction
}

func (r *RotateSettings) SetDirection(direction int) {
	r.direction = direction
}
