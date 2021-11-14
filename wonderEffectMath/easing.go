package wonderEffectMath

import "math"

// ApplyEasing Возможно данная фукнция выглядит ужасно
// Но это максимальная и быстрая реализация применения функций плавности
// И предполагаю, что здесь что-то очень редко будет меняться (кроме добавления еще нериализованных функций)
func ApplyEasing(t float32, easingFunction string) float32 {
	switch easingFunction {
	case "easeInSine":
		t = easeInSine(t)
	case "easeInCubic":
		t = easeInCubic(t)
	case "easeInQuint":
		t = easeInQuint(t)
	case "easeInCirc":
		t = easeInCirc(t)
	case "easeInElastic":
		t = easeInElastic(t)
	case "easeOutSine":
		t = easeOutSine(t)
	case "easeOutCubic":
		t = easeOutCubic(t)
	case "easeOutQuint":
		t = easeOutQuint(t)
	case "easeOutCirc":
		t = easeOutCirc(t)
	case "easeOutElastic":
		t = easeOutElastic(t)
	case "easeInOutSine":
		t = easeInOutSine(t)
	case "easeInOutCubic":
		t = easeInOutCubic(t)
	case "easeInOutQuint":
		t = easeInOutQuint(t)
	case "easeInOutCirc":
		t = easeInOutCirc(t)
	case "easeInOutElastic":
		t = easeInOutElastic(t)
	case "easeInOutBack":
		t = easeInOutBack(t)
	case "easeInBounce":
		t = easeInBounce(t)
	case "easeOutBounce":
		t = easeOutBounce(t)
	case "easeInOutBounce":
		t = easeInOutBounce(t)
	}

	return t
}

// Блок easeIn

// https://easings.net/ru#easeInSine
func easeInSine(t float32) float32 {
	return float32(1 - math.Cos((float64(t)*math.Pi)/2))
}

// https://easings.net/ru#easeInCubic
func easeInCubic(t float32) float32 {
	return t * t * t
}

// https://easings.net/ru#easeInQuint
func easeInQuint(t float32) float32 {
	return t * t * t * t * t
}

// https://easings.net/ru#easeInCirc
func easeInCirc(t float32) float32 {
	return 1 - float32(math.Sqrt(1-math.Pow(float64(t), 2)))
}

// https://easings.net/ru#easeInElastic
func easeInElastic(t float32) float32 {
	if 0 == t {
		return 0
	}

	if 1 == t {
		return 1
	}

	return float32(-math.Pow(2, 10*float64(t)-10) * math.Sin((float64(t)*10-10.75)*((2*math.Pi)/3)))
}

// Блок easeOut

// https://easings.net/ru#easeOutSine
func easeOutSine(t float32) float32 {
	return float32(math.Sin((float64(t) * math.Pi) / 2))
}

// https://easings.net/ru#easeOutCubic
func easeOutCubic(t float32) float32 {
	return 1 - float32(math.Pow(1-float64(t), 3))
}

// https://easings.net/ru#easeOutQuint
func easeOutQuint(t float32) float32 {
	return 1 - float32(math.Pow(1-float64(t), 5))
}

// https://easings.net/ru#easeOutCirc
func easeOutCirc(t float32) float32 {
	return float32(math.Sqrt(1 - math.Pow(float64(t)-1, 2)))
}

// https://easings.net/ru#easeOutElastic
func easeOutElastic(t float32) float32 {
	if 0 == t {
		return 0
	}

	if 1 == t {
		return 1
	}

	return float32(math.Pow(2, -10*float64(t))*math.Sin((float64(t)*10-0.75)*(2*math.Pi)/3)) + 1
}

// Блок easeInOut

// https://easings.net/ru#easeInOutSine
func easeInOutSine(t float32) float32 {
	return -float32(math.Cos(math.Pi*float64(t))-1) / 2
}

// https://easings.net/ru#easeInOutCubic
func easeInOutCubic(t float32) float32 {
	if 0.5 > t {
		return 4 * t * t * t
	}

	return 1 - float32(math.Pow(-2*float64(t)+2, 3)/2)
}

// https://easings.net/ru#easeInOutQuint
func easeInOutQuint(t float32) float32 {
	if 0.5 > t {
		return 16 * t * t * t * t * t
	}

	return 1 - float32(math.Pow(-2*float64(t)+2, 5)/2)
}

// https://easings.net/ru#easeInOutCirc
func easeInOutCirc(t float32) float32 {
	if 0.5 > t {
		return float32(1-math.Sqrt(1-math.Pow(2*float64(t), 2))) / 2
	}

	return float32(math.Sqrt(1-math.Pow(-2*float64(t)+2, 2))+1) / 2
}

// https://easings.net/ru#easeInOutElastic
func easeInOutElastic(t float32) float32 {
	const c5 = (2 * math.Pi) / 4.5

	if 0 == t {
		return 0
	}

	if 1 == t {
		return 1
	}

	if 0.5 > t {
		return -float32(math.Pow(2, 20*float64(t)-10)*math.Sin((20*float64(t)-11.125)*c5)) / 2
	}

	return float32(math.Pow(2, -20*float64(t)+10)*math.Sin((20*float64(t)-11.125)*c5))/2 + 1
}

// Блок Back

func easeInOutBack(t float32) float32 {
	const c1 = 1.70158
	const c2 = c1 * 1.525

	if 0.5 > t {
		return float32(math.Pow(2*float64(t), 2)*((c2+1)*2*float64(t)-c2)) / 2
	}

	return float32(math.Pow(2*float64(t)-2, 2)*((c2+1)*(float64(t)*2-2)+c2)+2) / 2
}

// Блок Bounce

// https://easings.net/ru#easeInBounce
func easeInBounce(t float32) float32 {
	return 1 - easeOutBounce(1-t)
}

// https://easings.net/ru#easeOutBounce
func easeOutBounce(t float32) float32 {
	const n1 = 7.5625
	const d1 = 2.75

	if 1/d1 > t {
		return n1 * t * t
	}

	if 2/d1 > t {
		t2 := t - 1.5/d1

		return n1*t2*t2 + 0.75
	}

	if 2.5/d1 > t {
		t2 := t - 2.25/d1

		return n1*t2*t2 + 0.9375
	}

	t2 := t - 2.625/d1

	return n1*t2*t2 + 0.984375
}

// https://easings.net/ru#easeInOutBounce
func easeInOutBounce(t float32) float32 {
	if 0.5 > t {
		return (1 - easeOutBounce(1-2*t)) / 2
	}

	return (1 + easeOutBounce(2*t-1)) / 2
}
