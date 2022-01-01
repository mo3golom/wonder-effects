package wonderEffectType

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

func NewComplexPoplavokOne() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &MoveEffect{},
				options: map[string]interface{}{
					"distance":       30,
					"easingFunction": "easeOutElastic",
					"direction":      "left",
				},
				startProgress: 0,
				stopProgress:  0.4,
				startValues: map[string]interface{}{
					"startX": 30,
				},
			},
			{
				effectType: &RotateEffect{},
				options: map[string]interface{}{
					"angle":          15,
					"easingFunction": "easeOutElastic",
					"direction":      "normal",
				},
				startProgress: 0,
				stopProgress:  0.4,
				startValues: map[string]interface{}{
					"startRotate": -15,
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
