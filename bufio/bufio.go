package bufio

import (
	"bufio"
	"errors"
	"fmt"
	//"io"
	"os"
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
	ErrBufferFull        = errors.New("bufio:buffer full")
	ErrNegativeCount     = errors.New("bufio:negative count") //负数
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

//ReadSlice读取直到输入中第一次出现delim为止，返回一个指向缓冲区中字节的切片。 字节在下一次读取时不再有效。 如果ReadSlice在找到定界符之前遇到错误，它将返回缓冲区中的所有数据以及错误本身（通常为io.EOF）。 如果缓冲区填充不带delim，ReadSlice将失败，并显示错误ErrBufferFull。 由于从ReadSlice返回的数据将被下一个I / O操作覆盖，因此大多数客户端应改用ReadBytes或ReadString。 当且仅当行不以delim结尾时，ReadSlice才返回err！= nil。
//func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

//ReadString读取直到输入中第一次出现delim为止，并返回一个字符串，其中包含直到定界符（包括定界符）的数据。 如果ReadString在找到定界符之前遇到错误，它将返回错误之前读取的数据和错误本身（通常为io.EOF）。 当且仅当返回的数据未以delim结尾时，ReadString才返回err！= nil。 对于简单的用途，扫描仪可能更方便。
//func (b *Reader) ReadString(delim byte) (string, error)

//重置将丢弃所有缓冲的数据，重置所有状态，并将缓冲的读取器切换为从r读取。
//func (b *Reader) Reset(r io.Reader)

//Size返回基础缓冲区的大小（以字节为单位）。
//func (b *Reader) Size() int

//UnreadByte unreads the last byte. Only the most recently read byte can be unread.

//UnreadByte returns an error if the most recent method called on the Reader was not a read operation. Notably, Peek is not considered a read operation.
//func (b *Reader) UnreadByte() error

//UnreadRune unreads the last rune. If the most recent method called on the Reader was not a ReadRune, UnreadRune returns an error. (In this regard it is stricter than UnreadByte, which will unread the last byte from any read operation.)
//func (b *Reader) UnreadRune() error

//WriteTo实现io.WriterTo。这可能会多次调用基础Reader的Read方法。如果基础Reader支持WriteTo方法，则此方法无需缓冲即可调用基础WriteTo。
//func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

//扫描仪提供了一个方便的界面来读取数据，例如用换行符分隔的文本行的文件。 连续调用Scan方法将逐步浏览文件的“令牌”，跳过令牌之间的字节。 令牌的规范是由SplitFunc类型的split函数定义的； 默认的split功能将输入分成几行，剥去了行终端。 在此软件包中定义了拆分功能，用于将文件扫描为行，字节，UTF-8编码的符文和空格分隔的单词。 客户端可以改为提供自定义拆分功能。
//扫描将在EOF，第一个I / O错误或令牌太大而无法容纳在缓冲区中时停止恢复。 扫描停止时，读取器可能会任意向前移动，超过最后一个令牌。 需要对错误处理或较大令牌进行更多控制的程序，或者必须在读取器上运行顺序扫描的程序，应改用bufio.Reader。
type Scanner struct {
}

//NewScanner返回一个新的Scanner以从r中读取。 拆分功能默认为ScanLines。
//func NewScanner(r io.Reader) *Scanner

//缓冲区设置扫描时要使用的初始缓冲区以及在扫描过程中可能分配的缓冲区的最大大小。 最大令牌大小是max和cap（buf）中的较大者。 如果max <= cap（buf），Scan将仅使用此缓冲区，并且不分配。
//默认情况下，扫描使用内部缓冲区并将最大令牌大小设置为MaxScanTokenSize。
//Buffer panics if it is called after scanning has started.
//func (s *Scanner) Buffer(buf []byte, max int)

//字节返回通过调用Scan生成的最新令牌。基础数组可能指向将由后续对Scan的调用覆盖的数据。它不分配。
//func (s *Scanner) Bytes() []byte

//Err返回扫描仪遇到的第一个非EOF错误。
//func (s *Scanner) Err() error

//扫描会将扫描程序前进到下一个令牌，然后可以通过Bytes或Text方法使用该令牌。 当扫描停止（到达输入结尾或错误）时，它将返回false。 在Scan返回false之后，Err方法将返回扫描期间发生的任何错误，除非它是io.EOF，否则Err将返回nil。 如果split函数在不提前输入的情况下返回了太多的空令牌，请扫描恐慌。 这是扫描仪的常见错误模式。
//func (s *Scanner) Scan() bool 相当于调用Next方法

//Split sets the split function for the Scanner. The default split function is ScanLines.
//Split panics if it is called after scanning has started.
//func (s *Scanner) Split(split SplitFunc)

//Text returns the most recent token generated by a call to Scan as a newly allocated string holding its bytes.
//func (s *Scanner) Text() string

//SplitFunc是用于对输入进行标记化的split函数的签名。参数是剩余的未处理数据的初始子字符串，以及标志atEOF，该标志报告Reader是否没有更多数据可提供。返回值是推进输入的字节数，以及返回给用户的下一个令牌（如果有）以及错误（如果有）。
//如果函数返回错误，扫描将停止，在这种情况下，某些输入可能会被丢弃。
//否则，扫描仪将前进输入。如果令牌不是nil，则扫描程序会将其返回给用户。如果令牌为零，则扫描程序将读取更多数据并继续扫描；否则，将继续扫描。如果没有更多数据-如果atEOF为true-扫描程序将返回。如果数据还没有完整的令牌，例如，如果在扫描行时没有换行符，则SplitFunc可以返回（0，nil，nil），以指示扫描程序将更多数据读取到切片中，然后尝试更长的时间切片从输入的同一点开始。
//除非atEOF为true，否则永远不要使用空数据片调用该函数。但是，如果atEOF为true，则数据可能为非空，并且一如既往地保留未处理的文本。
//type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

//Writer为io.Writer对象实现缓冲。 如果在写入Writer时发生错误，将不再接受更多数据，并且所有后续写入和Flush都将返回错误。 写入所有数据之后，客户端应调用Flush方法以确保所有数据都已转发到基础io.Writer。
type Writer struct {
}

//NewWriter returns a new Writer whose buffer has the default size.
//func NewWriter(w io.Writer) *Writer

//NewWriterSize returns a new Writer whose buffer has at least the specified size. If the argument io.Writer is already a Writer with large enough size, it returns the underlying Writer.
//func NewWriterSize(w io.Writer, size int) *Writer

//Available returns how many bytes are unused in the buffer.
//func (b *Writer) Available() int

//Buffered返回已写入当前缓冲区的字节数。
//func (b *Writer) Buffered() int

//Flush writes any buffered data to the underlying io.Writer.
//func (b *Writer) Flush() error

//ReadFrom实现io.ReaderFrom。如果基础编写器支持ReadFrom方法，并且b尚无缓冲数据，则这将调用基础ReadFrom而不进行缓冲。
//func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

//复位将丢弃所有未刷新的缓冲数据，清除所有错误，然后复位b将其输出写入w。
//func (b *Writer) Reset(w io.Writer)

//Size返回基础缓冲区的大小（以字节为单位）
//func (b *Writer) Size() int

//Write writes the contents of p into the buffer. It returns the number of bytes written. If nn < len(p), it also returns an error explaining why the write is short.
//func (b *Writer) Write(p []byte) (nn int, err error)

//WriteByte写入一个字节
//func (b *Writer) WriteByte(c byte) error

//WriteRune写入单个Unicode代码点，返回写入的字节数和任何错误。
//func (b *Writer) WriteRune(r rune) (size int, err error)

//WriteString writes a string. It returns the number of bytes written. If the count is less than len(s), it also returns an error explaining why the write is short.
//func (b *Writer) WriteString(s string) (int, error)

//demo 读取一个文件内容 写入内容到一个文件

func Read(f string) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//fr := bufio.NewReader(file)

	// for {
	// 	l, err := fr.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(l)
	// }
	fs := bufio.NewScanner(file)
	for {
		if fs.Scan() {
			fmt.Println(fs.Text())
		} else {
			break
		}
	}
}

func Write() {
	//相当于重置文件，没有则创建新文件，可读写
	//file, err := os.OpenFile("./bufio.log", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	file,err := os.OpenFile("./cache/bufio.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fw := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		fw.WriteString(fmt.Sprintf("%d 来呀", i))
	}
	err = fw.Flush()
	if err != nil {
		panic(err)
	}
}
