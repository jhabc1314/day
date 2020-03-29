package bytes

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

//包字节实现用于操作字节片的功能。它类似于字符串包的功能。
const (
	//MinRead是Buffer.ReadFrom传递给Read调用的最小切片大小。只要缓冲区的MinRead字节至少超出保留r的内容所需的字节，ReadFrom就不会增长底层缓冲区。
	MinRead = bytes.MinRead //512
)

//ErrTooLarge is passed to panic if memory cannot be allocated to store data in a buffer.
var ErrTooLarge = errors.New("Bytes.buffer:too large")

//Compare返回按字典顺序比较两个字节片的整数。如果a == b，结果将为0，如果a <b，则结果为-1，如果a> b，则结果为+1。 nil参数等效于一个空切片。
//func Compare(a, b []byte) int
//var a, b []byte
	// if bytes.Compare(a, b) < 0 {
	// 	// a less b
	// }

//包含报告子切片是否在b之内。
//func Contains(b, subslice []byte) bool
//eg fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))

//ContainsAny报告char中任何UTF-8编码的代码点是否在b之内。
//func ContainsAny(b []byte, chars string) bool GO1.7
//fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))

//ContainsRune报告该符文是否包含在UTF-8编码的字节片b中
//func ContainsRune(b []byte, r rune) bool
//eg fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))

//Count对s中的sep的非重叠实例进行计数。如果sep是一个空片，则Count返回1 +以s为单位的UTF-8编码的代码点数。 要求：统计非重叠，如果sep空则返回s的长度+1
//func Count(s, sep []byte) int
//fmt.Println(bytes.Count([]byte("cheese"), []byte("e"))) 3
//fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune 5
//fmt.Println(bytes.Count([]byte("ababa"),[]byte("aba"))) 1

//相等报告a和b是否长度相同且包含相同字节。 nil参数等效于一个空切片。
//func Equal(a, b []byte) bool

//EqualFold报告在Unicode大小写折叠下s和t（解释为UTF-8字符串）是否相等，这是不区分大小写的更通用形式。
//func EqualFold(s, t []byte) bool
//fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go"))) true 

//字段将s解释为一系列UTF-8编码的代码点。 它按照unicode.IsSpace的定义在一个或多个连续空白字符的每个实例周围分割slice s，返回s的子切片的切片；如果s仅包含空白，则返回空切片。
//等于按空格来分隔字节切片，多个空格都会剔除 返回分隔后的二维字节切片
//func Fields(s []byte) [][]byte
//fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
//Fields are: ["foo" "bar" "baz"]

//FieldsFunc将s解释为UTF-8编码的代码点的序列。 它将在满足f（c）的每个代码点c处分割切片s，并返回s的子切片的切片。 如果s中的所有代码点都满足f（c）或len（s）== 0，则返回一个空切片。 FieldsFunc不保证其调用f（c）的顺序。 如果f对于给定的c没有返回一致的结果，则FieldsFunc可能会崩溃。
//按给定的回调函数来分割字节切片
//func FieldsFunc(s []byte, f func(rune) bool) [][]byte

//HasPrefix tests whether the byte slice s begins with prefix.
//func HasPrefix(s, prefix []byte) bool
//fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go"))) true
//fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C"))) false
//fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte(""))) true !!!

//是否以指定切片结尾
//func HasSuffix(s, suffix []byte) bool

//Index返回s中sep的第一个实例的索引；如果s中不存在sep，则返回-1。
//func Index(s, sep []byte) int

//返回chars字符串中的任意一个字符第一次出现在s中的位置，如果没有或者chars为空则返回-1
//func IndexAny(s []byte, chars string) int
//fmt.Println(bytes.IndexAny([]byte("chicken"), "naheoiuy")) 1 h在chicken中最先出现，位置为1

//IndexByte返回b中c的第一个实例的索引；如果b中不存在c，则返回-1。
//func IndexByte(b []byte, c byte) int
//fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k'))) 直接'k' 也可以

//IndexFunc将s解释为UTF-8编码的代码点的序列。它返回满足f（c）的第一个Unicode代码点的s中的字节索引，如果没有，则返回-1。
//func IndexFunc(s []byte, f func(r rune) bool) int
//f := func(c rune) bool {
	//return unicode.Is(unicode.Han, c)
