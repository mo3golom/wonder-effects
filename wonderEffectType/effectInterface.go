package wonderEffectType

import (
	"github.com/mo3golom/wonder-effects.git/wonderEffectDTO"
)

type EffectInterface interface {
	Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error)
	TransformSettings(settings *map[string]string) EffectInterface
}
