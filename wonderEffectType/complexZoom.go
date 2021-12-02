package wonderEffectType

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

func NewComplexZoom() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &MoveEffect{},
				options: map[string]string{
					"distance":       "10",
					"easingFunction": "easeInSine",
					"direction":      "down",
				},
				startProgress: 0,
				stopProgress:  0.2,
				startValues: map[string]string{
					"startDistanceY": "-10",
				},
			},
			{
				effectType: &ScaleEffect{},
				options: map[string]string{
					"scaleOn":        "0.8",
					"easingFunction": "easeOutSine",
				},
				startProgress: 0,
				stopProgress:  0.2,
				startValues: map[string]string{
					"startScale":       "0.2",
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
