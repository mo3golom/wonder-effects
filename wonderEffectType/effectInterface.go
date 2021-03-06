package wonderEffectType

import (
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
)

type EffectInterface interface {
	Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error)
	TransformOptions(options *map[string]interface{}) EffectInterface
}
