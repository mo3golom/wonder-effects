package wonderEffectTransformer

import (
	"github.com/mo3golom/wonder-effects.git/wonderEffectSettings"
	"strconv"
)

func TransformMoveSettings(settings map[string]string) *wonderEffectSettings.MoveSettings {
	moveSettings := wonderEffectSettings.NewMoveSettings()

	setBaseSettings(moveSettings.BaseSettings, settings)

	distance, err := strconv.ParseFloat(settings["distance"], 64)

	if nil == err {
		moveSettings.SetDistance(distance)
	}

	return moveSettings
}

func TransformRotateSettings(settings map[string]string) *wonderEffectSettings.RotateSettings {
	rotateSettings := wonderEffectSettings.NewRotateSettings()

	setBaseSettings(rotateSettings.BaseSettings, settings)

	angle, err := strconv.ParseFloat(settings["angle"], 64)

	if nil == err {
		rotateSettings.SetAngle(angle)
	}

	return rotateSettings
}

func TransformOpacitySettings(settings map[string]string) *wonderEffectSettings.OpacitySettings {
	opacitySettings := wonderEffectSettings.NewOpacitySettings()

	setBaseSettings(opacitySettings.BaseSettings, settings)

	return opacitySettings
}

func TransformScaleSettings(settings map[string]string) *wonderEffectSettings.ScaleSettings {
	scaleSettings := wonderEffectSettings.NewScaleSettings()

	setBaseSettings(scaleSettings.BaseSettings, settings)

	scaleOn, err := strconv.ParseFloat(settings["scaleOn"], 64)

	if nil == err {
		scaleSettings.SetScaleOn(scaleOn)
	}

	return scaleSettings
}

// Устанваливает "базовые настройки"
func setBaseSettings(baseSettings *wonderEffectSettings.BaseSettings, settings map[string]string) {
	direction, err := strconv.Atoi(settings["direction"])

	if nil == err {
		baseSettings.SetDirection(direction)
	}

	easingFunction, ok := settings["easingFunction"]

	if ok {
		baseSettings.SetEasingFunction(easingFunction)
	}
}
