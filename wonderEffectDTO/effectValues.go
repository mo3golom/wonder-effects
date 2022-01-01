package wonderEffectDTO

// Можно было использовать йоты, но поскольку данные значения будут использоваться "клиентами",
// то лучше сделать их более читаемыми
const (
	RotatePointCenter      = "center"
	RotatePointerLeft      = "left"
	RotatePointTop         = "top"
	RotatePointRight       = "right"
	RotatePointBottom      = "bottom"
	RotatePointLeftTop     = "left-top"
	RotatePointRightTop    = "right-top"
	RotatePointLeftBottom  = "left-bottom"
	RotatePointRightBottom = "right-bottom"
)

type EffectValues struct {
	MoveOnX, MoveOnY float64
	RotateOn         float64
	RotatePoint      string `mapstructure:"rotatePoint"`
	OpacityOn        float32
	ScaleOn          float64

	StartX       float64 `mapstructure:"startX"`
	StartY       float64 `mapstructure:"startY"`
	StartRotate  float64 `mapstructure:"startRotate"`
	StartOpacity float32 `mapstructure:"startOpacity"`
	StartScale   float64 `mapstructure:"startScale"`
}

func NewEffectValues() *EffectValues {
	return &EffectValues{
		OpacityOn:    0,
		ScaleOn:      0,
		StartOpacity: 1,
		StartScale:   1,
	}
}

func (e *EffectValues) X() float64 {
	return e.StartX + e.MoveOnX
}

func (e *EffectValues) Y() float64 {
	return e.StartY + e.MoveOnY
}

func (e *EffectValues) Rotate() float64 {
	return e.StartRotate + e.RotateOn
}

func (e *EffectValues) Opacity() float32 {
	return e.StartOpacity + e.OpacityOn
}

func (e *EffectValues) Scale() float64 {
	return e.StartScale + e.ScaleOn
}
