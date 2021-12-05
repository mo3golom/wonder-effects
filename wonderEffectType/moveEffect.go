package wonderEffectType

import (
	"errors"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects/wonderEffectMath"
	"github.com/mo3golom/wonder-effects/wonderEffectOptions"
	"github.com/mo3golom/wonder-effects/wonderEffectTransformer"
	"math"
)

const pathMultiplier = 2

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
		effectValues.MoveOnY = -m.applyPathFunction(effectValues.MoveOnX)
	case wonderEffectOptions.DirectionUp:
		effectValues.MoveOnY = -(m.options.Distance() * progressEasing)
		effectValues.MoveOnX = m.applyPathFunction(effectValues.MoveOnY)
	case wonderEffectOptions.DirectionRight:
		effectValues.MoveOnX = m.options.Distance() * progressEasing
		effectValues.MoveOnY = m.applyPathFunction(effectValues.MoveOnX)
	case wonderEffectOptions.DirectionDown:
		effectValues.MoveOnY = m.options.Distance() * progressEasing
		effectValues.MoveOnX = -m.applyPathFunction(effectValues.MoveOnY)
	}

	return nil
}

func (m *MoveEffect) TransformOptions(options *map[string]string) EffectInterface {
	m.options = wonderEffectTransformer.TransformMoveOptions(*options)

	return m
}

func (m *MoveEffect) applyPathFunction(value float64) float64 {

	switch m.options.PathFunction() {
	case "sin":
		return math.Sin(0.5*value) * pathMultiplier

	case "cos":
		return math.Cos(0.5*value) * pathMultiplier

	case "circle":
		return m.options.Distance() - math.Sqrt(math.Pow(m.options.Distance(), 2)-math.Pow(value, 2))
	}

	return 0
}
