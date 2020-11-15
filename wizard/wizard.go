package wizard

import (
	"math"
	hero "turnBaseGame/heroCharacter"
)

type (
	Wizard struct {
		hero.Hero
		Manas map[string]int
		Mana  int
		Magic string
	}
)

func (w Wizard) Hit() int {
	if w.Mana >= w.Manas[w.Magic] {
		levelEffect := int(math.Ceil(float64(w.Level) * 0.1))
		return w.Attacks[w.Magic] + levelEffect
	} else {
		return 0
	}
}

func (w *Wizard) TakeDamage(damage int) {
	w.Blood = w.Blood - damage
}

func (w Wizard) GetInfo() (string, int) {
	return w.Name, w.Blood
}

func (w Wizard) GetMana() int {
	return w.Mana
}

func (w *Wizard) SpendMana() {
	if w.Mana >= w.Manas[w.Magic] {
		w.Mana = w.Mana - w.Manas[w.Magic]
	}
}

func (w Wizard) IsDeath() bool {
	return w.Blood <= 0
}

func (w *Wizard) SetRandomMagic() {

}

func CreateWizard(blood int, mana int) *Wizard {
	return &Wizard{Hero: hero.Hero{
		Attacks: map[string]int{"FireBall": 25, "Thunder": 18, "Ghost Attack": 30},
		//Blood:   80,
		Blood: blood,
	},
		//Mana:  100,
		Mana:  mana,
		Manas: map[string]int{"FireBall": 5, "Thunder": 10, "Ghost Attack": 30}}
}