//}
//fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f)) 7 汉字的第一次出现位置
//fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f)) -1

//IndexRune将s解释为UTF-8编码的代码点的序列。 它返回给定符文中s中第一次出现的字节索引。 如果s中不存在符文，则返回-1。 如果r为utf8.RuneError，它将返回任何无效UTF-8字节序列的第一个实例。
//func IndexRune(s []byte, r rune) int
//fmt.Println(bytes.IndexRune([]byte("chicke我n"), '我')) 6 注意单引号
//fmt.Println(bytes.IndexRune([]byte("chicken"), 'd')) -1 

//Join将s的元素连接起来以创建一个新的字节片。分隔符sep放置在所得切片中的元素之间。 以sep拼接s implode
//func Join(s [][]byte, sep []byte) []byte
//s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
//fmt.Printf("%s", bytes.Join(s, []byte(", "))) foo, bar, baz

//LastIndex返回s中sep的最后一个实例的索引；如果s中不存在sep，则返回-1。
//func LastIndex(s, sep []byte) int

//LastIndexAny将s解释为UTF-8编码的Unicode代码点的序列。它返回char中任何Unicode码点中s中最后一次出现的字节索引。如果char为空或没有共同的代码点，则返回-1。
//func LastIndexAny(s []byte, chars string) int

//LastIndexByte返回s中c的最后一个实例的索引；如果s中不存在c，则返回-1。
//func LastIndexByte(s []byte, c byte) int

//LastIndexFunc将s解释为UTF-8编码的代码点的序列。它返回满足f（c）的最后一个Unicode代码点的s中的字节索引，如果没有，则返回-1。
//func LastIndexFunc(s []byte, f func(r rune) bool) int

//Map返回字节片s的副本，其所有字符都根据映射函数进行了修改。 如果映射返回负值，则将字符从字节片中丢弃，并且不进行替换。 s中的字符和输出被解释为UTF-8编码的代码点。
//func Map(mapping func(r rune) rune, s []byte) []byte

//重复返回一个新的字节切片，该切片由b的计数副本组成。 如果count为负或（len（b）* count）的结果溢出，它会感到恐慌。
//func Repeat(b []byte, count int) []byte

//替换返回切片s的副本，其中旧的前n个非重叠实例被新的替换。 如果old为空，则它在切片的开头和每个UTF-8序列之后匹配，最多产生k个符文切片的k + 1个替换。 如果n <0，则替换次数没有限制。
//func Replace(s, old, new []byte, n int) []byte

//ReplaceAll返回slice的副本，其中所有旧的非重叠实例都被new替换。如果old为空，则它在切片的开头和每个UTF-8序列之后匹配，最多可产生k个符文切片的k + 1个替换。
//func ReplaceAll(s, old, new []byte) []byte

//符文将s解释为UTF-8编码的代码点的序列。它返回相当于s的一部分符文（Unicode代码点）。
//func Runes(s []byte) []rune

//将片段s分割为所有由sep分隔的子片段，并返回这些分隔符之间的子片段的片段。如果sep为空，则Split在每个UTF-8序列之后拆分。它等效于SplitN，计数为-1。
//func Split(s, sep []byte) [][]byte

//在Sep的每个实例之后，SplitAfter将s切片为所有子切片，并返回这些子切片的切片。如果sep为空，则SplitAfter在每个UTF-8序列之后拆分。它等效于SplitAfterN，计数为-1。 保留分隔用的字节片
//func SplitAfter(s, sep []byte) [][]byte

//SplitAfterN slices s into subslices after each instance of sep and returns a slice of those subslices. If sep is empty, SplitAfterN splits after each UTF-8 sequence. The count determines the number of subslices to return:
//n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//n == 0: the result is nil (zero subslices)
//n < 0: all subslices
//func SplitAfterN(s, sep []byte, n int) [][]byte

//SplitN slices s into subslices separated by sep and returns a slice of the subslices between those separators. If sep is empty, SplitN splits after each UTF-8 sequence. The count determines the number of subslices to return:
//n > 0: at most n subslices; the last subslice will be the unsplit remainder.
//n == 0: the result is nil (zero subslices)
//n < 0: all subslices
//func SplitN(s, sep []byte, n int) [][]byte

//Title treats s as UTF-8-encoded bytes and returns a copy with all Unicode letters that begin words mapped to their title case.
//所有单词首字母大写
//BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
//func Title(s []byte) []byte

