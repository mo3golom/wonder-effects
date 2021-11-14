package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects.git/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects.git/wonderEffectMath"
	"github.com/mo3golom/wonder-effects.git/wonderEffectSettings"
	"github.com/mo3golom/wonder-effects.git/wonderEffectTransformer"
)

type OpacityEffect struct {
	settings *wonderEffectSettings.OpacitySettings
}

func (o *OpacityEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == o.settings {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := wonderEffectMath.ApplyEasing(*progress, o.settings.EasingFunction())

	switch o.settings.Direction() {
	case wonderEffectSettings.DirectionNormal:
		effectValues.OpacityOn = progressEasing
	case wonderEffectSettings.DirectionReverse:
		effectValues.OpacityOn = -progressEasing
	}

	return nil
}

func (o *OpacityEffect) TransformSettings(settings *map[string]string) EffectInterface {
	o.settings = wonderEffectTransformer.TransformOpacitySettings(*settings)

	return o
}
