package pattern

/*
Можно использовать если:
вы хотите параметризовать объекты выполняемым действием
вы хотите ставить операции в очередь,
	выполнять их по расписанию или передавать по сети
вам нужна операция отмены

+:
Убирает прямую зависимость между объектами, вызывающими операции,
	и объектами, которые их непосредственно выполняют
Позволяет реализовать простую отмену и повтор операций
Позволяет реализовать отложенный запуск операций
Позволяет собирать сложные команды из простых
Реализует принцип открытости/закрытости

-:
 Усложняет код программы из-за введения множества дополнительных классов
*/
import "fmt"

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}
type onCommand struct {
	d device
}

func (o *onCommand) execute() {
	o.d.on()
}

type offCommand struct {
	d device
}

func (o *offCommand) execute() {
	o.d.off()
}

type muteCommand struct {
	d device
}

func (o *muteCommand) execute() {
	o.d.mute()
}

type unmuteCommand struct {
	d device
}

func (o *unmuteCommand) execute() {
	o.d.unmute()
}

type device interface {
	on()
	off()
	mute()
	unmute()
}
type TV struct {
	isPowered bool
	isMuted   bool
}

func (tv *TV) on() {
	tv.isPowered = true
	fmt.Printf("Телевизор включён!")
}
func (tv *TV) off() {
	tv.isPowered = false
	fmt.Printf("Телевизор выключен!")
}
func (tv *TV) mute() {
	if tv.isPowered {
		tv.isMuted = true
		fmt.Printf("Звук выключен!")
	} else {
		fmt.Printf("Телевизор выключен!")
	}

}
func (tv *TV) unmute() {
	if tv.isPowered {
		tv.isMuted = false
		fmt.Printf("Звук включен!")
	} else {
		fmt.Printf("Телевизор выключен!")
	}
}
func main() {
	myHomeTv := &TV{isMuted: false, isPowered: false}
	onCommand := &onCommand{myHomeTv}
	offCommand := &offCommand{myHomeTv}
	muteCommand := &muteCommand{myHomeTv}
	unmuteCommand := &unmuteCommand{myHomeTv}
	onCommand.execute()
	muteCommand.execute()
	unmuteCommand.execute()
	offCommand.execute()

	onButton := &button{onCommand}
	offButton := &button{offCommand}
	muteButton := &button{muteCommand}
	unmuteButton := &button{unmuteCommand}
	onButton.press()
	offButton.press()
	muteButton.press()
	unmuteButton.press()
}
