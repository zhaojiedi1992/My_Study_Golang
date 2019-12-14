case
========================

witch 要比C语言的更加通用，表达式不需要为常量，甚至不需要为整数，case 按照从上到下的顺序进行求值，
直到找到匹配的项，如果 switch 没有表达式，则对 true 进行匹配，

.. code-block:: go

    var a = "hello"
    switch a {
    case "hello":
        fmt.Println(1)
    case "world":
        fmt.Println(2)
    default:
        fmt.Println(0)
    }


    // 一个分支多支的

    var a = "mum"
    switch a {
    case "mum", "daddy":
        fmt.Println("family")
    }

    // 分支表达式

    var r int = 11
    switch {
    case r > 10 && r < 20:
        fmt.Println(r)
    }
    //case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能，

    var s = "hello"
    switch {
    case s == "hello":
        fmt.Println("hello")
        fallthrough
    case s != "world":
        fmt.Println("world")
    }
