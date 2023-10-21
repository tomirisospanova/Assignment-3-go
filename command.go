package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type Receiver struct {
	State string
}

func (r *Receiver) Action(action string) {
	r.State = action
}


type ConcreteCommand struct {
	Receiver *Receiver
	Action   string
}

func (cc *ConcreteCommand) Execute() {
	cc.Receiver.Action(cc.Action)
}

func (cc *ConcreteCommand) Undo() {
	
}

type CommandHistory struct {
	Commands []Command
}

func (ch *CommandHistory) AddCommand(command Command) {
	ch.Commands = append(ch.Commands, command)
}

func (ch *CommandHistory) ExecuteAllCommands() {
	for _, command := range ch.Commands {
		command.Execute()
	}
}

func main() {
	receiver := &Receiver{}
	commandHistory := &CommandHistory{}

	command1 := &ConcreteCommand{Receiver: receiver, Action: "Action 1"}
	command2 := &ConcreteCommand{Receiver: receiver, Action: "Action 2"}

	commandHistory.AddCommand(command1)
	commandHistory.AddCommand(command2)

	commandHistory.ExecuteAllCommands()

	fmt.Printf("Receiver state: %s\n", receiver.State)
	
	commandHistory.Commands = nil
}
