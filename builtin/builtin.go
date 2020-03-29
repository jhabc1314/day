package builtin

import "fmt"

const (
    T = iota
)
type ss struct {
    name string
    school string
}
//内置的方法常量等集合
func Builtin() {
    fmt.Println(T == 0)

    st := ss{"jh", "milly"}

    //fr := []string{"x", "b"}
    fmt.Printf("%v", st)
    //fmt.Println("world"...)
    s := append([]byte("hello"), "world"...)
    fmt.Println(s)

    //var m map[string]string //nil map 不可赋值
    //m["name"] = "jh" 会报错

    ma := make(map[string]string)
    delete(ma, "name")

    //print("hello world")

}

//close(c chan <- Type) 关闭通道，只能是双向或仅发送通道
//从已经关闭的通道接收数据会立即返回，不阻塞，并返回通道内类型的零值

//complex 复数 
//func imag(c ComplexType) FloatType 返回复数虚部

//内置复制功能将元素从源切片复制到目标切片。 （作为一种特殊情况，它还会将字节从字符串复制到字节的一部分。）源和目标可能会重叠。 复制返回复制的元素数量，该数量将是len（src）和len（dst）的最小值。
//func copy(dst, src []Type) int

//The delete built-in function deletes the element with the specified key (m[key]) from the map. If m is nil or there is no such element, delete is a no-op.
//func delete(m map[Type]Type1, key Type)

//make内置函数分配和初始化slice，map或chan类型的对象（仅）。像new一样，第一个参数是类型，而不是值。与new不同，make的返回类型与其参数的类型相同，而不是指向它的指针。
//func make(t Type, size ...IntegerType) Type

//新的内置函数分配内存。第一个参数是类型，而不是值，返回的值是指向该类型新分配的零值的指针。
//func new(Type) *Type

//实数内置函数返回复数c的实数部分。返回值将是对应于c类型的浮点类型。
//func real(c ComplexType) FloatType

//错误内置接口类型是用于表示错误状态的常规接口，其中nil值表示没有错误。
type error interface {
    Error() string
}

//uintptr是一个整数类型，其大小足以容纳任何指针的位模式。
//type uintptr uintptr
