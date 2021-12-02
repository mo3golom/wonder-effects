package wonderEffectOptions

type RotateOptions struct {
	angle float64
	*BaseOptions
}

func NewRotateOptions() *RotateOptions {
	return &RotateOptions{
		angle:       360,
		BaseOptions: NewBaseOptions(DirectionNormal),
	}
}

func (r *RotateOptions) Angle() float64 {
	return r.angle
}

func (r *RotateOptions) SetAngle(angle float64) {
	r.angle = angle
}
