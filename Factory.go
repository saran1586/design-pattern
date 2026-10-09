package main

import "fmt"

type Animal interface{
 speak()
}


type Dog struct {
}
type Cat struct{
}

func (Dog) speak(){
fmt.Println("roar")
}

func (Cat) speak(){
fmt.Println("meow")
}

func Factory(name string)Animal{
switch name{
case "dog" :
    return Dog{}
case "cat" :
    return Cat{}
default:
    return nil
}
}

func main(){

animal1:= Factory("dog")
animal1.speak()


animal2:= Factory("cat")
animal2.speak()

}
