package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectOptions"
	"github.com/mo3golom/wonder-effects/wonderEffectTransformer"
)

type MoveEffect struct {
	options *wonderEffectOptions.MoveOptions
}

func (m *MoveEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if nil == m.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	progressEasing := float64(wonderEffectMath.ApplyEasing(*progress, m.options.EasingFunction()))

	switch m.options.Direction() {
	case wonderEffectOptions.DirectionLeft:
		effectValues.MoveOnX = -(m.options.Distance() * progressEasing)
		effectValues.MoveOnY = 0
	case wonderEffectOptions.DirectionUp:
		effectValues.MoveOnX = 0
		effectValues.MoveOnY = -(m.options.Distance() * progressEasing)
	case wonderEffectOptions.DirectionRight:
		effectValues.MoveOnX = m.options.Distance() * progressEasing
		effectValues.MoveOnY = 0
	case wonderEffectOptions.DirectionDown:
		effectValues.MoveOnX = 0
		effectValues.MoveOnY = m.options.Distance() * progressEasing
	}

	return nil
}

func (m *MoveEffect) TransformOptions(options *map[string]string) EffectInterface {
	m.options = wonderEffectTransformer.TransformMoveOptions(*options)

	return m
}
