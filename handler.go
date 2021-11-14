package WonderEffects

import (
	"github.com/mo3golom/wonder-effects.git/wonderEffectDTO"
	"github.com/mo3golom/wonder-effects.git/wonderEffectType"
)

// Handler обработчик, котрый можно собрать в цепочку
// Для последовательного применения эффектов к значениям
type Handler struct {
	id          string
	nextHandler *Handler
	effect      wonderEffectType.EffectInterface
}

// NewEffectHandlerBus Конструктор обработчик, который собирает цепочку обработчиков эффектов
func NewEffectHandlerBus(effectTypes []EffectType) *Handler {
	var prevTypeHandler *Handler

	for _, effectType := range effectTypes {
		handler := &Handler{id: effectType.Id(), effect: effectType.Effect()}

		if nil != prevTypeHandler {
			handler.nextHandler = prevTypeHandler
		}

		prevTypeHandler = handler
	}

	return prevTypeHandler
}

// Handle Непосредственная обработка
func (th *Handler) Handle(effect *wonderEffectDTO.Effect, effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	if effect.Id == th.id {

		return th.effect.TransformSettings(&effect.Settings).Processing(effectValues, progress)
	}

	if nil != th.nextHandler {
		return th.nextHandler.Handle(effect, effectValues, progress)
	}

	return nil
}
