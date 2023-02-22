package pattern

import "fmt"

type Department interface {
	execute(*document)
	setNext(Department)
}

type accountant struct {
	next Department
}

func (r *accountant) execute(p *document) {
	if p.signedAcc {
		fmt.Printf("Документ %s уже подписан бухгалтером\n", p.name)
		r.next.execute(p)
		return
	}
	fmt.Printf("Документ %s подписывается бухгалтером\n\n", p.name)
	p.signedAcc = true
	r.next.execute(p)
}

func (r *accountant) setNext(next Department) {
	r.next = next
}

type secretary struct {
	next Department
}

func (r *secretary) execute(p *document) {
	if p.signedSecr {
		fmt.Printf("Документ %s уже подписан секретарём\n", p.name)
		r.next.execute(p)
		return
	}
	fmt.Printf("Документ %s подписывается секретарём\n\n", p.name)
	p.signedSecr = true
	r.next.execute(p)
}

func (r *secretary) setNext(next Department) {
	r.next = next
}

type director struct {
	next Department
}

func (r *director) execute(p *document) {
	if p.signedSecr {
		fmt.Printf("Документ %s уже подписан директором\n", p.name)
		r.next.execute(p)
		return
	}
	fmt.Printf("Документ %s подписывается директором\n\n", p.name)
	p.signedDir = true
	r.next.execute(p)
}

func (r *director) setNext(next Department) {
	r.next = next
}

type document struct {
	name       string
	signedAcc  bool
	signedSecr bool
	signedDir  bool
}

func main() {

	mrJones := &director{}
	mrJonesSecretary := &secretary{mrJones}
	firmAccountant := &accountant{mrJonesSecretary}

	newDocument := &document{name: "ДОГОВОР"}

	firmAccountant.execute(newDocument)
}
