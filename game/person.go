package game

import (
	"math/rand"
	"time"
)

type Job int

const (
	Break Job = iota
	Exploring
	Returning
	Waiting
	Rummaging
	Entrance
	JobPower
	JobWater
	JobFood
	JobMeal

	TrainingCharm
	TrainingAwareness
	TrainingPower
	TrainingSuccess
	TrainingTenacity
	// TrainingOutgoing cannot train, born with it
	TrainingNimble
	TrainingErudite
)

type PersonID int

type Person struct {
	ID   PersonID
	Name string
	// 0 to 60, hit cap at 120,000 total acquired skill points
	Level int
	// 0 to 100
	Mood       int
	CurrentJob Job

	// Cached Attributes
	TotalHP    int
	CurrentHP  int
	ReviveCost int

	// HP things
	IsDead        bool
	Expiration    time.Time
	DamageAmount  int
	RadiateAmount int
	HealthLevel   int

	// Items
	Clothing
	Weapon

	// Stats range from 1 to 10, don't include clothing effects
	Stats SPECIAL
	// Skills increase when you make a successful Stat check
	// Enough checks will increase the equivalent Stat
	// Stat increase amounts:
	//   500 stat of 2
	//   1000 stat of 3
	//   2000 stat of 4
	//   4000 stat of 5
	//   8000 stat of 6
	//   16000 stat of 7
	//   32000 stat of 8
	//   64000 stat of 9
	//   128000 stat of 9
	Skills SPECIAL
}

type SPECIAL struct {
	Charm     int // charisma
	Awareness int // perception
	Power     int // strength
	Success   int // luck
	Tenacity  int // endurance
	Outgoing  int // social
	Nimble    int // agility
	Erudite   int // intelligence

	Luck int
}

func (p *Person) Damage(d int) {
	p.DamageAmount += d
	p.applyDamage()
}

func (p *Person) TickHealth(d time.Duration) {
	switch p.HealthLevel {
	// bad food/water
	case 0:
		// add a random amount of radiation based on level
		// drop mood

	// poor food/water
	case 1:
		// add a random amount of radiation based on level
		// drop mood

	// wasteland
	case 2:
		// heal a bit of radiation and health if endurance check succeeds

	// ok food/water
	case 3:
		// reduce radiation if there's some
		// drop a bit of damage
		// increase mood approaching 75

	// great food/water
	case 4:
		// reduce radiation a larger amount
		// reduce damage a bit more
		// increase mood approaching 90
	}
	p.applyDamage()
}

func (p *Person) applyDamage() {
	if p.DamageAmount+p.RadiateAmount > p.TotalHP {
		p.CurrentHP = 0
		p.IsDead = true
		p.ReviveCost = 50 + rand.Intn(100+2*p.Level) + 10*p.Level
		return
	}
	p.CurrentHP = p.TotalHP - p.DamageAmount - p.RadiateAmount
}

func (p Person) SimpleDPS() int {
	return 0
}
