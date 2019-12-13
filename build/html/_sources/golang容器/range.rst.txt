range
====================

range需要配合for来完成多个相同类型的元素的迭代问题。

.. code-block:: go

    func main()  {
        slice := []int {10,20,30,40}
        for index, value := range slice {
            fmt.Printf("Index: %d Value: %d\n", index, value)
        }

        for index :=2 ; index < len(slice) ; index ++ {
            fmt.Printf("Value %d", slice[index])
        }
    }

