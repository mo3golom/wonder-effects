package wonderEffectOptions

type OpacityOptions struct {
	*BaseOptions
}

func NewOpacityOptions() *OpacityOptions {
	return &OpacityOptions{
		BaseOptions: NewBaseOptions(DirectionNormal),
	}
}
