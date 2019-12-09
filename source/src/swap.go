package main


import "fmt"

func swap(a *int , b *int ){
	var t = *a 
	*a = *b 
	*b = t 
}

func main() {
	var a = 20
	var b = 10 
	fmt.Println(a,b)
	swap(&a,&b)
	fmt.Println(a,b)
}