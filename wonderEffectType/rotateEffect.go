package wonderEffectType

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectOptions"
)

type RotateEffect struct {
	options *wonderEffectOptions.RotateOptions
}

func (r *RotateEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == r.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := float64(wonderEffectMath.ApplyEasing(*progress, r.options.EasingFunction))
	angle := r.options.Angle * progressEasing

	switch r.options.Direction {
	case wonderEffectOptions.DirectionNormal:
		effectValues.RotateOn = angle
	case wonderEffectOptions.DirectionReverse:
		effectValues.RotateOn = -angle
	}

	return nil
}

func (r *RotateEffect) TransformOptions(options *map[string]interface{}) EffectInterface {
	rotateOptions := wonderEffectOptions.NewRotateOptions()
	_ = mapstructure.Decode(*options, rotateOptions)
	_ = mapstructure.Decode(*options, rotateOptions.BaseOptions)

	r.options = rotateOptions

	return r
}
