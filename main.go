package main

import (
	"math/rand"
	"fmt"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const sizeBytesNameBook = 10
const sizeBytesContentBook = 10

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

type Book struct {
	Id int64;
	Name string;
	Content string;
}

func NewBook () Book {
	var b Book
    b.Name = RandStringBytes(sizeBytesNameBook)
	b.Content = RandStringBytes(sizeBytesContentBook)
	return b
}
/*
type Shelf struct {
	id int64;
	Books []book;//cada shelf vai ter 32 libros
}

type Wall struct {
	id int64;
	shelfs []shelf;//cada parede vai conter 5 prateleiras
}

type Room struct {
	id int64;
	position int64;
	Walls []wall;//cada Room so vai conter 4 paredes com livros 
}
*/



func main() {
	//inicio criar 

	for i := 0; i < 31; i++ {
		fmt.Print(NewBook())
	}
	
	//fim criar livros

}
