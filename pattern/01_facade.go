package pattern

import "fmt"

type shopFacade struct {
	item Item
	bank Bank
}

func newShopFacade(cardNum int) *shopFacade {
	fmt.Println("Создаем аккаунт")
	shopFacade := &shopFacade{
		bank: findUser(cardNum),
	}
	fmt.Println("Аккаунт создан!")
	return shopFacade
}
func (f *shopFacade) tryToBuy(itemId int) {
	f.item = findItem(itemId)
	enoughMoney := f.bank.checkMoney(f.item.price)
	if enoughMoney {
		f.bank.spendMoney(f.item.price)
		f.item.buy()
		println("Успешно куплено")
	} else {
		println("Денег не хватает")
	}

}

type Item struct {
	price    int
	itemID   int
	quantity int
}

func findItem(itemID int) Item {
	item := new(Item)
	switch itemID % 10 {
	case 1:
		item.itemID = itemID
		item.price = 10
		item.quantity = 5
	case 2:
		item.itemID = itemID
		item.price = 15
		item.quantity = 0
	case 3:
		item.itemID = itemID
		item.price = 100
		item.quantity = 1
	default:
		item.itemID = itemID
		item.price = 100000
		item.quantity = 0
	}
	return *item
}
func (i *Item) buy() {
	i.quantity = -1
}

type Bank struct {
	bankID     int
	userCardID int
	userWallet int
}

func findUser(cardNum int) Bank {
	bank := new(Bank)
	switch cardNum % 10 {
	case 1:
		bank.bankID = 101
		bank.userCardID = cardNum
		bank.userWallet = 50

	case 2:
		bank.bankID = 202
		bank.userCardID = cardNum
		bank.userWallet = 20
	default:
		bank.bankID = 123
		bank.userCardID = cardNum
		bank.userWallet = 0
	}
	return *bank
}
func (b *Bank) checkMoney(val int) bool {
	return val <= b.userWallet
}
func (b *Bank) spendMoney(val int) {
	b.userWallet -= val
}

func main() {
	adidas := newShopFacade(121)
	adidas.tryToBuy(111)
	nike := newShopFacade(122)
	nike.tryToBuy(113)

}
