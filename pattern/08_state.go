package pattern

import (
	"fmt"
	"log"
)

type State interface {
	takeMoney(int) error
	putMoney(int) error
	lookMoney() error
	useKey() error
}

type safe struct {
	opened       State
	closed       State
	currentState State

	money int
}

func newSafe(money int) *safe {
	v := &safe{
		money: money,
	}
	openedState := &openedState{
		safe: v,
	}
	closedState := &closedState{
		safe: v,
	}
	v.currentState = openedState
	v.closed = closedState
	v.opened = openedState

	return v
}

func (v *safe) takeMoney(money int) error {
	return v.currentState.takeMoney(money)
}

func (v *safe) putMoney(money int) error {
	return v.currentState.putMoney(money)
}

func (v *safe) lookMoney() error {
	return v.currentState.lookMoney()
}
func (v *safe) useKey() error {
	return v.currentState.useKey()
}

type openedState struct {
	safe *safe
}

func (v *openedState) takeMoney(money int) error {
	if v.safe.money < money {
		v.safe.money -= money
		return nil
	}
	return fmt.Errorf("денег нет но вы держиитесь")

}

func (v *openedState) putMoney(money int) error {
	v.safe.money += money
	return nil
}

func (v *openedState) lookMoney() error {
	fmt.Printf("В сейфе лежит %v рублей\n", v.safe.money)
	return nil
}
func (v *openedState) useKey() error {
	v.safe.currentState = v.safe.closed
	fmt.Printf("Закрываем сейф")
	return nil
}

type closedState struct {
	safe *safe
}

func (v *closedState) takeMoney(money int) error {
	return fmt.Errorf("сейф закрыт")

}

func (v *closedState) putMoney(money int) error {
	return fmt.Errorf("сейф закрыт")
}

func (v *closedState) lookMoney() error {
	return fmt.Errorf("сейф закрыт")
}
func (v *closedState) useKey() error {
	v.safe.currentState = v.safe.opened
	fmt.Printf("Открываем сейф")
	return nil
}

func main() {
	mySafe := newSafe(10)

	err := mySafe.takeMoney(5)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = mySafe.putMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = mySafe.useKey()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = mySafe.lookMoney()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
