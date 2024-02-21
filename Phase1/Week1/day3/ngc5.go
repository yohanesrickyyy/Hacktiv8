package main

import (
	"fmt"
	"math/rand"
	"time"
)

// variabel person
type Person struct {
	Name string
	Age  int
	Job  string
}

// Logic soal 1
func (p *Person) GetInfo() {
	fmt.Printf("Name: %s\nAge: %d\nJob: %s\n", p.Name, p.Age, p.Job)
}

func (p *Person) AddYear() {
	p.Age++
	if p.Age >= 50 {
		p.Job = "Retired"
	}
}

// Logic soal 2
type Weapon struct {
	Attack int
}

// Hero struct represents a hero with various properties including Weapon
type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

func (h *Hero) CountDamage() int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	critChance := rand.Intn(2)
	if critChance == 0 {
		return h.BaseAttack + h.Weapon.Attack + h.CriticalDamage
	}
	return h.BaseAttack + h.Weapon.Attack
}

func main() {
	p := Person{Name: "Bambang", Age: 45, Job: "Gambler"}

	fmt.Println("Initial Information:")
	p.GetInfo()

	for i := 0; i < 5; i++ {
		p.AddYear()
	}

	fmt.Println("\nInformation after 5 years:")
	p.GetInfo()

	// soal 2
	hero := Hero{
		Name:           "Axe",
		BaseAttack:     30,
		Defence:        20,
		CriticalDamage: 50,
		HealthPoint:    100,
		Weapon:         Weapon{Attack: 10},
	}

	totalDamage := hero.CountDamage()

	fmt.Printf("Total damage dealt by %s: %d\n", hero.Name, totalDamage)

}
