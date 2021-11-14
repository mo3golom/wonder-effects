package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectSettings"
	"github.com/mo3golom/wonder-effects/wonderEffectTransformer"
)

type ScaleEffect struct {
	settings *wonderEffectSettings.ScaleSettings
}

func (s *ScaleEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == s.settings {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := float64(wonderEffectMath.ApplyEasing(*progress, s.settings.EasingFunction()))
	additional := s.settings.ScaleOn() * progressEasing

	switch s.settings.Direction() {
	case wonderEffectSettings.DirectionNormal:
		effectValues.ScaleOn = additional
	case wonderEffectSettings.DirectionReverse:
		effectValues.ScaleOn = -additional
	}

	return nil
}

func (s *ScaleEffect) TransformSettings(settings *map[string]string) EffectInterface {
	s.settings = wonderEffectTransformer.TransformScaleSettings(*settings)

	return s
}
