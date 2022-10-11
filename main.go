package main

import (
	"fmt"
	"math/rand"
	"time"
)

const BreakPoint = 11

type GameRules struct {
	hits       int
	lastPlayer string
}

func main() {
	korekApi := make(chan *GameRules)
	done := make(chan *GameRules)
	defer finish(done)

	players := []string{"player 1", "player 2", "player 3", "player 4"}
	for _, p := range players {
		go play(p, korekApi, done)
	}
	korekApi <- new(GameRules)
}

func play(name string, korekApi chan *GameRules, done chan *GameRules) {
	for {
		select {
		case k := <-korekApi:
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(100-1) + 1
			k.lastPlayer = name
			k.hits++
			time.Sleep(800 * time.Millisecond)
			fmt.Println("korek ada di", k.lastPlayer, "pada hit ke", k.hits, "dan mempunyai nilai", randomNumber)

			if randomNumber%BreakPoint == 0 {
				done <- k
				return
			}
			korekApi <- k
		}
	}
}

func finish(done chan *GameRules) {
	for {
		select {
		case d := <-done:
			fmt.Println(d.lastPlayer, "kalah pada hit ke ====>", d.hits)
			return
		}
	}
}
