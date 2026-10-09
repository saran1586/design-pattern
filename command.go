package main

import "fmt"

type Command interface{
Execute()
}

type Light struct{}

func (l Light) On(){
 fmt.Println("Light is on")
}

type LightOnCommand struct {
light Light
}

func (c LightOnCommand) Execute(){
c.light.On()
}

func main(){
light := Light{}
command := LightOnCommand{light:light}
command.Execute()
}
