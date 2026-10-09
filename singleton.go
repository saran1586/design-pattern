package main

import (
	"fmt"
	"sync"
	)

type DataBase struct{

 Name string
}
var(
 	db *DataBase
	once sync.Once
)

func GetDatabase() *DataBase{

once.Do(func() {
db = &DataBase{
Name: "saran",
}
})
return db
}

func main() {
Name1 := GetDatabase()
Name2 := GetDatabase()

fmt.Println(Name1.Name)
fmt.Println(Name2.Name)
fmt.Println(Name1 == Name2)
}

