package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectSettings"
	"github.com/mo3golom/wonder-effects/wonderEffectTransformer"
)

type MoveEffect struct {
	settings *wonderEffectSettings.MoveSettings
}

func (m *MoveEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == m.settings {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := float64(wonderEffectMath.ApplyEasing(*progress, m.settings.EasingFunction()))

	switch m.settings.Direction() {
	case wonderEffectSettings.DirectionLeft:
		effectValues.MoveOnX = -(m.settings.Distance() * progressEasing)
		effectValues.MoveOnY = 0
	case wonderEffectSettings.DirectionUp:
		effectValues.MoveOnX = 0
		effectValues.MoveOnY = -(m.settings.Distance() * progressEasing)
	case wonderEffectSettings.DirectionRight:
		effectValues.MoveOnX = m.settings.Distance() * progressEasing
		effectValues.MoveOnY = 0
	case wonderEffectSettings.DirectionDown:
		effectValues.MoveOnX = 0
		effectValues.MoveOnY = m.settings.Distance() * progressEasing
	}

	return nil
}

func (m *MoveEffect) TransformSettings(settings *map[string]string) EffectInterface {
	m.settings = wonderEffectTransformer.TransformMoveSettings(*settings)

	return m
}
