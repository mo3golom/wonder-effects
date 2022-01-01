package wonderEffectOptions

const scaleOnDefault = 0.2

type ScaleOptions struct {
	ScaleOn float64 `mapstructure:"scaleOn"`
	*BaseOptions
}

func NewScaleOptions() *ScaleOptions {
	return &ScaleOptions{
		ScaleOn:     scaleOnDefault,
		BaseOptions: NewBaseOptions(DirectionNormal),
	}
}
