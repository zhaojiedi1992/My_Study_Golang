map
==========================

map 是一种特殊的数据结构，一种元素对（pair）的无序集合，pair 对应一个 key（索引）和一个 value（值），所以这个结构也称为关联数组或字典。


.. code-block:: go

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

Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，
不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。