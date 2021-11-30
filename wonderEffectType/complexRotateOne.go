package wonderEffectType

// NewComplexRotateOne Комплексный эффект вращения 01
// Начало: поворот на 45 градусов, масштаб 0.4
func NewComplexRotateOne() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &ScaleEffect{},
				settings: map[string]string{
					"scaleOn":        "0.4",
					"easingFunction": "easeOutQuint",
				},
				startProgress: 0,
				stopProgress:  0.3,
				startValues: map[string]string{
					"startScale":       "0.6",
					"startRotate":      "45",
					"startRotatePoint": "center",
				},
			},
			{
				effectType: &RotateEffect{},
				settings: map[string]string{
					"angle":          "45",
					"direction":      "1",
					"easingFunction": "easeOutElastic",
				},
				startProgress: 0.2,
				stopProgress:  0.4,
				startValues: map[string]string{
					"startRotate":      "45",
					"startRotatePoint": "center",
				},
			},
			{
				effectType: &MoveEffect{},
				settings: map[string]string{
					"distance":       "40",
					"easingFunction": "easeInOutBack",
				},
				startProgress: 0.28,
				stopProgress:  0.5,
			},
			{
				effectType: &MoveEffect{},
				settings: map[string]string{
					"distance": "40",
				},
				startProgress:    0.5,
				stopProgress:     1,
				useOnlyLastState: true,
			},
		},
		map[string]string{
			"startRotate": "0",
			"startScale":  "1",
		},
	)
}
