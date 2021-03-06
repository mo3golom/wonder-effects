package wonderEffectType

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

// NewComplexRotateTwo Комплексный эффект вращения 02
func NewComplexRotateTwo() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &ScaleEffect{},
				options: map[string]interface{}{
					"scaleOn":        -0.4,
					"easingFunction": "easeOutQuint",
				},
				startProgress: 0,
				stopProgress:  0.3,
				startValues: map[string]interface{}{
					"startScale":  1.4,
					"rotatePoint": wonderEffectDTO.RotatePointCenter,
				},
			},
			{
				effectType: &MoveEffect{},
				options: map[string]interface{}{
					"distance":       40,
					"easingFunction": "easeInOutBack",
				},
				startProgress: 0.05,
				stopProgress:  0.5,
			},
			{
				effectType: &RotateEffect{},
				options: map[string]interface{}{
					"angle":          45,
					"direction":      "reverse",
					"easingFunction": "easeOutElastic",
				},
				startProgress: 0.5,
				stopProgress:  0.8,
				startValues: map[string]interface{}{
					"startRotate": 0,
					"rotatePoint": wonderEffectDTO.RotatePointCenter,
				},
			},
			{
				effectType: &MoveEffect{},
				options: map[string]interface{}{
					"distance": 40,
				},
				startProgress:    0.5,
				stopProgress:     1,
				useOnlyLastState: true,
			},
			{
				effectType: &RotateEffect{},
				options: map[string]interface{}{
					"angle":     45,
					"direction": "reverse",
				},
				startProgress:    0.8,
				stopProgress:     1,
				useOnlyLastState: true,
				startValues: map[string]interface{}{
					"startRotate": 0,
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
