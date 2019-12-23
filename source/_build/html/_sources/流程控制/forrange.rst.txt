for range
=======================

for range 可以用于遍历数组、切片、字符串、map及通道， 类似其他的语言中的foreach语句。 

.. code-block:: go 

    for key, value := range []int{1, 2, 3, 4} {
        fmt.Printf("key:%d  value:%d\n", key, value)
    }

    m := map[string]int{
        "hello": 100,
        "world": 200,
    }
    for key, value := range m {
        fmt.Println(key, value)
    }

    c := make(chan int)
    go func() {
        c <- 1
        c <- 2
        c <- 3
        close(c)
    }()
    for v := range c {
        fmt.Println(v)
    }