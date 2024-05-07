package creational

import (
	"errors"
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	ak47, _ := GetGunFactory("ak47")
	musket, _ := GetGunFactory("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Power: %d\n", g.GetPower())
}

type IGun interface {
	SetName(name string)
	GetName() string
	SetPower(power int)
	GetPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) SetName(name string) {
	g.name = name
}

func (g *Gun) GetName() string {
	return g.name
}

func (g *Gun) SetPower(power int) {
	g.power = power
}

func (g *Gun) GetPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func NewAk47(name string, power int) IGun {
	return &Ak47{
		Gun{
			name:  name,
			power: power,
		},
	}
}

type Musket struct {
	Gun
}

func NewMusket(name string, power int) IGun {
	return &Musket{
		Gun{
			name:  name,
			power: power,
		},
	}
}

func GetGunFactory(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47("ak47", 4), nil
	}

	if gunType == "musket" {
		return NewMusket("musket", 1), nil
	}

	return nil, errors.New("invalid gun type")
}
