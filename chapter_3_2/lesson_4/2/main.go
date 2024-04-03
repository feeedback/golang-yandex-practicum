package main // Тема 2/6: Интерфейсы → Урок 4/5 "Рефлексия"
// https://practicum.yandex.ru/trainer/go-basics/lesson/169fa0e0-eff1-4398-92bf-7abada970e8f/

import (
	"fmt"
	"log"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {
	if recv, ok := object.(CastReceiver); ok {
		recv.ReceiveSpell(spell)
		return
	}

	val := reflect.ValueOf(object)

	// проверяем, что переданный объект — указатель на структуру
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return
	}

	// ищем в структуре нужную характеристику
	field := val.Elem().FieldByName(spell.Char())

	// не нашли
	if !field.IsValid() {
		return
	}

	// нашли, но изменить её нельзя
	if !field.CanSet() {
		return
	}

	// тип найденного поля — не целое число
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// тип найденного поля — целое число
		field.SetInt(field.Int() + int64(spell.Value()))

		log.Printf("Casted spell %s to %#v", spell.Name(), object)

	default:
		// тип найденного поля — не целое число
		return
	}

}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}

func main() {

	player := &Player{
		name:   "Player_1",
		health: 100,
	}

	enemies := []interface{}{
		&Zombie{Health: 1000},
		&Zombie{Health: 1000},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Daemon{Health: 1000},
		&Daemon{Health: 1000},
		&Wall{Durability: 100},
	}

	CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
	CastToAll(newSpell("heal", "Health", 190), append(enemies, player))

	fmt.Println(player)
}
