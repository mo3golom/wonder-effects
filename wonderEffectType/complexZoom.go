package wonderEffectType

func NewComplexZoom() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &MoveEffect{},
				settings: map[string]string{
					"distance":       "10",
					"easingFunction": "easeInSine",
					"direction":      "3",
				},
				startProgress: 0,
				stopProgress:  0.2,
				startValues: map[string]string{
					"startDistanceY": "-10",
				},
			},
			{
				effectType: &ScaleEffect{},
				settings: map[string]string{
					"scaleOn":        "0.8",
					"easingFunction": "easeOutSine",
				},
				startProgress: 0,
				stopProgress:  0.2,
				startValues: map[string]string{
					"startScale": "0.2",
				},
			},
		},
		map[string]string{
			"startRotate": "0",
			"startScale":  "1",
		},
	)
}
