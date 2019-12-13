package main
import (
	"fmt"
	"sync"
)

func main01()  {
	
	m := make(map[int]int)

	go func(){
		for {
			m[1] = 1
		}
	}()
	go func(){
		for {
			_ = m[1]
		}
	}()

	for{

	}
}

func main(){
	
	var m sync.Map

	m.Store("one",1)
	m.Store("two",2)

	fmt.Println(m.Load("one"))

	m.Delete("one")

	m.Range(func(k,v interface{}) bool{
		fmt.Println(k,v)
		return true 
	})

}