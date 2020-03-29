package compress

import "compress/flate"
//包gzip实现了RFC 1952中指定的gzip格式压缩文件的读写。

const (
    NoCompression      = flate.NoCompression
    BestSpeed          = flate.BestSpeed
    BestCompression    = flate.BestCompression
    DefaultCompression = flate.DefaultCompression
    HuffmanOnly        = flate.HuffmanOnly
)

var (
    // ErrChecksum is returned when reading GZIP data that has an invalid checksum.
    ErrChecksum = errors.New("gzip: invalid checksum")
    // ErrHeader is returned when reading GZIP data that has an invalid header.
    ErrHeader = errors.New("gzip: invalid header")
)
//gzip文件存储一个标头，该标头提供有关压缩文件的元数据。 该标头作为Writer和Reader结构的字段公开。

//字符串必须是UTF-8编码的，并且由于GZIP文件格式的限制，只能包含Unicode代码点U + 0001至U + 00FF。
type Header struct {
    Comment string    // comment
    Extra   []byte    // "extra data"
    ModTime time.Time // modification time
    Name    string    // file name
    OS      byte      // operating system type
}

//Reader是io.Reader，可以读取该文件以从gzip格式的压缩文件中检索未压缩的数据。

//通常，gzip文件可以是gzip文件的串联，每个文件都有自己的标头。 从Reader读取将返回每个未压缩数据的串联。 在阅读器字段中仅记录第一个标头。

//Gzip文件存储未压缩数据的长度和校验和。 如果读取的数据没有预期的长度或校验和，则当读取到达未压缩数据的末尾时，读取器将返回ErrChecksum。 在收到io.EOF标记数据结束之前，客户应将Read返回的数据视为临时的。
type Reader struct {
    Header // valid after NewReader or Reader.Reset
    // contains filtered or unexported fields
}

//NewReader创建一个读取给定阅读器的新阅读器。 如果r还没有实现io.ByteReader，则解压缩器可能会从r读取比所需更多的数据。

//完成后，调用方有责任在Reader上调用Close。

//Reader.Header字段在返回的Reader中有效。
//func NewReader(r io.Reader) (*Reader, error)

//func (z *Reader) Close() error

/**
多流控制阅读器是否支持多流文件。

如果启用（默认），则Reader希望输入是一系列单独压缩的数据流，每个数据流都有自己的标题和结尾，以EOF结尾。 效果是，压缩文件序列的串联被视为等同于该序列串联的gzip。 这是gzip阅读器的标准行为。

调用Multistream（false）会禁用此行为； 当读取区分单个gzip数据流或将gzip数据流与其他数据流区分开的文件格式时，禁用该行为可能很有用。 在这种模式下，当Reader到达数据流的末尾时，Read返回io.EOF。 基础阅读器必须实现io.ByteReader才能将其定位在gzip流之后。 要开始下一个流，请调用z.Reset（r），然后调用z.Multistream（false）。 如果没有下一个流，则z.Reset（r）将返回io.EOF。
*/
//func (z *Reader) Multistream(ok bool) GO1.4

//
//func (z *Reader) Read(p []byte) (n int, err error)

//重置将丢弃Reader z的状态，并使其等效于从NewReader读取其原始状态的结果，但改为从r读取。这允许重用Reader，而不是分配新的Reader。
//func (z *Reader) Reset(r io.Reader) error

//Writer是io.WriteCloser。对Writer的写入将被压缩并写入w。
type Writer struct {
    Header // written at first call to Write, Flush, or Close
    // contains filtered or unexported fields
}

//NewWriter返回一个新的Writer。 对返回的写入器的写入将被压缩并写入w。

//完成后，调用方有责任在Writer上调用Close。 写入可能会被缓冲，直到关闭才刷新。

//希望在Writer.Header中设置字段的调用者必须在首次调用Write，Flush或Close之前进行设置。
//func NewWriter(w io.Writer) *Writer

//NewWriterLevel类似于NewWriter，但是指定压缩级别，而不是假定DefaultCompression。

//压缩级别可以是DefaultCompression，NoCompression，HuffmanOnly或BestSpeed和BestCompression之间的任何整数值。 如果级别有效，则返回的错误将为nil。
//func NewWriterLevel(w io.Writer, level int) (*Writer, error)

//func (z *Writer) Close() error
//func (z *Writer) Flush() error
//func (z *Writer) Reset(w io.Writer)
//func (z *Writer) Write(p []byte) (int, error)
