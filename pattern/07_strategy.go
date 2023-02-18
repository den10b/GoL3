package pattern

import "fmt"

// Интерфейс стратегии
type strateg interface {
	oppTurn(g *game)
}

// Конкретная стратегия
type easy struct {
}

func (strat *easy) oppTurn(g *game) {
	g.score += 1
	fmt.Println("Легкий бот прибавил себе 1")
}

// Конкретная стратегия
type med struct {
}

func (strat *med) oppTurn(g *game) {
	g.score += 3
	fmt.Println("Средний бот прибавил себе 3")
}

// Конкретная стратегия
type hard struct {
}

func (strat *hard) oppTurn(g *game) {
	g.score += 5
	fmt.Println("Тяжелый бот прибавил себе 5")
}

// Контекст
type game struct {
	score        int
	opponentMode strateg
}

func initGame(strat strateg) *game {
	return &game{
		score:        0,
		opponentMode: strat,
	}
}

func (g *game) add(value int) {
	g.score += value
}

func (g *game) oppTurn() {
	g.opponentMode.oppTurn(g)
}

func main() {
	ezGame := initGame(&easy{})
	ezGame.oppTurn()
	medGame := initGame(&med{})
	medGame.oppTurn()
	hdGame := initGame(&hard{})
	hdGame.oppTurn()

}
