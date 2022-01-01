package wonderEffectType

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

func NewComplexZoom() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &MoveEffect{},
				options: map[string]interface{}{
					"distance":       10,
					"easingFunction": "easeInSine",
					"direction":      "down",
				},
				startProgress: 0,
				stopProgress:  0.2,
				startValues: map[string]interface{}{
					"startY": -10,
				},
			},
			{
				effectType: &ScaleEffect{},
				options: map[string]interface{}{
					"scaleOn":        0.8,
					"easingFunction": "easeOutSine",
				},
				startProgress: 0,
				stopProgress:  0.2,
				startValues: map[string]interface{}{
					"startScale":  0.2,
					"rotatePoint": wonderEffectDTO.RotatePointCenter,
				},
			},
		},
		map[string]interface{}{
			"startRotate": 0,
			"startScale":  1,
		},
	)
}
