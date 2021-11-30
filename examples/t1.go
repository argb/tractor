package main


import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i:=0; i<10; i++ {
		l.PushBack(i)
	}
	for elem := l.Front(); elem!=nil; elem = elem.Next() {
		fmt.Println(elem)
	}
}
