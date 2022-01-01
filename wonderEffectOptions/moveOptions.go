package wonderEffectOptions

const distanceDefault = 100

type MoveOptions struct {
	Distance     float64 `mapstructure:"distance"`
	PathFunction string  `mapstructure:"pathFunction"`
	*BaseOptions
}

func NewMoveOptions() *MoveOptions {
	return &MoveOptions{
		Distance:    distanceDefault,
		BaseOptions: NewBaseOptions(DirectionRight),
	}
}
