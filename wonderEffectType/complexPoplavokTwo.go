package wonderEffectType

func NewComplexPoplavokTwo() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &MoveEffect{},
				settings: map[string]string{
					"distance":       "30",
					"easingFunction": "easeOutElastic",
					"direction":      "0",
				},
				startProgress: 0,
				stopProgress:  0.4,
				startValues: map[string]string{
					"startDistanceX": "30",
				},
			},
			{
				effectType: &RotateEffect{},
				settings: map[string]string{
					"angle":          "15",
					"easingFunction": "easeOutElastic",
					"direction":      "1",
				},
				startProgress: 0,
				stopProgress:  0.4,
				startValues: map[string]string{
					"startRotate": "15",
				},
			},
		},
		map[string]string{
			"startRotate": "0",
			"startScale":  "1",
		},
	)
}
