package main 

import "fmt"


func main(){
	
	// 定义2个变量
	var a int = 1
	var b int = 2
	// 使用 &变量名  获取对应变量的地址 
	fmt.Printf("%p %p", &a, &b)


	var ptr_a = &a 
	//var ptr_b = &b 

	// 打印ptr的类型
    fmt.Printf("ptr type: %T\n", ptr_a)
    // 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr_a)

	aa := *ptr_a
	
    // 取值后的类型
    fmt.Printf("value type: %T\n", aa)
    // 指针取值后就是指向变量的值
    fmt.Printf("value: %d\n", aa)

}