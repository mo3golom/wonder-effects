package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectOptions"
	"github.com/mo3golom/wonder-effects/wonderEffectTransformer"
)

type ScaleEffect struct {
	options *wonderEffectOptions.ScaleOptions
}

func (s *ScaleEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == s.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := float64(wonderEffectMath.ApplyEasing(*progress, s.options.EasingFunction()))
	additional := s.options.ScaleOn() * progressEasing

	switch s.options.Direction() {
	case wonderEffectOptions.DirectionNormal:
		effectValues.ScaleOn = additional
	case wonderEffectOptions.DirectionReverse:
		effectValues.ScaleOn = -additional
	}

	return nil
}

func (s *ScaleEffect) TransformOptions(options *map[string]string) EffectInterface {
	s.options = wonderEffectTransformer.TransformScaleOptions(*options)

	return s
}
