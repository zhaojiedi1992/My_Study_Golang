package main 

import "fmt"
import "container/list"

func main()  {
	l := list.New()
	l.PushBack("fist")
	l.PushFront(67)
	element := l.PushBack("fist")
	l.InsertAfter("high", element)
	l.InsertBefore("noon", element)
	l.Remove(element)

	for i := l.Front() ; i !=nil ; i=i.Next(){
		fmt.Println(i.Value)
	}

}