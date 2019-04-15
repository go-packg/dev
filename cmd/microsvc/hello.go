package main

import "fmt"

type attacker struct {
		attackpower int
		dmgbonus int
}
type sword struct {
		attacker
		twohanded bool
}
type gun struct {
		attacker
		bulletsremaining int
}

func (s sword) Wield() bool {
		fmt.Println("You've wielded a sword!")
		return true
}
func (g gun) Wield() bool {
		fmt.Println("You've wielded a gun!")
		return true
}

type weapon interface {
		Wield() bool
}

func wielder(w weapon) bool {
	fmt.Println("Wielding...")
	return w.Wield()
}

type chair struct {
	legcount int
	leather bool
}	

func (c chair) Wield() bool {
		fmt.Println("You've wielded a chair!! You having a bad day?")
		return true
}

func main() {
	sword1 := sword{attacker: attacker{attackpower: 1, dmgbonus: 5}, twohanded: true}
	gun1 := gun{attacker: attacker{attackpower: 10, dmgbonus: 20}, bulletsremaining: 11}
	fmt.Printf("Weapons: sword: %v, gun: %v\n", sword1, gun1)
	
	sword1.Wield()
	gun1.Wield()
	
	wielder(sword1)
	wielder(gun1)
	
	chair1 := chair{legcount: 3, leather: true}
	wielder(chair1)
}
