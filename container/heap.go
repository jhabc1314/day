package container

//包堆为实现了heap.Interface的任何类型提供堆操作。 堆是一棵树，其属性为每个节点是其子树中的最小值节点。

//树中的最小元素是根，索引为0。

//堆是实现优先级队列的常用方法。 要构建优先级队列，请以（负）优先级实现Heap接口，作为Less方法的顺序，因此Push添加项，而Pop删除队列中优先级最高的项。 示例包括这样的实现； 文件example_pq_test.go包含完整的源代码。

//在索引i的元素更改其值之后，Fix会重新建立堆顺序。 更改索引i处元素的值，然后调用Fix等效于，但比调用Remove（h，i）紧随其后的是推入新值，但其开销要小。 复杂度为O（log n），其中n = h.Len（）。
//func Fix(h Interface, i int)

