package wonderEffectDTO

type EffectValues struct {
	MoveOnX, MoveOnY float64
	RotateOn         float64
	RotatePoint      string
	OpacityOn        float32
	ScaleOn          float64
	StartX, StartY   float64
	StartRotate      float64
	StartOpacity     float32
	StartScale       float64
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
