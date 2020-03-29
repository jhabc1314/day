package compress

import "io"

//flate 压缩算法

const (
	NoCompression = 0
	BestSpeed = 1
	BestCompression = 9
	DefaultCompression = -1
	//HuffmanOnly禁用Lempel-Ziv匹配搜索，仅执行Huffman熵编码。 此模式在压缩已使用缺乏熵编码器的LZ样式算法（例如Snappy或LZ4）压缩的数据时很有用。 当输入流中的某些字节比其他字节更频繁出现时，可获得压缩增益。
	//请注意，HuffmanOnly会生成符合RFC 1951的压缩输出。 也就是说，任何有效的DEFLATE解压缩器将继续能够解压缩此输出。
	HuffmanOnly = -2
)

//NewReader返回一个新的ReadCloser，可用于读取r的未压缩版本。 如果r还没有实现io.ByteReader，则解压缩器可能会从r读取比所需更多的数据。 完成阅读后，调用方有责任在ReadCloser上调用Close。
//NewReader返回的ReadCloser也实现了Resetter。
//func NewReader(r io.Reader) io.ReadCloser

//NewReaderDict类似于NewReader，但是使用预设字典初始化阅读器。 返回的Reader的行为就像未压缩的数据流以给定的字典（已被读取）开始。 NewReaderDict通常用于读取由NewWriterDict压缩的数据。

//NewReader返回的ReadCloser也实现了Resetter。
//func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser

//CorruptInputError 报告在给定偏移量下存在损坏的输入。
type CorruptInputError int64

//func (e CorruptInputError) Error() string

//一个InternalError报告Flate代码本身中的错误。
type InternalError string

//func (e InternalError) Error() string

//ReadError报告读取输入时遇到的错误。
type ReadError struct {
    Offset int64 // byte offset where error occurred
    Err    error // error returned by underlying Read
}
//func (e *ReadError) Error() string

//重置程序重置由NewReader或NewReaderDict返回的ReadCloser，以切换到新的基础Reader。这允许重新使用ReadCloser而不是分配新的ReadCloser。 GO1.4
type Resetter interface {
    // Reset discards any buffered data and resets the Resetter as if it was
    // newly initialized with the given reader.
    Reset(r io.Reader, dict []byte) error
}

//WriteError报告写入输出时遇到的错误 不推荐使用：不再返回。
type WriteError struct {
    Offset int64 // byte offset where error occurred
    Err    error // error returned by underlying Write
}

//Writer接收写入其中的数据，并将该数据的压缩形式写入基础写入器（请参见NewWriter）。
type Writer struct {

}

//NewWriter返回给定级别的新Writer压缩数据。 在zlib之后，级别范围从1（BestSpeed）到9（BestCompression）； 较高的级别通常运行较慢，但压缩程度更高。 级别0（NoCompression）不尝试任何压缩； 它仅添加必要的DEFLATE帧。 级别-1（DefaultCompression）使用默认压缩级别。 级别-2（仅HuffmanOnly）将仅使用Huffman压缩，可以对所有类型的输入进行非常快速的压缩，但会牺牲可观的压缩效率。

//如果级别在[-2，9]范围内，则返回的错误为零。 否则，返回的错误将为非零。
//func NewWriter(w io.Writer, level int) (*Writer, error)

//NewWriterDict类似于NewWriter，但使用预设字典初始化新Writer。 返回的Writer的行为就像已将字典写入其中而不产生任何压缩的输出。 写入w的压缩数据只能由使用同一词典初始化的Reader进行解压缩。
//func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)

//func (w *Writer) Close() error

//刷新将所有未决数据刷新到基础写入器。 主要在压缩网络协议中有用，以确保远程读取器具有足够的数据来重构数据包。 在写入数据之前，刷新不会返回。 当没有待处理的数据时调用Flush仍然会使Writer发出至少4个字节的同步标记。 如果基础编写器返回错误，则Flush返回该错误。

//在zlib库的术语中，Flush等效于Z_SYNC_FLUSH。
//func (w *Writer) Flush() error

//重置将丢弃编写器的状态，并使其等效于用dst和w的级别和字典调用的NewWriter或NewWriterDict的结果。
//func (w *Writer) Reset(dst io.Writer)

//Write将数据写入w，最终将压缩的数据形式写入其基础写入器。
//func (w *Writer) Write(data []byte) (n int, err error)
