package wonderEffectOptions

const (
	DirectionNormal  = "normal"
	DirectionReverse = "reverse"
)

const (
	DirectionLeft  = "left"
	DirectionUp    = "up"
	DirectionRight = "right"
	DirectionDown  = "down"
)

type BaseOptions struct {
	Direction      string `mapstructure:"direction"`
	EasingFunction string `mapstructure:"easingFunction"`
}

func NewBaseOptions(direction string) *BaseOptions {
	return &BaseOptions{
		Direction:      direction,
		EasingFunction: "linear",
	}
}
