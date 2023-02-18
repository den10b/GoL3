package pattern

import "fmt"

type event interface {
	getType() string
	accept(Visitor)
}
type artMuseum struct {
	nExps int
}

func (s *artMuseum) accept(v Visitor) {
	v.visitArt(s)
}

func (s *artMuseum) getType() string {
	return "ArtMuseum"
}

type scienceExpo struct {
	theme     string
	nOfPeople int
}

func (c *scienceExpo) accept(v Visitor) {
	v.visitExpo(c)
}

func (c *scienceExpo) getType() string {
	return "ScienceExpo"
}

type concert struct {
	tracksNum int
	headliner string
}

func (t *concert) accept(v Visitor) {
	v.visitConcert(t)
}

func (t *concert) getType() string {
	return "Concert"
}

type Visitor interface {
	visitArt(*artMuseum)
	visitExpo(*scienceExpo)
	visitConcert(*concert)
}
type critique struct {
	article string
}

func (a *critique) visitArt(s *artMuseum) {
	fmt.Printf("Статья про музей %v\n", s)
}

func (a *critique) visitExpo(s *scienceExpo) {
	fmt.Printf("Статья про выставку %v\n", s)
}
func (a *critique) visitConcert(s *concert) {
	fmt.Printf("Статья про концерт %v\n", s)
}

type photographer struct {
	nPhotos uint
}

func (a *photographer) visitArt(c *artMuseum) {
	fmt.Printf("Фото с концерта %v\n", c)
}

func (a *photographer) visitExpo(c *scienceExpo) {
	fmt.Printf("Фото с выставки %v\n", c)
}
func (a *photographer) visitConcert(c *concert) {
	fmt.Printf("фото с концерта %v\n", c)
}

func main() {
	newExpo := &scienceExpo{}
	newMuseum := &artMuseum{}
	newConcert := &concert{}
	weekendEvents := []event{newExpo, newMuseum, newConcert}
	mrJack := &critique{}
	for _, curr := range weekendEvents {
		curr.accept(mrJack)
	}

	fmt.Println()
	mrJones := &photographer{}
	for _, curr := range weekendEvents {
		curr.accept(mrJones)
	}
}
