package wonderEffectType

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectOptions"
)

type OpacityEffect struct {
	options *wonderEffectOptions.OpacityOptions
}

func (o *OpacityEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == o.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := wonderEffectMath.ApplyEasing(*progress, o.options.EasingFunction)

	switch o.options.Direction {
	case wonderEffectOptions.DirectionNormal:
		effectValues.OpacityOn = progressEasing
	case wonderEffectOptions.DirectionReverse:
		effectValues.OpacityOn = -progressEasing
	}

	return nil
}

func (o *OpacityEffect) TransformOptions(options *map[string]interface{}) EffectInterface {
	opacityOptions := wonderEffectOptions.NewOpacityOptions()
	_ = mapstructure.Decode(*options, opacityOptions)
	_ = mapstructure.Decode(*options, opacityOptions.BaseOptions)

	o.options = opacityOptions

	return o
}
