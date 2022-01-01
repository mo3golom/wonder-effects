package wonderEffectType

import (
	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
)

// Часть комплексного эффекта
type effectPart struct {
	effectType                  EffectInterface        // эффект, который нужно применить
	options                     map[string]interface{} // настройки для эффекта
	startProgress, stopProgress float32                // старт и конец на линии "глобального" прогресса
	startValues                 map[string]interface{} // стартовые значения, переопределяют Start значения в effectValues
	useOnlyLastState            bool                   // Использовать только конечное состояние в эффекте (100% прогресса)
}

// ComplexEffect Вспомогательный асбтрактный комплексный эффект
// Позволяет создавать конкретные реализации комплексных эффектов через конфигурацию в конструкторе
type ComplexEffect struct {
	effects     []effectPart
	startValues map[string]interface{}
}

func NewComplexEffect(effects []effectPart, startValues map[string]interface{}) *ComplexEffect {
	return &ComplexEffect{
		effects:     effects,
		startValues: startValues,
	}
}

func (a *ComplexEffect) Processing(effectValues *wonderEffectDTO.EffectValues, progress *float32) (err error) {
	a.setStartValues(effectValues, a.startValues)

	for _, effectPart := range a.effects {
		// Если текущий прогресс не входит в промежуток эффекта либо не является эффектом последнего состояния, то пропускаем эффект
		if effectPart.startProgress > *progress || *progress > effectPart.stopProgress {
			continue
		}

		// Преобразуем "глобальный" прогресс в "локальный" (по-умолчанию 100%)
		var effectPartProgress float32 = 1

		// Если не используем только конечное состояние, то вычисляем актуальный "локальный" прогресс
		if !effectPart.useOnlyLastState {
			effectPartProgress = (*progress - effectPart.startProgress) / (effectPart.stopProgress - effectPart.startProgress)
		}

		// Устанавливаем стартовые значения для конкретного эффекта
		a.setStartValues(effectValues, effectPart.startValues)

		// Применяем эффект
		err = effectPart.effectType.TransformOptions(&effectPart.options).Processing(effectValues, &effectPartProgress)

		if nil != err {
			return err
		}
	}

	return nil
}

func (a *ComplexEffect) TransformOptions(_ *map[string]interface{}) EffectInterface {
	return a
}

func (a *ComplexEffect) setStartValues(effectValues *wonderEffectDTO.EffectValues, startValues map[string]interface{}) {
	_ = mapstructure.Decode(startValues, effectValues)
}
