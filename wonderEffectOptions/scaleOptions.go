package wonderEffectOptions

const scaleOnDefault = 0.2

type ScaleOptions struct {
	scaleOn float64
	*BaseOptions
}

func NewScaleOptions() *ScaleOptions {
	return &ScaleOptions{
		scaleOn:     scaleOnDefault,
		BaseOptions: NewBaseOptions(DirectionNormal),
	}
}

func (s *ScaleOptions) ScaleOn() float64 {
	return s.scaleOn
}

func (s *ScaleOptions) SetScaleOn(scaleOn float64) {
	s.scaleOn = scaleOn
}
