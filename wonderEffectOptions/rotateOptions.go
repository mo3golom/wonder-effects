package wonderEffectOptions

type RotateOptions struct {
	Angle float64 `mapstructure:"angle"`
	*BaseOptions
}

func NewRotateOptions() *RotateOptions {
	return &RotateOptions{
		Angle:       360,
		BaseOptions: NewBaseOptions(DirectionNormal),
	}
}
