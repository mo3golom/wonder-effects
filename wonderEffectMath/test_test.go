package wonderEffectMath_test

import (
	"github.com/mo3golom/wonder-effects.git/wonderEffectMath"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEaseInSine(t *testing.T) {
	testEdgeValues(t, "easeInSine")
}

func TestEaseInCubic(t *testing.T) {
	testEdgeValues(t, "easeInCubic")
}

func TestEaseInQuint(t *testing.T) {
	testEdgeValues(t, "easeInQuint")
}

func TestEaseInCirc(t *testing.T) {
	testEdgeValues(t, "easeInCirc")
}

func TestEaseInElastic(t *testing.T) {
	testEdgeValues(t, "easeInElastic")
}

func TestEaseOutSine(t *testing.T) {
	testEdgeValues(t, "easeOutSine")
}

func TestEaseOutCubic(t *testing.T) {
	testEdgeValues(t, "easeOutCubic")
}

func TestEaseOutQuint(t *testing.T) {
	testEdgeValues(t, "easeOutQuint")
}

func TestEaseOutCirc(t *testing.T) {
	testEdgeValues(t, "easeOutCirc")
}

func TestEaseOutElastic(t *testing.T) {
	testEdgeValues(t, "easeOutElastic")
}

func TestEaseInOutSine(t *testing.T) {
	testEdgeValues(t, "easeInOutSine")
}

func TestEaseInOutCubic(t *testing.T) {
	testEdgeValues(t, "easeInOutCubic")
}

func TestEaseInOutQuint(t *testing.T) {
	testEdgeValues(t, "easeInOutQuint")
}

func TestEaseInOutCirc(t *testing.T) {
	testEdgeValues(t, "easeInOutCirc")
}

func TestEaseInOutBack(t *testing.T) {
	testEdgeValues(t, "easeInOutBack")
}

func TestEaseInBounce(t *testing.T) {
	testEdgeValues(t, "easeInBounce")
}

func TestEaseOutBounce(t *testing.T) {
	testEdgeValues(t, "easeOutBounce")
}

func TestEaseInOutBounce(t *testing.T) {
	testEdgeValues(t, "easeInOutBounce")
}

func testEdgeValues(t *testing.T, functionName string) {
	var val float32

	val = wonderEffectMath.ApplyEasing(0, functionName)
	assert.Equal(t, float32(0), val)

	val = wonderEffectMath.ApplyEasing(1, functionName)
	assert.Equal(t, float32(1), val)
}
