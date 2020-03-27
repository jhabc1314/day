package bufio

import (
	"bufio"
	"errors"
)

//包bufio实现了缓冲的I / O。它包装了一个io.Reader或io.Writer对象，创建了另一个对象（Reader或Writer），该对象也实现了该接口，但提供了缓冲和一些有关文本I / O的帮助。
//可以更高效率的读写文件

const (
	//MaxScanTokenSize是用于缓冲令牌的最大大小，除非用户使用Scanner.Buffer提供显式缓冲区。
	//由于缓冲区可能需要包含例如换行符，因此实际最大令牌大小可能会更小。
	MaxScanTokenSize = bufio.MaxScanTokenSize
)
var (
	ErrInvalidUnreadByte = errors.New("bufio:invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio:invalid use of unreadRune")
	ErrBufferFull = errors.New("bufio:buffer full")
	ErrNegativeCount = errors.New("bufio:negative count") //负数
)

//Errors returned by Scanner.
var (
	ErrTooLong = errors.New("bufio.Scanner: token too long")
	//SplitFunc返回负提前计数
	ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	//SplitFunc返回超出输入的提前计数
	ErrAdvanceTooBar = errors.New("bufio.Scanner: SplitFunc returns advance count beyond count")
)

var ErrFinalToken = errors.New("final token")

//ScanBytes is a split function for a Scanner that returns each byte as a token.
//func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)

//ScanLines是Scanner的拆分功能，它返回文本的每一行，并删除所有行尾标记。 返回的行可能为空。 行尾标记是一个可选的回车符，后跟一个强制换行符。 在正则表达式中，它是`\ r？\ n`。 即使没有换行符，也将返回输入的最后一个非空行。
//func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)

//rune 本质为 uint32 用来标识utf8编码 byte 用来表示原始数据 eg 39 为 '0'
//ScanRunes是扫描程序的拆分功能，它会将每个UTF-8编码的符文作为令牌返回。 返回的符文序列与输入中的范围循环作为字符串等效，这意味着错误的UTF-8编码会转换为U + FFFD =“ \ xef \ xbf \ xbd”。 由于具有扫描接口，因此客户端无法将正确编码的替换符文与编码错误区分开。
//func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)

//ScanWords是扫描程序的拆分功能，它返回每个空格分隔的文本单词，并删除周围的空格。它永远不会返回空字符串。空间的定义由unicode.IsSpace设置。
//func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)

//ReadWriter 存储指向读取器和写入器的指针。它实现io.ReadWriter。
type ReadWriter struct {
	*Writer
	*Reader
}

//Reader为io.Reader对象实现缓冲。
type Reader struct {

}

//func NewReadWriter(r *Reader, w *Writer) *ReadWriter


//func NewReader(rd io.Reader) *Reader

//NewReaderSize返回一个新的Reader，其缓冲区至少具有指定的大小。如果参数io.Reader已经是足够大的Reader，则它将返回基础Reader。
//func NewReaderSize(rd io.Reader, size int) *Reader

//Buffered returns the number of bytes that can be read from the current buffer.
//func (b *Reader) Buffered() int

//Discard跳过接下来的n个字节，返回丢弃的字节数。

//如果“放弃”跳过少于n个字节，则它还会返回错误。 如果0 <= n <= b.Buffered（），则确保Discard成功执行而无需从基础io.Reader中读取。
//func (b *Reader) Discard(n int) (discarded int, err error)

//Peek返回下一个n个字节，而不会使阅读器前进。 在下一个读取调用中，字节停止有效。 如果Peek返回的字节数少于n个字节，则它还会返回一个错误，解释读取短的原因。 如果n大于b的缓冲区大小，则错误为ErrBufferFull。

//调用Peek会阻止UnreadByte或UnreadRune调用成功，直到下一次读取操作为止。
//func (b *Reader) Peek(n int) ([]byte, error)

//读取将数据读入p。 它返回读入p的字节数。 这些字节是从基础读取器上的最多一个读取中获取的，因此n可能小于len（p）。 要精确读取len（p）个字节，请使用io.ReadFull（b，p）。 在EOF时，计数将为零，而err将为io.EOF。
//func (b *Reader) Read(p []byte) (n int, err error)

//ReadByte读取并返回一个字节。如果没有可用的字节，则返回错误
//func (b *Reader) ReadByte() (byte, error)

//ReadBytes读取直到输入中第一次出现delim为止，并返回一个切片，该切片包含直到定界符（包括定界符）的数据。 如果ReadBytes在找到定界符之前遇到错误，它将返回错误之前读取的数据和错误本身（通常为io.EOF）。 当且仅当返回的数据未以delim结尾时，ReadBytes返回err！= nil。 对于简单的用途，Scanner 可能更方便。
//func (b *Reader) ReadBytes(delim byte) ([]byte, error)

//ReadLine是低级别的行读取原语。 大多数调用者应改用ReadBytes（'\ n'）或ReadString（'\ n'）或使用扫描仪。
//ReadLine尝试返回单行，不包括行尾字节。 如果该行对于缓冲区而言太长，则设置isPrefix并返回该行的开头。 该行的其余部分将从以后的呼叫中返回。 返回行的最后一个片段时，isPrefix将为false。 返回的缓冲区仅在下一次调用ReadLine之前有效。 ReadLine要么返回非空行，要么返回错误，从不都返回。
//从ReadLine返回的文本不包含行尾（“ \ r \ n”或“ \ n”）。 如果输入在没有最后一行结束的情况下结束，则不会给出任何指示或错误。 在ReadLine之后调用UnreadByte将始终不读取最后一个读取的字节（可能是属于行尾的字符），即使该字节不属于ReadLine返回的行的一部分。
//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

//ReadRune读取单个UTF-8编码的Unicode字符，并返回符文及其大小（以字节为单位）。 如果编码的符文无效，则它将占用一个字节并返回unicode.ReplacementChar（U + FFFD），大小为1。
//func (b *Reader) ReadRune() (r rune, size int, err error)

//
//func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
