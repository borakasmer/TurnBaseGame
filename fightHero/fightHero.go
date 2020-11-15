package fightHero

type FightHero interface {
	Hit() int
	TakeDamage(damage int)
	GetInfo() (string, int)
	IsDeath() bool
}
