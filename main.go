package main

import (
	"math/rand"
	"fmt"
	"encoding/json"

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

type Shelf struct {
	id int64;
	books []Book;//cada shelf vai ter 32 libros
}

func NewShelf() Shelf {
	var s Shelf;
	for i := 0; i < 31; i++ {
		s.books = append(s.books, NewBook())
	}
	return s
} 

type Wall struct {
	id int64;
	shelfs []Shelf;//cada parede vai conter 5 prateleiras
}

func NewWall() Wall {
	var w Wall;
	for i := 0; i < 5; i++ {
		w.shelfs = append(w.shelfs, NewShelf())
	}
	return w
}

type Room struct {
	id int64;
	position int64;
	walls []Wall;//cada Room so vai conter 4 paredes com livros 
}

func NewRoom() Room {
	var r Room;
	for i := 0; i < 4; i++ {
		r.walls = append(r.walls, NewWall())
	}
	return r
}

func Babel() {
	var room Room = NewRoom()

	room2 := &room
	js, err := json.Marshal(room2)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return;
	}

	fmt.Printf("%s", js)
}

func main() {
	Babel()
}
