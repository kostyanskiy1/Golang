package pattern

import "fmt"

// Command интерфейс.
type Command interface {
	Execute() string
}

// ToggleOnCommand реализует Command интерфейс.
type ToggleOnCommand struct {
	receiver *Receiver
}

// выполнение команды.
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// ToggleOffCommand реализует Command интерфейс.
type ToggleOffCommand struct {
	receiver *Receiver
}

// выполнение команды.
func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Приемщик .
type Receiver struct {
}

// ToggleOn .
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

// ToggleOff .
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Invoker .
type Invoker struct {
	commands []Command
}

// добавляет команду.
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// удаляет команду.
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// выполняет команды
func (i *Invoker) Execute() string {
	var result string
	for _, command := range i.commands {
		result += command.Execute() + "\n"
	}
	return result
}

func CommandFunc() {
	invoker := &Invoker{} //инициализируем
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})
	invoker.UnStoreCommand()
	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})
	result := invoker.Execute()
	fmt.Println(result)
}