//全部转小写
//func ToLower(s []byte) []byte

//按指定规则转换 特定位置字母为小写
//func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte

//ToTitle将s视为UTF-8编码的字节，并返回一个副本，其中所有Unicode字母均映射到其标题大小写。
//func ToTitle(s []byte) []byte

//func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte

//全部转大写
//func ToUpper(s []byte) []byte

//func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte

//ToValidUTF8将s视为UTF-8编码的字节，并返回一个副本，其中每次运行的字节均表示无效的UTF-8，并替换为替换中的字节，该字节可以为空。
//func ToValidUTF8(s, replacement []byte) []byte

//Trim通过切掉cutset中包含的所有前导和尾随UTF-8编码的代码点来返回s的子片段。 去除所有包含的
//func Trim(s []byte, cutset string) []byte
//fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
//["Achtung! Achtung"] 前后的 "! " 全部被去掉,不限顺序啥的，单字节匹配

//TrimFunc通过分割满足f（c）的所有前导和尾随UTF-8编码的代码点c来返回s的子片段。
//func TrimFunc(s []byte, f func(r rune) bool) []byte

//TrimLeft通过切掉cutset中包含的所有前导UTF-8编码的代码点来返回s的子片段。 去除所有包含的
//func TrimLeft(s []byte, cutset string) []byte

//func TrimRight(s []byte, cutset string) []byte

//TrimPrefix返回s，而没有提供的前导前缀字符串。如果s不以前缀开头，则s不变返回。 完整匹配，和ltrim一致
//func TrimPrefix(s, prefix []byte) []byte
//和rtrim一致
//func TrimSuffix(s, suffix []byte) []byte

//去除前后的所有空白字符 包括任意 \n这种
//func TrimSpace(s []byte) []byte



//Buffer是具有Read和Write方法的可变大小的字节缓冲区。 Buffer的零值是准备使用的空缓冲区。
type Buffer struct {

} 
//NewBuffer使用buf作为其初始内容创建并初始化一个新的Buffer。 新的Buffer拥有buf的所有权，并且在此调用之后，调用方不应使用buf。 NewBuffer旨在准备一个Buffer以读取现有数据。 它也可以用来设置用于写入的内部缓冲区的初始大小。 为此，buf应该具有所需的容量，但长度为零。

//在大多数情况下，new（Buffer）（或仅声明一个Buffer变量）足以初始化Buffer。
//func NewBuffer(buf []byte) *Buffer

//NewBufferString使用字符串s作为其初始内容创建并初始化一个新的Buffer。 目的是准备一个缓冲区以读取现有的字符串。
//在大多数情况下，new（Buffer）（或仅声明一个Buffer变量）足以初始化Buffer。
//func NewBufferString(s string) *Buffer

//字节返回长度为b.Len（）的切片，其中包含缓冲区的未读部分。 该切片仅在下一次修改缓冲区之前有效（即，仅在下一次调用诸如Read，Write，Reset或Truncate之类的方法之前）才有效。 切片至少在下一次缓冲区修改之前就将缓冲区内容作为别名，因此对切片的立即更改将影响将来读取的结果。
//func (b *Buffer) Bytes() []byte

//容量
//func (b *Buffer) Cap() int
//如有必要，可以增加缓冲区的容量，以保证另外n个字节的空间。 在Grow（n）之后，至少可以将n个字节写入缓冲区，而无需进行其他分配。 如果n为负，则增长会惊慌。 如果缓冲区无法增长，则会因ErrTooLarge感到恐慌
//func (b *Buffer) Grow(n int)

//Len returns the number of bytes of the unread portion of the buffer; b.Len() == len(b.Bytes()).
//func (b *Buffer) Len() int

//如有必要，可以增加缓冲区的容量，以保证另外n个字节的空间。 在Grow（n）之后，至少可以将n个字节写入缓冲区，而无需进行其他分配。 如果n为负，则增长会惊慌。 如果缓冲区无法增长，则会因ErrTooLarge感到恐慌。
//func (b *Buffer) Next(n int) []byte

//读取从缓冲区中读取下一个len（p）字节，或者直到缓冲区耗尽为止。 返回值n是读取的字节数。 如果缓冲区没有数据要返回，则err为io.EOF（除非len（p）为零）；否则为。 否则为零。
//func (b *Buffer) Read(p []byte) (n int, err error)

