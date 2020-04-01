package container

//环 在循环列表上实现操作

//环是循环列表或环的元素。 戒指没有起点或终点； 指向任何环元素的指针都用作整个环的引用。 空环表示为nil环指针。 环的零值是一个零元素的环。

//New creates a ring of n elements.
//func New(n int) *Ring

//请按向前的顺序在环的每个元素上调用函数f。如果f改变* r，Do的行为是不确定的。
//func (r *Ring) Do(f func(interface{}))

type Ring struct {
    Value interface{} // for use by client; untouched by this library
    // contains filtered or unexported fields
}

//Len计算环r中的元素数。它按与元素数量成比例的时间执行。
//func (r *Ring) Len() int

//链接将环r与环s连接，使r.Next（）成为s并返回r.Next（）的原始值。 r不能为空。

//如果r和s指向同一个环，则将它们链接会从环中删除r和s之间的元素。 移除的元素形成一个子环，结果是对该子环的引用（如果未移除任何元素，则结果仍然是r.Next（）的原始值，而不是nil）。

//如果r和s指向不同的环，将它们链接会创建一个单个环，其中s的元素插入到r之后。 结果指向插入后s的最后一个元素之后的元素。
//func (r *Ring) Link(s *Ring) *Ring

//Move将n％r.Len（）元素在环中向后（n <0）或向前（n> = 0）移动，并返回该环元素。 r不能为空。
//func (r *Ring) Move(n int) *Ring

//Next returns the next ring element. r must not be empty.
//func (r *Ring) Next() *Ring

//Prev returns the previous ring element. r must not be empty.
//func (r *Ring) Prev() *Ring

//Unlink从r.Next（）开始，从环r中删除n％r.Len（）个元素。如果n％r.Len（）== 0，则r保持不变。结果是删除的子环。 r不能为空
//func (r *Ring) Unlink(n int) *Ring
