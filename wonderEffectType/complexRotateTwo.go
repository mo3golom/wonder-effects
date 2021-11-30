package wonderEffectType

// NewComplexRotateTwo Комплексный эффект вращения 02
func NewComplexRotateTwo() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &ScaleEffect{},
				settings: map[string]string{
					"scaleOn":        "-0.4",
					"easingFunction": "easeOutQuint",
				},
				startProgress: 0,
				stopProgress:  0.3,
				startValues: map[string]string{
					"startScale":       "1.4",
					"startRotatePoint": "center",
				},
			},
			{
				effectType: &MoveEffect{},
				settings: map[string]string{
					"distance":       "40",
					"easingFunction": "easeInOutBack",
				},
				startProgress: 0.05,
				stopProgress:  0.5,
			},
			{
				effectType: &RotateEffect{},
				settings: map[string]string{
					"angle":          "45",
					"direction":      "1",
					"easingFunction": "easeOutElastic",
				},
				startProgress: 0.5,
				stopProgress:  0.8,
				startValues: map[string]string{
					"startRotate":      "0",
					"startRotatePoint": "center",
				},
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
			{
				effectType: &RotateEffect{},
				settings: map[string]string{
					"angle":     "45",
					"direction": "1",
				},
				startProgress:    0.8,
				stopProgress:     1,
				useOnlyLastState: true,
				startValues: map[string]string{
					"startRotate":      "0",
					"startRotatePoint": "center",
				},
			},
		},
		map[string]string{
			"startRotate": "0",
			"startScale":  "1",
		},
	)
}
