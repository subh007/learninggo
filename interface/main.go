package main

import (
	"fmt"
	"math/rand"
)

type OptionFn func(*Options)

type Options struct {
	stamina, power int
}

func withStrengh(o *Options) {
	o.power = rand.Intn(5) + 50
}

type DefaultOptionPlayer interface {
	handleDefault(fns ...OptionFn)
}

type DefaultHandler struct {
	Options
}

func (d *DefaultHandler) handleDefault(fns ...OptionFn) {
	d.Options = Options{
		stamina: rand.Intn(10),
		power:   rand.Intn(10),
	}

	for _, fn := range fns {
		fn(&d.Options)
	}
}

// func

type Player interface {
	kickBall()
}

type CR7 struct {
	DefaultHandler
	SUI int
}

func NewCR7(fns ...OptionFn) *CR7 {
	c := &CR7{
		SUI: 10,
	}
	c.handleDefault(fns...)
	return c
}

func (c CR7) kickBall() {
	fmt.Printf("CR7 hitting ball with : %d \n", c.stamina*c.power+c.SUI)
}

type FootballPlayer struct {
	DefaultHandler
}

func NewFootballPlayer(fns ...OptionFn) *FootballPlayer {
	f := &FootballPlayer{}
	f.handleDefault(fns...)
	return f
}

func (f FootballPlayer) kickBall() {
	fmt.Printf("I'm hitting ball with : %d \n", f.stamina*f.power)
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-1; i++ {
		team[i] = NewFootballPlayer(withStrengh)
	}

	team[len(team)-1] = NewCR7(withStrengh)

	for _, player := range team {
		player.kickBall()
	}
}