//ReadByte reads and returns the next byte from the buffer. If no byte is available, it returns error io.EOF.
//func (b *Buffer) ReadByte() (byte, error)

//ReadBytes读取直到输入中第一次出现delim为止，并返回一个切片，该切片包含直到定界符（包括定界符）的数据。 如果ReadBytes在找到定界符之前遇到错误，它将返回错误之前读取的数据和错误本身（通常为io.EOF）。 当且仅当返回的数据未以delim结尾时，ReadBytes返回err！= nil。
//func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)

//ReadFrom reads data from r until EOF and appends it to the buffer, growing the buffer as needed. The return value n is the number of bytes read. Any error except io.EOF encountered during the read is also returned. If the buffer becomes too large, ReadFrom will panic with ErrTooLarge.
//func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)

//ReadRune从缓冲区读取并返回下一个UTF-8编码的Unicode代码点。如果没有可用的字节，则返回的错误是io.EOF。如果字节是错误的UTF-8编码，则它消耗一个字节并返回U + FFFD，1。
//func (b *Buffer) ReadRune() (r rune, size int, err error)

//ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter. If ReadString encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadString returns err != nil if and only if the returned data does not end in delim.
//func (b *Buffer) ReadString(delim byte) (line string, err error)

//重置会将缓冲区重置为空，但会保留基础存储以供将来的写入操作使用。重置与Truncate（0）相同。
//func (b *Buffer) Reset()

//字符串以字符串形式返回缓冲区未读部分的内容。如果Buffer是nil指针，则返回“ <nil>”。 要更有效地构建字符串，请参见strings.Builder类型。
//func (b *Buffer) String() string

//截断会丢弃缓冲区中除前n个未读取字节以外的所有字节，但会继续使用相同的已分配存储。如果n为负数或大于缓冲区的长度，则会发生恐慌。
//func (b *Buffer) Truncate(n int)

//塞回上一次读取的最后一个字节回缓冲区
//func (b *Buffer) UnreadByte() error
//最后一个读取的utf8
//func (b *Buffer) UnreadRune() error

//写操作会将p的内容附加到缓冲区，并根据需要扩展缓冲区。返回值n是p的长度；错误始终为零。如果缓冲区太大，则Write会因ErrTooLarge感到恐慌。
//func (b *Buffer) Write(p []byte) (n int, err error)
//写一个字节到缓冲区
//func (b *Buffer) WriteByte(c byte) error

//func (b *Buffer) WriteRune(r rune) (n int, err error)
//func (b *Buffer) WriteString(s string) (n int, err error)

//WriteTo将数据写入w，直到缓冲区耗尽或发生错误。 返回值n是写入的字节数。 它始终适合int，但与io.WriterTo接口匹配为int64。 写入期间遇到的任何错误也将返回。
//func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)

//Reader通过读取字节片来实现io.Reader，io.ReaderAt，io.WriterTo，io.Seeker，io.ByteScanner和io.RuneScanner接口。 与Buffer不同，Reader是只读的，并支持查找。 Reader的零值的操作类似于空切片的Reader。
type Reader struct {
	
}
//NewReader returns a new Reader reading from b.
//func NewReader(b []byte) *Reader

//Len returns the number of bytes of the unread portion of the slice.
//func (r *Reader) Len() int

//
//func (r *Reader) Read(b []byte) (n int, err error)

//func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

//
//func (r *Reader) ReadByte() (byte, error)
//func (r *Reader) ReadRune() (ch rune, size int, err error)
//重置将读取器重置为从b读取。
//func (r *Reader) Reset(b []byte)
//func (r *Reader) Seek(offset int64, whence int) (int64, error)
//func (r *Reader) Size() int64
//func (r *Reader) UnreadByte() error
//func (r *Reader) UnreadRune() error
//func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

func Byte() {
	fmt.Println(string(bytes.Trim([]byte(" hello world "), " h")))
	//新建一个缓冲区

	b := bytes.NewBuffer([]byte("你好啊，世界"))
	fmt.Println(b.String(),b.Cap(),b.Len())
	for {
		l,err := b.ReadString('\n');
		if err == io.EOF {
			break;
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(l)
	}

}
