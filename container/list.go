package container

//软件包列表实现了双向链接列表。

//元素是链接列表的元素。
type Element struct {

    // The value stored with this element.
    Value interface{}
    // contains filtered or unexported fields
}

//Next返回下一个列表元素或nil。
//func (e *Element) Next() *Element

//Prev returns the previous list element or nil.
//func (e *Element) Prev() *Element

//列表表示双链表。 List的零值是可以使用的空列表。
type List struct {
    // contains filtered or unexported fields
}

//New返回一个初始化列表。
//func New() *List

//Back返回列表l的最后一个元素；如果列表为空，则返回nil。
//func (l *List) Back() *Element

//如果列表为空，则Front返回列表l的第一个元素或nil。
//func (l *List) Front() *Element

//初始化初始化或清除列表l。
//func (l *List) Init() *List

//InsertAfter在标记之后立即插入一个值为v的新元素e并返回e。如果mark不是l的元素，则不修改列表。mark 不能为零。
//func (l *List) InsertAfter(v interface{}, mark *Element) *Element

//InsertBefore在标记之前插入一个值为v的新元素e并返回e。如果mark不是l的元素，则不修改列表。mark不能为零。
//func (l *List) InsertBefore(v interface{}, mark *Element) *Element

//Len返回列表l的元素数。复杂度为O（1）。
//func (l *List) Len() int

//MoveAfter将元素e移到mark之后的新位置。如果e或mark不是l的元素，或者e == mark，则不修改列表。元素和标记不得为零。
//func (l *List) MoveAfter(e, mark *Element)

//func (l *List) MoveBefore(e, mark *Element)

//MoveToBack moves element e to the back of list l. If e is not an element of l, the list is not modified. The element must not be nil.
//func (l *List) MoveToBack(e *Element)

//func (l *List) MoveToFront(e *Element)

//PushBack inserts a new element e with value v at the back of list l and returns e.
//func (l *List) PushBack(v interface{}) *Element

//PushBackList在列表l的后面插入另一个列表的副本。列表l和其他列表可以相同。他们一定不能为零。
//func (l *List) PushBackList(other *List)

//func (l *List) PushFront(v interface{}) *Element
//func (l *List) PushFrontList(other *List)

//func (l *List) Remove(e *Element) interface{}

