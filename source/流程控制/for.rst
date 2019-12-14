for
==================================

循环语句只支持 for 关键字，而不支持 while 和 do-while 结构，

.. code-block:: go

    // 第一种使用
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }

    // 第二种使用
    sum := 0
    for {
        sum++
        if sum > 100 {
            break
        }
    }

    // 第三种

    step := 2
    for ; step > 0; step-- {
        fmt.Println(step)
    }

    // 第四种

    for i <= 10 {
        i++
    }