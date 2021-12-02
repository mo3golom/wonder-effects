package wonderEffectType

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

func NewComplexPoplavokTwo() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &MoveEffect{},
				options: map[string]string{
					"distance":       "30",
					"easingFunction": "easeOutElastic",
					"direction":      "left",
				},
				startProgress: 0,
				stopProgress:  0.4,
				startValues: map[string]string{
					"startDistanceX": "30",
				},
			},
			{
				effectType: &RotateEffect{},
				options: map[string]string{
					"angle":          "15",
					"easingFunction": "easeOutElastic",
					"direction":      "reverse",
				},
				startProgress: 0,
				stopProgress:  0.4,
				startValues: map[string]string{
					"startRotate":      "15",
					"startRotatePoint": wonderEffectDTO.RotatePointCenter,
				},
			},
		},
		map[string]string{
			"startRotate": "0",
			"startScale":  "1",
		},
	)
}
