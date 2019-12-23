io操作
===================================

几乎所有的数据结构都围绕接口展开，接口是Go语言中所有数据结构的核心。在实际开发过程中，无论是实现 web 应用程序，还是控制台输入输出，
又或者是网络操作，都不可避免的会遇到 I/O 操作。

.. code-block:: go 

    package main

    import (
        "bufio"
        "bytes"
        "fmt"
    )

    func main() {
        wr := bytes.NewBuffer(nil)
        w := bufio.NewWriter(wr)
        s := "C语言中文网"
        n, err := w.WriteString(s)
        w.Flush()
        fmt.Println(string(wr.Bytes()), n, err)
    }
