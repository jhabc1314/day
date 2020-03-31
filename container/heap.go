package container

import (
	"container/heap"
	"fmt"
	"sort"
)

//包堆为实现了heap.Interface的任何类型提供堆操作。 堆是一棵树，其属性为每个节点是其子树中的最小值节点。

//树中的最小元素是根，索引为0。

//堆是实现优先级队列的常用方法。 要构建优先级队列，请以（负）优先级实现Heap接口，作为Less方法的顺序，因此Push添加项，而Pop删除队列中优先级最高的项。 示例包括这样的实现； 文件example_pq_test.go包含完整的源代码。

//在索引i的元素更改其值之后，Fix会重新建立堆顺序。 更改索引i处元素的值，然后调用Fix等效于，但比调用Remove（h，i）紧随其后的是推入新值，但其开销要小。 复杂度为O（log n），其中n = h.Len（）。
//func Fix(h Interface, i int)

//Init建立此程序包中其他例程所需的堆不变式。 关于堆不变式，Init是幂等的，只要使堆不变式无效，就可以调用它。 复杂度为O（n），其中n = h.Len（）。
//func Init(h Interface)

//Pop从堆中删除并返回最小元素（根据Less）。 复杂度为O（log n），其中n = h.Len（）。 Pop等效于Remove（h，0）。
//func Pop(h Interface) interface{}

//推将元素x推入堆。 复杂度为O（log n），其中n = h.Len（）。
//func Push(h Interface, x interface{})

//Remove从堆中删除并返回索引i处的元素。 复杂度为O（log n），其中n = h.Len（）。
//func Remove(h Interface, i int) interface{}

//接口类型使用此包中的例程描述对类型的要求。 任何实现它的类型都可以用作具有以下不变量的最小堆（在调用Init或数据为空或已排序之后建立）：
//!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()

//请注意，此接口中的Push和Pop是供程序包堆的实现调用的。 要从堆中添加和删除内容，请使用heap.Push和heap.Pop。

type Interface interface {
    sort.Interface
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1.
}

func HeapFunc(h *IntHeap) {
	//首先h要实现heap.Interface 接口

	heap.Init(h)
	heap.Push(h, 10)
	for h.Len() >0 {
		p := heap.Pop(h)
		fmt.Println(p)
	}
	
}

type IntHeap []int

func (h IntHeap)Less(i,j int) bool {
	return h[i] < h[j]
}

func (h IntHeap)Len() int {
	return len(h)
}

func (h IntHeap)Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *IntHeap)Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap)Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


