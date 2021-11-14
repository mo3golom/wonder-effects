package wonderEffectSettings

const (
	DirectionNormal int = iota
	DirectionReverse
)

const (
	DirectionLeft int = iota
	DirectionUp
	DirectionRight
	DirectionDown
)

type BaseSettings struct {
	direction      int
	easingFunction string
}

func NewBaseSettings(direction int) *BaseSettings {
	return &BaseSettings{
		direction:      direction,
		easingFunction: "linear",
	}
}

func (b *BaseSettings) Direction() int {
	return b.direction
}

func (b *BaseSettings) SetDirection(direction int) {
	b.direction = direction
}

func (b *BaseSettings) EasingFunction() string {
	return b.easingFunction
}

func (b *BaseSettings) SetEasingFunction(easingFunction string) {
	b.easingFunction = easingFunction
}
