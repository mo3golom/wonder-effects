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
	direction      string
	easingFunction string
}

func NewBaseOptions(direction string) *BaseOptions {
	return &BaseOptions{
		direction:      direction,
		easingFunction: "linear",
	}
}

func (b *BaseOptions) Direction() string {
	return b.direction
}

func (b *BaseOptions) SetDirection(direction string) {
	b.direction = direction
}

func (b *BaseOptions) EasingFunction() string {
	return b.easingFunction
}

func (b *BaseOptions) SetEasingFunction(easingFunction string) {
	b.easingFunction = easingFunction
}
