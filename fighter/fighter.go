package fighter

import (
	"math"
	hero "turnBaseGame/heroCharacter"
)

type Fighter struct {
	hero.Hero
	Staminas map[string]int
	Stamina  int
	Weapon   string
}

func (f Fighter) Hit() int {
	if f.Stamina >= f.Staminas[f.Weapon] {
		levelEffect := int(math.Ceil(float64(f.Level) * 0.1))
		return f.Attacks[f.Weapon] + levelEffect
	} else {
		return 0
	}
}

func (f *Fighter) TakeDamage(damage int) {
	f.Blood = f.Blood - damage
}

func (f Fighter) GetInfo() (string, int) {
	return f.Name, f.Blood
}

func (f Fighter) GetStamina() int {
	return f.Stamina
}

func (f *Fighter) SpendStamina() {
	if f.Stamina >= f.Staminas[f.Weapon] {
		f.Stamina = f.Stamina - f.Staminas[f.Weapon]
	}
}

func (f Fighter) IsDeath() bool {
	return f.Blood <= 0
}

func CreateFighter(blood int, stamina int) *Fighter {
	return &Fighter{Hero: hero.Hero{
		Attacks: map[string]int{"Stick": 15, "Axe": 30, "Sword": 40},
		//Blood:   100,
		Blood: blood,
	},
		//Stamina:  100,
		Stamina:  stamina,
		Staminas: map[string]int{"Axe": 10, "Stick": 5, "Sword": 30},
	}
}
