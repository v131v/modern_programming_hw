package main

import (
	"time"

	termbox "github.com/nsf/termbox-go"
	game "github.com/v131v/modern_programming_hw/game/src/game"
)

func controllerLoop(g *game.Game) {
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				g.Status = game.END
				return
			case termbox.KeyArrowLeft:
				g.PlayerX--
			case termbox.KeyArrowRight:
				g.PlayerX++
			}
		}
	}
}

func drawLoop(g *game.Game) {
	for {
		if g.Status == game.EXIT {
			return
		}

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		switch g.Status {
		case game.IN_GAME:
			g.Draw()
		case game.END:
			g.DrawEnd()
		}

		termbox.Flush()
	}
}

func logicLoop(g *game.Game) {
	for {
		if g.Status == game.END || g.Status == game.EXIT {
			return
		}

		g.Score++
		g.UpdateRoad()

		if g.PlayerX < g.Road[0].X || g.PlayerX > g.Road[0].X+g.Road[0].W {
			g.Status = game.END
		}

		time.Sleep(time.Duration(g.Speed) * time.Millisecond)

		if g.Speed > 50 {
			g.Speed--
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	w, h := termbox.Size()

	g := game.CreateGame(w, h)

	go drawLoop(g)
	go logicLoop(g)
	controllerLoop(g)
}
