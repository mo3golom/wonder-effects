package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects.git/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects.git/wonderEffectMath"
	"github.com/mo3golom/wonder-effects.git/wonderEffectSettings"
	"github.com/mo3golom/wonder-effects.git/wonderEffectTransformer"
)

type RotateEffect struct {
	settings *wonderEffectSettings.RotateSettings
}

func (r *RotateEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == r.settings {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := float64(wonderEffectMath.ApplyEasing(*progress, r.settings.EasingFunction()))
	angle := r.settings.Angle() * progressEasing

	switch r.settings.Direction() {
	case wonderEffectSettings.DirectionNormal:
		effectValues.RotateOn = angle
	case wonderEffectSettings.DirectionReverse:
		effectValues.RotateOn = -angle
	}

	return nil
}

func (r *RotateEffect) TransformSettings(settings *map[string]string) EffectInterface {
	r.settings = wonderEffectTransformer.TransformRotateSettings(*settings)

	return r
}
