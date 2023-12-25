package game

import (
	"fmt"
	"math/rand"

	termbox "github.com/nsf/termbox-go"
)

type Row struct {
	X int
	W int
}

type Status int

const (
	IN_GAME Status = iota
	END
	EXIT
)

type Game struct {
	W          int
	H          int
	PlayerX    int
	PlayerY    int
	PlayerChar rune
	BorderChar rune
	Road       []Row

	Speed  int
	Status Status
	Score  int
}

func CreateGame(w, h int) *Game {
	g := &Game{
		W:          w,
		H:          h,
		PlayerX:    0,
		PlayerY:    h - 1,
		PlayerChar: '^',
		BorderChar: '*',
		Road:       make([]Row, h+1),
		Speed:      400,
		Status:     IN_GAME,
	}

	roadWidth := 10

	for i := range g.Road {
		if i == 0 {
			g.Road[i] = Row{g.PlayerX - roadWidth/2, roadWidth}
			continue
		}
		g.Road[i] = g.Road[i-1]
		g.Road[i].X = g.Road[i-1].X + rand.Intn(3) - 1
	}

	return g
}

func (g *Game) Draw() {

	for i := 0; i < g.H; i++ {
		left, right := g.Road[i].X-1-g.PlayerX+g.W/2, g.Road[i].X+g.Road[i].W+1-g.PlayerX+g.W/2
		y := g.H - i - 1

		termbox.SetCell(left, y, g.BorderChar, termbox.ColorBlack, termbox.ColorYellow)
		termbox.SetCell(right, y, g.BorderChar, termbox.ColorBlack, termbox.ColorYellow)

		for j := left + 1; j < right; j++ {
			termbox.SetCell(j, y, ' ', termbox.ColorBlack, termbox.ColorYellow)
		}
	}

	termbox.SetCell(g.W/2, g.H-1, g.PlayerChar, termbox.ColorGreen, termbox.ColorBlack)
}

func (g *Game) DrawEnd() {
	messages := []string{
		"End game",
		fmt.Sprintf("Your score: %v", g.Score),
		"Press ESC to exit",
	}

	for i, s := range messages {
		x := g.W/2 - len(s)/2
		y := g.H/2 - (len(messages)/2 - i)

		for _, c := range s {
			termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}
	}
}

func (g *Game) UpdateRoad() {
	g.Road = g.Road[1:]
	last := g.Road[len(g.Road)-1]
	last.X += rand.Intn(3) - 1

	if last.X == g.Road[len(g.Road)-1].X {
		p := rand.Intn(5)
		if (p == 0 || p == 1) && last.W > 4 {
			last.W--
			last.X++
		}
		if p == 2 && last.W < 15 {
			last.W++
			last.X--
		}
	}

	g.Road = append(g.Road, last)
}
