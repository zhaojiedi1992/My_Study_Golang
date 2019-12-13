
package main
import "fmt"


func main()  {
	var demo01 map[string]int
	var demo02 map[string]int

	demo01 = map[string]int { "one":1,"two":2}

	demo02 = make(map[string]int, 10)
	
	demo02["one"] =1
	demo02["two"] =2

	for k,v :=range demo01 {
		fmt.Println(k, v)
	}

	for k,v :=range demo02 {
		fmt.Println(k, v)
	}

	fmt.Println("======")
	delete(demo02, "one")

	for k,v :=range demo02 {
		fmt.Println(k, v)
	}
}

