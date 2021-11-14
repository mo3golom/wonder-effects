package WonderEffects

import (
	"github.com/mo3golom/wonder-effects.git/wonderEffectType"
)

type EffectType struct {
	id     string
	name   string
	effect wonderEffectType.EffectInterface
}

func (e *EffectType) Id() string {
	return e.id
}

func (e *EffectType) Name() string {
	return e.name
}

func (e *EffectType) Effect() wonderEffectType.EffectInterface {
	return e.effect
}

// GetTypesList Возвращает список доступных эффектов
func GetTypesList() []EffectType {
	return []EffectType{
		{id: "move", name: "Движение", effect: &wonderEffectType.MoveEffect{}},
		{id: "rotate", name: "Вращение", effect: &wonderEffectType.RotateEffect{}},
		{id: "opacity", name: "Прозрачность", effect: &wonderEffectType.OpacityEffect{}},
		{id: "scale", name: "Масштаб", effect: &wonderEffectType.ScaleEffect{}},
		{id: "complexRotateOne", name: "Комплексный эффект вращения 01", effect: wonderEffectType.NewComplexRotateOne()},
		{id: "complexRotateTwo", name: "Комплексный эффект вращения 02", effect: wonderEffectType.NewComplexRotateTwo()},
		{id: "complexZoom", name: "Комплексный эффект увеличения", effect: wonderEffectType.NewComplexZoom()},
		{id: "complexPoplavokOne", name: "Комплексный эффект поплавок 01", effect: wonderEffectType.NewComplexPoplavokOne()},
		{id: "complexPoplavokTwo", name: "Комплексный эффект поплавок 02", effect: wonderEffectType.NewComplexPoplavokTwo()},
	}
}
