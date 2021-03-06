package wonderEffectType

import "github.com/mo3golom/wonder-effects/wonderEffectDTO"

// NewComplexRotateOne Комплексный эффект вращения 01
// Начало: поворот на 45 градусов, масштаб 0.4
func NewComplexRotateOne() *ComplexEffect {
	return NewComplexEffect(
		[]effectPart{
			{
				effectType: &ScaleEffect{},
				options: map[string]interface{}{
					"scaleOn":        0.4,
					"easingFunction": "easeOutQuint",
				},
				startProgress: 0,
				stopProgress:  0.3,
				startValues: map[string]interface{}{
					"startScale":  0.6,
					"startRotate": 45,
					"rotatePoint": wonderEffectDTO.RotatePointCenter,
				},
			},
			{
				effectType: &RotateEffect{},
				options: map[string]interface{}{
					"angle":          45,
					"direction":      "reverse",
					"easingFunction": "easeOutElastic",
				},
				startProgress: 0.2,
				stopProgress:  0.4,
				startValues: map[string]interface{}{
					"startRotate": 45,
					"rotatePoint": wonderEffectDTO.RotatePointCenter,
				},
			},
			{
				effectType: &MoveEffect{},
				options: map[string]interface{}{
					"distance":       40,
					"easingFunction": "easeInOutBack",
				},
				startProgress: 0.28,
				stopProgress:  0.5,
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
		},
		map[string]interface{}{
			"startRotate": 0,
			"startScale":  1,
		},
	)
}
