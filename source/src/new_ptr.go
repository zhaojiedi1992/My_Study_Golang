package main 

import  "fmt"

func main() {
	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}
