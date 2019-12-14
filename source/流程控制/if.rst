if
==================================

关键字 if 是用于测试某个条件（布尔型或逻辑型）的语句，如果该条件成立，则会执行 if 后由大括号{}括起来的代码块，否则就忽略该代码块继续执行后续的代码。

.. code-block:: go 

    // 第一种情况
    if condition {
        // do something
    }

    // 第二种情况
    if condition {
        // do something
    } else {
        // do something
    }

    if condition1 {
    // do something
    } else if condition2 {
        // do something else
    }else {
        // catch-all or default
    }


    // 特殊用法
    if err := Connect(); err != nil {
        fmt.Println(err)
        return
    }

.. note:: 关键字 if 和 else 之后的左大括号{必须和关键字在同一行，

