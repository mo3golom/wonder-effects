package wonderEffectTransformer

import (
	"github.com/mo3golom/wonder-effects/wonderEffectOptions"
	"strconv"
)

func TransformMoveOptions(options map[string]string) *wonderEffectOptions.MoveOptions {
	moveOptions := wonderEffectOptions.NewMoveOptions()

	setBaseOptions(moveOptions.BaseOptions, options)

	distance, err := strconv.ParseFloat(options["distance"], 64)

	if nil == err {
		moveOptions.SetDistance(distance)
	}

	pathFunction, ok := options["pathFunction"]

	if ok {
		moveOptions.SetPathFunction(pathFunction)
	}

	return moveOptions
}

func TransformRotateOptions(options map[string]string) *wonderEffectOptions.RotateOptions {
	rotateOptions := wonderEffectOptions.NewRotateOptions()

	setBaseOptions(rotateOptions.BaseOptions, options)

	angle, err := strconv.ParseFloat(options["angle"], 64)

	if nil == err {
		rotateOptions.SetAngle(angle)
	}

	return rotateOptions
}

func TransformOpacityOptions(options map[string]string) *wonderEffectOptions.OpacityOptions {
	opacityOptions := wonderEffectOptions.NewOpacityOptions()

	setBaseOptions(opacityOptions.BaseOptions, options)

	return opacityOptions
}

func TransformScaleOptions(options map[string]string) *wonderEffectOptions.ScaleOptions {
	scaleOptions := wonderEffectOptions.NewScaleOptions()

	setBaseOptions(scaleOptions.BaseOptions, options)

	scaleOn, err := strconv.ParseFloat(options["scaleOn"], 64)

	if nil == err {
		scaleOptions.SetScaleOn(scaleOn)
	}

	return scaleOptions
}

// Устанавливает "базовые настройки"
func setBaseOptions(baseOptions *wonderEffectOptions.BaseOptions, options map[string]string) {
	direction, ok := options["direction"]

	if ok {
		baseOptions.SetDirection(direction)
	}

	easingFunction, ok := options["easingFunction"]

	if ok {
		baseOptions.SetEasingFunction(easingFunction)
	}
}
