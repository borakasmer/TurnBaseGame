package druid

import (
	"math"
	hero "turnBaseGame/heroCharacter"
)

type (
	Druid struct {
		hero.Hero
		Manas  map[string]int
		Mana   int
		Animal string
	}
)

func (d Druid) Hit() int {
	if d.Mana >= d.Manas[d.Animal] {
		levelEffect := int(math.Ceil(float64(d.Level) * 0.1))
		return d.Attacks[d.Animal] + levelEffect
	} else {
		return 0
	}

}

func (d *Druid) TakeDamage(damage int) {
	d.Blood = d.Blood - damage
}

func (d Druid) GetInfo() (string, int) {
	return d.Name, d.Blood
}

func (d Druid) GetMana() int {
	return d.Mana
}

func (d *Druid) SpendMana() {
	if d.Mana >= d.Manas[d.Animal] {
		d.Mana = d.Mana - d.Manas[d.Animal]
	}
}

func (d Druid) IsDeath() bool {
	return d.Blood <= 0
}

func CreateDruid(blood int, mana int) *Druid {
	return &Druid{Hero: hero.Hero{
		Attacks: map[string]int{"Wolf": 15, "Bear": 35, "Sheap": 5},
		//Blood:   80,
		Blood: blood,
	},
		//Mana:  100,
		Mana:  mana,
		Manas: map[string]int{"Wolf": 5, "Bear": 15, "Sheap": 1}}
}
