package pingpong

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MaxPoint = 11

type Player struct {
	Name  string
	Point int
}

func newPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Point: 0,
	}
}

type Game struct {
	Player1, Player2 *Player
	wg               *sync.WaitGroup
	ch               chan string
	r                *rand.Rand
}

func NewGame(name1, name2 string) *Game {
	return &Game{
		Player1: newPlayer(name1),
		Player2: newPlayer(name2),
		wg:      new(sync.WaitGroup),
		ch:      make(chan string),
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Game) Start() {
	go g.Player1.play(g)
	go g.Player2.play(g)

	g.ch <- "begin"

	g.wg.Wait()

	fmt.Printf("%s: %v - %s: %v\n", g.Player1.Name, g.Player1.Point, g.Player2.Name, g.Player2.Point)
}

func (p *Player) play(g *Game) {
	g.wg.Add(1)
	defer g.wg.Done()

	for val := range g.ch {
		switch val {
		case "begin", "stop":
			g.ch <- nextKick(val)
		case "ping", "pong":
			fmt.Println(val)

			if g.checkGoal() {
				p.Point++

				if p.Point == MaxPoint {
					close(g.ch)
				} else {
					g.ch <- "stop"
				}
			} else {
				g.ch <- nextKick(val)
			}
		}
	}
}

func (g *Game) checkGoal() bool {
	return g.r.Intn(5) == 0
}

func nextKick(val string) string {
	if val == "ping" {
		return "pong"
	}

	return "ping"
}
